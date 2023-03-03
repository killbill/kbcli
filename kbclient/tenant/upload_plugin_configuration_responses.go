// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbcommon"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// UploadPluginConfigurationReader is a Reader for the UploadPluginConfiguration structure.
type UploadPluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadPluginConfigurationCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewUploadPluginConfigurationCreated creates a UploadPluginConfigurationCreated with default headers values
func NewUploadPluginConfigurationCreated() *UploadPluginConfigurationCreated {
	return &UploadPluginConfigurationCreated{}
}

/*
UploadPluginConfigurationCreated describes a response with status code 201, with default header values.

Plugin configuration uploaded successfully
*/
type UploadPluginConfigurationCreated struct {
	Payload      *kbmodel.TenantKeyValue
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload plugin configuration created response
func (o *UploadPluginConfigurationCreated) Code() int {
	return 201
}

// IsSuccess returns true when this upload plugin configuration created response has a 2xx status code
func (o *UploadPluginConfigurationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload plugin configuration created response has a 3xx status code
func (o *UploadPluginConfigurationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload plugin configuration created response has a 4xx status code
func (o *UploadPluginConfigurationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload plugin configuration created response has a 5xx status code
func (o *UploadPluginConfigurationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload plugin configuration created response a status code equal to that given
func (o *UploadPluginConfigurationCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UploadPluginConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationCreated  %+v", 201, o.Payload)
}

func (o *UploadPluginConfigurationCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationCreated  %+v", 201, o.Payload)
}

func (o *UploadPluginConfigurationCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *UploadPluginConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPluginConfigurationBadRequest creates a UploadPluginConfigurationBadRequest with default headers values
func NewUploadPluginConfigurationBadRequest() *UploadPluginConfigurationBadRequest {
	return &UploadPluginConfigurationBadRequest{}
}

/*
UploadPluginConfigurationBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type UploadPluginConfigurationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload plugin configuration bad request response
func (o *UploadPluginConfigurationBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this upload plugin configuration bad request response has a 2xx status code
func (o *UploadPluginConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload plugin configuration bad request response has a 3xx status code
func (o *UploadPluginConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload plugin configuration bad request response has a 4xx status code
func (o *UploadPluginConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload plugin configuration bad request response has a 5xx status code
func (o *UploadPluginConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload plugin configuration bad request response a status code equal to that given
func (o *UploadPluginConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadPluginConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationBadRequest ", 400)
}

func (o *UploadPluginConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationBadRequest ", 400)
}

func (o *UploadPluginConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
