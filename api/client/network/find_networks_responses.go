// Code generated by go-swagger; DO NOT EDIT.

package network

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

// FindNetworksReader is a Reader for the FindNetworks structure.
type FindNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindNetworksOK creates a FindNetworksOK with default headers values
func NewFindNetworksOK() *FindNetworksOK {
	return &FindNetworksOK{}
}

/*
FindNetworksOK describes a response with status code 200, with default header values.

OK
*/
type FindNetworksOK struct {
	Payload []*models.V1NetworkResponse
}

// IsSuccess returns true when this find networks o k response has a 2xx status code
func (o *FindNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find networks o k response has a 3xx status code
func (o *FindNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find networks o k response has a 4xx status code
func (o *FindNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find networks o k response has a 5xx status code
func (o *FindNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find networks o k response a status code equal to that given
func (o *FindNetworksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find networks o k response
func (o *FindNetworksOK) Code() int {
	return 200
}

func (o *FindNetworksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/find][%d] findNetworksOK %s", 200, payload)
}

func (o *FindNetworksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/find][%d] findNetworksOK %s", 200, payload)
}

func (o *FindNetworksOK) GetPayload() []*models.V1NetworkResponse {
	return o.Payload
}

func (o *FindNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindNetworksDefault creates a FindNetworksDefault with default headers values
func NewFindNetworksDefault(code int) *FindNetworksDefault {
	return &FindNetworksDefault{
		_statusCode: code,
	}
}

/*
FindNetworksDefault describes a response with status code -1, with default header values.

Error
*/
type FindNetworksDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this find networks default response has a 2xx status code
func (o *FindNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find networks default response has a 3xx status code
func (o *FindNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find networks default response has a 4xx status code
func (o *FindNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find networks default response has a 5xx status code
func (o *FindNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find networks default response a status code equal to that given
func (o *FindNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find networks default response
func (o *FindNetworksDefault) Code() int {
	return o._statusCode
}

func (o *FindNetworksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/find][%d] findNetworks default %s", o._statusCode, payload)
}

func (o *FindNetworksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/find][%d] findNetworks default %s", o._statusCode, payload)
}

func (o *FindNetworksDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
