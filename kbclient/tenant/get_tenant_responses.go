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

// GetTenantReader is a Reader for the GetTenant structure.
type GetTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTenantOK()
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

// NewGetTenantOK creates a GetTenantOK with default headers values
func NewGetTenantOK() *GetTenantOK {
	return &GetTenantOK{}
}

/*
GetTenantOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTenantOK struct {
	Payload      *kbmodel.Tenant
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get tenant o k response
func (o *GetTenantOK) Code() int {
	return 200
}

// IsSuccess returns true when this get tenant o k response has a 2xx status code
func (o *GetTenantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tenant o k response has a 3xx status code
func (o *GetTenantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant o k response has a 4xx status code
func (o *GetTenantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenant o k response has a 5xx status code
func (o *GetTenantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant o k response a status code equal to that given
func (o *GetTenantOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTenantOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantOK  %+v", 200, o.Payload)
}

func (o *GetTenantOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantOK  %+v", 200, o.Payload)
}

func (o *GetTenantOK) GetPayload() *kbmodel.Tenant {
	return o.Payload
}

func (o *GetTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantBadRequest creates a GetTenantBadRequest with default headers values
func NewGetTenantBadRequest() *GetTenantBadRequest {
	return &GetTenantBadRequest{}
}

/*
GetTenantBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type GetTenantBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get tenant bad request response
func (o *GetTenantBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get tenant bad request response has a 2xx status code
func (o *GetTenantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tenant bad request response has a 3xx status code
func (o *GetTenantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant bad request response has a 4xx status code
func (o *GetTenantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tenant bad request response has a 5xx status code
func (o *GetTenantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant bad request response a status code equal to that given
func (o *GetTenantBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTenantBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantBadRequest ", 400)
}

func (o *GetTenantBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantBadRequest ", 400)
}

func (o *GetTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTenantNotFound creates a GetTenantNotFound with default headers values
func NewGetTenantNotFound() *GetTenantNotFound {
	return &GetTenantNotFound{}
}

/*
GetTenantNotFound describes a response with status code 404, with default header values.

Tenant not found
*/
type GetTenantNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get tenant not found response
func (o *GetTenantNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get tenant not found response has a 2xx status code
func (o *GetTenantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tenant not found response has a 3xx status code
func (o *GetTenantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant not found response has a 4xx status code
func (o *GetTenantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tenant not found response has a 5xx status code
func (o *GetTenantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant not found response a status code equal to that given
func (o *GetTenantNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTenantNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantNotFound ", 404)
}

func (o *GetTenantNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/{tenantId}][%d] getTenantNotFound ", 404)
}

func (o *GetTenantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
