package commands

import (
	"context"
	"fmt"
	"reflect"

	"github.com/killbill/kbcli/kbclient/payment_method"

	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var paymentMethodFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NAME",
			Path: "$.pluginName",
		},
		{
			Name: "ID",
			Path: "$.paymentMethodId",
		},
		{
			Name: "IS_DEFAULT",
			Path: "$.isDefault",
		},
	},
}

func listAccountPaymentMethods(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	accIDOrKey := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	resp, err := o.Client().Account.GetPaymentMethodsForAccount(ctx, &account.GetPaymentMethodsForAccountParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return err
}

func addAccountPaymentMethod(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}

	accKey := o.Args[0]
	method := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accKey)
	if err != nil {
		return err
	}
	o.Log.Infof("account %s retrieved. Now adding payment method", acc.AccountID)

	_, err = o.Client().Account.CreatePaymentMethod(ctx, &account.CreatePaymentMethodParams{
		AccountID: acc.AccountID,
		Body: &kbmodel.PaymentMethod{
			PluginName: method,
			PluginInfo: &kbmodel.PaymentMethodPluginDetail{},
		},
	})

	return err
}

func removeAccountPaymentMethod(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}

	accKey := o.Args[0]
	method := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accKey)
	if err != nil {
		return err
	}
	resp, err := o.Client().Account.GetPaymentMethodsForAccount(ctx, &account.GetPaymentMethodsForAccountParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}

	var pm *kbmodel.PaymentMethod
	for _, m := range resp.Payload {
		if m.PluginName == method {
			pm = m
			break
		}
	}
	if pm == nil {
		return fmt.Errorf("payment method %s not found for account %s", method, accKey)
	}

	_, err = o.Client().PaymentMethod.DeletePaymentMethod(ctx, &payment_method.DeletePaymentMethodParams{
		PaymentMethodID: pm.PaymentMethodID,
	})

	return err
}

func registerAccountPaymentCommands(r *cmdlib.App) {
	// Payment method
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.PaymentMethod{}), paymentMethodFormatter)

	// List payment methods
	r.Register("accounts", cli.Command{
		Name:      "list-payment-methods",
		Usage:     "List payment methods for the given account",
		ArgsUsage: `ACCOUNT`,
	}, listAccountPaymentMethods)

	// Add payment method
	r.Register("accounts", cli.Command{
		Name:      "add-payment-method",
		Usage:     "Add new payment method",
		ArgsUsage: `ACCOUNT METHOD`,
	}, addAccountPaymentMethod)

	// Remove payment method
	r.Register("accounts", cli.Command{
		Name:      "remove-payment-method",
		Usage:     "Remove payment method",
		ArgsUsage: `ACCOUNT METHOD`,
	}, removeAccountPaymentMethod)
}
