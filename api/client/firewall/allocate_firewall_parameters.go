// Code generated by go-swagger; DO NOT EDIT.

package firewall

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

// NewAllocateFirewallParams creates a new AllocateFirewallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllocateFirewallParams() *AllocateFirewallParams {
	return &AllocateFirewallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateFirewallParamsWithTimeout creates a new AllocateFirewallParams object
// with the ability to set a timeout on a request.
func NewAllocateFirewallParamsWithTimeout(timeout time.Duration) *AllocateFirewallParams {
	return &AllocateFirewallParams{
		timeout: timeout,
	}
}

// NewAllocateFirewallParamsWithContext creates a new AllocateFirewallParams object
// with the ability to set a context for a request.
func NewAllocateFirewallParamsWithContext(ctx context.Context) *AllocateFirewallParams {
	return &AllocateFirewallParams{
		Context: ctx,
	}
}

// NewAllocateFirewallParamsWithHTTPClient creates a new AllocateFirewallParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllocateFirewallParamsWithHTTPClient(client *http.Client) *AllocateFirewallParams {
	return &AllocateFirewallParams{
		HTTPClient: client,
	}
}

/* AllocateFirewallParams contains all the parameters to send to the API endpoint
   for the allocate firewall operation.

   Typically these are written to a http.Request.
*/
type AllocateFirewallParams struct {

	// Body.
	Body *models.V1FirewallCreateRequest

	/* Try.

	   try allocation before actually doing so to get informed early about possible incompatible allocation parameters
	*/
	Try *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allocate firewall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateFirewallParams) WithDefaults() *AllocateFirewallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allocate firewall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateFirewallParams) SetDefaults() {
	var (
		tryDefault = bool(false)
	)

	val := AllocateFirewallParams{
		Try: &tryDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the allocate firewall params
func (o *AllocateFirewallParams) WithTimeout(timeout time.Duration) *AllocateFirewallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate firewall params
func (o *AllocateFirewallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate firewall params
func (o *AllocateFirewallParams) WithContext(ctx context.Context) *AllocateFirewallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate firewall params
func (o *AllocateFirewallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate firewall params
func (o *AllocateFirewallParams) WithHTTPClient(client *http.Client) *AllocateFirewallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate firewall params
func (o *AllocateFirewallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allocate firewall params
func (o *AllocateFirewallParams) WithBody(body *models.V1FirewallCreateRequest) *AllocateFirewallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allocate firewall params
func (o *AllocateFirewallParams) SetBody(body *models.V1FirewallCreateRequest) {
	o.Body = body
}

// WithTry adds the try to the allocate firewall params
func (o *AllocateFirewallParams) WithTry(try *bool) *AllocateFirewallParams {
	o.SetTry(try)
	return o
}

// SetTry adds the try to the allocate firewall params
func (o *AllocateFirewallParams) SetTry(try *bool) {
	o.Try = try
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateFirewallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Try != nil {

		// query param try
		var qrTry bool

		if o.Try != nil {
			qrTry = *o.Try
		}
		qTry := swag.FormatBool(qrTry)
		if qTry != "" {

			if err := r.SetQueryParam("try", qTry); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
