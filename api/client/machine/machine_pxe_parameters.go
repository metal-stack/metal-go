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

// NewMachinePxeParams creates a new MachinePxeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachinePxeParams() *MachinePxeParams {
	return &MachinePxeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachinePxeParamsWithTimeout creates a new MachinePxeParams object
// with the ability to set a timeout on a request.
func NewMachinePxeParamsWithTimeout(timeout time.Duration) *MachinePxeParams {
	return &MachinePxeParams{
		timeout: timeout,
	}
}

// NewMachinePxeParamsWithContext creates a new MachinePxeParams object
// with the ability to set a context for a request.
func NewMachinePxeParamsWithContext(ctx context.Context) *MachinePxeParams {
	return &MachinePxeParams{
		Context: ctx,
	}
}

// NewMachinePxeParamsWithHTTPClient creates a new MachinePxeParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachinePxeParamsWithHTTPClient(client *http.Client) *MachinePxeParams {
	return &MachinePxeParams{
		HTTPClient: client,
	}
}

/* MachinePxeParams contains all the parameters to send to the API endpoint
   for the machine pxe operation.

   Typically these are written to a http.Request.
*/
type MachinePxeParams struct {

	// Body.
	Body models.V1EmptyBody

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the machine pxe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinePxeParams) WithDefaults() *MachinePxeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machine pxe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachinePxeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machine pxe params
func (o *MachinePxeParams) WithTimeout(timeout time.Duration) *MachinePxeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine pxe params
func (o *MachinePxeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine pxe params
func (o *MachinePxeParams) WithContext(ctx context.Context) *MachinePxeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine pxe params
func (o *MachinePxeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine pxe params
func (o *MachinePxeParams) WithHTTPClient(client *http.Client) *MachinePxeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine pxe params
func (o *MachinePxeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the machine pxe params
func (o *MachinePxeParams) WithBody(body models.V1EmptyBody) *MachinePxeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the machine pxe params
func (o *MachinePxeParams) SetBody(body models.V1EmptyBody) {
	o.Body = body
}

// WithID adds the id to the machine pxe params
func (o *MachinePxeParams) WithID(id string) *MachinePxeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the machine pxe params
func (o *MachinePxeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MachinePxeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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