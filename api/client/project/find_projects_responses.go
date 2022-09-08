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

// FindProjectsReader is a Reader for the FindProjects structure.
type FindProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindProjectsOK creates a FindProjectsOK with default headers values
func NewFindProjectsOK() *FindProjectsOK {
	return &FindProjectsOK{}
}

/*
FindProjectsOK describes a response with status code 200, with default header values.

OK
*/
type FindProjectsOK struct {
	Payload []*models.V1ProjectResponse
}

// IsSuccess returns true when this find projects o k response has a 2xx status code
func (o *FindProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find projects o k response has a 3xx status code
func (o *FindProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find projects o k response has a 4xx status code
func (o *FindProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find projects o k response has a 5xx status code
func (o *FindProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find projects o k response a status code equal to that given
func (o *FindProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindProjectsOK) Error() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjectsOK  %+v", 200, o.Payload)
}

func (o *FindProjectsOK) String() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjectsOK  %+v", 200, o.Payload)
}

func (o *FindProjectsOK) GetPayload() []*models.V1ProjectResponse {
	return o.Payload
}

func (o *FindProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindProjectsDefault creates a FindProjectsDefault with default headers values
func NewFindProjectsDefault(code int) *FindProjectsDefault {
	return &FindProjectsDefault{
		_statusCode: code,
	}
}

/*
FindProjectsDefault describes a response with status code -1, with default header values.

Error
*/
type FindProjectsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find projects default response
func (o *FindProjectsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find projects default response has a 2xx status code
func (o *FindProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find projects default response has a 3xx status code
func (o *FindProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find projects default response has a 4xx status code
func (o *FindProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find projects default response has a 5xx status code
func (o *FindProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find projects default response a status code equal to that given
func (o *FindProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindProjectsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjects default  %+v", o._statusCode, o.Payload)
}

func (o *FindProjectsDefault) String() string {
	return fmt.Sprintf("[POST /v1/project/find][%d] findProjects default  %+v", o._statusCode, o.Payload)
}

func (o *FindProjectsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
