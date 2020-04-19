package accounts

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient"
	"github.com/killbill/kbcli/v2/kbmodel"

	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v2/kbcmd/kblib"
	"github.com/urfave/cli"
)

func getCustomFieldByNameOrID(
	ctx context.Context,
	c *kbclient.KillBill,
	accID strfmt.UUID,
	cfIDOrName string) (*kbmodel.CustomField, error) {

	cfKeyOrID, isID := kblib.ParseKeyOrID(cfIDOrName)

	resp, err := c.Account.GetAccountCustomFields(ctx, &account.GetAccountCustomFieldsParams{
		AccountID: accID,
	})
	if err != nil {
		return nil, err
	}

	for _, cf := range resp.Payload {
		if (isID && cfKeyOrID == string(cf.CustomFieldID)) || cfKeyOrID == *cf.Name {
			return cf, nil
		}
	}
	return nil, fmt.Errorf("custom field %s not found", cfIDOrName)
}

// listCustomFields - list all custom fields
func listCustomFields(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), o.Args[0])
	if err != nil {
		return err
	}

	resp, err := o.Client().Account.GetAccountCustomFields(ctx, &account.GetAccountCustomFieldsParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)

	return err
}

// addCustomField - Add custom field to account
func addCustomField(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 3 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrName := o.Args[0]
	cfName := o.Args[1]
	cfValue := o.Args[2]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrName)
	if err != nil {
		return err
	}
	cf := &kbmodel.CustomField{
		Name:  &cfName,
		Value: &cfValue,
	}

	resp, err := o.Client().Account.CreateAccountCustomFields(ctx, &account.CreateAccountCustomFieldsParams{
		AccountID:             acc.AccountID,
		Body:                  []*kbmodel.CustomField{cf},
		ProcessLocationHeader: true,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)

	return err
}

// deleteCustomField - Delete custom field to account
func deleteCustomField(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrName := o.Args[0]
	cfIDOrName := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrName)
	if err != nil {
		return err
	}

	cf, err := getCustomFieldByNameOrID(ctx, o.Client(), acc.AccountID, cfIDOrName)
	if err != nil {
		return nil
	}

	_, err = o.Client().Account.DeleteAccountCustomFields(ctx, &account.DeleteAccountCustomFieldsParams{
		AccountID:   acc.AccountID,
		CustomField: []strfmt.UUID{cf.CustomFieldID},
	})
	if err != nil {
		return err
	}

	return err
}

func registerAccountCustomFieldCommands(r *cmdlib.App) {
	r.Register("accounts", cli.Command{
		Name:    "custom-fields",
		Aliases: []string{"cf"},
		Usage:   "Custom field related commands",
	}, nil)

	// List custom fields
	r.Register("accounts.custom-fields", cli.Command{
		Name:      "list",
		Usage:     "List account custom fields",
		ArgsUsage: "ACCOUNT",
	}, listCustomFields)

	// Delete custom field
	r.Register("accounts.custom-fields", cli.Command{
		Name:    "delete",
		Aliases: []string{"del"},
		Usage:   "Delete account custom field",
		ArgsUsage: `ACCOUNT CF_NAME_OR_ID
   For e.g.,
      account custom-fields delete johndoe secondary-email`,
	}, deleteCustomField)

	// Add custom fields
	r.Register("accounts.custom-fields", cli.Command{
		Name:  "add",
		Usage: "Add account custom field",
		ArgsUsage: `ACCOUNT CF_ID VALUE

   For e.g.,
      accounts custom-fields add johndoe secondary-email john@john.com`,
	}, addCustomField)
}
