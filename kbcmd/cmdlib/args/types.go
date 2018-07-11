package args

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
)

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
	reflect.TypeOf(strfmt.Date{}): &typeHandler{
		Type:        reflect.TypeOf(strfmt.Date{}),
		UsageString: "DATE",
		SetterFn: func(f reflect.Value, val string) error {
			date, err := time.Parse(strfmt.RFC3339FullDate, val)
			if err != nil {
				return fmt.Errorf("Value %s is not in date format. %v", val, err)
			}
			strDate := strfmt.Date(date)
			if f.Type().Kind() == reflect.Ptr {
				f.Set(reflect.ValueOf(&strDate))
			} else {
				f.Set(reflect.ValueOf(strDate))
			}
			return nil
		},
	},
}

type iEnum interface {
	IsValid() bool
}

// Handles alias types (enums)
var stringAliasTypeHandler = &typeHandler{
	Type:        reflect.TypeOf(""),
	UsageString: "STRING",
	SetterFn: func(f reflect.Value, val string) error {
		var newVal reflect.Value
		if f.Type().Kind() == reflect.Ptr {
			newVal = reflect.New(f.Type().Elem())
		} else {
			newVal = reflect.New(f.Type())
		}
		newVal.Elem().SetString(val)

		if iEn, ok := newVal.Elem().Interface().(iEnum); ok {
			if !iEn.IsValid() {
				return fmt.Errorf("Invalid value %s for enum", val)
			}
		}

		if f.Type().Kind() == reflect.Ptr {
			f.Set(newVal)
		} else {
			f.SetString(val)
		}
		return nil
	},
}

func getTypeHandler(tp reflect.Type) *typeHandler {
	handler := supportedTypes[tp]
	if handler != nil {
		return handler
	}
	if tp.Kind() == reflect.String {
		return stringAliasTypeHandler
	}
	return nil
}
