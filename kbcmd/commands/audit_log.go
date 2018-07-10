package commands

import (
	"reflect"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbmodel"
)

var auditLogFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "CHANGE_DATE",
			Path: "$.changeDate",
		},
		{
			Name: "CHANGE_TYPE",
			Path: "$.changeType",
		},
		{
			Name: "REASON",
			Path: "$.reasonCode",
		},
		{
			Name: "COMMENTS",
			Path: "$.comments",
		},
	},
}

func registerAuditLogCommands(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.AuditLog{}), auditLogFormatter)
}
