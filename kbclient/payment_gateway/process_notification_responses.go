// Code generated by go-swagger; DO NOT EDIT.

package payment_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ProcessNotificationReader is a Reader for the ProcessNotification structure.
type ProcessNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProcessNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProcessNotificationOK()
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

// NewProcessNotificationOK creates a ProcessNotificationOK with default headers values
func NewProcessNotificationOK() *ProcessNotificationOK {
	return &ProcessNotificationOK{}
}

/*ProcessNotificationOK handles this case with default header values.

Successful
*/
type ProcessNotificationOK struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessNotificationOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentGateways/notification/{pluginName}][%d] processNotificationOK ", 200)
}

func (o *ProcessNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
