// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new machine API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for machine API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AbortReinstallMachine(params *AbortReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter) (*AbortReinstallMachineOK, error)

	AddProvisioningEvent(params *AddProvisioningEventParams, authInfo runtime.ClientAuthInfoWriter) (*AddProvisioningEventOK, error)

	AllocateMachine(params *AllocateMachineParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateMachineOK, error)

	ChassisIdentifyLEDOff(params *ChassisIdentifyLEDOffParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOffOK, error)

	ChassisIdentifyLEDOn(params *ChassisIdentifyLEDOnParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOnOK, error)

	FinalizeAllocation(params *FinalizeAllocationParams, authInfo runtime.ClientAuthInfoWriter) (*FinalizeAllocationOK, error)

	FindIPMIMachine(params *FindIPMIMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPMIMachineOK, error)

	FindIPMIMachines(params *FindIPMIMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPMIMachinesOK, error)

	FindMachine(params *FindMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachineOK, error)

	FindMachines(params *FindMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachinesOK, error)

	FreeMachine(params *FreeMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FreeMachineOK, error)

	GetProvisioningEventContainer(params *GetProvisioningEventContainerParams, authInfo runtime.ClientAuthInfoWriter) (*GetProvisioningEventContainerOK, error)

	IpmiReport(params *IpmiReportParams, authInfo runtime.ClientAuthInfoWriter) (*IpmiReportOK, error)

	ListMachines(params *ListMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*ListMachinesOK, error)

	MachineBios(params *MachineBiosParams, authInfo runtime.ClientAuthInfoWriter) (*MachineBiosOK, error)

	MachineDisk(params *MachineDiskParams, authInfo runtime.ClientAuthInfoWriter) (*MachineDiskOK, error)

	MachineOff(params *MachineOffParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOffOK, error)

	MachineOn(params *MachineOnParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOnOK, error)

	MachinePxe(params *MachinePxeParams, authInfo runtime.ClientAuthInfoWriter) (*MachinePxeOK, error)

	MachineReset(params *MachineResetParams, authInfo runtime.ClientAuthInfoWriter) (*MachineResetOK, error)

	RegisterMachine(params *RegisterMachineParams, authInfo runtime.ClientAuthInfoWriter) (*RegisterMachineOK, *RegisterMachineCreated, error)

	ReinstallMachine(params *ReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter) (*ReinstallMachineOK, error)

	SetChassisIdentifyLEDState(params *SetChassisIdentifyLEDStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetChassisIdentifyLEDStateOK, error)

	SetMachineState(params *SetMachineStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetMachineStateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AbortReinstallMachine aborts reinstall this machine
*/
func (a *Client) AbortReinstallMachine(params *AbortReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter) (*AbortReinstallMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAbortReinstallMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "abortReinstallMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/abort-reinstall",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AbortReinstallMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AbortReinstallMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AbortReinstallMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddProvisioningEvent adds a machine provisioning event
*/
func (a *Client) AddProvisioningEvent(params *AddProvisioningEventParams, authInfo runtime.ClientAuthInfoWriter) (*AddProvisioningEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProvisioningEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addProvisioningEvent",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddProvisioningEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddProvisioningEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddProvisioningEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AllocateMachine allocates a machine
*/
func (a *Client) AllocateMachine(params *AllocateMachineParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllocateMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AllocateMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChassisIdentifyLEDOff sends a power off to the chassis identify l e d
*/
func (a *Client) ChassisIdentifyLEDOff(params *ChassisIdentifyLEDOffParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChassisIdentifyLEDOffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "chassisIdentifyLEDOff",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/chassis-identify-led-off",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChassisIdentifyLEDOffReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChassisIdentifyLEDOffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChassisIdentifyLEDOffDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChassisIdentifyLEDOn sends a power on to the chassis identify l e d
*/
func (a *Client) ChassisIdentifyLEDOn(params *ChassisIdentifyLEDOnParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChassisIdentifyLEDOnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "chassisIdentifyLEDOn",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/chassis-identify-led-on",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChassisIdentifyLEDOnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChassisIdentifyLEDOnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChassisIdentifyLEDOnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FinalizeAllocation finalizes the allocation of the machine by reconfiguring the switch sent on successful image installation
*/
func (a *Client) FinalizeAllocation(params *FinalizeAllocationParams, authInfo runtime.ClientAuthInfoWriter) (*FinalizeAllocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFinalizeAllocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "finalizeAllocation",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/finalize-allocation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FinalizeAllocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FinalizeAllocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FinalizeAllocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindIPMIMachine returns a machine including the ipmi connection data
*/
func (a *Client) FindIPMIMachine(params *FindIPMIMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPMIMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPMIMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPMIMachine",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}/ipmi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPMIMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindIPMIMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindIPMIMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindIPMIMachines returns machines including the ipmi connection data
*/
func (a *Client) FindIPMIMachines(params *FindIPMIMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPMIMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPMIMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPMIMachines",
		Method:             "POST",
		PathPattern:        "/v1/machine/ipmi/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPMIMachinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindIPMIMachinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindIPMIMachinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindMachine gets machine by id
*/
func (a *Client) FindMachine(params *FindMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findMachine",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindMachines finds machines by multiple criteria
*/
func (a *Client) FindMachines(params *FindMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findMachines",
		Method:             "POST",
		PathPattern:        "/v1/machine/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindMachinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMachinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindMachinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FreeMachine frees a machine
*/
func (a *Client) FreeMachine(params *FreeMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FreeMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "freeMachine",
		Method:             "DELETE",
		PathPattern:        "/v1/machine/{id}/free",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FreeMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FreeMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetProvisioningEventContainer gets the current machine provisioning event container
*/
func (a *Client) GetProvisioningEventContainer(params *GetProvisioningEventContainerParams, authInfo runtime.ClientAuthInfoWriter) (*GetProvisioningEventContainerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvisioningEventContainerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProvisioningEventContainer",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProvisioningEventContainerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProvisioningEventContainerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProvisioningEventContainerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IpmiReport reports IP m i ip addresses leased by a management server for machines
*/
func (a *Client) IpmiReport(params *IpmiReportParams, authInfo runtime.ClientAuthInfoWriter) (*IpmiReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpmiReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipmiReport",
		Method:             "POST",
		PathPattern:        "/v1/machine/ipmi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IpmiReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IpmiReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IpmiReportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListMachines gets all known machines
*/
func (a *Client) ListMachines(params *ListMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*ListMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listMachines",
		Method:             "GET",
		PathPattern:        "/v1/machine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListMachinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMachinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListMachinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachineBios boots machine into b i o s
*/
func (a *Client) MachineBios(params *MachineBiosParams, authInfo runtime.ClientAuthInfoWriter) (*MachineBiosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineBiosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineBios",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/bios",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineBiosReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachineBiosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachineBiosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachineDisk boots machine from disk
*/
func (a *Client) MachineDisk(params *MachineDiskParams, authInfo runtime.ClientAuthInfoWriter) (*MachineDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineDiskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineDisk",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/disk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineDiskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachineDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachineDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachineOff sends a power off to the machine
*/
func (a *Client) MachineOff(params *MachineOffParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineOffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineOff",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/off",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineOffReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachineOffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachineOffDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachineOn sends a power on to the machine
*/
func (a *Client) MachineOn(params *MachineOnParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineOnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineOn",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/on",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineOnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachineOnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachineOnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachinePxe boots machine from p x e
*/
func (a *Client) MachinePxe(params *MachinePxeParams, authInfo runtime.ClientAuthInfoWriter) (*MachinePxeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachinePxeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machinePxe",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/pxe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachinePxeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachinePxeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachinePxeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MachineReset sends a reset to the machine
*/
func (a *Client) MachineReset(params *MachineResetParams, authInfo runtime.ClientAuthInfoWriter) (*MachineResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineReset",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineResetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MachineResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MachineResetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RegisterMachine registers a machine
*/
func (a *Client) RegisterMachine(params *RegisterMachineParams, authInfo runtime.ClientAuthInfoWriter) (*RegisterMachineOK, *RegisterMachineCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RegisterMachineOK:
		return value, nil, nil
	case *RegisterMachineCreated:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterMachineDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReinstallMachine reinstalls this machine
*/
func (a *Client) ReinstallMachine(params *ReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter) (*ReinstallMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReinstallMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reinstallMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/reinstall",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReinstallMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReinstallMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReinstallMachineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SetChassisIdentifyLEDState sets the state of a chassis identify l e d
*/
func (a *Client) SetChassisIdentifyLEDState(params *SetChassisIdentifyLEDStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetChassisIdentifyLEDStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetChassisIdentifyLEDStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setChassisIdentifyLEDState",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/chassis-identify-led-state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetChassisIdentifyLEDStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetChassisIdentifyLEDStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetChassisIdentifyLEDStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SetMachineState sets the state of a machine
*/
func (a *Client) SetMachineState(params *SetMachineStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetMachineStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMachineStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setMachineState",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetMachineStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetMachineStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetMachineStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
