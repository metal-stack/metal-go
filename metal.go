package metalgo

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/client"
	"github.com/metal-stack/metal-go/api/client/audit"
	"github.com/metal-stack/metal-go/api/client/filesystemlayout"
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
	"github.com/metal-stack/metal-go/api/client/sizeimageconstraint"
	"github.com/metal-stack/metal-go/api/client/switch_operations"
	"github.com/metal-stack/metal-go/api/client/tenant"
	"github.com/metal-stack/metal-go/api/client/user"
	"github.com/metal-stack/metal-go/api/client/version"
	"github.com/metal-stack/metal-go/api/client/vpn"
	"github.com/metal-stack/security"
)

const (
	defaultHMACAuthType = "Metal-Admin"
)

type Client interface {
	Audit() audit.ClientService
	Filesystemlayout() filesystemlayout.ClientService
	Firewall() firewall.ClientService
	Firmware() firmware.ClientService
	Health() health.ClientService
	Image() image.ClientService
	IP() ip.ClientService
	Machine() machine.ClientService
	Network() network.ClientService
	Partition() partition.ClientService
	Project() project.ClientService
	Size() size.ClientService
	Sizeimageconstraint() sizeimageconstraint.ClientService
	SwitchOperations() switch_operations.ClientService
	Tenant() tenant.ClientService
	User() user.ClientService
	Version() version.ClientService
	VPN() vpn.ClientService
}

// driver holds the client connection to the metal api
type driver struct {
	c *client.MetalAPI

	bearer       string
	hmacAuthType string
	hmac         *security.HMACAuth
}

// Option for config of Driver
type option func(driver *driver)
type clientOption func(transport *httptransport.Runtime)

// AuthType sets the authType for HMAC-Auth
func AuthType(authType string) clientOption {
	return func(transport *httptransport.Runtime) {
		transport.DefaultAuthentication = authType
	}
}

func BearerToken(bearer string) clientOption {
	return func(transport *httptransport.Runtime) {
		if bearer != "" {
			transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
				security.AddUserTokenToClientRequest(request, bearer)
				return nil
			})
		}
	}
}

func HMACAuth(hmac string, authType string) clientOption {
	return func(transport *httptransport.Runtime) {
		if hmac != "" {
			auth := security.NewHMACAuth(authType, []byte(hmac))

			transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
				auth.AddAuthToClientRequest(request, time.Now())
				return nil
			})
		}
	}
}

func Transport(transport http.RoundTripper) clientOption {
	return func(runtime *httptransport.Runtime) {
		runtime.Transport = transport
	}
}

// NewDriver Create a new Driver for Metal to given url. Either bearer OR hmacKey must be set.
// Deprecated: Use NewClient instead
func NewDriver(baseURL, bearer, hmacKey string, options ...option) (Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", baseURL)
	}

	transport := httptransport.New(parsedURL.Host, parsedURL.Path, []string{parsedURL.Scheme})
	c := client.New(transport, strfmt.Default)

	driver := &driver{
		c:            c,
		bearer:       bearer,
		hmacAuthType: defaultHMACAuthType,
	}

	for _, opt := range options {
		opt(driver)
	}

	if hmacKey != "" {
		auth := security.NewHMACAuth(driver.hmacAuthType, []byte(hmacKey))
		driver.hmac = &auth
	}

	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(driver.auther)

	return driver, nil
}

func NewClient(baseURL string, options ...clientOption) (Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", baseURL)
	}

	transport := httptransport.New(parsedURL.Host, parsedURL.Path, []string{parsedURL.Scheme})

	for _, opt := range options {
		opt(transport)
	}

	return &driver{
		c: client.New(transport, strfmt.Default),
	}, nil
}

func (d *driver) auther(rq runtime.ClientRequest, rg strfmt.Registry) error {
	if d.hmac != nil {
		d.hmac.AddAuthToClientRequest(rq, time.Now())
	} else if d.bearer != "" {
		security.AddUserTokenToClientRequest(rq, d.bearer)
	}
	return nil
}

func (d *driver) Audit() audit.ClientService {
	return d.c.Audit
}
func (d *driver) Filesystemlayout() filesystemlayout.ClientService {
	return d.c.Filesystemlayout
}
func (d *driver) Firewall() firewall.ClientService {
	return d.c.Firewall
}
func (d *driver) Firmware() firmware.ClientService {
	return d.c.Firmware
}
func (d *driver) Health() health.ClientService {
	return d.c.Health
}
func (d *driver) Image() image.ClientService {
	return d.c.Image
}
func (d *driver) IP() ip.ClientService {
	return d.c.IP
}
func (d *driver) Machine() machine.ClientService {
	return d.c.Machine
}
func (d *driver) Network() network.ClientService {
	return d.c.Network
}
func (d *driver) Partition() partition.ClientService {
	return d.c.Partition
}
func (d *driver) Project() project.ClientService {
	return d.c.Project
}
func (d *driver) Size() size.ClientService {
	return d.c.Size
}
func (d *driver) Sizeimageconstraint() sizeimageconstraint.ClientService {
	return d.c.Sizeimageconstraint
}
func (d *driver) SwitchOperations() switch_operations.ClientService {
	return d.c.SwitchOperations
}
func (d *driver) Tenant() tenant.ClientService {
	return d.c.Tenant
}
func (d *driver) User() user.ClientService {
	return d.c.User
}
func (d *driver) Version() version.ClientService {
	return d.c.Version
}
func (d *driver) VPN() vpn.ClientService {
	return d.c.Vpn
}
