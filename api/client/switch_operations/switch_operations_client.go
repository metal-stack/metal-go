// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new switch operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for switch operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSwitch(params *DeleteSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSwitchOK, error)

	FindSwitch(params *FindSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSwitchOK, error)

	ListSwitches(params *ListSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSwitchesOK, error)

	NotifySwitch(params *NotifySwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NotifySwitchOK, error)

	RegisterSwitch(params *RegisterSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterSwitchOK, *RegisterSwitchCreated, error)

	UpdateSwitch(params *UpdateSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSwitchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteSwitch deletes an switch and returns the deleted entity
*/
func (a *Client) DeleteSwitch(params *DeleteSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSwitchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSwitch",
		Method:             "DELETE",
		PathPattern:        "/v1/switch/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSwitchReader{formats: a.formats},
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
	success, ok := result.(*DeleteSwitchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSwitchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindSwitch gets switch by id
*/
func (a *Client) FindSwitch(params *FindSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSwitchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSwitch",
		Method:             "GET",
		PathPattern:        "/v1/switch/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSwitchReader{formats: a.formats},
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
	success, ok := result.(*FindSwitchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSwitchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSwitches gets all switches
*/
func (a *Client) ListSwitches(params *ListSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSwitchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSwitchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSwitches",
		Method:             "GET",
		PathPattern:        "/v1/switch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSwitchesReader{formats: a.formats},
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
	success, ok := result.(*ListSwitchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSwitchesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NotifySwitch notifies the metal api about a configuration change of a switch
*/
func (a *Client) NotifySwitch(params *NotifySwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NotifySwitchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotifySwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notifySwitch",
		Method:             "POST",
		PathPattern:        "/v1/switch/{id}/notify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NotifySwitchReader{formats: a.formats},
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
	success, ok := result.(*NotifySwitchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NotifySwitchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RegisterSwitch registers a switch
*/
func (a *Client) RegisterSwitch(params *RegisterSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterSwitchOK, *RegisterSwitchCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterSwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerSwitch",
		Method:             "POST",
		PathPattern:        "/v1/switch/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterSwitchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RegisterSwitchOK:
		return value, nil, nil
	case *RegisterSwitchCreated:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterSwitchDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSwitch updates a switch if the switch was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateSwitch(params *UpdateSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSwitchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSwitch",
		Method:             "POST",
		PathPattern:        "/v1/switch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSwitchReader{formats: a.formats},
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
	success, ok := result.(*UpdateSwitchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSwitchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
