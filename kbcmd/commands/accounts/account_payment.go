package accounts

import (
	"context"
	"reflect"
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/kbcmd/cmdlib/args"

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

var addAccountPaymentMethodProperties args.Properties

func listAccountPayments(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	accIDOrKey := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	withAttempts := true
	resp, err := o.Client().Account.GetPaymentsForAccount(ctx, &account.GetPaymentsForAccountParams{
		AccountID:    acc.AccountID,
		WithAttempts: &withAttempts,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return err
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

func getAccountPaymentMethod(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	paymentMethodID := o.Args[0]

	resp, err := o.Client().PaymentMethod.GetPaymentMethod(ctx, &payment_method.GetPaymentMethodParams{
		PaymentMethodID: strfmt.UUID(paymentMethodID),
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)
	return err
}

func addAccountPaymentMethod(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 4 {
		return cmdlib.ErrorInvalidArgs
	}

	accKey := o.Args[0]
	method := o.Args[1]
	externalKey := o.Args[2]
	isDefault, err := strconv.ParseBool(o.Args[3])
	if err != nil {
		return err
	}

	var payAllUnpaidInvoices bool
	if len(o.Args) > 4 {
		payAllUnpaidInvoices, err = strconv.ParseBool(o.Args[4])
		if err != nil {
			return err
		}
	} else {
		payAllUnpaidInvoices = false
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accKey)
	if err != nil {
		return err
	}

	var pluginProperties []args.Input
	if len(o.Args) > 5 {
		pluginProperties, err = args.ParseArgs(o.Args[5:])
		if err != nil {
			return err
		} else {
			pluginProperties = []args.Input{}
		}
	}

	pm := &kbmodel.PaymentMethod{
		ExternalKey: externalKey,
		PluginName:  method,
		PluginInfo:  &kbmodel.PaymentMethodPluginDetail{},
	}
	for _, pp := range pluginProperties {
		pm.PluginInfo.Properties = append(pm.PluginInfo.Properties, &kbmodel.PluginProperty{
			Key:   pp.Key,
			Value: pp.Value,
		})
	}

	resp, err := o.Client().Account.CreatePaymentMethod(ctx, &account.CreatePaymentMethodParams{
		AccountID:             acc.AccountID,
		Body:                  pm,
		IsDefault:             &isDefault,
		PayAllUnpaidInvoices:  &payAllUnpaidInvoices,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)

	return err
}

func removeAccountPaymentMethod(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}

	paymentMethodID := o.Args[0]
	var force bool
	if len(o.Args) > 1 {
		forceBool, err := strconv.ParseBool(o.Args[1])
		if err != nil {
			return err
		} else {
			force = forceBool
		}
	} else {
		force = false
	}
	_, err := o.Client().PaymentMethod.DeletePaymentMethod(ctx, &payment_method.DeletePaymentMethodParams{
		PaymentMethodID:        strfmt.UUID(paymentMethodID),
		ForceDefaultPmDeletion: &force,
	})

	return err
}

func registerAccountPaymentCommands(r *cmdlib.App) {
	// Payment method
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.PaymentMethod{}), paymentMethodFormatter)

	addAccountPaymentMethodProperties = args.GetProperties(&kbmodel.PaymentMethod{})

	// Payments
	r.Register("accounts", cli.Command{
		Name:    "payments",
		Aliases: []string{},
		Usage:   "Payments related commands",
	}, nil)

	// List payments
	r.Register("accounts.payments", cli.Command{
		Name:      "list",
		Aliases:   []string{"ls"},
		Usage:     "List payments for the given account",
		ArgsUsage: `ACCOUNT`,
	}, listAccountPayments)

	// payment methods
	r.Register("accounts", cli.Command{
		Name:    "payment-methods",
		Aliases: []string{"pm"},
		Usage:   "Payment method related commands",
	}, nil)

	// List payment methods
	r.Register("accounts.payment-methods", cli.Command{
		Name:      "list",
		Aliases:   []string{"ls"},
		Usage:     "List payment methods for the given account",
		ArgsUsage: `ACCOUNT`,
	}, listAccountPaymentMethods)

	// Get payment method by ID
	r.Register("accounts.payment-methods", cli.Command{
		Name:      "get",
		Usage:     "Get payment method for the given account",
		ArgsUsage: `PAYMENT_METHOD_ID`,
	}, getAccountPaymentMethod)

	// Add payment method
	r.Register("accounts.payment-methods", cli.Command{
		Name:  "add",
		Usage: "Add new payment method",
		ArgsUsage: `ACCOUNT PLUGIN_NAME EXTERNAL_KEY IS_DEFAULT PAY_ALL_UNPAID_INVOICES [Property1=Value1] ...

   For ex.,
      kbcmd accounts payment-methods add johndoe killbill-stripe visa true false token=tok_1CidZ7HGlIo9NLGOy7sPvbsz
		`,
	}, addAccountPaymentMethod)

	// Remove payment method
	r.Register("accounts.payment-methods", cli.Command{
		Name:      "remove",
		Aliases:   []string{"rm"},
		Usage:     "Remove payment method",
		ArgsUsage: `PM_METHOD_ID [FORCE]`,
	}, removeAccountPaymentMethod)
}
