package args

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var defaultProperties = map[string]bool{
	"xkillbillapikey":    true,
	"xkillbillapisecret": true,
	"xkillbillcomment":   true,
	"xkillbillcreatedby": true,
	"xkillbillreason":    true,
}

// Property represents a single key/value pair that is used to collect input
// from user. For ex., in a command, a property represents user input.
// Since many killbill APIs have tons of properties, it is tedious and error prone
// to define each property manually. So, we also provide helpers to load properties
// from a given struct.
type Property struct {
	Name     string // Name of the property. The name should match the struct field name.
	Required bool   // Specifies if the property is required.
	Default  string // Default value for the property.
}

// NameLower returns lowercase  name
func (p Property) NameLower() string {
	return strings.ToLower(p.Name)
}

// Properties - list of properties.
type Properties []Property

// ToMap returns map.
func (pl *Properties) ToMap() map[string]Property {
	m := map[string]Property{}
	for _, p := range *pl {
		m[p.NameLower()] = p
	}
	return m
}

// Remove removes the given property from the list
func (pl *Properties) Remove(propName string) {
	propList := []Property(*pl)
	for i, p := range propList {
		if p.NameLower() == strings.ToLower(propName) {
			*pl = append(propList[0:i], propList[i+1:]...)
			return
		}
	}
	panic(fmt.Errorf("property %s not found", propName))
}

// Get returns the property by name.
func (pl *Properties) Get(propName string) *Property {
	propList := (*[]Property)(pl)
	for i := range *propList {
		p := &(*propList)[i]
		if p.NameLower() == strings.ToLower(propName) {
			return p
		}
	}
	panic(fmt.Sprintf("property %s not found", propName))
}

// Sort sorts the properties
func (pl *Properties) Sort(moveRequiredPropertiesFirst bool, alphabeticSort bool) {
	propList := []Property(*pl)
	sort.SliceStable(propList, func(i, j int) bool {
		x := propList[i]
		y := propList[j]
		if moveRequiredPropertiesFirst {
			if x.Required && !y.Required {
				return j < i
			} else if y.Required && !x.Required {
				return i < j
			}
		}
		if alphabeticSort {
			return strings.Compare(x.NameLower(), y.NameLower()) < 0
		}
		return i < j
	})
}

// loadProperty loads property value to the target struct.
func loadProperty(val reflect.Value, propName string, valueToSet string) bool {
	for i := 0; i < val.Type().NumField(); i++ {
		ft := val.Type().Field(i)
		f := val.FieldByIndex(ft.Index)

		if strings.ToLower(ft.Name) == strings.ToLower(propName) {
			if !f.CanSet() {
				panic(fmt.Errorf("readonly property %s", ft.Name))
			}
			err := setFieldValue(f, ft, valueToSet)
			if err != nil {
				return false
			}
			return true
		}

		if ft.Type.Kind() == reflect.Struct {
			success := loadProperty(f, propName, valueToSet)
			if success {
				return success
			}
		}
	}
	return false
}

// loadPropertiesFromInput loads key value pairs into target object.
// If a required property is not found, then error is raised.
func loadPropertiesFromInput(obj interface{}, properties Properties, inputs Inputs) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	propMap := properties.ToMap()

	// Set up defaults
	for _, p := range properties {
		if p.Default != "" {
			if !loadProperty(val, p.Name, p.Default) {
				return fmt.Errorf("property %s not found in type %s", p.Name, val.Type().String())
			}
		}
	}

	suppliedProperties := map[string]bool{}
	for _, inp := range inputs {
		suppliedProperties[inp.KeyLower()] = true
		p, ok := propMap[inp.KeyLower()]
		if !ok {
			return fmt.Errorf("property %s not found", inp.Key)
		}
		if !loadProperty(val, p.Name, inp.Value) {
			return fmt.Errorf("property %s not found in type %s", p.Name, val.Type().String())
		}
	}

	missingProperties := []string{}
	for _, p := range properties {
		if p.Required && !suppliedProperties[p.NameLower()] {
			missingProperties = append(missingProperties, p.Name)
		}
	}

	if len(missingProperties) > 0 {
		return fmt.Errorf("Required properties are missing: %s", strings.Join(missingProperties, ","))
	}

	return nil
}

// LoadProperties loads properties from input arguments
func LoadProperties(obj interface{}, properties Properties, argsList []string) error {
	inputs, err := ParseArgs(argsList)
	if err != nil {
		return err
	}
	return loadPropertiesFromInput(obj, properties, inputs)
}

// GenerateUsageString generates usage string for given list of properties.
// This can be used to generate usage strings for commands.
//
// For ex., given struct
// type account struct
//   AccountID string
//   CompanyName string
//   Email string
// }
//
// and call like this
//   GenerateUsageString(&account{}, []Property{{Name: "AccountID", Required: true},
//	{Name: "CompanyName"}})
//
// Returns:
//   AccountID=STRING [CompanyName=STRING]
//
func GenerateUsageString(obj interface{}, properties Properties) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	fieldMap := toFieldsMap(getStructFields(val.Type()))

	var result = []string{""}
	for _, p := range properties {
		f, ok := fieldMap[p.NameLower()]
		if !ok {
			panic(fmt.Errorf("property %s not found in object %s", p.Name, val.Type().String()))
		}

		fType := f.Type
		if fType.Kind() == reflect.Ptr {
			fType = fType.Elem()
		}
		handler, ok := supportedTypes[fType]
		if !ok {
			panic(fmt.Errorf("unsupported type %s for property %s", fType.String(), f.Name))
		}

		var line string
		if p.Required {
			line = fmt.Sprintf("%s=%s", f.Name, handler.UsageString)
		} else {
			line = fmt.Sprintf("[%s=%s]", f.Name, handler.UsageString)
		}
		if p.Default != "" {
			line += fmt.Sprintf("    Default: %s", p.Default)
		}
		result = append(result, line)

	}
	return strings.Join(result, "\n         ")
}

// GetProperties returns all properties in a given object
func GetProperties(obj interface{}) Properties {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	fieldList := getStructFields(val.Type())
	var result []Property
	for _, f := range fieldList {
		if _, ok := defaultProperties[strings.ToLower(f.Name)]; !ok {
			result = append(result, Property{Name: f.Name})
		}
	}

	return result
}

// ValidatePropertyList validates given list is valid for a given object
func ValidatePropertyList(obj interface{}, props Properties) {
	propMap := map[string]bool{}
	for _, p := range props {
		propMap[p.NameLower()] = true
	}

	for _, p := range GetProperties(obj) {
		if !propMap[p.NameLower()] {
			panic(fmt.Errorf("Property '%s' doesn't exist in type %s", p.Name, reflect.TypeOf(obj).String()))
		}
	}
}

// getStructFields returns list of struct fields in a given struct.
// All nested anonymous types are also processed
func getStructFields(tp reflect.Type) []reflect.StructField {
	if tp.Kind() != reflect.Struct {
		panic(fmt.Sprintf("invalid object. expecting struct type, got %v instead", tp.Kind().String()))
	}

	var fields []reflect.StructField
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		if isTypeSupported(f.Type) {
			fields = append(fields, f)
		}
		if f.Type.Kind() == reflect.Struct && f.Anonymous {
			fields = append(fields, getStructFields(f.Type)...)
		}
	}

	return fields
}

// toFieldsMap converts list of fields into map.
func toFieldsMap(fields []reflect.StructField) map[string]reflect.StructField {
	m := map[string]reflect.StructField{}
	for _, f := range fields {
		fn := strings.ToLower(f.Name)
		m[fn] = f
	}
	return m
}

// Returns true if the type is supported.
func isTypeSupported(tp reflect.Type) bool {
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return supportedTypes[tp] != nil
}

// setFieldValue parses the given value and sets the value to the target field.
func setFieldValue(f reflect.Value, ft reflect.StructField, val string) error {
	if !isTypeSupported(f.Type()) {
		return fmt.Errorf("type %s not supported", ft.Name)
	}

	tp := f.Type()
	isPtr := f.Type().Kind() == reflect.Ptr
	if isPtr {
		tp = tp.Elem()
	}
	handler := supportedTypes[tp]
	if handler == nil {
		return fmt.Errorf("unknown type %s", tp.String())
	}
	return handler.SetterFn(f, val)
}
