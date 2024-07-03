// Code generated by go-swagger; DO NOT EDIT.

package project

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

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 412:
		result := NewUpdateProjectPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*
UpdateProjectOK describes a response with status code 200, with default header values.

Updated
*/
type UpdateProjectOK struct {
	Payload *models.V1ProjectResponse
}

// IsSuccess returns true when this update project o k response has a 2xx status code
func (o *UpdateProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project o k response has a 3xx status code
func (o *UpdateProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project o k response has a 4xx status code
func (o *UpdateProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project o k response has a 5xx status code
func (o *UpdateProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project o k response a status code equal to that given
func (o *UpdateProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update project o k response
func (o *UpdateProjectOK) Code() int {
	return 200
}

func (o *UpdateProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectOK %s", 200, payload)
}

func (o *UpdateProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectOK %s", 200, payload)
}

func (o *UpdateProjectOK) GetPayload() *models.V1ProjectResponse {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectPreconditionFailed creates a UpdateProjectPreconditionFailed with default headers values
func NewUpdateProjectPreconditionFailed() *UpdateProjectPreconditionFailed {
	return &UpdateProjectPreconditionFailed{}
}

/*
UpdateProjectPreconditionFailed describes a response with status code 412, with default header values.

OptimisticLock
*/
type UpdateProjectPreconditionFailed struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update project precondition failed response has a 2xx status code
func (o *UpdateProjectPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project precondition failed response has a 3xx status code
func (o *UpdateProjectPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project precondition failed response has a 4xx status code
func (o *UpdateProjectPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project precondition failed response has a 5xx status code
func (o *UpdateProjectPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this update project precondition failed response a status code equal to that given
func (o *UpdateProjectPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the update project precondition failed response
func (o *UpdateProjectPreconditionFailed) Code() int {
	return 412
}

func (o *UpdateProjectPreconditionFailed) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectPreconditionFailed %s", 412, payload)
}

func (o *UpdateProjectPreconditionFailed) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectPreconditionFailed %s", 412, payload)
}

func (o *UpdateProjectPreconditionFailed) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateProjectPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDefault creates a UpdateProjectDefault with default headers values
func NewUpdateProjectDefault(code int) *UpdateProjectDefault {
	return &UpdateProjectDefault{
		_statusCode: code,
	}
}

/*
UpdateProjectDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update project default response has a 2xx status code
func (o *UpdateProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update project default response has a 3xx status code
func (o *UpdateProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update project default response has a 4xx status code
func (o *UpdateProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update project default response has a 5xx status code
func (o *UpdateProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update project default response a status code equal to that given
func (o *UpdateProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProject default %s", o._statusCode, payload)
}

func (o *UpdateProjectDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project][%d] updateProject default %s", o._statusCode, payload)
}

func (o *UpdateProjectDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
