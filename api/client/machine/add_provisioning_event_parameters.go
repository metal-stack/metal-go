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

// NewAddProvisioningEventParams creates a new AddProvisioningEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddProvisioningEventParams() *AddProvisioningEventParams {
	return &AddProvisioningEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddProvisioningEventParamsWithTimeout creates a new AddProvisioningEventParams object
// with the ability to set a timeout on a request.
func NewAddProvisioningEventParamsWithTimeout(timeout time.Duration) *AddProvisioningEventParams {
	return &AddProvisioningEventParams{
		timeout: timeout,
	}
}

// NewAddProvisioningEventParamsWithContext creates a new AddProvisioningEventParams object
// with the ability to set a context for a request.
func NewAddProvisioningEventParamsWithContext(ctx context.Context) *AddProvisioningEventParams {
	return &AddProvisioningEventParams{
		Context: ctx,
	}
}

// NewAddProvisioningEventParamsWithHTTPClient creates a new AddProvisioningEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddProvisioningEventParamsWithHTTPClient(client *http.Client) *AddProvisioningEventParams {
	return &AddProvisioningEventParams{
		HTTPClient: client,
	}
}

/*
AddProvisioningEventParams contains all the parameters to send to the API endpoint

	for the add provisioning event operation.

	Typically these are written to a http.Request.
*/
type AddProvisioningEventParams struct {

	// Body.
	Body *models.V1MachineProvisioningEvent

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add provisioning event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProvisioningEventParams) WithDefaults() *AddProvisioningEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add provisioning event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProvisioningEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add provisioning event params
func (o *AddProvisioningEventParams) WithTimeout(timeout time.Duration) *AddProvisioningEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add provisioning event params
func (o *AddProvisioningEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add provisioning event params
func (o *AddProvisioningEventParams) WithContext(ctx context.Context) *AddProvisioningEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add provisioning event params
func (o *AddProvisioningEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add provisioning event params
func (o *AddProvisioningEventParams) WithHTTPClient(client *http.Client) *AddProvisioningEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add provisioning event params
func (o *AddProvisioningEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add provisioning event params
func (o *AddProvisioningEventParams) WithBody(body *models.V1MachineProvisioningEvent) *AddProvisioningEventParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add provisioning event params
func (o *AddProvisioningEventParams) SetBody(body *models.V1MachineProvisioningEvent) {
	o.Body = body
}

// WithID adds the id to the add provisioning event params
func (o *AddProvisioningEventParams) WithID(id string) *AddProvisioningEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add provisioning event params
func (o *AddProvisioningEventParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddProvisioningEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
