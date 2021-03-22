package metalgo

import (
	"fmt"
	"github.com/metal-stack/metal-go/api/client/firmware"
	"github.com/metal-stack/metal-go/api/client/tenant"
	"net/url"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/metal-stack/metal-go/api/client"
	"github.com/metal-stack/metal-go/api/client/firewall"
	"github.com/metal-stack/metal-go/api/client/image"
	"github.com/metal-stack/metal-go/api/client/ip"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/client/network"
	"github.com/metal-stack/metal-go/api/client/partition"
	"github.com/metal-stack/metal-go/api/client/project"
	"github.com/metal-stack/metal-go/api/client/size"
	sw "github.com/metal-stack/metal-go/api/client/switch_operations"
	"github.com/metal-stack/security"
)

const (
	defaultHMACAuthType = "Metal-Admin"
)

// Driver holds the client connection to the metal api
type Driver struct {
	client       *client.MetalAPI
	image        image.ClientService
	firmware     firmware.ClientService
	machine      machine.ClientService
	firewall     firewall.ClientService
	partition    partition.ClientService
	project      project.ClientService
	tenant       tenant.ClientService
	size         size.ClientService
	sw           sw.ClientService
	network      network.ClientService
	ip           ip.ClientService
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
func NewDriver(baseURL, bearer, hmacKey string, options ...option) (*Driver, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid url:%s, must be in the form scheme://host[:port]/basepath", baseURL)
	}

	transport := httptransport.New(parsedURL.Host, parsedURL.Path, []string{parsedURL.Scheme})
	c := client.New(transport, strfmt.Default)

	driver := &Driver{
		client:       c,
		firmware:     c.Firmware,
		machine:      c.Machine,
		firewall:     c.Firewall,
		size:         c.Size,
		image:        c.Image,
		project:      c.Project,
		partition:    c.Partition,
		sw:           c.SwitchOperations,
		network:      c.Network,
		ip:           c.IP,
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
