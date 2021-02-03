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
)

// NewFindFirewallParams creates a new FindFirewallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindFirewallParams() *FindFirewallParams {
	return &FindFirewallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindFirewallParamsWithTimeout creates a new FindFirewallParams object
// with the ability to set a timeout on a request.
func NewFindFirewallParamsWithTimeout(timeout time.Duration) *FindFirewallParams {
	return &FindFirewallParams{
		timeout: timeout,
	}
}

// NewFindFirewallParamsWithContext creates a new FindFirewallParams object
// with the ability to set a context for a request.
func NewFindFirewallParamsWithContext(ctx context.Context) *FindFirewallParams {
	return &FindFirewallParams{
		Context: ctx,
	}
}

// NewFindFirewallParamsWithHTTPClient creates a new FindFirewallParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindFirewallParamsWithHTTPClient(client *http.Client) *FindFirewallParams {
	return &FindFirewallParams{
		HTTPClient: client,
	}
}

/* FindFirewallParams contains all the parameters to send to the API endpoint
   for the find firewall operation.

   Typically these are written to a http.Request.
*/
type FindFirewallParams struct {

	/* ID.

	   identifier of the firewall
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find firewall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindFirewallParams) WithDefaults() *FindFirewallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find firewall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindFirewallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find firewall params
func (o *FindFirewallParams) WithTimeout(timeout time.Duration) *FindFirewallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find firewall params
func (o *FindFirewallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find firewall params
func (o *FindFirewallParams) WithContext(ctx context.Context) *FindFirewallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find firewall params
func (o *FindFirewallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find firewall params
func (o *FindFirewallParams) WithHTTPClient(client *http.Client) *FindFirewallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find firewall params
func (o *FindFirewallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find firewall params
func (o *FindFirewallParams) WithID(id string) *FindFirewallParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find firewall params
func (o *FindFirewallParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindFirewallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
