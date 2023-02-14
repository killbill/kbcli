// Code generated by go-swagger; DO NOT EDIT.

package custom_field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchCustomFieldsParams creates a new SearchCustomFieldsParams object
// with the default values initialized.
func NewSearchCustomFieldsParams() *SearchCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchCustomFieldsParamsWithTimeout creates a new SearchCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchCustomFieldsParamsWithTimeout(timeout time.Duration) *SearchCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewSearchCustomFieldsParamsWithContext creates a new SearchCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchCustomFieldsParamsWithContext(ctx context.Context) *SearchCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewSearchCustomFieldsParamsWithHTTPClient creates a new SearchCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchCustomFieldsParamsWithHTTPClient(client *http.Client) *SearchCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchCustomFieldsParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*SearchCustomFieldsParams contains all the parameters to send to the API endpoint
for the search custom fields operation typically these are written to a http.Request
*/
type SearchCustomFieldsParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*SearchKey*/
	SearchKey string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the search custom fields params
func (o *SearchCustomFieldsParams) WithTimeout(timeout time.Duration) *SearchCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search custom fields params
func (o *SearchCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search custom fields params
func (o *SearchCustomFieldsParams) WithContext(ctx context.Context) *SearchCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search custom fields params
func (o *SearchCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search custom fields params
func (o *SearchCustomFieldsParams) WithHTTPClient(client *http.Client) *SearchCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search custom fields params
func (o *SearchCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the search custom fields params
func (o *SearchCustomFieldsParams) WithAudit(audit *string) *SearchCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search custom fields params
func (o *SearchCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search custom fields params
func (o *SearchCustomFieldsParams) WithLimit(limit *int64) *SearchCustomFieldsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search custom fields params
func (o *SearchCustomFieldsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search custom fields params
func (o *SearchCustomFieldsParams) WithOffset(offset *int64) *SearchCustomFieldsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search custom fields params
func (o *SearchCustomFieldsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchKey adds the searchKey to the search custom fields params
func (o *SearchCustomFieldsParams) WithSearchKey(searchKey string) *SearchCustomFieldsParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search custom fields params
func (o *SearchCustomFieldsParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WriteToRequest writes these params to a swagger request
func (o *SearchCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
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
