// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetBlockingStatesReader is a Reader for the GetBlockingStates structure.
type GetBlockingStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockingStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlockingStatesOK()
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

// NewGetBlockingStatesOK creates a GetBlockingStatesOK with default headers values
func NewGetBlockingStatesOK() *GetBlockingStatesOK {
	return &GetBlockingStatesOK{}
}

/*GetBlockingStatesOK handles this case with default header values.

successful operation
*/
type GetBlockingStatesOK struct {
	Payload []*kbmodel.BlockingState

	HttpResponse runtime.ClientResponse
}

func (o *GetBlockingStatesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/block][%d] getBlockingStatesOK  %+v", 200, o.Payload)
}

func (o *GetBlockingStatesOK) GetPayload() []*kbmodel.BlockingState {
	return o.Payload
}

func (o *GetBlockingStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockingStatesBadRequest creates a GetBlockingStatesBadRequest with default headers values
func NewGetBlockingStatesBadRequest() *GetBlockingStatesBadRequest {
	return &GetBlockingStatesBadRequest{}
}

/*GetBlockingStatesBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetBlockingStatesBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetBlockingStatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/block][%d] getBlockingStatesBadRequest ", 400)
}

func (o *GetBlockingStatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
