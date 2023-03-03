// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPaymentMethodCustomFieldsParams creates a new GetPaymentMethodCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPaymentMethodCustomFieldsParams() *GetPaymentMethodCustomFieldsParams {
	return &GetPaymentMethodCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithTimeout creates a new GetPaymentMethodCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewGetPaymentMethodCustomFieldsParamsWithTimeout(timeout time.Duration) *GetPaymentMethodCustomFieldsParams {
	return &GetPaymentMethodCustomFieldsParams{
		timeout: timeout,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithContext creates a new GetPaymentMethodCustomFieldsParams object
// with the ability to set a context for a request.
func NewGetPaymentMethodCustomFieldsParamsWithContext(ctx context.Context) *GetPaymentMethodCustomFieldsParams {
	return &GetPaymentMethodCustomFieldsParams{
		Context: ctx,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithHTTPClient creates a new GetPaymentMethodCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPaymentMethodCustomFieldsParamsWithHTTPClient(client *http.Client) *GetPaymentMethodCustomFieldsParams {
	return &GetPaymentMethodCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
GetPaymentMethodCustomFieldsParams contains all the parameters to send to the API endpoint

	for the get payment method custom fields operation.

	Typically these are written to a http.Request.
*/
type GetPaymentMethodCustomFieldsParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentMethodCustomFieldsParams) WithDefaults() *GetPaymentMethodCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentMethodCustomFieldsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")
	)

	val := GetPaymentMethodCustomFieldsParams{
		Audit: &auditDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithTimeout(timeout time.Duration) *GetPaymentMethodCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithContext(ctx context.Context) *GetPaymentMethodCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithHTTPClient(client *http.Client) *GetPaymentMethodCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithAudit(audit *string) *GetPaymentMethodCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPaymentMethodID adds the paymentMethodID to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *GetPaymentMethodCustomFieldsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
