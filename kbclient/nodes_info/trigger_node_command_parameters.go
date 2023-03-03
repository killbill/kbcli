// Code generated by go-swagger; DO NOT EDIT.

package nodes_info

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

// NewTriggerNodeCommandParams creates a new TriggerNodeCommandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerNodeCommandParams() *TriggerNodeCommandParams {
	return &TriggerNodeCommandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerNodeCommandParamsWithTimeout creates a new TriggerNodeCommandParams object
// with the ability to set a timeout on a request.
func NewTriggerNodeCommandParamsWithTimeout(timeout time.Duration) *TriggerNodeCommandParams {
	return &TriggerNodeCommandParams{
		timeout: timeout,
	}
}

// NewTriggerNodeCommandParamsWithContext creates a new TriggerNodeCommandParams object
// with the ability to set a context for a request.
func NewTriggerNodeCommandParamsWithContext(ctx context.Context) *TriggerNodeCommandParams {
	return &TriggerNodeCommandParams{
		Context: ctx,
	}
}

// NewTriggerNodeCommandParamsWithHTTPClient creates a new TriggerNodeCommandParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerNodeCommandParamsWithHTTPClient(client *http.Client) *TriggerNodeCommandParams {
	return &TriggerNodeCommandParams{
		HTTPClient: client,
	}
}

/*
TriggerNodeCommandParams contains all the parameters to send to the API endpoint

	for the trigger node command operation.

	Typically these are written to a http.Request.
*/
type TriggerNodeCommandParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.NodeCommand

	// LocalNodeOnly.
	LocalNodeOnly *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the trigger node command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerNodeCommandParams) WithDefaults() *TriggerNodeCommandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger node command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerNodeCommandParams) SetDefaults() {
	var (
		localNodeOnlyDefault = bool(false)
	)

	val := TriggerNodeCommandParams{
		LocalNodeOnly: &localNodeOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the trigger node command params
func (o *TriggerNodeCommandParams) WithTimeout(timeout time.Duration) *TriggerNodeCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger node command params
func (o *TriggerNodeCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger node command params
func (o *TriggerNodeCommandParams) WithContext(ctx context.Context) *TriggerNodeCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger node command params
func (o *TriggerNodeCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger node command params
func (o *TriggerNodeCommandParams) WithHTTPClient(client *http.Client) *TriggerNodeCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger node command params
func (o *TriggerNodeCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the trigger node command params
func (o *TriggerNodeCommandParams) WithXKillbillComment(xKillbillComment *string) *TriggerNodeCommandParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the trigger node command params
func (o *TriggerNodeCommandParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the trigger node command params
func (o *TriggerNodeCommandParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *TriggerNodeCommandParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the trigger node command params
func (o *TriggerNodeCommandParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the trigger node command params
func (o *TriggerNodeCommandParams) WithXKillbillReason(xKillbillReason *string) *TriggerNodeCommandParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the trigger node command params
func (o *TriggerNodeCommandParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the trigger node command params
func (o *TriggerNodeCommandParams) WithBody(body *kbmodel.NodeCommand) *TriggerNodeCommandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trigger node command params
func (o *TriggerNodeCommandParams) SetBody(body *kbmodel.NodeCommand) {
	o.Body = body
}

// WithLocalNodeOnly adds the localNodeOnly to the trigger node command params
func (o *TriggerNodeCommandParams) WithLocalNodeOnly(localNodeOnly *bool) *TriggerNodeCommandParams {
	o.SetLocalNodeOnly(localNodeOnly)
	return o
}

// SetLocalNodeOnly adds the localNodeOnly to the trigger node command params
func (o *TriggerNodeCommandParams) SetLocalNodeOnly(localNodeOnly *bool) {
	o.LocalNodeOnly = localNodeOnly
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerNodeCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LocalNodeOnly != nil {

		// query param localNodeOnly
		var qrLocalNodeOnly bool

		if o.LocalNodeOnly != nil {
			qrLocalNodeOnly = *o.LocalNodeOnly
		}
		qLocalNodeOnly := swag.FormatBool(qrLocalNodeOnly)
		if qLocalNodeOnly != "" {

			if err := r.SetQueryParam("localNodeOnly", qLocalNodeOnly); err != nil {
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
