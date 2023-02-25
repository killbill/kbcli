package commands

import (
	"context"
	"fmt"
	"github.com/killbill/kbcli/v3/kbcommon"
	"reflect"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbclient/account"
	"github.com/killbill/kbcli/v3/kbclient/invoice"

	"github.com/killbill/kbcli/v3/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v3/kbcmd/cmdlib/args"
	"github.com/killbill/kbcli/v3/kbcmd/kblib"
	"github.com/killbill/kbcli/v3/kbmodel"
	"github.com/urfave/cli"
)

var (
	getInvoiceProperties           args.Properties
	getInvoicePaymentsProperties   args.Properties
	listAccountInvoicesProperties  args.Properties
	dryRunInvoiceProperties        args.Properties
	targetInvoiceProperties        args.Properties
	payInvoiceProperties           args.Properties
	createExternalChargeProperties args.Properties
)

var invoicePaymentFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NUMBER",
			Path: "$.paymentNumber",
		},
		{
			Name: "INVOICE",
			Path: "$.targetInvoiceId",
		},
		{
			Name: "PURCHASE_AMOUNT",
			Path: "$.purchasedAmount",
		},
		{
			Name: "REFUND_AMOUNT",
			Path: "$.refundedAmount",
		},
	},
}

var invoiceItemFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "INVOICE_ITEM_ID",
			Path: "$.invoiceItemId",
		},
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

func getInvoicePayments(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	invoiceIDOrNumber := o.Args[0]

	var invoiceId strfmt.UUID
	if strfmt.IsUUID(invoiceIDOrNumber) {
		invoiceId = strfmt.UUID(invoiceIDOrNumber)
	} else {
		invoiceNumber, err := strconv.ParseInt(invoiceIDOrNumber, 10, 64)
		if err != nil {
			return err
		}
		invoiceByNumberParams := &invoice.GetInvoiceByNumberParams{
			InvoiceNumber: int32(invoiceNumber),
		}

		resp, err := o.Client().Invoice.GetInvoiceByNumber(ctx, invoiceByNumberParams)
		if err != nil {
			return err
		}
		invoiceId = resp.Payload.InvoiceID
	}

	withAttempts := true
	params := &invoice.GetPaymentsForInvoiceParams{
		InvoiceID:    invoiceId,
		WithAttempts: &withAttempts,
	}
	err := args.LoadProperties(params, getInvoicePaymentsProperties, o.Args[1:])
	if err != nil {
		return err
	}
	resp, err := o.Client().Invoice.GetPaymentsForInvoice(ctx, params)
	if err != nil {
		return err
	}

	o.Print(resp.Payload)

	return nil
}

type dryRunInvoiceParams struct {
	TargetDate     string
	SubscriptionID string
}

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

type targetInvoiceParams struct {
	TargetDate     string
	SubscriptionID string
}

func targetInvoice(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	var inputParams targetInvoiceParams
	if err := args.LoadProperties(&inputParams, targetInvoiceProperties, o.Args[1:]); err != nil {
		return err
	}

	p := &invoice.CreateFutureInvoiceParams{}

	if inputParams.TargetDate != "" {
		targetDate, err := time.Parse("2006-01-02", inputParams.TargetDate)
		if err != nil {
			return fmt.Errorf("unable to parse date %s. %v", inputParams.TargetDate, err)
		}
		p.TargetDate = (*strfmt.Date)(&targetDate)
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}
	p.AccountID = acc.AccountID
	p.ProcessLocationHeader = true

	invoice, err := o.Client().Invoice.CreateFutureInvoice(ctx, p)
	if err != nil {
		kberr, ok := err.(*kbcommon.KillbillError)
		if ok && kberr.HTTPCode == 404 {
			o.Output("No invoices to generate\n")
			return nil
		} else {
			return err
		}
	}

	if invoice != nil {
		o.Print(invoice.Payload)
	}

	return nil
}

type payInvoiceParams struct {
	Amount          string
	PaymentMethodID string
}

func payInvoice(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	invoiceIDOrNumber := o.Args[0]

	var invoiceId strfmt.UUID
	var invoiceBalance float64
	var accountId strfmt.UUID
	if !strfmt.IsUUID(invoiceIDOrNumber) {
		invoiceNumber, err := strconv.ParseInt(invoiceIDOrNumber, 10, 64)
		if err != nil {
			return err
		}
		invoiceByNumberParams := &invoice.GetInvoiceByNumberParams{
			InvoiceNumber: int32(invoiceNumber),
		}

		resp, err := o.Client().Invoice.GetInvoiceByNumber(ctx, invoiceByNumberParams)
		if err != nil {
			return err
		}
		invoiceId = resp.Payload.InvoiceID
		invoiceBalance = resp.Payload.Balance
		accountId = resp.Payload.AccountID
	} else {
		invoiceId = strfmt.UUID(invoiceIDOrNumber)
		invoiceByIdParams := &invoice.GetInvoiceParams{
			InvoiceID: invoiceId,
		}

		resp, err := o.Client().Invoice.GetInvoice(ctx, invoiceByIdParams)
		if err != nil {
			return err
		}
		invoiceBalance = resp.Payload.Balance
		accountId = resp.Payload.AccountID
	}

	var inputParams payInvoiceParams
	if err := args.LoadProperties(&inputParams, payInvoiceProperties, o.Args[1:]); err != nil {
		return err
	}

	createInstantPaymentParams := &invoice.CreateInstantPaymentParams{
		InvoiceID: invoiceId,
		Body: &kbmodel.InvoicePayment{
			AccountID:       accountId,
			TargetInvoiceID: invoiceId,
		},
		// TODO Not respected because of https://github.com/killbill/kbcli/issues/12
		ProcessLocationHeader: true,
	}

	if inputParams.Amount == "" {
		if invoiceBalance == 0 {
			o.Output("Nothing to pay\n")
			return nil
		}
		createInstantPaymentParams.Body.PurchasedAmount = invoiceBalance
	} else {
		amount, err := strconv.ParseFloat(inputParams.Amount, 64)
		if err != nil {
			return err
		}
		createInstantPaymentParams.Body.PurchasedAmount = amount
	}

	if inputParams.PaymentMethodID != "" {
		createInstantPaymentParams.Body.PaymentMethodID = strfmt.UUID(inputParams.PaymentMethodID)
	}

	resp, _, err := o.Client().Invoice.CreateInstantPayment(ctx, createInstantPaymentParams)
	if err != nil {
		return err
	}

	if resp != nil {
		o.Print(resp.Payload)
	} else {
		o.Log.Warningf("Unable to trigger payment (missing payment method?)")
	}

	return nil
}

type createExternalChargeParams struct {
	Amount     float64
	AutoCommit bool
}

func createExternalCharge(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 1 {
		return cmdlib.ErrorInvalidArgs
	}
	accIDOrKey := o.Args[0]

	var inputParams createExternalChargeParams
	if err := args.LoadProperties(&inputParams, createExternalChargeProperties, o.Args[1:]); err != nil {
		return err
	}

	p := &invoice.CreateExternalChargesParams{
		Body: []*kbmodel.InvoiceItem{&kbmodel.InvoiceItem{
			Amount: inputParams.Amount,
		}},
		AutoCommit: &inputParams.AutoCommit,
	}

	acc, err := kblib.GetAccountByKeyOrID(ctx, o.Client(), accIDOrKey)
	if err != nil {
		return err
	}
	p.AccountID = acc.AccountID

	items, err := o.Client().Invoice.CreateExternalCharges(ctx, p)
	if err != nil {
		return err
	}

	o.Print(items.Payload)

	return nil
}

func registerInvoicesCommands(r *cmdlib.App) {
	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.Invoice{}), invoiceFormatter)

	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.InvoiceItem{}), invoiceItemFormatter)

	// Register formatters
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.InvoicePayment{}), invoicePaymentFormatter)

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

	getInvoiceUsage := args.GenerateUsageString(&invoice.GetInvoiceParams{}, getInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "get",
		Usage: "get invoice",
		ArgsUsage: fmt.Sprintf(`INVOICE_ID %s

   For e.g.,
	   kbcmd invoices get 57f3da8e-6125-43a5-9a38-7b448b15da83
	   kbcmd invoices get 2`, getInvoiceUsage),
	}, getInvoice)

	// get invoice payment
	getInvoicePaymentsProperties = args.GetProperties(&invoice.GetPaymentsForInvoiceParams{})
	getInvoicePaymentsUsage := args.GenerateUsageString(&invoice.GetPaymentsForInvoiceParams{}, getInvoicePaymentsProperties)
	r.Register("invoices", cli.Command{
		Name:  "payments",
		Usage: "payments invoice",
		ArgsUsage: fmt.Sprintf(`INVOICE_ID %s

   For e.g.,
	   kbcmd invoices payments 57f3da8e-6125-43a5-9a38-7b448b15da83
	   kbcmd invoices payments 2`, getInvoicePaymentsUsage),
	}, getInvoicePayments)

	// DryRun invoices
	dryRunInvoiceProperties = args.GetProperties(&dryRunInvoiceParams{})
	dryRunInvoiceUsage := args.GenerateUsageString(&dryRunInvoiceParams{}, dryRunInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "dry-run",
		Usage: "Dry run invoice generation for a given account",
		ArgsUsage: fmt.Sprintf(`ACCOUNT %s
For ex.,
kbcmd invoices dry-run account3 TargetDate=2018-05-05

will generate a dry-run invoice for the given date. If date is omitted, next upcoming invoice will be generated.
`, dryRunInvoiceUsage),
	}, dryRunInvoice)

	// Target invoices
	targetInvoiceProperties = args.GetProperties(&targetInvoiceParams{})
	targetInvoiceUsage := args.GenerateUsageString(&targetInvoiceParams{}, targetInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "target-run",
		Usage: "Future invoice generation for a given account",
		ArgsUsage: fmt.Sprintf(`ACCOUNT %s
For ex.,
kbcmd invoices target-run account3 TargetDate=2018-05-05

will generate an invoice for the given date.
`, targetInvoiceUsage),
	}, targetInvoice)

	// Pay invoice
	payInvoiceProperties = args.GetProperties(&payInvoiceParams{})
	payInvoiceUsage := args.GenerateUsageString(&payInvoiceParams{}, payInvoiceProperties)
	r.Register("invoices", cli.Command{
		Name:  "pay",
		Usage: "Trigger a payment for a given invoice",
		ArgsUsage: fmt.Sprintf(`INVOICE_ID %s
For ex.,
kbcmd invoices pay 7fcc7b69-9fdb-4143-98e6-94f3bad4842f Amount=5 PaymentMethodId=d7c14d0e-2368-4214-a522-92ce6e00f535
`, payInvoiceUsage),
	}, payInvoice)

	// Create external charge
	createExternalChargeProperties = args.GetProperties(&createExternalChargeParams{})
	createExternalChargeUsage := args.GenerateUsageString(&createExternalChargeParams{}, createExternalChargeProperties)
	r.Register("invoices", cli.Command{
		Name:  "charge",
		Usage: "Create an external charge for a given account",
		ArgsUsage: fmt.Sprintf(`ACCOUNT %s
For ex.,
kbcmd invoices charge account3 Amount=10 AutoCommit=false

will generate a DRAFT invoice for $10.
`, createExternalChargeUsage),
	}, createExternalCharge)
}
