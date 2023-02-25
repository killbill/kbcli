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

// GetAccountBundlesPaginatedReader is a Reader for the GetAccountBundlesPaginated structure.
type GetAccountBundlesPaginatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountBundlesPaginatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountBundlesPaginatedOK()
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

// NewGetAccountBundlesPaginatedOK creates a GetAccountBundlesPaginatedOK with default headers values
func NewGetAccountBundlesPaginatedOK() *GetAccountBundlesPaginatedOK {
	return &GetAccountBundlesPaginatedOK{}
}

/*GetAccountBundlesPaginatedOK handles this case with default header values.

successful operation
*/
type GetAccountBundlesPaginatedOK struct {
	Payload []*kbmodel.Bundle

	HttpResponse runtime.ClientResponse
}

func (o *GetAccountBundlesPaginatedOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/bundles/pagination][%d] getAccountBundlesPaginatedOK  %+v", 200, o.Payload)
}

func (o *GetAccountBundlesPaginatedOK) GetPayload() []*kbmodel.Bundle {
	return o.Payload
}

func (o *GetAccountBundlesPaginatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBundlesPaginatedBadRequest creates a GetAccountBundlesPaginatedBadRequest with default headers values
func NewGetAccountBundlesPaginatedBadRequest() *GetAccountBundlesPaginatedBadRequest {
	return &GetAccountBundlesPaginatedBadRequest{}
}

/*GetAccountBundlesPaginatedBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetAccountBundlesPaginatedBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountBundlesPaginatedBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/bundles/pagination][%d] getAccountBundlesPaginatedBadRequest ", 400)
}

func (o *GetAccountBundlesPaginatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountBundlesPaginatedNotFound creates a GetAccountBundlesPaginatedNotFound with default headers values
func NewGetAccountBundlesPaginatedNotFound() *GetAccountBundlesPaginatedNotFound {
	return &GetAccountBundlesPaginatedNotFound{}
}

/*GetAccountBundlesPaginatedNotFound handles this case with default header values.

Account not found
*/
type GetAccountBundlesPaginatedNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountBundlesPaginatedNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/bundles/pagination][%d] getAccountBundlesPaginatedNotFound ", 404)
}

func (o *GetAccountBundlesPaginatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
