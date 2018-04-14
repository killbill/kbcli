package commands

import (
	"context"
	"reflect"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbclient/bundle"
	"github.com/killbill/kbcli/kbclient/subscription"
	"github.com/killbill/kbcli/kbmodel"
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

func createSubscription(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 5 {
		return cmdlib.ErrorInvalidArgs
	}

	accIDOrKey := o.Args[0]
	key := o.Args[1]
	product := o.Args[2]
	plan := o.Args[3]
	priceList := o.Args[4]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	_, err = o.Client().Subscription.CreateSubscription(ctx, &subscription.CreateSubscriptionParams{
		Body: &kbmodel.Subscription{
			ExternalKey:     key,
			AccountID:       acc.AccountID,
			StartDate:       strfmt.Date(time.Now()),
			ProductName:     &product,
			ProductCategory: "BASE",
			PlanName:        &plan,
			PriceList:       &priceList,
		},
	})

	if err != nil {
		return err
	}

	resp, err := o.Client().Bundle.GetBundleByKey(ctx, &bundle.GetBundleByKeyParams{
		ExternalKey: key,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return nil
}

func cancelSubscription(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	subscID := o.Args[0]

	_, err := o.Client().Subscription.CancelSubscriptionPlan(ctx, &subscription.CancelSubscriptionPlanParams{
		SubscriptionID: strfmt.UUID(subscID),
	})
	if err != nil {
		return err
	}

	resp, err := o.Client().Subscription.GetSubscription(ctx, &subscription.GetSubscriptionParams{
		SubscriptionID: strfmt.UUID(subscID),
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return nil
}

func registerSubscriptionCommands(r *cmdlib.App) {
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

	// Create subscription
	r.Register("subscriptions", cli.Command{
		Name:      "create",
		Usage:     "create new subscription",
		ArgsUsage: "ACC_ID UNIQUE_NAME PRODUCT PLAN PRICE_LIST",
	}, createSubscription)

	// Cancel subscription
	r.Register("subscriptions", cli.Command{
		Name:      "cancel",
		Usage:     "cancel a subscription",
		ArgsUsage: "SUBSC_ID",
	}, cancelSubscription)
}
