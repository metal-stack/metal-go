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

// UpdateSizeReservationReader is a Reader for the UpdateSizeReservation structure.
type UpdateSizeReservationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSizeReservationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSizeReservationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateSizeReservationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSizeReservationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSizeReservationOK creates a UpdateSizeReservationOK with default headers values
func NewUpdateSizeReservationOK() *UpdateSizeReservationOK {
	return &UpdateSizeReservationOK{}
}

/*
UpdateSizeReservationOK describes a response with status code 200, with default header values.

OK
*/
type UpdateSizeReservationOK struct {
	Payload *models.V1SizeReservationResponse
}

// IsSuccess returns true when this update size reservation o k response has a 2xx status code
func (o *UpdateSizeReservationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update size reservation o k response has a 3xx status code
func (o *UpdateSizeReservationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update size reservation o k response has a 4xx status code
func (o *UpdateSizeReservationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update size reservation o k response has a 5xx status code
func (o *UpdateSizeReservationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update size reservation o k response a status code equal to that given
func (o *UpdateSizeReservationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update size reservation o k response
func (o *UpdateSizeReservationOK) Code() int {
	return 200
}

func (o *UpdateSizeReservationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservationOK %s", 200, payload)
}

func (o *UpdateSizeReservationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservationOK %s", 200, payload)
}

func (o *UpdateSizeReservationOK) GetPayload() *models.V1SizeReservationResponse {
	return o.Payload
}

func (o *UpdateSizeReservationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SizeReservationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSizeReservationConflict creates a UpdateSizeReservationConflict with default headers values
func NewUpdateSizeReservationConflict() *UpdateSizeReservationConflict {
	return &UpdateSizeReservationConflict{}
}

/*
UpdateSizeReservationConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateSizeReservationConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update size reservation conflict response has a 2xx status code
func (o *UpdateSizeReservationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update size reservation conflict response has a 3xx status code
func (o *UpdateSizeReservationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update size reservation conflict response has a 4xx status code
func (o *UpdateSizeReservationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update size reservation conflict response has a 5xx status code
func (o *UpdateSizeReservationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update size reservation conflict response a status code equal to that given
func (o *UpdateSizeReservationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update size reservation conflict response
func (o *UpdateSizeReservationConflict) Code() int {
	return 409
}

func (o *UpdateSizeReservationConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservationConflict %s", 409, payload)
}

func (o *UpdateSizeReservationConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservationConflict %s", 409, payload)
}

func (o *UpdateSizeReservationConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateSizeReservationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSizeReservationDefault creates a UpdateSizeReservationDefault with default headers values
func NewUpdateSizeReservationDefault(code int) *UpdateSizeReservationDefault {
	return &UpdateSizeReservationDefault{
		_statusCode: code,
	}
}

/*
UpdateSizeReservationDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateSizeReservationDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update size reservation default response has a 2xx status code
func (o *UpdateSizeReservationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update size reservation default response has a 3xx status code
func (o *UpdateSizeReservationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update size reservation default response has a 4xx status code
func (o *UpdateSizeReservationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update size reservation default response has a 5xx status code
func (o *UpdateSizeReservationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update size reservation default response a status code equal to that given
func (o *UpdateSizeReservationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update size reservation default response
func (o *UpdateSizeReservationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSizeReservationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservation default %s", o._statusCode, payload)
}

func (o *UpdateSizeReservationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/size/reservations][%d] updateSizeReservation default %s", o._statusCode, payload)
}

func (o *UpdateSizeReservationDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateSizeReservationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}