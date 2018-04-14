// Code generated by go-swagger; DO NOT EDIT.

package nodes_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TriggerNodeCommandReader is a Reader for the TriggerNodeCommand structure.
type TriggerNodeCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerNodeCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewTriggerNodeCommandAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggerNodeCommandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggerNodeCommandAccepted creates a TriggerNodeCommandAccepted with default headers values
func NewTriggerNodeCommandAccepted() *TriggerNodeCommandAccepted {
	return &TriggerNodeCommandAccepted{}
}

/*TriggerNodeCommandAccepted handles this case with default header values.

Successful operation
*/
type TriggerNodeCommandAccepted struct {
}

func (o *TriggerNodeCommandAccepted) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/nodesInfo][%d] triggerNodeCommandAccepted ", 202)
}

func (o *TriggerNodeCommandAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerNodeCommandBadRequest creates a TriggerNodeCommandBadRequest with default headers values
func NewTriggerNodeCommandBadRequest() *TriggerNodeCommandBadRequest {
	return &TriggerNodeCommandBadRequest{}
}

/*TriggerNodeCommandBadRequest handles this case with default header values.

Invalid node command supplied
*/
type TriggerNodeCommandBadRequest struct {
}

func (o *TriggerNodeCommandBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/nodesInfo][%d] triggerNodeCommandBadRequest ", 400)
}

func (o *TriggerNodeCommandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}