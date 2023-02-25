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
)

// RebalanceExistingCBAOnAccountReader is a Reader for the RebalanceExistingCBAOnAccount structure.
type RebalanceExistingCBAOnAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebalanceExistingCBAOnAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRebalanceExistingCBAOnAccountNoContent()
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

// NewRebalanceExistingCBAOnAccountNoContent creates a RebalanceExistingCBAOnAccountNoContent with default headers values
func NewRebalanceExistingCBAOnAccountNoContent() *RebalanceExistingCBAOnAccountNoContent {
	return &RebalanceExistingCBAOnAccountNoContent{}
}

/*RebalanceExistingCBAOnAccountNoContent handles this case with default header values.

Successful operation
*/
type RebalanceExistingCBAOnAccountNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *RebalanceExistingCBAOnAccountNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/cbaRebalancing][%d] rebalanceExistingCBAOnAccountNoContent ", 204)
}

func (o *RebalanceExistingCBAOnAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebalanceExistingCBAOnAccountBadRequest creates a RebalanceExistingCBAOnAccountBadRequest with default headers values
func NewRebalanceExistingCBAOnAccountBadRequest() *RebalanceExistingCBAOnAccountBadRequest {
	return &RebalanceExistingCBAOnAccountBadRequest{}
}

/*RebalanceExistingCBAOnAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type RebalanceExistingCBAOnAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *RebalanceExistingCBAOnAccountBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/cbaRebalancing][%d] rebalanceExistingCBAOnAccountBadRequest ", 400)
}

func (o *RebalanceExistingCBAOnAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
