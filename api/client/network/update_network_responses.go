// Code generated by go-swagger; DO NOT EDIT.

package network

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

// UpdateNetworkReader is a Reader for the UpdateNetwork structure.
type UpdateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetworkOK creates a UpdateNetworkOK with default headers values
func NewUpdateNetworkOK() *UpdateNetworkOK {
	return &UpdateNetworkOK{}
}

/*
	UpdateNetworkOK describes a response with status code 200, with default header values.

OK
*/
type UpdateNetworkOK struct {
	Payload *models.V1NetworkResponse
}

func (o *UpdateNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/network][%d] updateNetworkOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkOK) GetPayload() *models.V1NetworkResponse {
	return o.Payload
}

func (o *UpdateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkConflict creates a UpdateNetworkConflict with default headers values
func NewUpdateNetworkConflict() *UpdateNetworkConflict {
	return &UpdateNetworkConflict{}
}

/*
	UpdateNetworkConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateNetworkConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *UpdateNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/network][%d] updateNetworkConflict  %+v", 409, o.Payload)
}
func (o *UpdateNetworkConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkDefault creates a UpdateNetworkDefault with default headers values
func NewUpdateNetworkDefault(code int) *UpdateNetworkDefault {
	return &UpdateNetworkDefault{
		_statusCode: code,
	}
}

/*
	UpdateNetworkDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateNetworkDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the update network default response
func (o *UpdateNetworkDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/network][%d] updateNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateNetworkDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
