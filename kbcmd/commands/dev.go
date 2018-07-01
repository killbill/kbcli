package commands

import (
	"context"
	"reflect"

	"github.com/killbill/kbcli/kbmodel"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient/debug"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/urfave/cli"
)

var clockFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "DATE_TIME",
			Path: "$.currentUtcTime",
		},
		{
			Name: "TIME_ZONE",
			Path: "$.timeZone",
		},
		{
			Name: "LOCAL DATE",
			Path: "$.localDate",
		},
	},
}

func getClock(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 0 {
		return cmdlib.ErrorInvalidArgs
	}

	cl := debug.New(o.Client().Transport, strfmt.Default, httptransport.BasicAuth(o.Username, o.Password), o.Client().Defaults())
	resp, err := cl.GetClock(ctx, &debug.GetClockParams{})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

func setClock(ctx context.Context, o *cmdlib.Options) error {
	var err error

	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	dateStr := o.Args[0]
	var date strfmt.DateTime
	if strfmt.IsDate(dateStr) {
		dateStr += "T00:00:00Z"
	}
	date, err = strfmt.ParseDateTime(dateStr)
	if err != nil {
		return err
	}

	cl := debug.New(o.Client().Transport, strfmt.Default, httptransport.BasicAuth(o.Username, o.Password), o.Client().Defaults())

	resp, err := cl.SetClock(ctx, &debug.SetClockParams{
		RequestedDate: date,
	})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

func registerDevCommands(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Clock{}), clockFormatter)
	// Register top level command
	r.Register("", cli.Command{
		Name:  "dev",
		Usage: "development/debugging related commands",
	}, nil)

	r.Register("dev", cli.Command{
		Name:  "get-clock",
		Usage: "print current test clock time",
	}, getClock)

	// add stripe plugin
	r.Register("dev", cli.Command{
		Name:  "set-clock",
		Usage: "set test clock",
		ArgsUsage: `DATE 
		
		For e.g., 
			dev set-clock 2018-03-04`,
	}, setClock)
}
