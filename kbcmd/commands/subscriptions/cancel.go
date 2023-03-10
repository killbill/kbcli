package subscriptions

import (
	"context"
	"fmt"

	"github.com/killbill/kbcli/v3/kbcmd/cmdlib/args"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbclient/subscription"
	"github.com/killbill/kbcli/v3/kbcmd/cmdlib"
	"github.com/urfave/cli"
)

var (
	subscriptionCancelProperties args.Properties
)

func cancelSubscription(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}

	subscID := o.Args[0]
	csa := &subscription.CancelSubscriptionPlanParams{}
	err := args.LoadProperties(csa, subscriptionCancelProperties, o.Args[1:])
	if err != nil {
		return err
	}
	csa.SubscriptionID = strfmt.UUID(subscID)

	_, err = o.Client().Subscription.CancelSubscriptionPlan(ctx, csa)
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

func registerCancelCommand(r *cmdlib.App) {
	subscriptionCancelProperties = args.GetProperties(&subscription.CancelSubscriptionPlanParams{})
	subscriptionCancelProperties.Remove("SubscriptionID")
	usageString := args.GenerateUsageString(&subscription.CancelSubscriptionPlanParams{}, subscriptionCancelProperties)

	// Cancel subscription
	r.Register("subscriptions", cli.Command{
		Name:      "cancel",
		Usage:     "cancel a subscription",
		ArgsUsage: fmt.Sprintf(`SUBSC_ID %s`, usageString),
	}, cancelSubscription)
}
