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
)

// RegisterSwitchReader is a Reader for the RegisterSwitch structure.
type RegisterSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRegisterSwitchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterSwitchOK creates a RegisterSwitchOK with default headers values
func NewRegisterSwitchOK() *RegisterSwitchOK {
	return &RegisterSwitchOK{}
}

/*RegisterSwitchOK handles this case with default header values.

OK
*/
type RegisterSwitchOK struct {
	Payload *models.V1SwitchResponse
}

func (o *RegisterSwitchOK) Error() string {
	return fmt.Sprintf("[POST /v1/switch/register][%d] registerSwitchOK  %+v", 200, o.Payload)
}

func (o *RegisterSwitchOK) GetPayload() *models.V1SwitchResponse {
	return o.Payload
}

func (o *RegisterSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterSwitchCreated creates a RegisterSwitchCreated with default headers values
func NewRegisterSwitchCreated() *RegisterSwitchCreated {
	return &RegisterSwitchCreated{}
}

/*RegisterSwitchCreated handles this case with default header values.

Created
*/
type RegisterSwitchCreated struct {
	Payload *models.V1SwitchResponse
}

func (o *RegisterSwitchCreated) Error() string {
	return fmt.Sprintf("[POST /v1/switch/register][%d] registerSwitchCreated  %+v", 201, o.Payload)
}

func (o *RegisterSwitchCreated) GetPayload() *models.V1SwitchResponse {
	return o.Payload
}

func (o *RegisterSwitchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterSwitchDefault creates a RegisterSwitchDefault with default headers values
func NewRegisterSwitchDefault(code int) *RegisterSwitchDefault {
	return &RegisterSwitchDefault{
		_statusCode: code,
	}
}

/*RegisterSwitchDefault handles this case with default header values.

Error
*/
type RegisterSwitchDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the register switch default response
func (o *RegisterSwitchDefault) Code() int {
	return o._statusCode
}

func (o *RegisterSwitchDefault) Error() string {
	return fmt.Sprintf("[POST /v1/switch/register][%d] registerSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterSwitchDefault) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *RegisterSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
