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
	"github.com/go-openapi/swag"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewAdjustInvoiceItemParams creates a new AdjustInvoiceItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdjustInvoiceItemParams() *AdjustInvoiceItemParams {
	return &AdjustInvoiceItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdjustInvoiceItemParamsWithTimeout creates a new AdjustInvoiceItemParams object
// with the ability to set a timeout on a request.
func NewAdjustInvoiceItemParamsWithTimeout(timeout time.Duration) *AdjustInvoiceItemParams {
	return &AdjustInvoiceItemParams{
		timeout: timeout,
	}
}

// NewAdjustInvoiceItemParamsWithContext creates a new AdjustInvoiceItemParams object
// with the ability to set a context for a request.
func NewAdjustInvoiceItemParamsWithContext(ctx context.Context) *AdjustInvoiceItemParams {
	return &AdjustInvoiceItemParams{
		Context: ctx,
	}
}

// NewAdjustInvoiceItemParamsWithHTTPClient creates a new AdjustInvoiceItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdjustInvoiceItemParamsWithHTTPClient(client *http.Client) *AdjustInvoiceItemParams {
	return &AdjustInvoiceItemParams{
		HTTPClient: client,
	}
}

/*
AdjustInvoiceItemParams contains all the parameters to send to the API endpoint

	for the adjust invoice item operation.

	Typically these are written to a http.Request.
*/
type AdjustInvoiceItemParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.InvoiceItem

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the adjust invoice item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdjustInvoiceItemParams) WithDefaults() *AdjustInvoiceItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the adjust invoice item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdjustInvoiceItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithTimeout(timeout time.Duration) *AdjustInvoiceItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithContext(ctx context.Context) *AdjustInvoiceItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithHTTPClient(client *http.Client) *AdjustInvoiceItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithXKillbillComment(xKillbillComment *string) *AdjustInvoiceItemParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AdjustInvoiceItemParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithXKillbillReason(xKillbillReason *string) *AdjustInvoiceItemParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithBody(body *kbmodel.InvoiceItem) *AdjustInvoiceItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetBody(body *kbmodel.InvoiceItem) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithInvoiceID(invoiceID strfmt.UUID) *AdjustInvoiceItemParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WithPluginProperty adds the pluginProperty to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithPluginProperty(pluginProperty []string) *AdjustInvoiceItemParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the adjust invoice item params
func (o *AdjustInvoiceItemParams) WithRequestedDate(requestedDate *strfmt.Date) *AdjustInvoiceItemParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the adjust invoice item params
func (o *AdjustInvoiceItemParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *AdjustInvoiceItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}
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

// bindParamAdjustInvoiceItem binds the parameter pluginProperty
func (o *AdjustInvoiceItemParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
