// Code generated by go-swagger; DO NOT EDIT.

package project

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

/*UpdateProjectOK handles this case with default header values.

Updated
*/
type UpdateProjectOK struct {
	Payload *models.V1ProjectResponse
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectOK  %+v", 200, o.Payload)
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

/*UpdateProjectPreconditionFailed handles this case with default header values.

OptimisticLock
*/
type UpdateProjectPreconditionFailed struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *UpdateProjectPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectPreconditionFailed  %+v", 412, o.Payload)
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

/*UpdateProjectDefault handles this case with default header values.

Error
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectDefault) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProject default  %+v", o._statusCode, o.Payload)
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
