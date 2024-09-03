// Code generated by go-swagger; DO NOT EDIT.

package network

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
	"github.com/go-openapi/swag"

	"github.com/metal-stack/metal-go/api/models"
)

// NewUpdateNetworkParams creates a new UpdateNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkParams() *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkParamsWithTimeout creates a new UpdateNetworkParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkParamsWithTimeout(timeout time.Duration) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkParamsWithContext creates a new UpdateNetworkParams object
// with the ability to set a context for a request.
func NewUpdateNetworkParamsWithContext(ctx context.Context) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		Context: ctx,
	}
}

// NewUpdateNetworkParamsWithHTTPClient creates a new UpdateNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkParamsWithHTTPClient(client *http.Client) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkParams contains all the parameters to send to the API endpoint

	for the update network operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkParams struct {

	// Body.
	Body *models.V1NetworkUpdateRequest

	/* Force.

	   if true update forcefully
	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) WithDefaults() *UpdateNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := UpdateNetworkParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) WithTimeout(timeout time.Duration) *UpdateNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network params
func (o *UpdateNetworkParams) WithContext(ctx context.Context) *UpdateNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network params
func (o *UpdateNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) WithHTTPClient(client *http.Client) *UpdateNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update network params
func (o *UpdateNetworkParams) WithBody(body *models.V1NetworkUpdateRequest) *UpdateNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update network params
func (o *UpdateNetworkParams) SetBody(body *models.V1NetworkUpdateRequest) {
	o.Body = body
}

// WithForce adds the force to the update network params
func (o *UpdateNetworkParams) WithForce(force *bool) *UpdateNetworkParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the update network params
func (o *UpdateNetworkParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
