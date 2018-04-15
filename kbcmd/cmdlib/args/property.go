package args

import (
	"fmt"
	"reflect"
	"strings"
)

var defaultProperties = map[string]bool{
	"xkillbillapikey":    true,
	"xkillbillapisecret": true,
	"xkillbillcomment":   true,
	"xkillbillcreatedby": true,
	"xkillbillreason":    true,
}

// loadPropertiesFromInput loads key value pairs into target object.
// property names must match
func loadPropertiesFromInput(obj interface{}, properties Properties, inputs Inputs) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	fieldMap := toFieldsMap(getStructFields(val.Type()))
	propMap := properties.ToMap()

	for _, inp := range inputs {

		_, ok := propMap[inp.KeyLower()]
		if !ok {
			return fmt.Errorf("property %s not found", inp.Key)
		}

		ft, ok := fieldMap[inp.KeyLower()]
		if !ok {
			panic(fmt.Errorf("property %s not found in the object", inp.Key))
		}

		f := val.FieldByIndex(ft.Index)
		if !f.CanSet() {
			panic(fmt.Errorf("readonly property %s", ft.Name))
		}

		err := setFieldValue(f, ft, inp.Value)
		if err != nil {
			return err
		}
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

// GenerateUsageString parses given object and returns the usage string for properties.
// For ex., given struct
// type account struct
//   AccountID string
//   CompanyName string
// }
// and call like this
//   GenerateUsageString(&account{}, []string{"AccountID"})
// Returns, AccountID=STRING [CompanyName=STRING]
//
// filter can be
//   +PROP_NAME - include, required
//   -PROP_NAME - exclude
//   -          - exclude all but the included ones
func GenerateUsageString(obj interface{}, properties Properties) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	fieldMap := toFieldsMap(getStructFields(val.Type()))

	var result []string
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

		if p.Required {
			result = append(result, fmt.Sprintf("%s=%s", f.Name, handler.UsageString))
		} else {
			result = append(result, fmt.Sprintf("[%s=%s]", f.Name, handler.UsageString))
		}
	}
	return strings.Join(result, " ")
}

// GetPropetyList returns all properties in a given object
func GetPropetyList(obj interface{}) Properties {
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

// getStructFields returns list of struct fields that can be set through reflection.
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

func isTypeSupported(tp reflect.Type) bool {
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return supportedTypes[tp] != nil
}

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
