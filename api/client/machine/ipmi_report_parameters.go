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

// NewIpmiReportParams creates a new IpmiReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpmiReportParams() *IpmiReportParams {
	return &IpmiReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpmiReportParamsWithTimeout creates a new IpmiReportParams object
// with the ability to set a timeout on a request.
func NewIpmiReportParamsWithTimeout(timeout time.Duration) *IpmiReportParams {
	return &IpmiReportParams{
		timeout: timeout,
	}
}

// NewIpmiReportParamsWithContext creates a new IpmiReportParams object
// with the ability to set a context for a request.
func NewIpmiReportParamsWithContext(ctx context.Context) *IpmiReportParams {
	return &IpmiReportParams{
		Context: ctx,
	}
}

// NewIpmiReportParamsWithHTTPClient creates a new IpmiReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpmiReportParamsWithHTTPClient(client *http.Client) *IpmiReportParams {
	return &IpmiReportParams{
		HTTPClient: client,
	}
}

/*
IpmiReportParams contains all the parameters to send to the API endpoint

	for the ipmi report operation.

	Typically these are written to a http.Request.
*/
type IpmiReportParams struct {

	// Body.
	Body *models.V1MachineIpmiReports

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipmi report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpmiReportParams) WithDefaults() *IpmiReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipmi report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpmiReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipmi report params
func (o *IpmiReportParams) WithTimeout(timeout time.Duration) *IpmiReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipmi report params
func (o *IpmiReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipmi report params
func (o *IpmiReportParams) WithContext(ctx context.Context) *IpmiReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipmi report params
func (o *IpmiReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipmi report params
func (o *IpmiReportParams) WithHTTPClient(client *http.Client) *IpmiReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipmi report params
func (o *IpmiReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ipmi report params
func (o *IpmiReportParams) WithBody(body *models.V1MachineIpmiReports) *IpmiReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ipmi report params
func (o *IpmiReportParams) SetBody(body *models.V1MachineIpmiReports) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IpmiReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
