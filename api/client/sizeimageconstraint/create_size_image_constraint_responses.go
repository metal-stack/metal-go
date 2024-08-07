// Code generated by go-swagger; DO NOT EDIT.

package sizeimageconstraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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

/*
CreateSizeImageConstraintCreated describes a response with status code 201, with default header values.

Created
*/
type CreateSizeImageConstraintCreated struct {
	Payload *models.V1SizeImageConstraintResponse
}

// IsSuccess returns true when this create size image constraint created response has a 2xx status code
func (o *CreateSizeImageConstraintCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create size image constraint created response has a 3xx status code
func (o *CreateSizeImageConstraintCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create size image constraint created response has a 4xx status code
func (o *CreateSizeImageConstraintCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create size image constraint created response has a 5xx status code
func (o *CreateSizeImageConstraintCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create size image constraint created response a status code equal to that given
func (o *CreateSizeImageConstraintCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create size image constraint created response
func (o *CreateSizeImageConstraintCreated) Code() int {
	return 201
}

func (o *CreateSizeImageConstraintCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraintCreated %s", 201, payload)
}

func (o *CreateSizeImageConstraintCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraintCreated %s", 201, payload)
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

/*
CreateSizeImageConstraintConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateSizeImageConstraintConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this create size image constraint conflict response has a 2xx status code
func (o *CreateSizeImageConstraintConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create size image constraint conflict response has a 3xx status code
func (o *CreateSizeImageConstraintConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create size image constraint conflict response has a 4xx status code
func (o *CreateSizeImageConstraintConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create size image constraint conflict response has a 5xx status code
func (o *CreateSizeImageConstraintConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create size image constraint conflict response a status code equal to that given
func (o *CreateSizeImageConstraintConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create size image constraint conflict response
func (o *CreateSizeImageConstraintConflict) Code() int {
	return 409
}

func (o *CreateSizeImageConstraintConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraintConflict %s", 409, payload)
}

func (o *CreateSizeImageConstraintConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraintConflict %s", 409, payload)
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

/*
CreateSizeImageConstraintDefault describes a response with status code -1, with default header values.

Error
*/
type CreateSizeImageConstraintDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this create size image constraint default response has a 2xx status code
func (o *CreateSizeImageConstraintDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create size image constraint default response has a 3xx status code
func (o *CreateSizeImageConstraintDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create size image constraint default response has a 4xx status code
func (o *CreateSizeImageConstraintDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create size image constraint default response has a 5xx status code
func (o *CreateSizeImageConstraintDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create size image constraint default response a status code equal to that given
func (o *CreateSizeImageConstraintDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create size image constraint default response
func (o *CreateSizeImageConstraintDefault) Code() int {
	return o._statusCode
}

func (o *CreateSizeImageConstraintDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraint default %s", o._statusCode, payload)
}

func (o *CreateSizeImageConstraintDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/size-image-constraint][%d] createSizeImageConstraint default %s", o._statusCode, payload)
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
