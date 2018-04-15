package commands

import (
	"context"
	"reflect"

	"github.com/killbill/kbcli/kbclient/tenant"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

func createTenant(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 3 {
		return cmdlib.ErrorInvalidArgs
	}
	name := o.Args[0]
	apiKey := o.Args[1]
	apiSecret := o.Args[2]

	t := &kbmodel.Tenant{
		ExternalKey: name,
		APIKey:      &apiKey,
		APISecret:   &apiSecret,
	}
	_, err := o.Client().Tenant.CreateTenant(ctx, &tenant.CreateTenantParams{
		Body: t,
	})
	if err != nil {
		return err
	}
	o.Print(t)
	return nil
}

// func configurePlugin(ctx context.Context, o *cmdlib.Options) error {
// 	if len(o.Args) < 1 {
// 		return cmdlib.ErrorInvalidArgs
// 	}
// 	pluginName := o.Args[0]
// 	kv, err := cmdlib.ParseKeyValuePairs(o.Args[1:])
// 	if err != nil {
// 		return err
// 	}

// 	t := &kbmodel.Tenant{
// 		ExternalKey: name,
// 		APIKey:      &apiKey,
// 		APISecret:   &apiSecret,
// 	}
// 	_, err := o.Client().Tenant.CreateTenant(ctx, &tenant.CreateTenantParams{
// 		Body: t,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	o.Print(t)
// 	return nil
// }

func registerTenantCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Tenant{}), cmdlib.Formatter{
		Columns: []cmdlib.Column{
			{
				Name: "EXTERNAL_KEY",
				Path: "$.externalKey",
			},
			{
				Name: "API_KEY",
				Path: "$.apiKey",
			},
			{
				Name: "API_SECRET",
				Path: "$.apiSecret",
			},
		},
	})

	// Register top level command
	r.Register("", cli.Command{
		Name:    "tenants",
		Aliases: []string{"ten"},
		Usage:   "tenant related commands",
	}, nil)

	r.Register("tenants", cli.Command{
		Name:        "create",
		Usage:       "create new tenant",
		ArgsUsage:   "UNIQUE_NAME API_KEY API_SECRET",
		Description: "Creates new tenant",
	}, createTenant)
}
