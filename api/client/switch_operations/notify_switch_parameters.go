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

// NewNotifySwitchParams creates a new NotifySwitchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotifySwitchParams() *NotifySwitchParams {
	return &NotifySwitchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotifySwitchParamsWithTimeout creates a new NotifySwitchParams object
// with the ability to set a timeout on a request.
func NewNotifySwitchParamsWithTimeout(timeout time.Duration) *NotifySwitchParams {
	return &NotifySwitchParams{
		timeout: timeout,
	}
}

// NewNotifySwitchParamsWithContext creates a new NotifySwitchParams object
// with the ability to set a context for a request.
func NewNotifySwitchParamsWithContext(ctx context.Context) *NotifySwitchParams {
	return &NotifySwitchParams{
		Context: ctx,
	}
}

// NewNotifySwitchParamsWithHTTPClient creates a new NotifySwitchParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotifySwitchParamsWithHTTPClient(client *http.Client) *NotifySwitchParams {
	return &NotifySwitchParams{
		HTTPClient: client,
	}
}

/*
NotifySwitchParams contains all the parameters to send to the API endpoint

	for the notify switch operation.

	Typically these are written to a http.Request.
*/
type NotifySwitchParams struct {

	// Body.
	Body *models.V1SwitchNotifyRequest

	/* ID.

	   identifier of the switch
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notify switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifySwitchParams) WithDefaults() *NotifySwitchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notify switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifySwitchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notify switch params
func (o *NotifySwitchParams) WithTimeout(timeout time.Duration) *NotifySwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify switch params
func (o *NotifySwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify switch params
func (o *NotifySwitchParams) WithContext(ctx context.Context) *NotifySwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify switch params
func (o *NotifySwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify switch params
func (o *NotifySwitchParams) WithHTTPClient(client *http.Client) *NotifySwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify switch params
func (o *NotifySwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the notify switch params
func (o *NotifySwitchParams) WithBody(body *models.V1SwitchNotifyRequest) *NotifySwitchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notify switch params
func (o *NotifySwitchParams) SetBody(body *models.V1SwitchNotifyRequest) {
	o.Body = body
}

// WithID adds the id to the notify switch params
func (o *NotifySwitchParams) WithID(id string) *NotifySwitchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the notify switch params
func (o *NotifySwitchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NotifySwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
