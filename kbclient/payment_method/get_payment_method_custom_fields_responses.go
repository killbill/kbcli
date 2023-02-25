// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// GetPaymentMethodCustomFieldsReader is a Reader for the GetPaymentMethodCustomFields structure.
type GetPaymentMethodCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodCustomFieldsOK()
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

// NewGetPaymentMethodCustomFieldsOK creates a GetPaymentMethodCustomFieldsOK with default headers values
func NewGetPaymentMethodCustomFieldsOK() *GetPaymentMethodCustomFieldsOK {
	return &GetPaymentMethodCustomFieldsOK{}
}

/*GetPaymentMethodCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetPaymentMethodCustomFieldsOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] getPaymentMethodCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetPaymentMethodCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodCustomFieldsBadRequest creates a GetPaymentMethodCustomFieldsBadRequest with default headers values
func NewGetPaymentMethodCustomFieldsBadRequest() *GetPaymentMethodCustomFieldsBadRequest {
	return &GetPaymentMethodCustomFieldsBadRequest{}
}

/*GetPaymentMethodCustomFieldsBadRequest handles this case with default header values.

Invalid payment method id supplied
*/
type GetPaymentMethodCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] getPaymentMethodCustomFieldsBadRequest ", 400)
}

func (o *GetPaymentMethodCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
