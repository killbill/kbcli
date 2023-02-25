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

// GetAllTagsReader is a Reader for the GetAllTags structure.
type GetAllTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTagsOK()
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

// NewGetAllTagsOK creates a GetAllTagsOK with default headers values
func NewGetAllTagsOK() *GetAllTagsOK {
	return &GetAllTagsOK{}
}

/*GetAllTagsOK handles this case with default header values.

successful operation
*/
type GetAllTagsOK struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *GetAllTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetAllTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsBadRequest creates a GetAllTagsBadRequest with default headers values
func NewGetAllTagsBadRequest() *GetAllTagsBadRequest {
	return &GetAllTagsBadRequest{}
}

/*GetAllTagsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetAllTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAllTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsBadRequest ", 400)
}

func (o *GetAllTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTagsNotFound creates a GetAllTagsNotFound with default headers values
func NewGetAllTagsNotFound() *GetAllTagsNotFound {
	return &GetAllTagsNotFound{}
}

/*GetAllTagsNotFound handles this case with default header values.

Account not found
*/
type GetAllTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAllTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsNotFound ", 404)
}

func (o *GetAllTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
