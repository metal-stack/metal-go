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
)

// NewGetProvisioningEventContainerParams creates a new GetProvisioningEventContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProvisioningEventContainerParams() *GetProvisioningEventContainerParams {
	return &GetProvisioningEventContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvisioningEventContainerParamsWithTimeout creates a new GetProvisioningEventContainerParams object
// with the ability to set a timeout on a request.
func NewGetProvisioningEventContainerParamsWithTimeout(timeout time.Duration) *GetProvisioningEventContainerParams {
	return &GetProvisioningEventContainerParams{
		timeout: timeout,
	}
}

// NewGetProvisioningEventContainerParamsWithContext creates a new GetProvisioningEventContainerParams object
// with the ability to set a context for a request.
func NewGetProvisioningEventContainerParamsWithContext(ctx context.Context) *GetProvisioningEventContainerParams {
	return &GetProvisioningEventContainerParams{
		Context: ctx,
	}
}

// NewGetProvisioningEventContainerParamsWithHTTPClient creates a new GetProvisioningEventContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProvisioningEventContainerParamsWithHTTPClient(client *http.Client) *GetProvisioningEventContainerParams {
	return &GetProvisioningEventContainerParams{
		HTTPClient: client,
	}
}

/* GetProvisioningEventContainerParams contains all the parameters to send to the API endpoint
   for the get provisioning event container operation.

   Typically these are written to a http.Request.
*/
type GetProvisioningEventContainerParams struct {

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get provisioning event container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProvisioningEventContainerParams) WithDefaults() *GetProvisioningEventContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get provisioning event container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProvisioningEventContainerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) WithTimeout(timeout time.Duration) *GetProvisioningEventContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) WithContext(ctx context.Context) *GetProvisioningEventContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) WithHTTPClient(client *http.Client) *GetProvisioningEventContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) WithID(id string) *GetProvisioningEventContainerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get provisioning event container params
func (o *GetProvisioningEventContainerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvisioningEventContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
