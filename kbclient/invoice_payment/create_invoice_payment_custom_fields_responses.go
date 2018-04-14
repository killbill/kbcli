// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// CreateInvoicePaymentCustomFieldsReader is a Reader for the CreateInvoicePaymentCustomFields structure.
type CreateInvoicePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoicePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateInvoicePaymentCustomFieldsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateInvoicePaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInvoicePaymentCustomFieldsCreated creates a CreateInvoicePaymentCustomFieldsCreated with default headers values
func NewCreateInvoicePaymentCustomFieldsCreated() *CreateInvoicePaymentCustomFieldsCreated {
	return &CreateInvoicePaymentCustomFieldsCreated{}
}

/*CreateInvoicePaymentCustomFieldsCreated handles this case with default header values.

Custom field created successfully
*/
type CreateInvoicePaymentCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField
}

func (o *CreateInvoicePaymentCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoicePaymentCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoicePaymentCustomFieldsBadRequest creates a CreateInvoicePaymentCustomFieldsBadRequest with default headers values
func NewCreateInvoicePaymentCustomFieldsBadRequest() *CreateInvoicePaymentCustomFieldsBadRequest {
	return &CreateInvoicePaymentCustomFieldsBadRequest{}
}

/*CreateInvoicePaymentCustomFieldsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type CreateInvoicePaymentCustomFieldsBadRequest struct {
}

func (o *CreateInvoicePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoicePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}