package accounts

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v2/kbcmd/kblib"
	"github.com/killbill/kbcli/v2/kbmodel"
	"github.com/urfave/cli"
)

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

	resp, err := o.Client().Account.CreateAccountTags(ctx, &account.CreateAccountTagsParams{
		AccountID:             acc.AccountID,
		Body:                  []strfmt.UUID{tag.ID},
		ProcessLocationHeader: true,
	})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)

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

func registerAccountTagCommands(r *cmdlib.App) {
	// Tag
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Tag{}), tagFormatter)

	// Account tag related functions
	r.Register("accounts", cli.Command{
		Name:  "tags",
		Usage: "Tag related commands",
	}, nil)

	// List account tags
	r.Register("accounts.tags", cli.Command{
		Name:      "list",
		Aliases:   []string{"ls"},
		Usage:     "List all account tags",
		ArgsUsage: "ACCOUNT",
	}, listAccountTags)

	// Add account tag
	r.Register("accounts.tags", cli.Command{
		Name:      "add",
		Usage:     "Add tag to the account",
		ArgsUsage: "ACCOUNT TAG_NAME",
	}, addAccountTag)

	// Remove account tag
	r.Register("accounts.tags", cli.Command{
		Name:      "remove",
		Aliases:   []string{"rm"},
		Usage:     "Remove tag from the account",
		ArgsUsage: "ACCOUNT TAG_NAME",
	}, removeAccountTag)
}
