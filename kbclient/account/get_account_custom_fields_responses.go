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

// GetAccountCustomFieldsReader is a Reader for the GetAccountCustomFields structure.
type GetAccountCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountCustomFieldsOK()
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

// NewGetAccountCustomFieldsOK creates a GetAccountCustomFieldsOK with default headers values
func NewGetAccountCustomFieldsOK() *GetAccountCustomFieldsOK {
	return &GetAccountCustomFieldsOK{}
}

/*GetAccountCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetAccountCustomFieldsOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *GetAccountCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/customFields][%d] getAccountCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetAccountCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetAccountCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountCustomFieldsBadRequest creates a GetAccountCustomFieldsBadRequest with default headers values
func NewGetAccountCustomFieldsBadRequest() *GetAccountCustomFieldsBadRequest {
	return &GetAccountCustomFieldsBadRequest{}
}

/*GetAccountCustomFieldsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetAccountCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/customFields][%d] getAccountCustomFieldsBadRequest ", 400)
}

func (o *GetAccountCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
