// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// CreateInvoiceItemCustomFieldsReader is a Reader for the CreateInvoiceItemCustomFields structure.
type CreateInvoiceItemCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceItemCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateInvoiceItemCustomFieldsCreated()
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

// NewCreateInvoiceItemCustomFieldsCreated creates a CreateInvoiceItemCustomFieldsCreated with default headers values
func NewCreateInvoiceItemCustomFieldsCreated() *CreateInvoiceItemCustomFieldsCreated {
	return &CreateInvoiceItemCustomFieldsCreated{}
}

/*CreateInvoiceItemCustomFieldsCreated handles this case with default header values.

Custom field created successfully
*/
type CreateInvoiceItemCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *CreateInvoiceItemCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceItemCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateInvoiceItemCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoiceItemCustomFieldsBadRequest creates a CreateInvoiceItemCustomFieldsBadRequest with default headers values
func NewCreateInvoiceItemCustomFieldsBadRequest() *CreateInvoiceItemCustomFieldsBadRequest {
	return &CreateInvoiceItemCustomFieldsBadRequest{}
}

/*CreateInvoiceItemCustomFieldsBadRequest handles this case with default header values.

Invalid invoice item id supplied
*/
type CreateInvoiceItemCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateInvoiceItemCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoiceItemCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
