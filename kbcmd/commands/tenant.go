package commands

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/killbill/kbcli/kbclient/tenant"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var tenantKeyValueFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "KEY",
		},
	},
	CustomFn: customTenantKeyValueFormatter,
}

func customTenantKeyValueFormatter(v interface{}, fo cmdlib.FormatOptions) cmdlib.Output {
	kv := v.(*kbmodel.TenantKeyValue)
	out := cmdlib.Output{
		Title:   "Keys",
		Columns: []string{"KEY"},
	}
	for _, v := range kv.Values {
		out.Rows = append(out.Rows, cmdlib.OutputRow{
			Values: []string{v},
		})
	}
	return out
}

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

func configureStripePlugin(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 3 {
		return cmdlib.ErrorInvalidArgs
	}
	publicKey := o.Args[0]
	privateKey := o.Args[1]
	feesOrPercent := o.Args[2]

	var isPercent bool
	if strings.HasSuffix(feesOrPercent, "%") {
		isPercent = true
		feesOrPercent = feesOrPercent[0 : len(feesOrPercent)-1]
	}
	fees, err := strconv.ParseFloat(feesOrPercent, 64)
	if err != nil {
		return fmt.Errorf("unable to parse fees. %v", err)
	}
	if isPercent {
		fees /= 100
	}

	stripeConfig := fmt.Sprintf(`:stripe:
  :api_secret_key: "%s"
  :api_publishable_key: "%s"
`,
		privateKey, publicKey)

	if isPercent {
		stripeConfig += fmt.Sprintf("  :fees_percent: %.3f", fees)
	} else {
		stripeConfig += fmt.Sprintf("  :fees_amount: %f", fees)
	}

	_, err = o.Client().Tenant.UploadPluginConfiguration(ctx, &tenant.UploadPluginConfigurationParams{
		PluginName: "killbill-stripe",
		Body:       &stripeConfig,
	})
	if err != nil {
		return err
	}

	config, err := o.Client().Tenant.GetPluginConfiguration(ctx, &tenant.GetPluginConfigurationParams{
		PluginName: "killbill-stripe",
	})
	o.Print(config.Payload)
	return nil
}

func registerTenantCommands(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.TenantKeyValue{}), tenantKeyValueFormatter)
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
		Name:      "create",
		Usage:     "create new tenant",
		ArgsUsage: "UNIQUE_NAME API_KEY API_SECRET",
	}, createTenant)

	// add stripe plugin
	r.Register("tenants", cli.Command{
		Name:  "configure-stripe-plugin",
		Usage: "Adds stripe plugin configuration",
		ArgsUsage: `PUBLIC_KEY PRIVATE_KEY FEES 
		
		For e.g., 
			tenants configure-stripe-plugin pk_test_WgswzvAG2ssFcB7Xk7tBi0XL sk_test_aM1UY1u411tPNI4LEL2tTFay 3%`,
	}, configureStripePlugin)
}
