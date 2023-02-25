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
)

// UploadOverdueConfigXMLReader is a Reader for the UploadOverdueConfigXML structure.
type UploadOverdueConfigXMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadOverdueConfigXMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadOverdueConfigXMLCreated()
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

// NewUploadOverdueConfigXMLCreated creates a UploadOverdueConfigXMLCreated with default headers values
func NewUploadOverdueConfigXMLCreated() *UploadOverdueConfigXMLCreated {
	return &UploadOverdueConfigXMLCreated{}
}

/*UploadOverdueConfigXMLCreated handles this case with default header values.

Successfully uploaded overdue config
*/
type UploadOverdueConfigXMLCreated struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *UploadOverdueConfigXMLCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/overdue/xml][%d] uploadOverdueConfigXmlCreated  %+v", 201, o.Payload)
}

func (o *UploadOverdueConfigXMLCreated) GetPayload() string {
	return o.Payload
}

func (o *UploadOverdueConfigXMLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadOverdueConfigXMLBadRequest creates a UploadOverdueConfigXMLBadRequest with default headers values
func NewUploadOverdueConfigXMLBadRequest() *UploadOverdueConfigXMLBadRequest {
	return &UploadOverdueConfigXMLBadRequest{}
}

/*UploadOverdueConfigXMLBadRequest handles this case with default header values.

Invalid node command supplied
*/
type UploadOverdueConfigXMLBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *UploadOverdueConfigXMLBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/overdue/xml][%d] uploadOverdueConfigXmlBadRequest ", 400)
}

func (o *UploadOverdueConfigXMLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
