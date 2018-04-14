// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoicePaymentTagsParams creates a new GetInvoicePaymentTagsParams object
// with the default values initialized.
func NewGetInvoicePaymentTagsParams() *GetInvoicePaymentTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoicePaymentTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicePaymentTagsParamsWithTimeout creates a new GetInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicePaymentTagsParamsWithTimeout(timeout time.Duration) *GetInvoicePaymentTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoicePaymentTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetInvoicePaymentTagsParamsWithContext creates a new GetInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicePaymentTagsParamsWithContext(ctx context.Context) *GetInvoicePaymentTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoicePaymentTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetInvoicePaymentTagsParamsWithHTTPClient creates a new GetInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicePaymentTagsParamsWithHTTPClient(client *http.Client) *GetInvoicePaymentTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoicePaymentTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetInvoicePaymentTagsParams contains all the parameters to send to the API endpoint
for the get invoice payment tags operation typically these are written to a http.Request
*/
type GetInvoicePaymentTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*IncludedDeleted*/
	IncludedDeleted *bool
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithTimeout(timeout time.Duration) *GetInvoicePaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithContext(ctx context.Context) *GetInvoicePaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithHTTPClient(client *http.Client) *GetInvoicePaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoicePaymentTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoicePaymentTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithAudit(audit *string) *GetInvoicePaymentTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetInvoicePaymentTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithPaymentID adds the paymentID to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *GetInvoicePaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) WithPluginProperty(pluginProperty []string) *GetInvoicePaymentTagsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get invoice payment tags params
func (o *GetInvoicePaymentTagsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicePaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
		}

	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}