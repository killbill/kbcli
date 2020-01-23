package accounts

import (
	"context"
	"fmt"
	"github.com/killbill/kbcli/kbcmd/cmdlib/args"
	"reflect"
	"time"

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

	acc, err := kblib.GetAccountByKeyOrIDWithBalanceAndCBA(ctx, o.Client(), o.Args[0])
	if err == nil {
		o.Print(acc)
	}

	return err
}

func createAccount(ctx context.Context, o *cmdlib.Options) error {
	accToCreate := &kbmodel.Account{}
	err := args.LoadProperties(accToCreate, createAccountPropertyList, o.Args)
	if err != nil {
		return err
	}

	accCreated, err := o.Client().Account.CreateAccount(ctx, &account.CreateAccountParams{
		Body: accToCreate,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return err
	}
	o.Print(accCreated.Payload)

	return nil
}

func updateAccount(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 2 {
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

// RegisterAccountCommands registers all account commands.
func RegisterAccountCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Account{}), accountFormatter)

	// Register top level command
	r.Register("", cli.Command{
		Name:    "accounts",
		Aliases: []string{"acc"},
		Usage:   "Account related commands",
	}, nil)

	// Get account
	r.Register("accounts", cli.Command{
		Name:        "get",
		Usage:       "Get account information",
		ArgsUsage:   "ACCOUNT",
		Description: "ACCOUNT can be the account id or external key. An optional + may be prefixed to external key for disambiguation",
	}, getAccount)

	// List all accounts
	r.Register("accounts", cli.Command{
		Name:      "list",
		Usage:     "List all accounts",
		ArgsUsage: "",
	}, listAccounts)

	// Create account
	createAccountPropertyList = args.GetProperties(&kbmodel.Account{})
	createAccountPropertyList.Get("ReferenceTime").Default = time.Now().Format(time.RFC3339)
	createAccountPropertyList.Get("TimeZone").Default = "UTC"
	createAccountPropertyList.Get("Currency").Default = string(kbmodel.AccountCurrencyUSD)
	createAccountPropertyList.Sort(true, true)

	createAccountsUsage := fmt.Sprintf(`%s

		For ex.,:
				kbcmd accounts create ExternalKey=prem1 Name="Prem Ramanathan" Email=prem@prem.com Currency=USD
				`,
		args.GenerateUsageString(&kbmodel.Account{}, createAccountPropertyList))

	r.Register("accounts", cli.Command{
		Name:        "create",
		Usage:       "Create new account",
		ArgsUsage:   createAccountsUsage,
		Description: "Creates new account",
	}, createAccount)

	// Update account
	updateAccountPropertyList = args.GetProperties(&kbmodel.Account{})
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
		Usage:       "Updates existing account",
		ArgsUsage:   updateAccountsUsage,
		Description: "Updates existing account",
	}, updateAccount)

	registerAccountPaymentCommands(r)
	registerAccountTagCommands(r)
	registerAccountCustomFieldCommands(r)
}
