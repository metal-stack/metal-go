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

/*DeleteSwitchOK handles this case with default header values.

OK
*/
type DeleteSwitchOK struct {
	Payload *models.V1SwitchResponse
}

func (o *DeleteSwitchOK) Error() string {
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

/*DeleteSwitchDefault handles this case with default header values.

Error
*/
type DeleteSwitchDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete switch default response
func (o *DeleteSwitchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSwitchDefault) Error() string {
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
