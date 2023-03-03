// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewDeletePushNotificationCallbacksParams creates a new DeletePushNotificationCallbacksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePushNotificationCallbacksParams() *DeletePushNotificationCallbacksParams {
	return &DeletePushNotificationCallbacksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePushNotificationCallbacksParamsWithTimeout creates a new DeletePushNotificationCallbacksParams object
// with the ability to set a timeout on a request.
func NewDeletePushNotificationCallbacksParamsWithTimeout(timeout time.Duration) *DeletePushNotificationCallbacksParams {
	return &DeletePushNotificationCallbacksParams{
		timeout: timeout,
	}
}

// NewDeletePushNotificationCallbacksParamsWithContext creates a new DeletePushNotificationCallbacksParams object
// with the ability to set a context for a request.
func NewDeletePushNotificationCallbacksParamsWithContext(ctx context.Context) *DeletePushNotificationCallbacksParams {
	return &DeletePushNotificationCallbacksParams{
		Context: ctx,
	}
}

// NewDeletePushNotificationCallbacksParamsWithHTTPClient creates a new DeletePushNotificationCallbacksParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePushNotificationCallbacksParamsWithHTTPClient(client *http.Client) *DeletePushNotificationCallbacksParams {
	return &DeletePushNotificationCallbacksParams{
		HTTPClient: client,
	}
}

/*
DeletePushNotificationCallbacksParams contains all the parameters to send to the API endpoint

	for the delete push notification callbacks operation.

	Typically these are written to a http.Request.
*/
type DeletePushNotificationCallbacksParams struct {

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

// WithDefaults hydrates default values in the delete push notification callbacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePushNotificationCallbacksParams) WithDefaults() *DeletePushNotificationCallbacksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete push notification callbacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePushNotificationCallbacksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithTimeout(timeout time.Duration) *DeletePushNotificationCallbacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithContext(ctx context.Context) *DeletePushNotificationCallbacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithHTTPClient(client *http.Client) *DeletePushNotificationCallbacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithXKillbillComment(xKillbillComment *string) *DeletePushNotificationCallbacksParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePushNotificationCallbacksParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) WithXKillbillReason(xKillbillReason *string) *DeletePushNotificationCallbacksParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete push notification callbacks params
func (o *DeletePushNotificationCallbacksParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePushNotificationCallbacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
