// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyPaymentMethodCustomFieldsReader is a Reader for the ModifyPaymentMethodCustomFields structure.
type ModifyPaymentMethodCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyPaymentMethodCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyPaymentMethodCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyPaymentMethodCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyPaymentMethodCustomFieldsNoContent creates a ModifyPaymentMethodCustomFieldsNoContent with default headers values
func NewModifyPaymentMethodCustomFieldsNoContent() *ModifyPaymentMethodCustomFieldsNoContent {
	return &ModifyPaymentMethodCustomFieldsNoContent{}
}

/*ModifyPaymentMethodCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyPaymentMethodCustomFieldsNoContent struct {
}

func (o *ModifyPaymentMethodCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] modifyPaymentMethodCustomFieldsNoContent ", 204)
}

func (o *ModifyPaymentMethodCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPaymentMethodCustomFieldsBadRequest creates a ModifyPaymentMethodCustomFieldsBadRequest with default headers values
func NewModifyPaymentMethodCustomFieldsBadRequest() *ModifyPaymentMethodCustomFieldsBadRequest {
	return &ModifyPaymentMethodCustomFieldsBadRequest{}
}

/*ModifyPaymentMethodCustomFieldsBadRequest handles this case with default header values.

Invalid payment method id supplied
*/
type ModifyPaymentMethodCustomFieldsBadRequest struct {
}

func (o *ModifyPaymentMethodCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] modifyPaymentMethodCustomFieldsBadRequest ", 400)
}

func (o *ModifyPaymentMethodCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}