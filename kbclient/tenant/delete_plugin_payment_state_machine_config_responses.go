// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePluginPaymentStateMachineConfigReader is a Reader for the DeletePluginPaymentStateMachineConfig structure.
type DeletePluginPaymentStateMachineConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginPaymentStateMachineConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePluginPaymentStateMachineConfigNoContent()
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

// NewDeletePluginPaymentStateMachineConfigNoContent creates a DeletePluginPaymentStateMachineConfigNoContent with default headers values
func NewDeletePluginPaymentStateMachineConfigNoContent() *DeletePluginPaymentStateMachineConfigNoContent {
	return &DeletePluginPaymentStateMachineConfigNoContent{}
}

/*DeletePluginPaymentStateMachineConfigNoContent handles this case with default header values.

Successful operation
*/
type DeletePluginPaymentStateMachineConfigNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeletePluginPaymentStateMachineConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] deletePluginPaymentStateMachineConfigNoContent ", 204)
}

func (o *DeletePluginPaymentStateMachineConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePluginPaymentStateMachineConfigBadRequest creates a DeletePluginPaymentStateMachineConfigBadRequest with default headers values
func NewDeletePluginPaymentStateMachineConfigBadRequest() *DeletePluginPaymentStateMachineConfigBadRequest {
	return &DeletePluginPaymentStateMachineConfigBadRequest{}
}

/*DeletePluginPaymentStateMachineConfigBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type DeletePluginPaymentStateMachineConfigBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeletePluginPaymentStateMachineConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] deletePluginPaymentStateMachineConfigBadRequest ", 400)
}

func (o *DeletePluginPaymentStateMachineConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
