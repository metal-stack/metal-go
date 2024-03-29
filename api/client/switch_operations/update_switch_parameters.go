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

// NewUpdateSwitchParams creates a new UpdateSwitchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSwitchParams() *UpdateSwitchParams {
	return &UpdateSwitchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSwitchParamsWithTimeout creates a new UpdateSwitchParams object
// with the ability to set a timeout on a request.
func NewUpdateSwitchParamsWithTimeout(timeout time.Duration) *UpdateSwitchParams {
	return &UpdateSwitchParams{
		timeout: timeout,
	}
}

// NewUpdateSwitchParamsWithContext creates a new UpdateSwitchParams object
// with the ability to set a context for a request.
func NewUpdateSwitchParamsWithContext(ctx context.Context) *UpdateSwitchParams {
	return &UpdateSwitchParams{
		Context: ctx,
	}
}

// NewUpdateSwitchParamsWithHTTPClient creates a new UpdateSwitchParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSwitchParamsWithHTTPClient(client *http.Client) *UpdateSwitchParams {
	return &UpdateSwitchParams{
		HTTPClient: client,
	}
}

/*
UpdateSwitchParams contains all the parameters to send to the API endpoint

	for the update switch operation.

	Typically these are written to a http.Request.
*/
type UpdateSwitchParams struct {

	// Body.
	Body *models.V1SwitchUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSwitchParams) WithDefaults() *UpdateSwitchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSwitchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update switch params
func (o *UpdateSwitchParams) WithTimeout(timeout time.Duration) *UpdateSwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update switch params
func (o *UpdateSwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update switch params
func (o *UpdateSwitchParams) WithContext(ctx context.Context) *UpdateSwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update switch params
func (o *UpdateSwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update switch params
func (o *UpdateSwitchParams) WithHTTPClient(client *http.Client) *UpdateSwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update switch params
func (o *UpdateSwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update switch params
func (o *UpdateSwitchParams) WithBody(body *models.V1SwitchUpdateRequest) *UpdateSwitchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update switch params
func (o *UpdateSwitchParams) SetBody(body *models.V1SwitchUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
