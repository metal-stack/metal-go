// Code generated by go-swagger; DO NOT EDIT.

package size

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// FindSizeReader is a Reader for the FindSize structure.
type FindSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindSizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindSizeOK creates a FindSizeOK with default headers values
func NewFindSizeOK() *FindSizeOK {
	return &FindSizeOK{}
}

/*FindSizeOK handles this case with default header values.

OK
*/
type FindSizeOK struct {
	Payload *models.V1SizeResponse
}

func (o *FindSizeOK) Error() string {
	return fmt.Sprintf("[GET /v1/size/{id}][%d] findSizeOK  %+v", 200, o.Payload)
}

func (o *FindSizeOK) GetPayload() *models.V1SizeResponse {
	return o.Payload
}

func (o *FindSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindSizeDefault creates a FindSizeDefault with default headers values
func NewFindSizeDefault(code int) *FindSizeDefault {
	return &FindSizeDefault{
		_statusCode: code,
	}
}

/*FindSizeDefault handles this case with default header values.

Error
*/
type FindSizeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find size default response
func (o *FindSizeDefault) Code() int {
	return o._statusCode
}

func (o *FindSizeDefault) Error() string {
	return fmt.Sprintf("[GET /v1/size/{id}][%d] findSize default  %+v", o._statusCode, o.Payload)
}

func (o *FindSizeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindSizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
