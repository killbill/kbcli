// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewModifyTransactionCustomFieldsParams creates a new ModifyTransactionCustomFieldsParams object
// with the default values initialized.
func NewModifyTransactionCustomFieldsParams() *ModifyTransactionCustomFieldsParams {
	var ()
	return &ModifyTransactionCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyTransactionCustomFieldsParamsWithTimeout creates a new ModifyTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyTransactionCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyTransactionCustomFieldsParams {
	var ()
	return &ModifyTransactionCustomFieldsParams{

		timeout: timeout,
	}
}

// NewModifyTransactionCustomFieldsParamsWithContext creates a new ModifyTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyTransactionCustomFieldsParamsWithContext(ctx context.Context) *ModifyTransactionCustomFieldsParams {
	var ()
	return &ModifyTransactionCustomFieldsParams{

		Context: ctx,
	}
}

// NewModifyTransactionCustomFieldsParamsWithHTTPClient creates a new ModifyTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyTransactionCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyTransactionCustomFieldsParams {
	var ()
	return &ModifyTransactionCustomFieldsParams{
		HTTPClient: client,
	}
}

/*ModifyTransactionCustomFieldsParams contains all the parameters to send to the API endpoint
for the modify transaction custom fields operation typically these are written to a http.Request
*/
type ModifyTransactionCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*TransactionID*/
	TransactionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyTransactionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithContext(ctx context.Context) *ModifyTransactionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyTransactionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyTransactionCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyTransactionCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyTransactionCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyTransactionCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithTransactionID adds the transactionID to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) WithTransactionID(transactionID strfmt.UUID) *ModifyTransactionCustomFieldsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the modify transaction custom fields params
func (o *ModifyTransactionCustomFieldsParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyTransactionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
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
