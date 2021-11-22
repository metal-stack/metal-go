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

// CreateSizeImageConstraintReader is a Reader for the CreateSizeImageConstraint structure.
type CreateSizeImageConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSizeImageConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSizeImageConstraintCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateSizeImageConstraintConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSizeImageConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSizeImageConstraintCreated creates a CreateSizeImageConstraintCreated with default headers values
func NewCreateSizeImageConstraintCreated() *CreateSizeImageConstraintCreated {
	return &CreateSizeImageConstraintCreated{}
}

/* CreateSizeImageConstraintCreated describes a response with status code 201, with default header values.

Created
*/
type CreateSizeImageConstraintCreated struct {
	Payload *models.V1SizeImageConstraintResponse
}

func (o *CreateSizeImageConstraintCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/size/sizeimageconstraints][%d] createSizeImageConstraintCreated  %+v", 201, o.Payload)
}
func (o *CreateSizeImageConstraintCreated) GetPayload() *models.V1SizeImageConstraintResponse {
	return o.Payload
}

func (o *CreateSizeImageConstraintCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeImageConstraintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSizeImageConstraintConflict creates a CreateSizeImageConstraintConflict with default headers values
func NewCreateSizeImageConstraintConflict() *CreateSizeImageConstraintConflict {
	return &CreateSizeImageConstraintConflict{}
}

/* CreateSizeImageConstraintConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateSizeImageConstraintConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *CreateSizeImageConstraintConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/size/sizeimageconstraints][%d] createSizeImageConstraintConflict  %+v", 409, o.Payload)
}
func (o *CreateSizeImageConstraintConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateSizeImageConstraintConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSizeImageConstraintDefault creates a CreateSizeImageConstraintDefault with default headers values
func NewCreateSizeImageConstraintDefault(code int) *CreateSizeImageConstraintDefault {
	return &CreateSizeImageConstraintDefault{
		_statusCode: code,
	}
}

/* CreateSizeImageConstraintDefault describes a response with status code -1, with default header values.

Error
*/
type CreateSizeImageConstraintDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the create size image constraint default response
func (o *CreateSizeImageConstraintDefault) Code() int {
	return o._statusCode
}

func (o *CreateSizeImageConstraintDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/size/sizeimageconstraints][%d] createSizeImageConstraint default  %+v", o._statusCode, o.Payload)
}
func (o *CreateSizeImageConstraintDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateSizeImageConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
