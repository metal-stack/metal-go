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

// FindSizeImageConstraintReader is a Reader for the FindSizeImageConstraint structure.
type FindSizeImageConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindSizeImageConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindSizeImageConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindSizeImageConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindSizeImageConstraintOK creates a FindSizeImageConstraintOK with default headers values
func NewFindSizeImageConstraintOK() *FindSizeImageConstraintOK {
	return &FindSizeImageConstraintOK{}
}

/*
FindSizeImageConstraintOK describes a response with status code 200, with default header values.

OK
*/
type FindSizeImageConstraintOK struct {
	Payload *models.V1SizeImageConstraintResponse
}

// IsSuccess returns true when this find size image constraint o k response has a 2xx status code
func (o *FindSizeImageConstraintOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find size image constraint o k response has a 3xx status code
func (o *FindSizeImageConstraintOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find size image constraint o k response has a 4xx status code
func (o *FindSizeImageConstraintOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find size image constraint o k response has a 5xx status code
func (o *FindSizeImageConstraintOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find size image constraint o k response a status code equal to that given
func (o *FindSizeImageConstraintOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindSizeImageConstraintOK) Error() string {
	return fmt.Sprintf("[GET /v1/size-image-constraint/{id}][%d] findSizeImageConstraintOK  %+v", 200, o.Payload)
}

func (o *FindSizeImageConstraintOK) String() string {
	return fmt.Sprintf("[GET /v1/size-image-constraint/{id}][%d] findSizeImageConstraintOK  %+v", 200, o.Payload)
}

func (o *FindSizeImageConstraintOK) GetPayload() *models.V1SizeImageConstraintResponse {
	return o.Payload
}

func (o *FindSizeImageConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeImageConstraintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindSizeImageConstraintDefault creates a FindSizeImageConstraintDefault with default headers values
func NewFindSizeImageConstraintDefault(code int) *FindSizeImageConstraintDefault {
	return &FindSizeImageConstraintDefault{
		_statusCode: code,
	}
}

/*
FindSizeImageConstraintDefault describes a response with status code -1, with default header values.

Error
*/
type FindSizeImageConstraintDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find size image constraint default response
func (o *FindSizeImageConstraintDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find size image constraint default response has a 2xx status code
func (o *FindSizeImageConstraintDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find size image constraint default response has a 3xx status code
func (o *FindSizeImageConstraintDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find size image constraint default response has a 4xx status code
func (o *FindSizeImageConstraintDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find size image constraint default response has a 5xx status code
func (o *FindSizeImageConstraintDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find size image constraint default response a status code equal to that given
func (o *FindSizeImageConstraintDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindSizeImageConstraintDefault) Error() string {
	return fmt.Sprintf("[GET /v1/size-image-constraint/{id}][%d] findSizeImageConstraint default  %+v", o._statusCode, o.Payload)
}

func (o *FindSizeImageConstraintDefault) String() string {
	return fmt.Sprintf("[GET /v1/size-image-constraint/{id}][%d] findSizeImageConstraint default  %+v", o._statusCode, o.Payload)
}

func (o *FindSizeImageConstraintDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindSizeImageConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
