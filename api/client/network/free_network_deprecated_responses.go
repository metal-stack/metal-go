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

// FreeNetworkDeprecatedReader is a Reader for the FreeNetworkDeprecated structure.
type FreeNetworkDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FreeNetworkDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFreeNetworkDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewFreeNetworkDeprecatedConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFreeNetworkDeprecatedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFreeNetworkDeprecatedOK creates a FreeNetworkDeprecatedOK with default headers values
func NewFreeNetworkDeprecatedOK() *FreeNetworkDeprecatedOK {
	return &FreeNetworkDeprecatedOK{}
}

/*
FreeNetworkDeprecatedOK describes a response with status code 200, with default header values.

OK
*/
type FreeNetworkDeprecatedOK struct {
	Payload *models.V1NetworkResponse
}

// IsSuccess returns true when this free network deprecated o k response has a 2xx status code
func (o *FreeNetworkDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this free network deprecated o k response has a 3xx status code
func (o *FreeNetworkDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this free network deprecated o k response has a 4xx status code
func (o *FreeNetworkDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this free network deprecated o k response has a 5xx status code
func (o *FreeNetworkDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this free network deprecated o k response a status code equal to that given
func (o *FreeNetworkDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the free network deprecated o k response
func (o *FreeNetworkDeprecatedOK) Code() int {
	return 200
}

func (o *FreeNetworkDeprecatedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecatedOK %s", 200, payload)
}

func (o *FreeNetworkDeprecatedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecatedOK %s", 200, payload)
}

func (o *FreeNetworkDeprecatedOK) GetPayload() *models.V1NetworkResponse {
	return o.Payload
}

func (o *FreeNetworkDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeNetworkDeprecatedConflict creates a FreeNetworkDeprecatedConflict with default headers values
func NewFreeNetworkDeprecatedConflict() *FreeNetworkDeprecatedConflict {
	return &FreeNetworkDeprecatedConflict{}
}

/*
FreeNetworkDeprecatedConflict describes a response with status code 409, with default header values.

Conflict
*/
type FreeNetworkDeprecatedConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this free network deprecated conflict response has a 2xx status code
func (o *FreeNetworkDeprecatedConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this free network deprecated conflict response has a 3xx status code
func (o *FreeNetworkDeprecatedConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this free network deprecated conflict response has a 4xx status code
func (o *FreeNetworkDeprecatedConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this free network deprecated conflict response has a 5xx status code
func (o *FreeNetworkDeprecatedConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this free network deprecated conflict response a status code equal to that given
func (o *FreeNetworkDeprecatedConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the free network deprecated conflict response
func (o *FreeNetworkDeprecatedConflict) Code() int {
	return 409
}

func (o *FreeNetworkDeprecatedConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecatedConflict %s", 409, payload)
}

func (o *FreeNetworkDeprecatedConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecatedConflict %s", 409, payload)
}

func (o *FreeNetworkDeprecatedConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FreeNetworkDeprecatedConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeNetworkDeprecatedDefault creates a FreeNetworkDeprecatedDefault with default headers values
func NewFreeNetworkDeprecatedDefault(code int) *FreeNetworkDeprecatedDefault {
	return &FreeNetworkDeprecatedDefault{
		_statusCode: code,
	}
}

/*
FreeNetworkDeprecatedDefault describes a response with status code -1, with default header values.

Error
*/
type FreeNetworkDeprecatedDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this free network deprecated default response has a 2xx status code
func (o *FreeNetworkDeprecatedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this free network deprecated default response has a 3xx status code
func (o *FreeNetworkDeprecatedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this free network deprecated default response has a 4xx status code
func (o *FreeNetworkDeprecatedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this free network deprecated default response has a 5xx status code
func (o *FreeNetworkDeprecatedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this free network deprecated default response a status code equal to that given
func (o *FreeNetworkDeprecatedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the free network deprecated default response
func (o *FreeNetworkDeprecatedDefault) Code() int {
	return o._statusCode
}

func (o *FreeNetworkDeprecatedDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecated default %s", o._statusCode, payload)
}

func (o *FreeNetworkDeprecatedDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkDeprecated default %s", o._statusCode, payload)
}

func (o *FreeNetworkDeprecatedDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FreeNetworkDeprecatedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
