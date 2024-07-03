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

// DeleteSizeImageConstraintReader is a Reader for the DeleteSizeImageConstraint structure.
type DeleteSizeImageConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSizeImageConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSizeImageConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSizeImageConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSizeImageConstraintOK creates a DeleteSizeImageConstraintOK with default headers values
func NewDeleteSizeImageConstraintOK() *DeleteSizeImageConstraintOK {
	return &DeleteSizeImageConstraintOK{}
}

/*
DeleteSizeImageConstraintOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSizeImageConstraintOK struct {
	Payload *models.V1SizeImageConstraintResponse
}

// IsSuccess returns true when this delete size image constraint o k response has a 2xx status code
func (o *DeleteSizeImageConstraintOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete size image constraint o k response has a 3xx status code
func (o *DeleteSizeImageConstraintOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete size image constraint o k response has a 4xx status code
func (o *DeleteSizeImageConstraintOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete size image constraint o k response has a 5xx status code
func (o *DeleteSizeImageConstraintOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete size image constraint o k response a status code equal to that given
func (o *DeleteSizeImageConstraintOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete size image constraint o k response
func (o *DeleteSizeImageConstraintOK) Code() int {
	return 200
}

func (o *DeleteSizeImageConstraintOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size-image-constraint/{id}][%d] deleteSizeImageConstraintOK %s", 200, payload)
}

func (o *DeleteSizeImageConstraintOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size-image-constraint/{id}][%d] deleteSizeImageConstraintOK %s", 200, payload)
}

func (o *DeleteSizeImageConstraintOK) GetPayload() *models.V1SizeImageConstraintResponse {
	return o.Payload
}

func (o *DeleteSizeImageConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeImageConstraintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSizeImageConstraintDefault creates a DeleteSizeImageConstraintDefault with default headers values
func NewDeleteSizeImageConstraintDefault(code int) *DeleteSizeImageConstraintDefault {
	return &DeleteSizeImageConstraintDefault{
		_statusCode: code,
	}
}

/*
DeleteSizeImageConstraintDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteSizeImageConstraintDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete size image constraint default response has a 2xx status code
func (o *DeleteSizeImageConstraintDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete size image constraint default response has a 3xx status code
func (o *DeleteSizeImageConstraintDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete size image constraint default response has a 4xx status code
func (o *DeleteSizeImageConstraintDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete size image constraint default response has a 5xx status code
func (o *DeleteSizeImageConstraintDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete size image constraint default response a status code equal to that given
func (o *DeleteSizeImageConstraintDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete size image constraint default response
func (o *DeleteSizeImageConstraintDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSizeImageConstraintDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size-image-constraint/{id}][%d] deleteSizeImageConstraint default %s", o._statusCode, payload)
}

func (o *DeleteSizeImageConstraintDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size-image-constraint/{id}][%d] deleteSizeImageConstraint default %s", o._statusCode, payload)
}

func (o *DeleteSizeImageConstraintDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteSizeImageConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
