package subscriptions

import (
	"context"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbclient/subscription"
	"github.com/killbill/kbcli/v3/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v3/kbmodel"
	"github.com/urfave/cli"
)

var subscriptionEventFormatter = cmdlib.Formatter{
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

var subscriptionFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "EXTERNAL_KEY",
			Path: "$.externalKey",
		},
		{
			Name: "SUBSC_ID",
			Path: "$.subscriptionId",
		},
		{
			Name: "STATE",
			Path: "$.state",
		},
		{
			Name: "PRODUCT",
			Path: "$.productName",
		},
		{
			Name: "EVENT_COUNT",
			Getter: func(v interface{}) interface{} {
				s := v.(*kbmodel.Subscription)
				return len(s.Events)
			},
		},
	},
	SubItems: []cmdlib.SubItem{
		{
			Name:      "Events",
			FieldName: "Events",
		},
	},
}

func getSubscription(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	subscID := o.Args[0]

	resp, err := o.Client().Subscription.GetSubscription(ctx, &subscription.GetSubscriptionParams{
		SubscriptionID: strfmt.UUID(subscID),
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return nil
}

// RegisterCommands registers all subscription commands
func RegisterCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Subscription{}), subscriptionFormatter)
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.EventSubscription{}), subscriptionEventFormatter)

	// Register top level command
	r.Register("", cli.Command{
		Name:    "subscriptions",
		Aliases: []string{"sub"},
		Usage:   "subscriptions related commands",
	}, nil)

	// Get subscription
	r.Register("subscriptions", cli.Command{
		Name:      "get",
		Usage:     "get subscription",
		ArgsUsage: "SUBSC_ID",
	}, getSubscription)

	registerCreateCommand(r)
	registerCancelCommand(r)
}
