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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetInvoiceByItemIDReader is a Reader for the GetInvoiceByItemID structure.
type GetInvoiceByItemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceByItemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceByItemIDOK()
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

// NewGetInvoiceByItemIDOK creates a GetInvoiceByItemIDOK with default headers values
func NewGetInvoiceByItemIDOK() *GetInvoiceByItemIDOK {
	return &GetInvoiceByItemIDOK{}
}

/*
GetInvoiceByItemIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceByItemIDOK struct {
	Payload      *kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice by item Id o k response
func (o *GetInvoiceByItemIDOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice by item Id o k response has a 2xx status code
func (o *GetInvoiceByItemIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice by item Id o k response has a 3xx status code
func (o *GetInvoiceByItemIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice by item Id o k response has a 4xx status code
func (o *GetInvoiceByItemIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice by item Id o k response has a 5xx status code
func (o *GetInvoiceByItemIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice by item Id o k response a status code equal to that given
func (o *GetInvoiceByItemIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceByItemIDOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/byItemId/{itemId}][%d] getInvoiceByItemIdOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceByItemIDOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/byItemId/{itemId}][%d] getInvoiceByItemIdOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceByItemIDOK) GetPayload() *kbmodel.Invoice {
	return o.Payload
}

func (o *GetInvoiceByItemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceByItemIDNotFound creates a GetInvoiceByItemIDNotFound with default headers values
func NewGetInvoiceByItemIDNotFound() *GetInvoiceByItemIDNotFound {
	return &GetInvoiceByItemIDNotFound{}
}

/*
GetInvoiceByItemIDNotFound describes a response with status code 404, with default header values.

Invoice not found
*/
type GetInvoiceByItemIDNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice by item Id not found response
func (o *GetInvoiceByItemIDNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get invoice by item Id not found response has a 2xx status code
func (o *GetInvoiceByItemIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice by item Id not found response has a 3xx status code
func (o *GetInvoiceByItemIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice by item Id not found response has a 4xx status code
func (o *GetInvoiceByItemIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice by item Id not found response has a 5xx status code
func (o *GetInvoiceByItemIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice by item Id not found response a status code equal to that given
func (o *GetInvoiceByItemIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoiceByItemIDNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/byItemId/{itemId}][%d] getInvoiceByItemIdNotFound ", 404)
}

func (o *GetInvoiceByItemIDNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/byItemId/{itemId}][%d] getInvoiceByItemIdNotFound ", 404)
}

func (o *GetInvoiceByItemIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
