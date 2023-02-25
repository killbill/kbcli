// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// CreateSubscriptionsWithAddOnsReader is a Reader for the CreateSubscriptionsWithAddOns structure.
type CreateSubscriptionsWithAddOnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionsWithAddOnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateSubscriptionsWithAddOnsCreated()
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

// NewCreateSubscriptionsWithAddOnsCreated creates a CreateSubscriptionsWithAddOnsCreated with default headers values
func NewCreateSubscriptionsWithAddOnsCreated() *CreateSubscriptionsWithAddOnsCreated {
	return &CreateSubscriptionsWithAddOnsCreated{}
}

/*CreateSubscriptionsWithAddOnsCreated handles this case with default header values.

Subscriptions created successfully
*/
type CreateSubscriptionsWithAddOnsCreated struct {
	Payload []*kbmodel.Bundle

	HttpResponse runtime.ClientResponse
}

func (o *CreateSubscriptionsWithAddOnsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/createSubscriptionsWithAddOns][%d] createSubscriptionsWithAddOnsCreated  %+v", 201, o.Payload)
}

func (o *CreateSubscriptionsWithAddOnsCreated) GetPayload() []*kbmodel.Bundle {
	return o.Payload
}

func (o *CreateSubscriptionsWithAddOnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
