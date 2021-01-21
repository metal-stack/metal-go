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

// DeleteSizeReader is a Reader for the DeleteSize structure.
type DeleteSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSizeOK creates a DeleteSizeOK with default headers values
func NewDeleteSizeOK() *DeleteSizeOK {
	return &DeleteSizeOK{}
}

/* DeleteSizeOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSizeOK struct {
	Payload *models.V1SizeResponse
}

func (o *DeleteSizeOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/size/{id}][%d] deleteSizeOK  %+v", 200, o.Payload)
}
func (o *DeleteSizeOK) GetPayload() *models.V1SizeResponse {
	return o.Payload
}

func (o *DeleteSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSizeDefault creates a DeleteSizeDefault with default headers values
func NewDeleteSizeDefault(code int) *DeleteSizeDefault {
	return &DeleteSizeDefault{
		_statusCode: code,
	}
}

/* DeleteSizeDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteSizeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete size default response
func (o *DeleteSizeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSizeDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/size/{id}][%d] deleteSize default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteSizeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteSizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
