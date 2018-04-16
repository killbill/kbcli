package cmdlib

import (
	"reflect"
)

func registerDefaultFormatters() {
	var stringFormatter = Formatter{
		Columns: []Column{
			{
				Name: "VALUE",
				Path: "$.",
			},
		},
	}
	AddFormatter(reflect.TypeOf(""), stringFormatter)
}
