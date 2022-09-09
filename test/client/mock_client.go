package client

import (
	"testing"

	metalgo "github.com/metal-stack/metal-go"

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

	filesystemlayoutmocks "github.com/metal-stack/metal-go/test/mocks/filesystemlayout"
	firewallmocks "github.com/metal-stack/metal-go/test/mocks/firewall"
	firmwaremocks "github.com/metal-stack/metal-go/test/mocks/firmware"
	healthmocks "github.com/metal-stack/metal-go/test/mocks/health"
	imagemocks "github.com/metal-stack/metal-go/test/mocks/image"
	ipmocks "github.com/metal-stack/metal-go/test/mocks/ip"
	machinemocks "github.com/metal-stack/metal-go/test/mocks/machine"
	networkmocks "github.com/metal-stack/metal-go/test/mocks/network"
	partitionmocks "github.com/metal-stack/metal-go/test/mocks/partition"
	projectmocks "github.com/metal-stack/metal-go/test/mocks/project"
	sizemocks "github.com/metal-stack/metal-go/test/mocks/size"
	sizeimageconstraintmocks "github.com/metal-stack/metal-go/test/mocks/sizeimageconstraint"
	switchoperationsmocks "github.com/metal-stack/metal-go/test/mocks/switch_operations"
	tenantmocks "github.com/metal-stack/metal-go/test/mocks/tenant"
	usermocks "github.com/metal-stack/metal-go/test/mocks/user"
	versionmocks "github.com/metal-stack/metal-go/test/mocks/version"
	vpnmocks "github.com/metal-stack/metal-go/test/mocks/vpn"

	"github.com/stretchr/testify/mock"
)

type MetalMockFns struct {
	Filesystemlayout    func(mock *mock.Mock)
	Firewall            func(mock *mock.Mock)
	Firmware            func(mock *mock.Mock)
	Health              func(mock *mock.Mock)
	Image               func(mock *mock.Mock)
	IP                  func(mock *mock.Mock)
	Machine             func(mock *mock.Mock)
	Network             func(mock *mock.Mock)
	Partition           func(mock *mock.Mock)
	Project             func(mock *mock.Mock)
	Size                func(mock *mock.Mock)
	SizeImageConstraint func(mock *mock.Mock)
	Switch              func(mock *mock.Mock)
	Tenant              func(mock *mock.Mock)
	User                func(mock *mock.Mock)
	Version             func(mock *mock.Mock)
	VPN                 func(mock *mock.Mock)
}

type MetalMockClient struct {
	filesystemlayout    filesystemlayoutmocks.ClientService
	firewall            firewallmocks.ClientService
	firmware            firmwaremocks.ClientService
	health              healthmocks.ClientService
	image               imagemocks.ClientService
	ip                  ipmocks.ClientService
	machine             machinemocks.ClientService
	network             networkmocks.ClientService
	partition           partitionmocks.ClientService
	project             projectmocks.ClientService
	size                sizemocks.ClientService
	sizeimageconstraint sizeimageconstraintmocks.ClientService
	switchoperations    switchoperationsmocks.ClientService
	tenant              tenantmocks.ClientService
	user                usermocks.ClientService
	version             versionmocks.ClientService
	vpn                 vpnmocks.ClientService
}

func NewMetalMockClient(mockFns *MetalMockFns) (*MetalMockClient, metalgo.Client) {
	client := &MetalMockClient{
		filesystemlayout:    filesystemlayoutmocks.ClientService{},
		firewall:            firewallmocks.ClientService{},
		firmware:            firmwaremocks.ClientService{},
		health:              healthmocks.ClientService{},
		image:               imagemocks.ClientService{},
		ip:                  ipmocks.ClientService{},
		machine:             machinemocks.ClientService{},
		network:             networkmocks.ClientService{},
		partition:           partitionmocks.ClientService{},
		project:             projectmocks.ClientService{},
		size:                sizemocks.ClientService{},
		sizeimageconstraint: sizeimageconstraintmocks.ClientService{},
		switchoperations:    switchoperationsmocks.ClientService{},
		tenant:              tenantmocks.ClientService{},
		user:                usermocks.ClientService{},
		version:             versionmocks.ClientService{},
		vpn:                 vpnmocks.ClientService{},
	}

	if mockFns == nil {
		return client, client
	}

	if mockFns.Filesystemlayout != nil {
		mockFns.Filesystemlayout(&client.filesystemlayout.Mock)
	}
	if mockFns.Firewall != nil {
		mockFns.Firewall(&client.firewall.Mock)
	}
	if mockFns.Firmware != nil {
		mockFns.Firmware(&client.firmware.Mock)
	}
	if mockFns.Health != nil {
		mockFns.Health(&client.health.Mock)
	}
	if mockFns.Image != nil {
		mockFns.Image(&client.image.Mock)
	}
	if mockFns.IP != nil {
		mockFns.IP(&client.ip.Mock)
	}
	if mockFns.Machine != nil {
		mockFns.Machine(&client.machine.Mock)
	}
	if mockFns.Network != nil {
		mockFns.Network(&client.network.Mock)
	}
	if mockFns.Partition != nil {
		mockFns.Partition(&client.partition.Mock)
	}
	if mockFns.Project != nil {
		mockFns.Project(&client.project.Mock)
	}
	if mockFns.Size != nil {
		mockFns.Size(&client.size.Mock)
	}
	if mockFns.SizeImageConstraint != nil {
		mockFns.SizeImageConstraint(&client.sizeimageconstraint.Mock)
	}
	if mockFns.Switch != nil {
		mockFns.Switch(&client.switchoperations.Mock)
	}
	if mockFns.Tenant != nil {
		mockFns.Tenant(&client.tenant.Mock)
	}
	if mockFns.User != nil {
		mockFns.User(&client.user.Mock)
	}
	if mockFns.Version != nil {
		mockFns.Version(&client.version.Mock)
	}
	if mockFns.VPN != nil {
		mockFns.VPN(&client.vpn.Mock)
	}

	return client, client
}

func (c *MetalMockClient) Filesystemlayout() filesystemlayout.ClientService {
	return &c.filesystemlayout
}

func (c *MetalMockClient) Firewall() firewall.ClientService {
	return &c.firewall
}

func (c *MetalMockClient) Firmware() firmware.ClientService {
	return &c.firmware
}

func (c *MetalMockClient) Health() health.ClientService {
	return &c.health
}

func (c *MetalMockClient) Image() image.ClientService {
	return &c.image
}

func (c *MetalMockClient) IP() ip.ClientService {
	return &c.ip
}

func (c *MetalMockClient) Machine() machine.ClientService {
	return &c.machine
}

func (c *MetalMockClient) Network() network.ClientService {
	return &c.network
}

func (c *MetalMockClient) Partition() partition.ClientService {
	return &c.partition
}

func (c *MetalMockClient) Project() project.ClientService {
	return &c.project
}

func (c *MetalMockClient) Size() size.ClientService {
	return &c.size
}

func (c *MetalMockClient) Sizeimageconstraint() sizeimageconstraint.ClientService {
	return &c.sizeimageconstraint
}

func (c *MetalMockClient) SwitchOperations() switch_operations.ClientService {
	return &c.switchoperations
}

func (c *MetalMockClient) Tenant() tenant.ClientService {
	return &c.tenant
}

func (c *MetalMockClient) User() user.ClientService {
	return &c.user
}

func (c *MetalMockClient) Version() version.ClientService {
	return &c.version
}

func (c *MetalMockClient) VPN() vpn.ClientService {
	return &c.vpn
}

func (c *MetalMockClient) AssertExpectations(t *testing.T) {
	_ = c.filesystemlayout.AssertExpectations(t)
	_ = c.firewall.AssertExpectations(t)
	_ = c.firmware.AssertExpectations(t)
	_ = c.health.AssertExpectations(t)
	_ = c.image.AssertExpectations(t)
	_ = c.ip.AssertExpectations(t)
	_ = c.machine.AssertExpectations(t)
	_ = c.network.AssertExpectations(t)
	_ = c.partition.AssertExpectations(t)
	_ = c.project.AssertExpectations(t)
	_ = c.size.AssertExpectations(t)
	_ = c.sizeimageconstraint.AssertExpectations(t)
	_ = c.switchoperations.AssertExpectations(t)
	_ = c.tenant.AssertExpectations(t)
	_ = c.user.AssertExpectations(t)
	_ = c.version.AssertExpectations(t)
	_ = c.vpn.AssertExpectations(t)
}
