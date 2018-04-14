// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBundleParams creates a new GetBundleParams object
// with the default values initialized.
func NewGetBundleParams() *GetBundleParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBundleParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleParamsWithTimeout creates a new GetBundleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBundleParamsWithTimeout(timeout time.Duration) *GetBundleParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBundleParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetBundleParamsWithContext creates a new GetBundleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBundleParamsWithContext(ctx context.Context) *GetBundleParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBundleParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetBundleParamsWithHTTPClient creates a new GetBundleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBundleParamsWithHTTPClient(client *http.Client) *GetBundleParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBundleParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetBundleParams contains all the parameters to send to the API endpoint
for the get bundle operation typically these are written to a http.Request
*/
type GetBundleParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*BundleID*/
	BundleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bundle params
func (o *GetBundleParams) WithTimeout(timeout time.Duration) *GetBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle params
func (o *GetBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle params
func (o *GetBundleParams) WithContext(ctx context.Context) *GetBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle params
func (o *GetBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle params
func (o *GetBundleParams) WithHTTPClient(client *http.Client) *GetBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle params
func (o *GetBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get bundle params
func (o *GetBundleParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetBundleParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get bundle params
func (o *GetBundleParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get bundle params
func (o *GetBundleParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetBundleParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get bundle params
func (o *GetBundleParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get bundle params
func (o *GetBundleParams) WithAudit(audit *string) *GetBundleParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundle params
func (o *GetBundleParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithBundleID adds the bundleID to the get bundle params
func (o *GetBundleParams) WithBundleID(bundleID strfmt.UUID) *GetBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get bundle params
func (o *GetBundleParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}