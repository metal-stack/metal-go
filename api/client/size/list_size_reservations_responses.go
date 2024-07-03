// Code generated by go-swagger; DO NOT EDIT.

package size

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

// ListSizeReservationsReader is a Reader for the ListSizeReservations structure.
type ListSizeReservationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSizeReservationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSizeReservationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSizeReservationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSizeReservationsOK creates a ListSizeReservationsOK with default headers values
func NewListSizeReservationsOK() *ListSizeReservationsOK {
	return &ListSizeReservationsOK{}
}

/*
ListSizeReservationsOK describes a response with status code 200, with default header values.

OK
*/
type ListSizeReservationsOK struct {
	Payload []*models.V1SizeReservationResponse
}

// IsSuccess returns true when this list size reservations o k response has a 2xx status code
func (o *ListSizeReservationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list size reservations o k response has a 3xx status code
func (o *ListSizeReservationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list size reservations o k response has a 4xx status code
func (o *ListSizeReservationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list size reservations o k response has a 5xx status code
func (o *ListSizeReservationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list size reservations o k response a status code equal to that given
func (o *ListSizeReservationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list size reservations o k response
func (o *ListSizeReservationsOK) Code() int {
	return 200
}

func (o *ListSizeReservationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] listSizeReservationsOK %s", 200, payload)
}

func (o *ListSizeReservationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] listSizeReservationsOK %s", 200, payload)
}

func (o *ListSizeReservationsOK) GetPayload() []*models.V1SizeReservationResponse {
	return o.Payload
}

func (o *ListSizeReservationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSizeReservationsDefault creates a ListSizeReservationsDefault with default headers values
func NewListSizeReservationsDefault(code int) *ListSizeReservationsDefault {
	return &ListSizeReservationsDefault{
		_statusCode: code,
	}
}

/*
ListSizeReservationsDefault describes a response with status code -1, with default header values.

Error
*/
type ListSizeReservationsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this list size reservations default response has a 2xx status code
func (o *ListSizeReservationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list size reservations default response has a 3xx status code
func (o *ListSizeReservationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list size reservations default response has a 4xx status code
func (o *ListSizeReservationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list size reservations default response has a 5xx status code
func (o *ListSizeReservationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list size reservations default response a status code equal to that given
func (o *ListSizeReservationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list size reservations default response
func (o *ListSizeReservationsDefault) Code() int {
	return o._statusCode
}

func (o *ListSizeReservationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] listSizeReservations default %s", o._statusCode, payload)
}

func (o *ListSizeReservationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] listSizeReservations default %s", o._statusCode, payload)
}

func (o *ListSizeReservationsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListSizeReservationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
