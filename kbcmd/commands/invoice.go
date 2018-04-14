package commands

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbclient/invoice"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var invoiceItemFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "AMOUNT",
			Path: "$.amount",
		},
		{
			Name: "RATE",
			Path: "$.rate",
		},
		{
			Name: "START_DATE",
			Path: "$.startDate",
		},
		{
			Name: "PLAN",
			Path: "$.planName",
		},
	},
}

var invoiceFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "AMOUNT",
			Path: "$.amount",
		},
		{
			Name: "BALANCE",
			Path: "$.balance",
		},
		{
			Name: "INVOICE_ID",
			Path: "$.invoiceId",
		},
		{
			Name: "TARGET_DATE",
			Path: "$.targetDate",
		},
	},
	SubItems: []cmdlib.SubItem{
		{
			Name:      "Invoice Items",
			FieldName: "Items",
		},
	},
}

func listAccountInvoices(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	resp, err := o.Client().Account.GetInvoicesForAccount(ctx, &account.GetInvoicesForAccountParams{
		AccountID: acc.AccountID,
	})
	if err != nil {
		return err
	}

	o.Print(resp.Payload)

	return nil
}

func dryRunInvoice(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 2 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]
	dateStr := o.Args[1]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	targetDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("unable to parse date %s. %v", dateStr, err)
	}

	invoice, _, err := o.Client().Invoice.GenerateDryRunInvoice(ctx, &invoice.GenerateDryRunInvoiceParams{
		AccountID:  acc.AccountID,
		TargetDate: (*strfmt.Date)(&targetDate),
	})
	if err != nil {
		return err
	}

	if invoice != nil {
		o.Print(invoice.Payload)
	} else {
		o.Output("No invoices to generate\n")
	}

	return nil
}

func registerInvoicesCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Invoice{}), invoiceFormatter)

	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.InvoiceItem{}), invoiceItemFormatter)

	// Register top level command
	r.Register("", cli.Command{
		Name:    "invoices",
		Aliases: []string{"inv"},
		Usage:   "invoices related commands",
	}, nil)

	// List invoices
	r.Register("invoices", cli.Command{
		Name:      "list",
		Usage:     "list all invoices for a given account",
		ArgsUsage: "ACCOUNT",
	}, listAccountInvoices)

	// DryRun invoices
	r.Register("invoices", cli.Command{
		Name:  "dry-run",
		Usage: "Dry run invoice generation for a given account",
		ArgsUsage: `ACCOUNT TARGET_DATE
For ex.,
kbcmd invoices dry-run account3 2018-05-05
`,
	}, dryRunInvoice)
}
