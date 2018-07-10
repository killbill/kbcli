package commands

import (
	"context"
	"reflect"

	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbclient/bundle"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var bundleTimelineFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "BILLING_PERIOD",
			Path: "$.billingPeriod",
		},
		{
			Name: "EFFECTIVE_DATE",
			Path: "$.effectiveDate",
		},
		{
			Name: "EVENT_TYPE",
			Path: "$.eventType",
		},
		{
			Name: "PRODUCT",
			Path: "$.product",
		},
	},
}

var bundleFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "EXTERNAL_KEY",
			Path: "$.externalKey",
		},
		{
			Name: "BUNDLE_ID",
			Path: "$.bundleId",
		},
		{
			Name: "ACTIVE_SUBSCRIPTIONS",
			Getter: func(v interface{}) interface{} {
				b := v.(*kbmodel.Bundle)
				var count int
				for _, s := range b.Subscriptions {
					if s.State == "ACTIVE" {
						count++
					}
				}
				return count
			},
		},
	},
	SubItems: []cmdlib.SubItem{
		{
			Name:      "Subscriptions",
			FieldName: "Subscriptions",
		},
		{
			Name:      "Timeline",
			FieldName: "Timeline.Events",
		},
	},
}

func listBundles(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accID := o.Args[0]
	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accID)
	if err != nil {
		return err
	}

	resp, err := o.Client().Account.GetAccountBundles(ctx, &account.GetAccountBundlesParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return nil
}

func getBundle(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	bundleKey := o.Args[0]

	resp, err := o.Client().Bundle.GetBundleByKey(ctx, &bundle.GetBundleByKeyParams{
		ExternalKey: bundleKey,
	})
	if err != nil {
		return err
	}
	if len(resp.Payload) == 1 {
		o.Print(resp.Payload[0])
	} else {
		o.Print(resp.Payload)
		o.Log.Warningf("why is getbundle returning multiple items?")
	}
	return nil
}

func registerBundleCommands(r *cmdlib.App) {

	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Bundle{}), bundleFormatter)
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.BundleTimeline{}), bundleTimelineFormatter)

	// Register top level command
	r.Register("", cli.Command{
		Name:    "bundles",
		Aliases: []string{"bun"},
		Usage:   "bundles related commands",
	}, nil)

	// List bundles
	r.Register("bundles", cli.Command{
		Name:      "list",
		Usage:     "list bundles for the account",
		ArgsUsage: "ACCOUNT",
	}, listBundles)

	// Get bundle by key
	r.Register("bundles", cli.Command{
		Name:      "get",
		Usage:     "get bundle by key",
		ArgsUsage: "UNIQUE_KEY",
	}, getBundle)
}
