// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewInvalidateUserParams creates a new InvalidateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvalidateUserParams() *InvalidateUserParams {
	return &InvalidateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidateUserParamsWithTimeout creates a new InvalidateUserParams object
// with the ability to set a timeout on a request.
func NewInvalidateUserParamsWithTimeout(timeout time.Duration) *InvalidateUserParams {
	return &InvalidateUserParams{
		timeout: timeout,
	}
}

// NewInvalidateUserParamsWithContext creates a new InvalidateUserParams object
// with the ability to set a context for a request.
func NewInvalidateUserParamsWithContext(ctx context.Context) *InvalidateUserParams {
	return &InvalidateUserParams{
		Context: ctx,
	}
}

// NewInvalidateUserParamsWithHTTPClient creates a new InvalidateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvalidateUserParamsWithHTTPClient(client *http.Client) *InvalidateUserParams {
	return &InvalidateUserParams{
		HTTPClient: client,
	}
}

/*
InvalidateUserParams contains all the parameters to send to the API endpoint

	for the invalidate user operation.

	Typically these are written to a http.Request.
*/
type InvalidateUserParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Username.
	Username string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the invalidate user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvalidateUserParams) WithDefaults() *InvalidateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invalidate user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvalidateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invalidate user params
func (o *InvalidateUserParams) WithTimeout(timeout time.Duration) *InvalidateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidate user params
func (o *InvalidateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidate user params
func (o *InvalidateUserParams) WithContext(ctx context.Context) *InvalidateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidate user params
func (o *InvalidateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidate user params
func (o *InvalidateUserParams) WithHTTPClient(client *http.Client) *InvalidateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidate user params
func (o *InvalidateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the invalidate user params
func (o *InvalidateUserParams) WithXKillbillComment(xKillbillComment *string) *InvalidateUserParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the invalidate user params
func (o *InvalidateUserParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the invalidate user params
func (o *InvalidateUserParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *InvalidateUserParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the invalidate user params
func (o *InvalidateUserParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the invalidate user params
func (o *InvalidateUserParams) WithXKillbillReason(xKillbillReason *string) *InvalidateUserParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the invalidate user params
func (o *InvalidateUserParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithUsername adds the username to the invalidate user params
func (o *InvalidateUserParams) WithUsername(username string) *InvalidateUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the invalidate user params
func (o *InvalidateUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
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
