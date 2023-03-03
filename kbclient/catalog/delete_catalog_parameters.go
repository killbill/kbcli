// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewDeleteCatalogParams creates a new DeleteCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCatalogParams() *DeleteCatalogParams {
	return &DeleteCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogParamsWithTimeout creates a new DeleteCatalogParams object
// with the ability to set a timeout on a request.
func NewDeleteCatalogParamsWithTimeout(timeout time.Duration) *DeleteCatalogParams {
	return &DeleteCatalogParams{
		timeout: timeout,
	}
}

// NewDeleteCatalogParamsWithContext creates a new DeleteCatalogParams object
// with the ability to set a context for a request.
func NewDeleteCatalogParamsWithContext(ctx context.Context) *DeleteCatalogParams {
	return &DeleteCatalogParams{
		Context: ctx,
	}
}

// NewDeleteCatalogParamsWithHTTPClient creates a new DeleteCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCatalogParamsWithHTTPClient(client *http.Client) *DeleteCatalogParams {
	return &DeleteCatalogParams{
		HTTPClient: client,
	}
}

/*
DeleteCatalogParams contains all the parameters to send to the API endpoint

	for the delete catalog operation.

	Typically these are written to a http.Request.
*/
type DeleteCatalogParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the delete catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogParams) WithDefaults() *DeleteCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete catalog params
func (o *DeleteCatalogParams) WithTimeout(timeout time.Duration) *DeleteCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog params
func (o *DeleteCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog params
func (o *DeleteCatalogParams) WithContext(ctx context.Context) *DeleteCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog params
func (o *DeleteCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalog params
func (o *DeleteCatalogParams) WithHTTPClient(client *http.Client) *DeleteCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalog params
func (o *DeleteCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete catalog params
func (o *DeleteCatalogParams) WithXKillbillComment(xKillbillComment *string) *DeleteCatalogParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete catalog params
func (o *DeleteCatalogParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete catalog params
func (o *DeleteCatalogParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteCatalogParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete catalog params
func (o *DeleteCatalogParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete catalog params
func (o *DeleteCatalogParams) WithXKillbillReason(xKillbillReason *string) *DeleteCatalogParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete catalog params
func (o *DeleteCatalogParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
