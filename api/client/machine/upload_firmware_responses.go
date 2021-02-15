// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/httperrors"
)

// UploadFirmwareReader is a Reader for the UploadFirmware structure.
type UploadFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFirmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadFirmwareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadFirmwareOK creates a UploadFirmwareOK with default headers values
func NewUploadFirmwareOK() *UploadFirmwareOK {
	return &UploadFirmwareOK{}
}

/*UploadFirmwareOK handles this case with default header values.

OK
*/
type UploadFirmwareOK struct {
}

func (o *UploadFirmwareOK) Error() string {
	return fmt.Sprintf("[PUT /v1/machine/upload-firmware/{vendor}/{board}/{revision}][%d] uploadFirmwareOK ", 200)
}

func (o *UploadFirmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadFirmwareDefault creates a UploadFirmwareDefault with default headers values
func NewUploadFirmwareDefault(code int) *UploadFirmwareDefault {
	return &UploadFirmwareDefault{
		_statusCode: code,
	}
}

/*UploadFirmwareDefault handles this case with default header values.

Error
*/
type UploadFirmwareDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the upload firmware default response
func (o *UploadFirmwareDefault) Code() int {
	return o._statusCode
}

func (o *UploadFirmwareDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/machine/upload-firmware/{vendor}/{board}/{revision}][%d] uploadFirmware default  %+v", o._statusCode, o.Payload)
}

func (o *UploadFirmwareDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UploadFirmwareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}