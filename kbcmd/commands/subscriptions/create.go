package subscriptions

import (
	"context"
	"fmt"
	"time"

	"github.com/killbill/kbcli/v2/kbclient/subscription"
	"github.com/killbill/kbcli/v2/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v2/kbcmd/cmdlib/args"
	"github.com/killbill/kbcli/v2/kbcmd/kblib"
	"github.com/killbill/kbcli/v2/kbmodel"
	"github.com/urfave/cli"
)

var subscriptionProperties args.Properties

type createSubscriptionArgs struct {
	Account string
	subscription.CreateSubscriptionParams
	kbmodel.Subscription
}

func createSubscription(ctx context.Context, o *cmdlib.Options) error {

	csa := &createSubscriptionArgs{}
	err := args.LoadProperties(csa, subscriptionProperties, o.Args)
	if err != nil {
		return err
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), csa.Account)
	if err != nil {
		return err
	}

	var createParams subscription.CreateSubscriptionParams
	createParams = csa.CreateSubscriptionParams
	createParams.Body = &csa.Subscription
	createParams.Body.AccountID = acc.AccountID
	createParams.ProcessLocationHeader = true
	resp, err := o.Client().Subscription.CreateSubscription(ctx, &createParams)

	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return nil
}

func registerCreateCommand(r *cmdlib.App) {
	subscriptionProperties = args.GetProperties(&createSubscriptionArgs{})
	subscriptionProperties.Get("ExternalKey").Required = true
	subscriptionProperties.Get("Account").Required = true
	subscriptionProperties.Remove("AccountID")
	subscriptionProperties.Get("StartDate").Default = time.Now().Format("2006-01-02")
	subscriptionProperties.Get("PlanName").Required = true
	subscriptionProperties.Get("BillingPeriod").Enums = kbmodel.SubscriptionBillingPeriodEnumValues
	subscriptionProperties.Sort(true, true)
	usageString := args.GenerateUsageString(&createSubscriptionArgs{}, subscriptionProperties)

	r.Register("subscriptions", cli.Command{
		Name:  "create",
		Usage: "create new subscription",
		ArgsUsage: fmt.Sprintf(`%s

       For e.g.,
         kbcmd subscriptions create Account=johndoe1 ExternalKey=bundle1 PlanName=simple-monthly`,
			usageString),
	}, createSubscription)
}
