// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewUploadInvoiceTemplateParams creates a new UploadInvoiceTemplateParams object
// with the default values initialized.
func NewUploadInvoiceTemplateParams() *UploadInvoiceTemplateParams {
	var (
		deleteIfExistsDefault = bool(false)
	)
	return &UploadInvoiceTemplateParams{
		DeleteIfExists: &deleteIfExistsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadInvoiceTemplateParamsWithTimeout creates a new UploadInvoiceTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadInvoiceTemplateParamsWithTimeout(timeout time.Duration) *UploadInvoiceTemplateParams {
	var (
		deleteIfExistsDefault = bool(false)
	)
	return &UploadInvoiceTemplateParams{
		DeleteIfExists: &deleteIfExistsDefault,

		timeout: timeout,
	}
}

// NewUploadInvoiceTemplateParamsWithContext creates a new UploadInvoiceTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadInvoiceTemplateParamsWithContext(ctx context.Context) *UploadInvoiceTemplateParams {
	var (
		deleteIfExistsDefault = bool(false)
	)
	return &UploadInvoiceTemplateParams{
		DeleteIfExists: &deleteIfExistsDefault,

		Context: ctx,
	}
}

// NewUploadInvoiceTemplateParamsWithHTTPClient creates a new UploadInvoiceTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadInvoiceTemplateParamsWithHTTPClient(client *http.Client) *UploadInvoiceTemplateParams {
	var (
		deleteIfExistsDefault = bool(false)
	)
	return &UploadInvoiceTemplateParams{
		DeleteIfExists: &deleteIfExistsDefault,
		HTTPClient:     client,
	}
}

/*UploadInvoiceTemplateParams contains all the parameters to send to the API endpoint
for the upload invoice template operation typically these are written to a http.Request
*/
type UploadInvoiceTemplateParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body string
	/*DeleteIfExists*/
	DeleteIfExists *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithTimeout(timeout time.Duration) *UploadInvoiceTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithContext(ctx context.Context) *UploadInvoiceTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithHTTPClient(client *http.Client) *UploadInvoiceTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithXKillbillComment(xKillbillComment *string) *UploadInvoiceTemplateParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadInvoiceTemplateParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithXKillbillReason(xKillbillReason *string) *UploadInvoiceTemplateParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithBody(body string) *UploadInvoiceTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetBody(body string) {
	o.Body = body
}

// WithDeleteIfExists adds the deleteIfExists to the upload invoice template params
func (o *UploadInvoiceTemplateParams) WithDeleteIfExists(deleteIfExists *bool) *UploadInvoiceTemplateParams {
	o.SetDeleteIfExists(deleteIfExists)
	return o
}

// SetDeleteIfExists adds the deleteIfExists to the upload invoice template params
func (o *UploadInvoiceTemplateParams) SetDeleteIfExists(deleteIfExists *bool) {
	o.DeleteIfExists = deleteIfExists
}

// WriteToRequest writes these params to a swagger request
func (o *UploadInvoiceTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.DeleteIfExists != nil {

		// query param deleteIfExists
		var qrDeleteIfExists bool
		if o.DeleteIfExists != nil {
			qrDeleteIfExists = *o.DeleteIfExists
		}
		qDeleteIfExists := swag.FormatBool(qrDeleteIfExists)
		if qDeleteIfExists != "" {
			if err := r.SetQueryParam("deleteIfExists", qDeleteIfExists); err != nil {
				return err
			}
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
