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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCustomFieldAuditLogsWithHistoryParams creates a new GetCustomFieldAuditLogsWithHistoryParams object
// with the default values initialized.
func NewGetCustomFieldAuditLogsWithHistoryParams() *GetCustomFieldAuditLogsWithHistoryParams {
	var ()
	return &GetCustomFieldAuditLogsWithHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomFieldAuditLogsWithHistoryParamsWithTimeout creates a new GetCustomFieldAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomFieldAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetCustomFieldAuditLogsWithHistoryParams {
	var ()
	return &GetCustomFieldAuditLogsWithHistoryParams{

		timeout: timeout,
	}
}

// NewGetCustomFieldAuditLogsWithHistoryParamsWithContext creates a new GetCustomFieldAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomFieldAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetCustomFieldAuditLogsWithHistoryParams {
	var ()
	return &GetCustomFieldAuditLogsWithHistoryParams{

		Context: ctx,
	}
}

// NewGetCustomFieldAuditLogsWithHistoryParamsWithHTTPClient creates a new GetCustomFieldAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomFieldAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetCustomFieldAuditLogsWithHistoryParams {
	var ()
	return &GetCustomFieldAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/*GetCustomFieldAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
for the get custom field audit logs with history operation typically these are written to a http.Request
*/
type GetCustomFieldAuditLogsWithHistoryParams struct {

	/*CustomFieldID*/
	CustomFieldID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetCustomFieldAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetCustomFieldAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetCustomFieldAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomFieldID adds the customFieldID to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) WithCustomFieldID(customFieldID strfmt.UUID) *GetCustomFieldAuditLogsWithHistoryParams {
	o.SetCustomFieldID(customFieldID)
	return o
}

// SetCustomFieldID adds the customFieldId to the get custom field audit logs with history params
func (o *GetCustomFieldAuditLogsWithHistoryParams) SetCustomFieldID(customFieldID strfmt.UUID) {
	o.CustomFieldID = customFieldID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomFieldAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customFieldId
	if err := r.SetPathParam("customFieldId", o.CustomFieldID.String()); err != nil {
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
