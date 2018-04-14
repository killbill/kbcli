package commands

import (
	"context"
	"fmt"
	"reflect"

	"github.com/killbill/kbcli/kbclient/payment_method"

	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var accountFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NAME",
			Path: "$.name",
		},
		{
			Name: "EXTERNAL_KEY",
			Path: "$.externalKey",
		},
		{
			Name: "ACCOUNT_ID",
			Path: "$.accountId",
		},
		{
			Name: "EMAIL",
			Path: "$.email",
		},
		{
			Name: "BALANCE",
			Path: "$.accountBalance",
		},
		{
			Name: "CURRENCY",
			Path: "$.currency",
		},
	},
}

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

var tagFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NAME",
			Path: "$.tagDefinitionName",
		},
		{
			Name: "TAG_ID",
			Path: "$.tagId",
		},
		{
			Name: "TAG_DEFINITION_ID",
			Path: "$.tagDefinitionId",
		},
	},
}

func listAccounts(ctx context.Context, o *cmdlib.Options) error {
	var err error
	resp, err := o.Client().Account.GetAccounts(ctx, &account.GetAccountsParams{})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

// getAccount - get account information command
func getAccount(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), o.Args[0])
	if err == nil {
		o.Print(acc)
	}

	return err
}

func createAccount(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 3 {
		return cmdlib.ErrorInvalidArgs
	}
	key := o.Args[0]
	name := o.Args[1]
	email := o.Args[2]

	_, err := o.Client().Account.CreateAccount(ctx, &account.CreateAccountParams{
		Body: &kbmodel.Account{
			ExternalKey: key,
			Name:        name,
			Email:       email,
			Currency:    "USD",
		},
	})
	if err != nil {
		return err
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), key)
	if err != nil {
		return err
	}
	o.Print(acc)
	return nil
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

func listAccountTags(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	resp, err := o.Client().Account.GetAccountTags(ctx, &account.GetAccountTagsParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

func addAccountTag(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]
	tagName := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	tags, err := kblib.GetTagDefinitions(ctx, o.Client())
	if err != nil {
		return err
	}
	tag := tags[tagName]
	if tag == nil {
		return fmt.Errorf("tag %s not found", tagName)
	}

	_, err = o.Client().Account.CreateAccountTags(ctx, &account.CreateAccountTagsParams{
		AccountID: acc.AccountID,
		Body:      []strfmt.UUID{tag.ID},
	})
	if err != nil {
		return err
	}

	return nil
}

func removeAccountTag(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]
	tagName := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	tags, err := kblib.GetTagDefinitions(ctx, o.Client())
	if err != nil {
		return err
	}
	tag := tags[tagName]
	if tag == nil {
		return fmt.Errorf("tag %s not found", tagName)
	}

	_, err = o.Client().Account.DeleteAccountTags(ctx, &account.DeleteAccountTagsParams{
		AccountID: acc.AccountID,
		TagDef:    []strfmt.UUID{tag.ID},
	})
	return err
}

func registerAccountCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Account{}), accountFormatter)

	// Payment method
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.PaymentMethod{}), paymentMethodFormatter)

	// Tag
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Tag{}), tagFormatter)

	// Register top level command
	r.Register("", cli.Command{
		Name:    "accounts",
		Aliases: []string{"acc"},
		Usage:   "account related commands",
	}, nil)

	// Get account
	r.Register("accounts", cli.Command{
		Name:        "get",
		Usage:       "get account information",
		ArgsUsage:   "{ACC_ID | [+]UNIQUE_KEY}",
		Description: "If the external key is also UUID, then an optional + may be prefixed with external key for disambiguation",
	}, getAccount)

	// List all accounts
	r.Register("accounts", cli.Command{
		Name:      "list",
		Usage:     "list all accounts",
		ArgsUsage: "",
	}, listAccounts)

	// Create account
	r.Register("accounts", cli.Command{
		Name:  "create",
		Usage: "create new account",
		ArgsUsage: `UNIQUE_KEY USER_NAME EMAIL

For ex.,:
		kbcmd accounts create prem1 "Prem Ramanathan" prem@prem.com
		`,
		Description: "Creates new account",
	}, createAccount)

	// List account tags
	r.Register("accounts", cli.Command{
		Name:      "list-tags",
		Usage:     "List all account tags",
		ArgsUsage: "ACCOUNT",
	}, listAccountTags)

	// Add account tag
	r.Register("accounts", cli.Command{
		Name:      "add-tag",
		Usage:     "Add tag to the account",
		ArgsUsage: "ACCOUNT TAG_NAME",
	}, addAccountTag)

	// Remove account tag
	r.Register("accounts", cli.Command{
		Name:      "remove-tag",
		Usage:     "Remove tag from the account",
		ArgsUsage: "ACCOUNT TAG_NAME",
	}, removeAccountTag)

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
