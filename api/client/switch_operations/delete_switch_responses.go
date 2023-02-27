// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// DeleteSwitchReader is a Reader for the DeleteSwitch structure.
type DeleteSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSwitchOK creates a DeleteSwitchOK with default headers values
func NewDeleteSwitchOK() *DeleteSwitchOK {
	return &DeleteSwitchOK{}
}

/*
DeleteSwitchOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSwitchOK struct {
	Payload *models.V1SwitchResponse
}

// IsSuccess returns true when this delete switch o k response has a 2xx status code
func (o *DeleteSwitchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete switch o k response has a 3xx status code
func (o *DeleteSwitchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete switch o k response has a 4xx status code
func (o *DeleteSwitchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete switch o k response has a 5xx status code
func (o *DeleteSwitchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete switch o k response a status code equal to that given
func (o *DeleteSwitchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete switch o k response
func (o *DeleteSwitchOK) Code() int {
	return 200
}

func (o *DeleteSwitchOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/switch/{id}][%d] deleteSwitchOK  %+v", 200, o.Payload)
}

func (o *DeleteSwitchOK) String() string {
	return fmt.Sprintf("[DELETE /v1/switch/{id}][%d] deleteSwitchOK  %+v", 200, o.Payload)
}

func (o *DeleteSwitchOK) GetPayload() *models.V1SwitchResponse {
	return o.Payload
}

func (o *DeleteSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSwitchDefault creates a DeleteSwitchDefault with default headers values
func NewDeleteSwitchDefault(code int) *DeleteSwitchDefault {
	return &DeleteSwitchDefault{
		_statusCode: code,
	}
}

/*
DeleteSwitchDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteSwitchDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this delete switch default response has a 2xx status code
func (o *DeleteSwitchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete switch default response has a 3xx status code
func (o *DeleteSwitchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete switch default response has a 4xx status code
func (o *DeleteSwitchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete switch default response has a 5xx status code
func (o *DeleteSwitchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete switch default response a status code equal to that given
func (o *DeleteSwitchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete switch default response
func (o *DeleteSwitchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSwitchDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/switch/{id}][%d] deleteSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSwitchDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/switch/{id}][%d] deleteSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSwitchDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
