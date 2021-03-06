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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPushNotificationCallbacksParams creates a new GetPushNotificationCallbacksParams object
// with the default values initialized.
func NewGetPushNotificationCallbacksParams() *GetPushNotificationCallbacksParams {

	return &GetPushNotificationCallbacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPushNotificationCallbacksParamsWithTimeout creates a new GetPushNotificationCallbacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPushNotificationCallbacksParamsWithTimeout(timeout time.Duration) *GetPushNotificationCallbacksParams {

	return &GetPushNotificationCallbacksParams{

		timeout: timeout,
	}
}

// NewGetPushNotificationCallbacksParamsWithContext creates a new GetPushNotificationCallbacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPushNotificationCallbacksParamsWithContext(ctx context.Context) *GetPushNotificationCallbacksParams {

	return &GetPushNotificationCallbacksParams{

		Context: ctx,
	}
}

// NewGetPushNotificationCallbacksParamsWithHTTPClient creates a new GetPushNotificationCallbacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPushNotificationCallbacksParamsWithHTTPClient(client *http.Client) *GetPushNotificationCallbacksParams {

	return &GetPushNotificationCallbacksParams{
		HTTPClient: client,
	}
}

/*GetPushNotificationCallbacksParams contains all the parameters to send to the API endpoint
for the get push notification callbacks operation typically these are written to a http.Request
*/
type GetPushNotificationCallbacksParams struct {
	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) WithTimeout(timeout time.Duration) *GetPushNotificationCallbacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) WithContext(ctx context.Context) *GetPushNotificationCallbacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) WithHTTPClient(client *http.Client) *GetPushNotificationCallbacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get push notification callbacks params
func (o *GetPushNotificationCallbacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPushNotificationCallbacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
