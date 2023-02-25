// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// GetInvoiceAsHTMLReader is a Reader for the GetInvoiceAsHTML structure.
type GetInvoiceAsHTMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceAsHTMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceAsHTMLOK()
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

// NewGetInvoiceAsHTMLOK creates a GetInvoiceAsHTMLOK with default headers values
func NewGetInvoiceAsHTMLOK() *GetInvoiceAsHTMLOK {
	return &GetInvoiceAsHTMLOK{}
}

/*GetInvoiceAsHTMLOK handles this case with default header values.

successful operation
*/
type GetInvoiceAsHTMLOK struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceAsHTMLOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/html][%d] getInvoiceAsHtmlOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceAsHTMLOK) GetPayload() string {
	return o.Payload
}

func (o *GetInvoiceAsHTMLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceAsHTMLNotFound creates a GetInvoiceAsHTMLNotFound with default headers values
func NewGetInvoiceAsHTMLNotFound() *GetInvoiceAsHTMLNotFound {
	return &GetInvoiceAsHTMLNotFound{}
}

/*GetInvoiceAsHTMLNotFound handles this case with default header values.

Invoice not found
*/
type GetInvoiceAsHTMLNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceAsHTMLNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/html][%d] getInvoiceAsHtmlNotFound ", 404)
}

func (o *GetInvoiceAsHTMLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
