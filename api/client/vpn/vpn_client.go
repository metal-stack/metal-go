// Code generated by go-swagger; DO NOT EDIT.

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vpn API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vpn API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetVPNAuthKey(params *GetVPNAuthKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVPNAuthKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetVPNAuthKey creates auth key to connect to project s v p n
*/
func (a *Client) GetVPNAuthKey(params *GetVPNAuthKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVPNAuthKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVPNAuthKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVPNAuthKey",
		Method:             "POST",
		PathPattern:        "/v1/vpn/authkey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVPNAuthKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVPNAuthKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVPNAuthKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
