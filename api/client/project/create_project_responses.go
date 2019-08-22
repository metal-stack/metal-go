// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/*CreateProjectCreated handles this case with default header values.

Created
*/
type CreateProjectCreated struct {
	Payload *models.V1ProjectResponse
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProjectCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectConflict creates a CreateProjectConflict with default headers values
func NewCreateProjectConflict() *CreateProjectConflict {
	return &CreateProjectConflict{}
}

/*CreateProjectConflict handles this case with default header values.

Conflict
*/
type CreateProjectConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *CreateProjectConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProjectConflict  %+v", 409, o.Payload)
}

func (o *CreateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectDefault creates a CreateProjectDefault with default headers values
func NewCreateProjectDefault(code int) *CreateProjectDefault {
	return &CreateProjectDefault{
		_statusCode: code,
	}
}

/*CreateProjectDefault handles this case with default header values.

Error
*/
type CreateProjectDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the create project default response
func (o *CreateProjectDefault) Code() int {
	return o._statusCode
}

func (o *CreateProjectDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProject default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
