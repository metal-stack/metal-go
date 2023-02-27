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

// CreateSizeReader is a Reader for the CreateSize structure.
type CreateSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSizeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateSizeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSizeCreated creates a CreateSizeCreated with default headers values
func NewCreateSizeCreated() *CreateSizeCreated {
	return &CreateSizeCreated{}
}

/*
CreateSizeCreated describes a response with status code 201, with default header values.

Created
*/
type CreateSizeCreated struct {
	Payload *models.V1SizeResponse
}

// IsSuccess returns true when this create size created response has a 2xx status code
func (o *CreateSizeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create size created response has a 3xx status code
func (o *CreateSizeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create size created response has a 4xx status code
func (o *CreateSizeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create size created response has a 5xx status code
func (o *CreateSizeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create size created response a status code equal to that given
func (o *CreateSizeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create size created response
func (o *CreateSizeCreated) Code() int {
	return 201
}

func (o *CreateSizeCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSizeCreated  %+v", 201, o.Payload)
}

func (o *CreateSizeCreated) String() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSizeCreated  %+v", 201, o.Payload)
}

func (o *CreateSizeCreated) GetPayload() *models.V1SizeResponse {
	return o.Payload
}

func (o *CreateSizeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSizeConflict creates a CreateSizeConflict with default headers values
func NewCreateSizeConflict() *CreateSizeConflict {
	return &CreateSizeConflict{}
}

/*
CreateSizeConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateSizeConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this create size conflict response has a 2xx status code
func (o *CreateSizeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create size conflict response has a 3xx status code
func (o *CreateSizeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create size conflict response has a 4xx status code
func (o *CreateSizeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create size conflict response has a 5xx status code
func (o *CreateSizeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create size conflict response a status code equal to that given
func (o *CreateSizeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create size conflict response
func (o *CreateSizeConflict) Code() int {
	return 409
}

func (o *CreateSizeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSizeConflict  %+v", 409, o.Payload)
}

func (o *CreateSizeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSizeConflict  %+v", 409, o.Payload)
}

func (o *CreateSizeConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateSizeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSizeDefault creates a CreateSizeDefault with default headers values
func NewCreateSizeDefault(code int) *CreateSizeDefault {
	return &CreateSizeDefault{
		_statusCode: code,
	}
}

/*
CreateSizeDefault describes a response with status code -1, with default header values.

Error
*/
type CreateSizeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this create size default response has a 2xx status code
func (o *CreateSizeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create size default response has a 3xx status code
func (o *CreateSizeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create size default response has a 4xx status code
func (o *CreateSizeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create size default response has a 5xx status code
func (o *CreateSizeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create size default response a status code equal to that given
func (o *CreateSizeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create size default response
func (o *CreateSizeDefault) Code() int {
	return o._statusCode
}

func (o *CreateSizeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSize default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSizeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/size][%d] createSize default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSizeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateSizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
