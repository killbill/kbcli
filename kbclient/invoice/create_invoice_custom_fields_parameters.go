// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewCreateInvoiceCustomFieldsParams creates a new CreateInvoiceCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInvoiceCustomFieldsParams() *CreateInvoiceCustomFieldsParams {
	return &CreateInvoiceCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithTimeout creates a new CreateInvoiceCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewCreateInvoiceCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateInvoiceCustomFieldsParams {
	return &CreateInvoiceCustomFieldsParams{
		timeout: timeout,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithContext creates a new CreateInvoiceCustomFieldsParams object
// with the ability to set a context for a request.
func NewCreateInvoiceCustomFieldsParamsWithContext(ctx context.Context) *CreateInvoiceCustomFieldsParams {
	return &CreateInvoiceCustomFieldsParams{
		Context: ctx,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithHTTPClient creates a new CreateInvoiceCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInvoiceCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateInvoiceCustomFieldsParams {
	return &CreateInvoiceCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
CreateInvoiceCustomFieldsParams contains all the parameters to send to the API endpoint

	for the create invoice custom fields operation.

	Typically these are written to a http.Request.
*/
type CreateInvoiceCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create invoice custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceCustomFieldsParams) WithDefaults() *CreateInvoiceCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create invoice custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateInvoiceCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithContext(ctx context.Context) *CreateInvoiceCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateInvoiceCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateInvoiceCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithInvoiceID(invoiceID strfmt.UUID) *CreateInvoiceCustomFieldsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoiceCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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
