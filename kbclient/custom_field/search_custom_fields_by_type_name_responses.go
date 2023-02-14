// Code generated by go-swagger; DO NOT EDIT.

package custom_field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// SearchCustomFieldsByTypeNameReader is a Reader for the SearchCustomFieldsByTypeName structure.
type SearchCustomFieldsByTypeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCustomFieldsByTypeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchCustomFieldsByTypeNameOK()
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

// NewSearchCustomFieldsByTypeNameOK creates a SearchCustomFieldsByTypeNameOK with default headers values
func NewSearchCustomFieldsByTypeNameOK() *SearchCustomFieldsByTypeNameOK {
	return &SearchCustomFieldsByTypeNameOK{}
}

/*SearchCustomFieldsByTypeNameOK handles this case with default header values.

successful operation
*/
type SearchCustomFieldsByTypeNameOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *SearchCustomFieldsByTypeNameOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/search][%d] searchCustomFieldsByTypeNameOK  %+v", 200, o.Payload)
}

func (o *SearchCustomFieldsByTypeNameOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *SearchCustomFieldsByTypeNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
