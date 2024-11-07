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

// NewMigrateSwitchParams creates a new MigrateSwitchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMigrateSwitchParams() *MigrateSwitchParams {
	return &MigrateSwitchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMigrateSwitchParamsWithTimeout creates a new MigrateSwitchParams object
// with the ability to set a timeout on a request.
func NewMigrateSwitchParamsWithTimeout(timeout time.Duration) *MigrateSwitchParams {
	return &MigrateSwitchParams{
		timeout: timeout,
	}
}

// NewMigrateSwitchParamsWithContext creates a new MigrateSwitchParams object
// with the ability to set a context for a request.
func NewMigrateSwitchParamsWithContext(ctx context.Context) *MigrateSwitchParams {
	return &MigrateSwitchParams{
		Context: ctx,
	}
}

// NewMigrateSwitchParamsWithHTTPClient creates a new MigrateSwitchParams object
// with the ability to set a custom HTTPClient for a request.
func NewMigrateSwitchParamsWithHTTPClient(client *http.Client) *MigrateSwitchParams {
	return &MigrateSwitchParams{
		HTTPClient: client,
	}
}

/*
MigrateSwitchParams contains all the parameters to send to the API endpoint

	for the migrate switch operation.

	Typically these are written to a http.Request.
*/
type MigrateSwitchParams struct {

	// Body.
	Body *models.V1SwitchMigrateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the migrate switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrateSwitchParams) WithDefaults() *MigrateSwitchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the migrate switch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrateSwitchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the migrate switch params
func (o *MigrateSwitchParams) WithTimeout(timeout time.Duration) *MigrateSwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the migrate switch params
func (o *MigrateSwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the migrate switch params
func (o *MigrateSwitchParams) WithContext(ctx context.Context) *MigrateSwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the migrate switch params
func (o *MigrateSwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the migrate switch params
func (o *MigrateSwitchParams) WithHTTPClient(client *http.Client) *MigrateSwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the migrate switch params
func (o *MigrateSwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the migrate switch params
func (o *MigrateSwitchParams) WithBody(body *models.V1SwitchMigrateRequest) *MigrateSwitchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the migrate switch params
func (o *MigrateSwitchParams) SetBody(body *models.V1SwitchMigrateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MigrateSwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
