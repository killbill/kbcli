package subscriptions

import (
	"context"
	"fmt"
	"time"

	"github.com/killbill/kbcli/kbclient/bundle"
	"github.com/killbill/kbcli/kbclient/subscription"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/cmdlib/args"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbmodel"
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

	_, err = o.Client().Subscription.CreateSubscription(ctx, &createParams)

	if err != nil {
		return err
	}

	resp, err := o.Client().Bundle.GetBundleByKey(ctx, &bundle.GetBundleByKeyParams{
		ExternalKey: createParams.Body.ExternalKey,
	})
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
	subscriptionProperties.Get("ProductName").Required = true
	subscriptionProperties.Get("ProductCategory").Default = "BASE"
	subscriptionProperties.Get("PlanName").Required = true
	subscriptionProperties.Get("PriceList").Required = true
	subscriptionProperties.Sort(true, true)
	usageString := args.GenerateUsageString(&createSubscriptionArgs{}, subscriptionProperties)

	r.Register("subscriptions", cli.Command{
		Name:  "create",
		Usage: "create new subscription",
		ArgsUsage: fmt.Sprintf(`%s
      
       For e.g.,
		kbcmd subscriptions create ExternalKey=bundle1 AccountID=johndoe1 ProductName=simple PlanName=simple-monthly PriceList=default`,
			usageString),
	}, createSubscription)

}
