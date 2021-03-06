// Code generated by go-swagger; DO NOT EDIT.

package machine

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

// IpmiReportReader is a Reader for the IpmiReport structure.
type IpmiReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpmiReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpmiReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpmiReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpmiReportOK creates a IpmiReportOK with default headers values
func NewIpmiReportOK() *IpmiReportOK {
	return &IpmiReportOK{}
}

/* IpmiReportOK describes a response with status code 200, with default header values.

OK
*/
type IpmiReportOK struct {
	Payload *models.V1MachineIpmiReportResponse
}

func (o *IpmiReportOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReportOK  %+v", 200, o.Payload)
}
func (o *IpmiReportOK) GetPayload() *models.V1MachineIpmiReportResponse {
	return o.Payload
}

func (o *IpmiReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineIpmiReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpmiReportDefault creates a IpmiReportDefault with default headers values
func NewIpmiReportDefault(code int) *IpmiReportDefault {
	return &IpmiReportDefault{
		_statusCode: code,
	}
}

/* IpmiReportDefault describes a response with status code -1, with default header values.

Error
*/
type IpmiReportDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the ipmi report default response
func (o *IpmiReportDefault) Code() int {
	return o._statusCode
}

func (o *IpmiReportDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReport default  %+v", o._statusCode, o.Payload)
}
func (o *IpmiReportDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *IpmiReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
