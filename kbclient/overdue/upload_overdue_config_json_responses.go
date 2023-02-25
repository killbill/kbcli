// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// UploadOverdueConfigJSONReader is a Reader for the UploadOverdueConfigJSON structure.
type UploadOverdueConfigJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadOverdueConfigJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadOverdueConfigJSONCreated()
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

// NewUploadOverdueConfigJSONCreated creates a UploadOverdueConfigJSONCreated with default headers values
func NewUploadOverdueConfigJSONCreated() *UploadOverdueConfigJSONCreated {
	return &UploadOverdueConfigJSONCreated{}
}

/*UploadOverdueConfigJSONCreated handles this case with default header values.

Successfully uploaded overdue config
*/
type UploadOverdueConfigJSONCreated struct {
	Payload *kbmodel.Overdue

	HttpResponse runtime.ClientResponse
}

func (o *UploadOverdueConfigJSONCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/overdue][%d] uploadOverdueConfigJsonCreated  %+v", 201, o.Payload)
}

func (o *UploadOverdueConfigJSONCreated) GetPayload() *kbmodel.Overdue {
	return o.Payload
}

func (o *UploadOverdueConfigJSONCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Overdue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadOverdueConfigJSONBadRequest creates a UploadOverdueConfigJSONBadRequest with default headers values
func NewUploadOverdueConfigJSONBadRequest() *UploadOverdueConfigJSONBadRequest {
	return &UploadOverdueConfigJSONBadRequest{}
}

/*UploadOverdueConfigJSONBadRequest handles this case with default header values.

Invalid node command supplied
*/
type UploadOverdueConfigJSONBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *UploadOverdueConfigJSONBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/overdue][%d] uploadOverdueConfigJsonBadRequest ", 400)
}

func (o *UploadOverdueConfigJSONBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
