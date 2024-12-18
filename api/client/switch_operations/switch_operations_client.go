// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new switch operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new switch operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new switch operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for switch operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSwitch(params *DeleteSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSwitchOK, error)

	FindSwitch(params *FindSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSwitchOK, error)

	FindSwitches(params *FindSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSwitchesOK, error)

	ListSwitches(params *ListSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSwitchesOK, error)

	MigrateSwitch(params *MigrateSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MigrateSwitchOK, error)

	NotifySwitch(params *NotifySwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NotifySwitchOK, error)

	RegisterSwitch(params *RegisterSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterSwitchOK, *RegisterSwitchCreated, error)

	ToggleSwitchPort(params *ToggleSwitchPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleSwitchPortOK, error)

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
FindSwitches gets all switches that match given properties
*/
func (a *Client) FindSwitches(params *FindSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSwitchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSwitchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSwitches",
		Method:             "POST",
		PathPattern:        "/v1/switch/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSwitchesReader{formats: a.formats},
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
	success, ok := result.(*FindSwitchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSwitchesDefault)
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
MigrateSwitch migrates machine connections from one switch to another
*/
func (a *Client) MigrateSwitch(params *MigrateSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MigrateSwitchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMigrateSwitchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "migrateSwitch",
		Method:             "POST",
		PathPattern:        "/v1/switch/migrate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MigrateSwitchReader{formats: a.formats},
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
	success, ok := result.(*MigrateSwitchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MigrateSwitchDefault)
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
ToggleSwitchPort toggles the port of the switch with a nicname to the given state
*/
func (a *Client) ToggleSwitchPort(params *ToggleSwitchPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleSwitchPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleSwitchPortParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "toggleSwitchPort",
		Method:             "POST",
		PathPattern:        "/v1/switch/{id}/port",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ToggleSwitchPortReader{formats: a.formats},
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
	success, ok := result.(*ToggleSwitchPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ToggleSwitchPortDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
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
