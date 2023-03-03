// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbcommon"
)

// GetInvoiceMPTemplateReader is a Reader for the GetInvoiceMPTemplate structure.
type GetInvoiceMPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceMPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceMPTemplateOK()
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

// NewGetInvoiceMPTemplateOK creates a GetInvoiceMPTemplateOK with default headers values
func NewGetInvoiceMPTemplateOK() *GetInvoiceMPTemplateOK {
	return &GetInvoiceMPTemplateOK{}
}

/*
GetInvoiceMPTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceMPTemplateOK struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice m p template o k response
func (o *GetInvoiceMPTemplateOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice m p template o k response has a 2xx status code
func (o *GetInvoiceMPTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice m p template o k response has a 3xx status code
func (o *GetInvoiceMPTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice m p template o k response has a 4xx status code
func (o *GetInvoiceMPTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice m p template o k response has a 5xx status code
func (o *GetInvoiceMPTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice m p template o k response a status code equal to that given
func (o *GetInvoiceMPTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceMPTemplateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceMPTemplateOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceMPTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *GetInvoiceMPTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceMPTemplateNotFound creates a GetInvoiceMPTemplateNotFound with default headers values
func NewGetInvoiceMPTemplateNotFound() *GetInvoiceMPTemplateNotFound {
	return &GetInvoiceMPTemplateNotFound{}
}

/*
GetInvoiceMPTemplateNotFound describes a response with status code 404, with default header values.

Template not found
*/
type GetInvoiceMPTemplateNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice m p template not found response
func (o *GetInvoiceMPTemplateNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get invoice m p template not found response has a 2xx status code
func (o *GetInvoiceMPTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice m p template not found response has a 3xx status code
func (o *GetInvoiceMPTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice m p template not found response has a 4xx status code
func (o *GetInvoiceMPTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice m p template not found response has a 5xx status code
func (o *GetInvoiceMPTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice m p template not found response a status code equal to that given
func (o *GetInvoiceMPTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoiceMPTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateNotFound ", 404)
}

func (o *GetInvoiceMPTemplateNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateNotFound ", 404)
}

func (o *GetInvoiceMPTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
