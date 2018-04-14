package cmdlib

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"
)

var defaultProperties = map[string]bool{
	"xkillbillapikey":    true,
	"xkillbillapisecret": true,
	"xkillbillcomment":   true,
	"xkillbillcreatedby": true,
	"xkillbillreason":    true,
}

type typeSetterFn func(f reflect.Value, val string) error

type typeHandler struct {
	Type        reflect.Type
	UsageString string
	SetterFn    typeSetterFn
}

var supportedTypes = map[reflect.Type]*typeHandler{
	reflect.TypeOf(""): &typeHandler{
		Type:        reflect.TypeOf(""),
		UsageString: "STRING",
		SetterFn: func(f reflect.Value, val string) error {
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&val))
			} else {
				f.SetString(val)
			}
			return nil
		},
	},
	reflect.TypeOf(false): &typeHandler{
		Type:        reflect.TypeOf(false),
		UsageString: "{True|False}",
		SetterFn: func(f reflect.Value, val string) error {
			boolVal, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&boolVal))
			} else {
				f.SetBool(boolVal)
			}
			return nil
		},
	},
	reflect.TypeOf(float64(0)): &typeHandler{
		Type:        reflect.TypeOf(float64(0)),
		UsageString: "NUMBER",
		SetterFn: func(f reflect.Value, val string) error {
			floatVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&floatVal))
			} else {
				f.SetFloat(floatVal)
			}
			return nil
		},
	},
	reflect.TypeOf(int32(0)): &typeHandler{
		Type:        reflect.TypeOf(int32(0)),
		UsageString: "INTEGER",
		SetterFn: func(f reflect.Value, val string) error {
			intVal, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return err
			}
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&intVal))
			} else {
				f.SetInt(intVal)
			}
			return nil
		},
	},
	reflect.TypeOf(strfmt.UUID("")): &typeHandler{
		Type:        reflect.TypeOf(strfmt.UUID("")),
		UsageString: "UUID",
		SetterFn: func(f reflect.Value, val string) error {
			if !strfmt.IsUUID(val) {
				return fmt.Errorf("Value %s is not UUID", val)
			}
			uuid := strfmt.UUID(val)
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&uuid))
			} else {
				f.Set(reflect.ValueOf(uuid))
			}
			return nil
		},
	},
	reflect.TypeOf(strfmt.DateTime{}): &typeHandler{
		Type:        reflect.TypeOf(strfmt.DateTime{}),
		UsageString: "DATETIME",
		SetterFn: func(f reflect.Value, val string) error {
			dateTime, err := strfmt.ParseDateTime(val)
			if err != nil {
				return fmt.Errorf("Value %s is not in date time format. %v", val, err)
			}
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&dateTime))
			} else {
				f.Set(reflect.ValueOf(dateTime))
			}
			return nil
		},
	},
}

// ParseProperties parses key value pairs from argument
func ParseProperties(args []string) (map[string]string, error) {
	result := map[string]string{}
	for _, a := range args {
		comps := strings.Split(a, "=")
		if len(comps) != 2 {
			return nil, fmt.Errorf("invalid input %s. expecting var=value", a)
		}
		result[comps[0]] = comps[1]
	}
	return result, nil
}

// ListProperties lists the settable properties of the type
func ListProperties(obj interface{}) []string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	tp := reflect.Indirect(reflect.ValueOf(obj)).Type()

	var result []string
	fields := getStructFields(tp)
	for _, f := range fields {
		fn := strings.ToLower(f.Name)
		if !defaultProperties[fn] {
			result = append(result, f.Name)
		}
	}

	return result
}

// LoadProperties loads key value pairs into target object.
// property names must match
func LoadProperties(properties map[string]string, obj interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))
	fieldMap := toFieldsMap(getStructFields(val.Type()))

	for k, v := range properties {
		propName := strings.ToLower(k)

		ft, ok := fieldMap[propName]
		if !ok {
			return fmt.Errorf("property %s not found", k)
		}
		f := val.FieldByIndex(ft.Index)
		if !f.CanSet() {
			return fmt.Errorf("readonly property %s", k)
		}
		err := setFieldValue(f, ft, v)
		if err != nil {
			return err
		}
	}
	return nil
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
func GenerateUsageString(obj interface{}, filter []string) []string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("invalid object. expecting pointer to struct. got - %#v", reflect.TypeOf(obj)))
	}

	val := reflect.Indirect(reflect.ValueOf(obj))

	includedProperties := map[string]bool{}
	excludedProperties := map[string]bool{}
	requiredProperties := map[string]bool{}
	for _, f := range filter {
		f := strings.ToLower(f)
		prefix := f[0:1]
		if prefix == "+" {
			requiredProperties[f[1:]] = true
		} else if prefix == "-" {
			excludedProperties[f[1:]] = true
		} else {
			includedProperties[f] = true
		}
	}
	fieldList := getStructFields(val.Type())

	var result []string
	for _, f := range fieldList {
		fType := f.Type
		if fType.Kind() == reflect.Ptr {
			fType = fType.Elem()
		}
		fn := strings.ToLower(f.Name)
		if defaultProperties[fn] {
			// Skip all default properties
			continue
		}
		handler := supportedTypes[fType]

		if requiredProperties[fn] {
			if handler == nil {
				panic(fmt.Errorf("required property %s type not supported", f.Name))
			}
			result = append(result, fmt.Sprintf("%s=%s", f.Name, handler.UsageString))
		} else if !excludedProperties[fn] {
			if handler == nil {
				panic(fmt.Errorf("property %s type not supported", f.Name))
			}
			result = append(result, fmt.Sprintf("[%s=%s]", f.Name, handler.UsageString))
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
	switch tp {
	case reflect.TypeOf(""),
		reflect.TypeOf(false),
		reflect.TypeOf(float64(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(strfmt.UUID("")),
		reflect.TypeOf(strfmt.DateTime{}):

		return true
	}

	return false
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
