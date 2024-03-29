// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
)

// NewUpdateFirmwareParams creates a new UpdateFirmwareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFirmwareParams() *UpdateFirmwareParams {
	return &UpdateFirmwareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFirmwareParamsWithTimeout creates a new UpdateFirmwareParams object
// with the ability to set a timeout on a request.
func NewUpdateFirmwareParamsWithTimeout(timeout time.Duration) *UpdateFirmwareParams {
	return &UpdateFirmwareParams{
		timeout: timeout,
	}
}

// NewUpdateFirmwareParamsWithContext creates a new UpdateFirmwareParams object
// with the ability to set a context for a request.
func NewUpdateFirmwareParamsWithContext(ctx context.Context) *UpdateFirmwareParams {
	return &UpdateFirmwareParams{
		Context: ctx,
	}
}

// NewUpdateFirmwareParamsWithHTTPClient creates a new UpdateFirmwareParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFirmwareParamsWithHTTPClient(client *http.Client) *UpdateFirmwareParams {
	return &UpdateFirmwareParams{
		HTTPClient: client,
	}
}

/*
UpdateFirmwareParams contains all the parameters to send to the API endpoint

	for the update firmware operation.

	Typically these are written to a http.Request.
*/
type UpdateFirmwareParams struct {

	// Body.
	Body *models.V1MachineUpdateFirmwareRequest

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update firmware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFirmwareParams) WithDefaults() *UpdateFirmwareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update firmware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFirmwareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update firmware params
func (o *UpdateFirmwareParams) WithTimeout(timeout time.Duration) *UpdateFirmwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update firmware params
func (o *UpdateFirmwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update firmware params
func (o *UpdateFirmwareParams) WithContext(ctx context.Context) *UpdateFirmwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update firmware params
func (o *UpdateFirmwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update firmware params
func (o *UpdateFirmwareParams) WithHTTPClient(client *http.Client) *UpdateFirmwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update firmware params
func (o *UpdateFirmwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update firmware params
func (o *UpdateFirmwareParams) WithBody(body *models.V1MachineUpdateFirmwareRequest) *UpdateFirmwareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update firmware params
func (o *UpdateFirmwareParams) SetBody(body *models.V1MachineUpdateFirmwareRequest) {
	o.Body = body
}

// WithID adds the id to the update firmware params
func (o *UpdateFirmwareParams) WithID(id string) *UpdateFirmwareParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update firmware params
func (o *UpdateFirmwareParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFirmwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
