// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewToggleSwitchPortParams creates a new ToggleSwitchPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewToggleSwitchPortParams() *ToggleSwitchPortParams {
	return &ToggleSwitchPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewToggleSwitchPortParamsWithTimeout creates a new ToggleSwitchPortParams object
// with the ability to set a timeout on a request.
func NewToggleSwitchPortParamsWithTimeout(timeout time.Duration) *ToggleSwitchPortParams {
	return &ToggleSwitchPortParams{
		timeout: timeout,
	}
}

// NewToggleSwitchPortParamsWithContext creates a new ToggleSwitchPortParams object
// with the ability to set a context for a request.
func NewToggleSwitchPortParamsWithContext(ctx context.Context) *ToggleSwitchPortParams {
	return &ToggleSwitchPortParams{
		Context: ctx,
	}
}

// NewToggleSwitchPortParamsWithHTTPClient creates a new ToggleSwitchPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewToggleSwitchPortParamsWithHTTPClient(client *http.Client) *ToggleSwitchPortParams {
	return &ToggleSwitchPortParams{
		HTTPClient: client,
	}
}

/*
ToggleSwitchPortParams contains all the parameters to send to the API endpoint

	for the toggle switch port operation.

	Typically these are written to a http.Request.
*/
type ToggleSwitchPortParams struct {

	// Body.
	Body *models.V1SwitchPortToggleRequest

	/* ID.

	   identifier of the switch
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the toggle switch port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleSwitchPortParams) WithDefaults() *ToggleSwitchPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the toggle switch port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleSwitchPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the toggle switch port params
func (o *ToggleSwitchPortParams) WithTimeout(timeout time.Duration) *ToggleSwitchPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle switch port params
func (o *ToggleSwitchPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle switch port params
func (o *ToggleSwitchPortParams) WithContext(ctx context.Context) *ToggleSwitchPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle switch port params
func (o *ToggleSwitchPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle switch port params
func (o *ToggleSwitchPortParams) WithHTTPClient(client *http.Client) *ToggleSwitchPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle switch port params
func (o *ToggleSwitchPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the toggle switch port params
func (o *ToggleSwitchPortParams) WithBody(body *models.V1SwitchPortToggleRequest) *ToggleSwitchPortParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the toggle switch port params
func (o *ToggleSwitchPortParams) SetBody(body *models.V1SwitchPortToggleRequest) {
	o.Body = body
}

// WithID adds the id to the toggle switch port params
func (o *ToggleSwitchPortParams) WithID(id string) *ToggleSwitchPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the toggle switch port params
func (o *ToggleSwitchPortParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleSwitchPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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