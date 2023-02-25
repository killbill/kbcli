package commands

import (
	"reflect"

	"github.com/killbill/kbcli/v3/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v3/kbmodel"
)

var customFieldFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NAME",
			Path: "$.name",
		},
		{
			Name: "VALUE",
			Path: "$.value",
		},
		{
			Name: "ID",
			Path: "$.customFieldId",
		},
		{
			Name: "OBJECT_ID",
			Path: "$.objectId",
		},
		{
			Name: "OBJ_TYPE",
			Path: "$.objectType",
		},
	},
}

func registerCustomFieldFunctions(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.CustomField{}), customFieldFormatter)
}
