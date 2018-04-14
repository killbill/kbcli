package commands

import (
	"context"
	"fmt"
	"reflect"

	"github.com/killbill/kbcli/kbcmd/cmdlib/args"

	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
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

var (
	createAccountPropertyList args.Properties
	updateAccountPropertyList args.Properties
)

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
	if len(o.Args) < 4 {
		return cmdlib.ErrorInvalidArgs
	}

	accToCreate := &kbmodel.Account{}
	err := args.LoadProperties(accToCreate, createAccountPropertyList, o.Args)
	if err != nil {
		return err
	}

	_, err = o.Client().Account.CreateAccount(ctx, &account.CreateAccountParams{
		Body: accToCreate,
	})
	if err != nil {
		return err
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accToCreate.ExternalKey)
	if err != nil {
		return err
	}
	o.Print(acc)
	return nil
}

func updateAccount(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	key := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), key)
	if err != nil {
		return err
	}
	err = args.LoadProperties(acc, updateAccountPropertyList, o.Args[1:])
	if err != nil {
		return err
	}

	_, err = o.Client().Account.UpdateAccount(ctx, &account.UpdateAccountParams{
		AccountID: acc.AccountID,
		Body:      acc,
	})
	if err != nil {
		return err
	}
	acc, err = kblib.GetAccountByKeyOrID(ctx, o.Client(), key)
	if err != nil {
		return err
	}
	o.Print(acc)
	return nil
}

func registerAccountCommands(r *cmdlib.App) {
	registerAccountPaymentCommands(r)
	registerAccountTagCommands(r)

	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Account{}), accountFormatter)

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
	createAccountPropertyList = args.GetPropetyList(&kbmodel.Account{})
	createAccountPropertyList.Get("ExternalKey").Required = true
	createAccountPropertyList.Get("Email").Required = true
	createAccountPropertyList.Get("Name").Required = true
	createAccountPropertyList.Sort(true, true)

	createAccountsUsage := fmt.Sprintf(`%s

		For ex.,:
				kbcmd accounts create ExternalKey=prem1 Name="Prem Ramanathan" Email=prem@prem.com Currency=USD
				`,
		args.GenerateUsageString(&kbmodel.Account{}, createAccountPropertyList))

	r.Register("accounts", cli.Command{
		Name:        "create",
		Usage:       "create new account",
		ArgsUsage:   createAccountsUsage,
		Description: "Creates new account",
	}, createAccount)

	// Update account
	updateAccountPropertyList = args.GetPropetyList(&kbmodel.Account{})
	// Following properties can't change
	updateAccountPropertyList.Remove("ExternalKey")
	updateAccountPropertyList.Remove("AccountID")
	updateAccountPropertyList.Remove("Currency")
	updateAccountPropertyList.Remove("BillCycleDayLocal")
	updateAccountPropertyList.Remove("TimeZone")
	updateAccountPropertyList.Remove("AccountBalance")
	updateAccountPropertyList.Sort(true, true)

	updateAccountsUsage := fmt.Sprintf(`ACCOUNT %s

		For ex.,:
				kbcmd accounts create ExternalKey=prem1 Name="Prem Ramanathan" Email=prem@prem.com Currency=USD
				`,
		args.GenerateUsageString(&kbmodel.Account{}, updateAccountPropertyList))

	r.Register("accounts", cli.Command{
		Name:        "update",
		Usage:       "updates existing account",
		ArgsUsage:   updateAccountsUsage,
		Description: "Updates existing account",
	}, updateAccount)

}
