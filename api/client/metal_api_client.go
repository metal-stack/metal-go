// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/client/firewall"
	"github.com/metal-stack/metal-go/api/client/firmware"
	"github.com/metal-stack/metal-go/api/client/health"
	"github.com/metal-stack/metal-go/api/client/image"
	"github.com/metal-stack/metal-go/api/client/ip"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/client/network"
	"github.com/metal-stack/metal-go/api/client/partition"
	"github.com/metal-stack/metal-go/api/client/project"
	"github.com/metal-stack/metal-go/api/client/size"
	"github.com/metal-stack/metal-go/api/client/switch_operations"
	"github.com/metal-stack/metal-go/api/client/tenant"
	"github.com/metal-stack/metal-go/api/client/version"
)

// Default metal API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new metal API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MetalAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new metal API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MetalAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new metal API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MetalAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MetalAPI)
	cli.Transport = transport
	cli.Firewall = firewall.New(transport, formats)
	cli.Firmware = firmware.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Image = image.New(transport, formats)
	cli.IP = ip.New(transport, formats)
	cli.Machine = machine.New(transport, formats)
	cli.Network = network.New(transport, formats)
	cli.Partition = partition.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Size = size.New(transport, formats)
	cli.SwitchOperations = switch_operations.New(transport, formats)
	cli.Tenant = tenant.New(transport, formats)
	cli.Version = version.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// MetalAPI is a client for metal API
type MetalAPI struct {
	Firewall firewall.ClientService

	Firmware firmware.ClientService

	Health health.ClientService

	Image image.ClientService

	IP ip.ClientService

	Machine machine.ClientService

	Network network.ClientService

	Partition partition.ClientService

	Project project.ClientService

	Size size.ClientService

	SwitchOperations switch_operations.ClientService

	Tenant tenant.ClientService

	Version version.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MetalAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Firewall.SetTransport(transport)
	c.Firmware.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Image.SetTransport(transport)
	c.IP.SetTransport(transport)
	c.Machine.SetTransport(transport)
	c.Network.SetTransport(transport)
	c.Partition.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Size.SetTransport(transport)
	c.SwitchOperations.SetTransport(transport)
	c.Tenant.SetTransport(transport)
	c.Version.SetTransport(transport)
}
