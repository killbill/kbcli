// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewUploadPluginConfigurationParams creates a new UploadPluginConfigurationParams object
// with the default values initialized.
func NewUploadPluginConfigurationParams() *UploadPluginConfigurationParams {
	var ()
	return &UploadPluginConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadPluginConfigurationParamsWithTimeout creates a new UploadPluginConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadPluginConfigurationParamsWithTimeout(timeout time.Duration) *UploadPluginConfigurationParams {
	var ()
	return &UploadPluginConfigurationParams{

		timeout: timeout,
	}
}

// NewUploadPluginConfigurationParamsWithContext creates a new UploadPluginConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadPluginConfigurationParamsWithContext(ctx context.Context) *UploadPluginConfigurationParams {
	var ()
	return &UploadPluginConfigurationParams{

		Context: ctx,
	}
}

// NewUploadPluginConfigurationParamsWithHTTPClient creates a new UploadPluginConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadPluginConfigurationParamsWithHTTPClient(client *http.Client) *UploadPluginConfigurationParams {
	var ()
	return &UploadPluginConfigurationParams{
		HTTPClient: client,
	}
}

/*UploadPluginConfigurationParams contains all the parameters to send to the API endpoint
for the upload plugin configuration operation typically these are written to a http.Request
*/
type UploadPluginConfigurationParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body string
	/*PluginName*/
	PluginName string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithTimeout(timeout time.Duration) *UploadPluginConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithContext(ctx context.Context) *UploadPluginConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithHTTPClient(client *http.Client) *UploadPluginConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithXKillbillComment(xKillbillComment *string) *UploadPluginConfigurationParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadPluginConfigurationParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithXKillbillReason(xKillbillReason *string) *UploadPluginConfigurationParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithBody(body string) *UploadPluginConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetBody(body string) {
	o.Body = body
}

// WithPluginName adds the pluginName to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) WithPluginName(pluginName string) *UploadPluginConfigurationParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the upload plugin configuration params
func (o *UploadPluginConfigurationParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *UploadPluginConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
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
