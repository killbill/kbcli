package killbill

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient/invoice"
	"github.com/killbill/kbcli/v2/kbmodel"
)

func (cli *RawClient) triggerDryRunAction(ctx context.Context, accountId strfmt.UUID, dryRunAction kbmodel.InvoiceDryRunDryRunActionEnum, bundleId strfmt.UUID, subscriptionId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	invOk, invNA, err := cli.TenantClient.Invoice.GenerateDryRunInvoice(ctx, &invoice.GenerateDryRunInvoiceParams{
		AccountID: accountId,
		Body: &kbmodel.InvoiceDryRun{
			DryRunAction:    dryRunAction,
			DryRunType:      kbmodel.InvoiceDryRunDryRunTypeSUBSCRIPTIONACTION,
			EffectiveDate:   effDt,
			BundleID:        bundleId,
			SubscriptionID:  subscriptionId,
			BillingPeriod:   kbmodel.InvoiceDryRunBillingPeriodMONTHLY,
			ProductCategory: kbmodel.InvoiceDryRunProductCategoryBASE,
			ProductName:     product,
			PriceListName:   priceList,
		},
		ProcessLocationHeader: true,
		TargetDate:            &targetDt,
	})
	if err != nil {
		return nil, err
	}

	if invNA != nil {
		return nil, nil
	} else {
		return invOk.Payload, nil
	}
}

func (cli *RawClient) TriggerDryRunCreateSubscription(ctx context.Context, accountId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionSTARTBILLING, "", "", product, priceList, effDt, targetDt)
}

func (cli *RawClient) TriggerDryRunChangePlan(ctx context.Context, accountId strfmt.UUID, bundleId strfmt.UUID, subscriptionId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionCHANGE, bundleId, subscriptionId, product, priceList, effDt, targetDt)
}

func (cli *RawClient) TriggerDryRunCancelSubscription(ctx context.Context, accountId strfmt.UUID, bundleId strfmt.UUID, subscriptionId strfmt.UUID, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionSTOPBILLING, bundleId, subscriptionId, "", "", effDt, targetDt)
}

// GetInvoice fetch an invoice using the given invoiceID.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Invoice.GetInvoice(ctx, &invoice.GetInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetAllInvoices fetchs all invoice for the given account.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetAllInvoices(ctx context.Context, accountID strfmt.UUID) ([]*kbmodel.Invoice, error) {
	return cli.GetInvoicesBetweenDates(ctx, accountID, nil, nil)
}

// TriggerTargetInvoice creates a DRAFT invoice for the given account,
// using the given target date.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) TriggerTargetInvoice(ctx context.Context, accountID strfmt.UUID, targetDate strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	inv, err := cli.TenantClient.Invoice.CreateFutureInvoice(ctx, &invoice.CreateFutureInvoiceParams{
		AccountID:             accountID,
		TargetDate:            &targetDate,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return nil, err
	}

	return inv.Payload, nil
}

// VoidInvoice voids an invoice. If the invoice was already paid, voiding it fails.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) VoidInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Invoice.VoidInvoice(ctx, &invoice.VoidInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// CommitInvoice commits an invoice. If the current invoice status is different from DRAFT, it returns an error.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) CommitInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Invoice.CommitInvoice(ctx, &invoice.CommitInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// PayInvoice adds a full-amount payment to the given invoice.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) PayInvoice(ctx context.Context, invoiceID strfmt.UUID, accountID strfmt.UUID, amount float64, externalPayment bool) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, _, err := cli.TenantClient.Invoice.CreateInstantPayment(ctx, &invoice.CreateInstantPaymentParams{
		InvoiceID:       invoiceID,
		ExternalPayment: &externalPayment,
		Body: &kbmodel.InvoicePayment{
			AccountID:       accountID,
			PurchasedAmount: amount,
		},
	})

	return err
}

// IsInvoiceWrittenOff checks if the given invoice is written-off, returning true or false
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) IsInvoiceWrittenOff(ctx context.Context, invoiceID strfmt.UUID) (bool, string, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()
	reason := ""
	audit := "FULL"

	// Get the Invoice tags with the audit information
	res, err := cli.TenantClient.Invoice.GetInvoiceTags(ctx, &invoice.GetInvoiceTagsParams{
		InvoiceID: invoiceID,
		Audit:     &audit,
	})
	if err != nil {
		return false, reason, err
	}

	for _, tag := range res.Payload {
		if tag.TagDefinitionID == WRITTEN_OFF {
			// Gets tag definition audit logs and returns Comments field value from the latest one
			if len(tag.AuditLogs) >= 1 {
				reason = tag.AuditLogs[len(tag.AuditLogs)-1].Comments
			}
			return true, reason, nil
		}
	}

	return false, reason, nil
}

// WriteOffInvoice writes-off the given invoice.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) WriteOffInvoice(ctx context.Context, invoiceID strfmt.UUID, reason string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := []strfmt.UUID{WRITTEN_OFF}

	_, err := cli.TenantClient.Invoice.CreateInvoiceTags(ctx, &invoice.CreateInvoiceTagsParams{
		InvoiceID:        invoiceID,
		Body:             fmtTags,
		XKillbillComment: &reason,
	})

	return err
}
