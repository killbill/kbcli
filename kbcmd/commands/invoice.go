package commands

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbclient/invoice"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/cmdlib/args"
	"github.com/killbill/kbcli/kbcmd/kblib"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var (
	getInvoiceProperties          args.Properties
	listAccountInvoicesProperties args.Properties
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
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}

	params := &account.GetInvoicesForAccountParams{
		AccountID: acc.AccountID,
	}

	if err := args.LoadProperties(params, listAccountInvoicesProperties, o.Args[1:]); err != nil {
		return err
	}

	resp, err := o.Client().Account.GetInvoicesForAccount(ctx, params)
	if err != nil {
		return err
	}

	o.Print(resp.Payload)

	return nil
}

func getInvoice(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	invoiceIDOrNumber := o.Args[0]

	if strfmt.IsUUID(invoiceIDOrNumber) {
		params := &invoice.GetInvoiceParams{
			InvoiceID: strfmt.UUID(invoiceIDOrNumber),
		}
		err := args.LoadProperties(params, getInvoiceProperties, o.Args[1:])
		if err != nil {
			return err
		}
		resp, err := o.Client().Invoice.GetInvoice(ctx, params)
		if err != nil {
			return err
		}

		o.Print(resp.Payload)
	} else {
		params := &invoice.GetInvoiceByNumberParams{}
		invoiceNumber, err := strconv.ParseInt(invoiceIDOrNumber, 10, 64)
		if err != nil {
			return err
		}
		params.InvoiceNumber = int32(invoiceNumber)
		err = args.LoadProperties(params, getInvoiceProperties, o.Args[1:])
		if err != nil {
			return err
		}
		resp, err := o.Client().Invoice.GetInvoiceByNumber(ctx, params)
		if err != nil {
			return err
		}

		o.Print(resp.Payload)
	}

	return nil
}

type dryRunInvoiceParams struct {
	TargetDate     string
	SubscriptionID string
}

var dryRunInvoiceProperties args.Properties

func dryRunInvoice(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	var inputParams dryRunInvoiceParams
	if err := args.LoadProperties(&inputParams, dryRunInvoiceProperties, o.Args[1:]); err != nil {
		return err
	}

	p := &invoice.GenerateDryRunInvoiceParams{
		Body: &kbmodel.InvoiceDryRun{
			DryRunType:     kbmodel.InvoiceDryRunDryRunTypeUPCOMINGINVOICE,
			SubscriptionID: strfmt.UUID(inputParams.SubscriptionID),
		},
	}

	if inputParams.TargetDate != "" {
		targetDate, err := time.Parse("2006-01-02", inputParams.TargetDate)
		if err != nil {
			return fmt.Errorf("unable to parse date %s. %v", inputParams.TargetDate, err)
		}
		p.TargetDate = (*strfmt.Date)(&targetDate)
		p.Body.DryRunType = kbmodel.InvoiceDryRunDryRunTypeTARGETDATE
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}
	p.AccountID = acc.AccountID

	invoice, _, err := o.Client().Invoice.GenerateDryRunInvoice(ctx, p)
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
	listAccountInvoicesProperties = args.GetProperties(&account.GetInvoicesForAccountParams{})
	listAccountInvoicesProperties.Remove("AccountID")
	getInvoicesForAccountUsage := args.GenerateUsageString(&account.GetInvoicesForAccountParams{}, listAccountInvoicesProperties)
	r.Register("invoices", cli.Command{
		Name:  "list",
		Usage: "list all invoices for a given account",
		ArgsUsage: fmt.Sprintf(`ACCOUNT %s
For ex.,
kbcmd invoices list account3 UnpaidInvoicesOnly=true
`, getInvoicesForAccountUsage),
	}, listAccountInvoices)

	// get invoice (Both getInvoice and getInvoiceByNumber share the same properties)
	getInvoiceProperties = args.GetProperties(&invoice.GetInvoiceParams{})
	getInvoiceProperties.Remove("InvoiceID")
	getInvoiceProperties.Get("WithItems").Default = "True"

	getInvoiceUsage := args.GenerateUsageString(&invoice.GetInvoiceParams{}, getInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "get",
		Usage: "get invoice",
		ArgsUsage: fmt.Sprintf(`INVOICE_ID %s

   For e.g.,
	   kbcmd invoices get 57f3da8e-6125-43a5-9a38-7b448b15da83
	   kbcmd invoices get 2`, getInvoiceUsage),
	}, getInvoice)

	// DryRun invoices
	dryRunInvoiceProperties = args.GetProperties(&dryRunInvoiceParams{})
	dryRunInvoiceUsage := args.GenerateUsageString(&dryRunInvoiceParams{}, dryRunInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "dry-run",
		Usage: "Dry run invoice generation for a given account",
		ArgsUsage: fmt.Sprintf(`ACCOUNT %s
For ex.,
kbcmd invoices dry-run account3 TargetDate=2018-05-05

will generate invoice for the given date. If date is omitted, next upcoming invoice will be generated.
`, dryRunInvoiceUsage),
	}, dryRunInvoice)
}
