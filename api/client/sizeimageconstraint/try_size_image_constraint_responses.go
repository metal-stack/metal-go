// Code generated by go-swagger; DO NOT EDIT.

package sizeimageconstraint

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

// TrySizeImageConstraintReader is a Reader for the TrySizeImageConstraint structure.
type TrySizeImageConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrySizeImageConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTrySizeImageConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTrySizeImageConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTrySizeImageConstraintOK creates a TrySizeImageConstraintOK with default headers values
func NewTrySizeImageConstraintOK() *TrySizeImageConstraintOK {
	return &TrySizeImageConstraintOK{}
}

/* TrySizeImageConstraintOK describes a response with status code 200, with default header values.

OK
*/
type TrySizeImageConstraintOK struct {
	Payload models.V1EmptyBody
}

func (o *TrySizeImageConstraintOK) Error() string {
	return fmt.Sprintf("[POST /v1/size-image-constraint/try][%d] trySizeImageConstraintOK  %+v", 200, o.Payload)
}
func (o *TrySizeImageConstraintOK) GetPayload() models.V1EmptyBody {
	return o.Payload
}

func (o *TrySizeImageConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrySizeImageConstraintDefault creates a TrySizeImageConstraintDefault with default headers values
func NewTrySizeImageConstraintDefault(code int) *TrySizeImageConstraintDefault {
	return &TrySizeImageConstraintDefault{
		_statusCode: code,
	}
}

/* TrySizeImageConstraintDefault describes a response with status code -1, with default header values.

Error
*/
type TrySizeImageConstraintDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the try size image constraint default response
func (o *TrySizeImageConstraintDefault) Code() int {
	return o._statusCode
}

func (o *TrySizeImageConstraintDefault) Error() string {
	return fmt.Sprintf("[POST /v1/size-image-constraint/try][%d] trySizeImageConstraint default  %+v", o._statusCode, o.Payload)
}
func (o *TrySizeImageConstraintDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *TrySizeImageConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}