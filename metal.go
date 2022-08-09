package metalgo

import (
	"fmt"
	"net/url"
	"time"

	"github.com/metal-stack/metal-go/api/client/vpn"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/metal-stack/security"

	"github.com/metal-stack/metal-go/api/client"
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
)

const (
	defaultHMACAuthType = "Metal-Admin"
)

type Client interface {
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
	VPN() vpn.ClientService
	Version() version.ClientService
}

// Driver holds the client connection to the metal api
type Driver struct {
	c *client.MetalAPI

	bearer       string
	hmacAuthType string
	hmac         *security.HMACAuth
}

// Option for config of Driver
type option func(driver *Driver)

// AuthType sets the authType for HMAC-Auth
func AuthType(authType string) option {
	return func(driver *Driver) {
		driver.hmacAuthType = authType
	}
}

// NewDriver Create a new Driver for Metal to given url. Either bearer OR hmacKey must be set.
// The returned *Driver will be deprecated at some point in time, please migrate to use the Client interface instead.
func NewDriver(baseURL, bearer, hmacKey string, options ...option) (Client, *Driver, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, nil, err
	}
	if parsedURL.Host == "" {
		return nil, nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", baseURL)
	}

	transport := httptransport.New(parsedURL.Host, parsedURL.Path, []string{parsedURL.Scheme})
	c := client.New(transport, strfmt.Default)

	driver := &Driver{
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

	// TODO: remove *Driver return at some point in the future in order to get rid off the handwritten wrappers
	// see: https://github.com/metal-stack/metal-go/issues/33
	return driver, driver, nil
}

func (d *Driver) auther(rq runtime.ClientRequest, rg strfmt.Registry) error {
	if d.hmac != nil {
		d.hmac.AddAuthToClientRequest(rq, time.Now())
	} else if d.bearer != "" {
		security.AddUserTokenToClientRequest(rq, d.bearer)
	}
	return nil
}

func StrDeref(s *string) string {
	if s == nil {
		var res string
		return res
	}
	return *s
}

func BoolDeref(b *bool) bool {
	if b == nil {
		var res bool
		return res
	}
	return *b
}

func Int64Deref(i *int64) int64 {
	if i == nil {
		var res int64
		return res
	}
	return *i
}

func (d *Driver) Filesystemlayout() filesystemlayout.ClientService {
	return d.c.Filesystemlayout
}
func (d *Driver) Firewall() firewall.ClientService {
	return d.c.Firewall
}
func (d *Driver) Firmware() firmware.ClientService {
	return d.c.Firmware
}
func (d *Driver) Health() health.ClientService {
	return d.c.Health
}
func (d *Driver) Image() image.ClientService {
	return d.c.Image
}
func (d *Driver) IP() ip.ClientService {
	return d.c.IP
}
func (d *Driver) Machine() machine.ClientService {
	return d.c.Machine
}
func (d *Driver) Network() network.ClientService {
	return d.c.Network
}
func (d *Driver) Partition() partition.ClientService {
	return d.c.Partition
}
func (d *Driver) Project() project.ClientService {
	return d.c.Project
}
func (d *Driver) Size() size.ClientService {
	return d.c.Size
}
func (d *Driver) Sizeimageconstraint() sizeimageconstraint.ClientService {
	return d.c.Sizeimageconstraint
}
func (d *Driver) VPN() vpn.ClientService {
	return d.c.Vpn
}
func (d *Driver) SwitchOperations() switch_operations.ClientService {
	return d.c.SwitchOperations
}
func (d *Driver) Tenant() tenant.ClientService {
	return d.c.Tenant
}
func (d *Driver) User() user.ClientService {
	return d.c.User
}
func (d *Driver) Version() version.ClientService {
	return d.c.Version
}
