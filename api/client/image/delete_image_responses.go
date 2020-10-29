// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
	httperrors "github.com/metal-stack/metal-lib/httperrors"
)

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteImageOK creates a DeleteImageOK with default headers values
func NewDeleteImageOK() *DeleteImageOK {
	return &DeleteImageOK{}
}

/*DeleteImageOK handles this case with default header values.

OK
*/
type DeleteImageOK struct {
	Payload *models.V1ImageResponse
}

func (o *DeleteImageOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/image/{id}][%d] deleteImageOK  %+v", 200, o.Payload)
}

func (o *DeleteImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ImageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageDefault creates a DeleteImageDefault with default headers values
func NewDeleteImageDefault(code int) *DeleteImageDefault {
	return &DeleteImageDefault{
		_statusCode: code,
	}
}

/*DeleteImageDefault handles this case with default header values.

Error
*/
type DeleteImageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete image default response
func (o *DeleteImageDefault) Code() int {
	return o._statusCode
}

func (o *DeleteImageDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/image/{id}][%d] deleteImage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
