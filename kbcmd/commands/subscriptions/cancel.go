package subscriptions

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient/subscription"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/urfave/cli"
)

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

func registerCancelCommand(r *cmdlib.App) {
	// Cancel subscription
	r.Register("subscriptions", cli.Command{
		Name:      "cancel",
		Usage:     "cancel a subscription",
		ArgsUsage: "SUBSC_ID",
	}, cancelSubscription)
}
