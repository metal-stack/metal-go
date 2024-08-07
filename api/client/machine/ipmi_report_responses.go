// Code generated by go-swagger; DO NOT EDIT.

package machine

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

/*
IpmiReportOK describes a response with status code 200, with default header values.

OK
*/
type IpmiReportOK struct {
	Payload *models.V1MachineIpmiReportResponse
}

// IsSuccess returns true when this ipmi report o k response has a 2xx status code
func (o *IpmiReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipmi report o k response has a 3xx status code
func (o *IpmiReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipmi report o k response has a 4xx status code
func (o *IpmiReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipmi report o k response has a 5xx status code
func (o *IpmiReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipmi report o k response a status code equal to that given
func (o *IpmiReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipmi report o k response
func (o *IpmiReportOK) Code() int {
	return 200
}

func (o *IpmiReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReportOK %s", 200, payload)
}

func (o *IpmiReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReportOK %s", 200, payload)
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

/*
IpmiReportDefault describes a response with status code -1, with default header values.

Error
*/
type IpmiReportDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this ipmi report default response has a 2xx status code
func (o *IpmiReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipmi report default response has a 3xx status code
func (o *IpmiReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipmi report default response has a 4xx status code
func (o *IpmiReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipmi report default response has a 5xx status code
func (o *IpmiReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipmi report default response a status code equal to that given
func (o *IpmiReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipmi report default response
func (o *IpmiReportDefault) Code() int {
	return o._statusCode
}

func (o *IpmiReportDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReport default %s", o._statusCode, payload)
}

func (o *IpmiReportDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/ipmi][%d] ipmiReport default %s", o._statusCode, payload)
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
