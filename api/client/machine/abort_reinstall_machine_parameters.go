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

// NewAbortReinstallMachineParams creates a new AbortReinstallMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAbortReinstallMachineParams() *AbortReinstallMachineParams {
	return &AbortReinstallMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAbortReinstallMachineParamsWithTimeout creates a new AbortReinstallMachineParams object
// with the ability to set a timeout on a request.
func NewAbortReinstallMachineParamsWithTimeout(timeout time.Duration) *AbortReinstallMachineParams {
	return &AbortReinstallMachineParams{
		timeout: timeout,
	}
}

// NewAbortReinstallMachineParamsWithContext creates a new AbortReinstallMachineParams object
// with the ability to set a context for a request.
func NewAbortReinstallMachineParamsWithContext(ctx context.Context) *AbortReinstallMachineParams {
	return &AbortReinstallMachineParams{
		Context: ctx,
	}
}

// NewAbortReinstallMachineParamsWithHTTPClient creates a new AbortReinstallMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewAbortReinstallMachineParamsWithHTTPClient(client *http.Client) *AbortReinstallMachineParams {
	return &AbortReinstallMachineParams{
		HTTPClient: client,
	}
}

/* AbortReinstallMachineParams contains all the parameters to send to the API endpoint
   for the abort reinstall machine operation.

   Typically these are written to a http.Request.
*/
type AbortReinstallMachineParams struct {

	// Body.
	Body *models.V1MachineAbortReinstallRequest

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the abort reinstall machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AbortReinstallMachineParams) WithDefaults() *AbortReinstallMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the abort reinstall machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AbortReinstallMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the abort reinstall machine params
func (o *AbortReinstallMachineParams) WithTimeout(timeout time.Duration) *AbortReinstallMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the abort reinstall machine params
func (o *AbortReinstallMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the abort reinstall machine params
func (o *AbortReinstallMachineParams) WithContext(ctx context.Context) *AbortReinstallMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the abort reinstall machine params
func (o *AbortReinstallMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the abort reinstall machine params
func (o *AbortReinstallMachineParams) WithHTTPClient(client *http.Client) *AbortReinstallMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the abort reinstall machine params
func (o *AbortReinstallMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the abort reinstall machine params
func (o *AbortReinstallMachineParams) WithBody(body *models.V1MachineAbortReinstallRequest) *AbortReinstallMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the abort reinstall machine params
func (o *AbortReinstallMachineParams) SetBody(body *models.V1MachineAbortReinstallRequest) {
	o.Body = body
}

// WithID adds the id to the abort reinstall machine params
func (o *AbortReinstallMachineParams) WithID(id string) *AbortReinstallMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the abort reinstall machine params
func (o *AbortReinstallMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AbortReinstallMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
