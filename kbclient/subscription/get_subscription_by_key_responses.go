// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// GetSubscriptionByKeyReader is a Reader for the GetSubscriptionByKey structure.
type GetSubscriptionByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionByKeyOK()
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

// NewGetSubscriptionByKeyOK creates a GetSubscriptionByKeyOK with default headers values
func NewGetSubscriptionByKeyOK() *GetSubscriptionByKeyOK {
	return &GetSubscriptionByKeyOK{}
}

/*
GetSubscriptionByKeyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSubscriptionByKeyOK struct {
	Payload      *kbmodel.Subscription
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription by key o k response
func (o *GetSubscriptionByKeyOK) Code() int {
	return 200
}

// IsSuccess returns true when this get subscription by key o k response has a 2xx status code
func (o *GetSubscriptionByKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subscription by key o k response has a 3xx status code
func (o *GetSubscriptionByKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription by key o k response has a 4xx status code
func (o *GetSubscriptionByKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription by key o k response has a 5xx status code
func (o *GetSubscriptionByKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription by key o k response a status code equal to that given
func (o *GetSubscriptionByKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSubscriptionByKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionByKeyOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionByKeyOK) GetPayload() *kbmodel.Subscription {
	return o.Payload
}

func (o *GetSubscriptionByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionByKeyNotFound creates a GetSubscriptionByKeyNotFound with default headers values
func NewGetSubscriptionByKeyNotFound() *GetSubscriptionByKeyNotFound {
	return &GetSubscriptionByKeyNotFound{}
}

/*
GetSubscriptionByKeyNotFound describes a response with status code 404, with default header values.

Subscription not found
*/
type GetSubscriptionByKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription by key not found response
func (o *GetSubscriptionByKeyNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get subscription by key not found response has a 2xx status code
func (o *GetSubscriptionByKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription by key not found response has a 3xx status code
func (o *GetSubscriptionByKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription by key not found response has a 4xx status code
func (o *GetSubscriptionByKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription by key not found response has a 5xx status code
func (o *GetSubscriptionByKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription by key not found response a status code equal to that given
func (o *GetSubscriptionByKeyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSubscriptionByKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyNotFound ", 404)
}

func (o *GetSubscriptionByKeyNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyNotFound ", 404)
}

func (o *GetSubscriptionByKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
