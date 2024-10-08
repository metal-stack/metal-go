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

// DeleteSizeReservationReader is a Reader for the DeleteSizeReservation structure.
type DeleteSizeReservationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSizeReservationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSizeReservationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSizeReservationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSizeReservationOK creates a DeleteSizeReservationOK with default headers values
func NewDeleteSizeReservationOK() *DeleteSizeReservationOK {
	return &DeleteSizeReservationOK{}
}

/*
DeleteSizeReservationOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSizeReservationOK struct {
	Payload *models.V1SizeReservationResponse
}

// IsSuccess returns true when this delete size reservation o k response has a 2xx status code
func (o *DeleteSizeReservationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete size reservation o k response has a 3xx status code
func (o *DeleteSizeReservationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete size reservation o k response has a 4xx status code
func (o *DeleteSizeReservationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete size reservation o k response has a 5xx status code
func (o *DeleteSizeReservationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete size reservation o k response a status code equal to that given
func (o *DeleteSizeReservationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete size reservation o k response
func (o *DeleteSizeReservationOK) Code() int {
	return 200
}

func (o *DeleteSizeReservationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size/reservations/{id}][%d] deleteSizeReservationOK %s", 200, payload)
}

func (o *DeleteSizeReservationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size/reservations/{id}][%d] deleteSizeReservationOK %s", 200, payload)
}

func (o *DeleteSizeReservationOK) GetPayload() *models.V1SizeReservationResponse {
	return o.Payload
}

func (o *DeleteSizeReservationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeReservationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSizeReservationDefault creates a DeleteSizeReservationDefault with default headers values
func NewDeleteSizeReservationDefault(code int) *DeleteSizeReservationDefault {
	return &DeleteSizeReservationDefault{
		_statusCode: code,
	}
}

/*
DeleteSizeReservationDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteSizeReservationDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete size reservation default response has a 2xx status code
func (o *DeleteSizeReservationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete size reservation default response has a 3xx status code
func (o *DeleteSizeReservationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete size reservation default response has a 4xx status code
func (o *DeleteSizeReservationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete size reservation default response has a 5xx status code
func (o *DeleteSizeReservationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete size reservation default response a status code equal to that given
func (o *DeleteSizeReservationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete size reservation default response
func (o *DeleteSizeReservationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSizeReservationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size/reservations/{id}][%d] deleteSizeReservation default %s", o._statusCode, payload)
}

func (o *DeleteSizeReservationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/size/reservations/{id}][%d] deleteSizeReservation default %s", o._statusCode, payload)
}

func (o *DeleteSizeReservationDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteSizeReservationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
