package metalgo

import (
	"fmt"

	"github.com/metal-pod/metal-go/api/client/ip"
	"github.com/metal-pod/metal-go/api/client/network"
	"github.com/metal-pod/metal-go/api/models"
)

const TagIPClusterID = "cluster.metal-pod.io/clusterid"

// NetworkGetResponse contains the network get result
type NetworkGetResponse struct {
	Network *models.V1NetworkResponse
}

// NetworkListResponse is the response of a NetworkList action
type NetworkListResponse struct {
	Networks []*models.V1NetworkResponse
}

// NetworkAcquireRequest is the request to acquire a new private network
type NetworkAcquireRequest struct {
	// a description for this entity
	Description string `json:"description,omitempty"`

	// the readable name
	Name string `json:"name,omitempty"`

	// the partition this network belongs to, TODO: can be empty ?
	// Required: true
	PartitionID string `json:"partitionid"`

	// the project this network belongs to, can be empty if globally available.
	// Required: true
	ProjectID string `json:"projectid,omitempty"`

	// A map of key/value pairs treated as labels.
	// Required: false
	Labels map[string]string `json:"labels"`
}

// NetworkCreateRequest is the request for create a new network
type NetworkCreateRequest struct {
	ID *string `json:"id"`
	// a description for this entity
	Description string `json:"description,omitempty"`

	// the readable name
	Name string `json:"name,omitempty"`

	// if set to true, packets leaving this network get masqueraded behind interface ip.
	// Required: true
	Nat bool `json:"nat"`

	// the partition this network belongs to, TODO: can be empty ?
	// Required: true
	Partitionid string `json:"partitionid"`

	// the prefixes of this network, required.
	// Required: true
	Prefixes []string `json:"prefixes"`

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes"`

	// if set to true, this network acts a supernetwork for private networks
	// Required: true
	PrivateSuper bool `json:"privatesuper"`

	// the project this network belongs to, can be empty if globally available.
	// Required: true
	Projectid string `json:"projectid"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay bool `json:"underlay"`

	// the vrf this network is associated with
	Vrf int64 `json:"vrf,omitempty"`

	// if set to true, the network can share the vrf with other networks
	// Required: false
	VrfShared bool `json:"vrfshared,omitempty"`
}

// NetworkDetailResponse is the response of a NetworkList action
type NetworkDetailResponse struct {
	Network *models.V1NetworkResponse
}

// NetworkUpdateRequest is the request to update the Network
type NetworkUpdateRequest struct {
	// the network id for this update request.
	Networkid string `json:"networkid"`
	// Prefix the prefix to add/remove
	Prefix string
}

// IPUpdateRequest is the request to update an IP
type IPUpdateRequest struct {
	// the ip address for this ip update request.
	IPAddress string `json:"ipaddress"`
	// a description for this entity
	Description string `json:"description,omitempty"`
	// the readable name
	Name string `json:"name,omitempty"`
	// the type of the ip
	Type string `json:"type,omitempty"`
	// tags for the ip
	Tags []string `json:"tags,omitempty"`
}

// IPAssociateRequest is the request to associate an IP with a cluster
type IPAssociateRequest struct {
	// the ip address for this ip update request.
	IPAddress string `json:"ipaddress"`
	// the cluster id to associate the ip address with.
	ClusterID string `json:"clusterid"`
	// tags to add to the ip
	Tags []string `json:"tags,omitempty"`
}

// IPDeassociateRequest is the request to deassociate an IP from a cluster
type IPDeassociateRequest struct {
	// the ip address for this ip update request.
	IPAddress string `json:"ipaddress"`
	// the cluster id to deassociate the ip address with.
	ClusterID string `json:"clusterid"`
	// tags to remove from the ip
	Tags []string `json:"tags,omitempty"`
}

// IPListResponse is the response when ips are listed
type IPListResponse struct {
	IPs []*models.V1IPResponse
}

// IPAcquireRequest is the request to acquire an IP
type IPAcquireRequest struct {

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the readable name
	Name string `json:"name,omitempty"`

	// the network this ip acquire request belongs to, required.
	// Required: true
	Networkid string `json:"networkid"`

	// the project this ip acquire request belongs to, required.
	// Required: true
	Projectid string `json:"projectid"`

	// the machine this ip acquire request belongs to
	Machineid *string `json:"machineid"`

	// the cluster this ip acquire request belongs to
	Clusterid *string `json:"clusterid"`

	// SpecificIP tries to acquire this ip.
	// Required: false
	SpecificIP string `json:"specificip"`

	// the type of the ip
	Type string `json:"type,omitempty"`

	// tags for the ip
	Tags []string `json:"tags,omitempty"`
}

// NetworkFindRequest contains criteria for a network listing
type NetworkFindRequest struct {
	ID                  *string
	Name                *string
	PartitionID         *string
	ProjectID           *string
	Prefixes            []string
	DestinationPrefixes []string
	Nat                 *bool
	PrivateSuper        *bool
	Underlay            *bool
	Vrf                 *int64
	ParentNetworkID     *string
}

// IPFindRequest contains criteria for a ip listing
type IPFindRequest struct {
	IPAddress        *string
	ProjectID        *string
	ParentPrefixCidr *string
	NetworkID        *string
	MachineID        *string
	ClusterID        *string
	Type             *string
	Tags             []string
}

// IPDetailResponse is the response to an IP detail request.
type IPDetailResponse struct {
	IP *models.V1IPResponse
}

// NetworkGet returns the network with the given ID
func (d *Driver) NetworkGet(id string) (*NetworkGetResponse, error) {
	findNetwork := network.NewFindNetworkParams()
	findNetwork.ID = id

	response := &NetworkGetResponse{}
	resp, err := d.network.FindNetwork(findNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload

	return response, nil
}

// NetworkList returns all networks
func (d *Driver) NetworkList() (*NetworkListResponse, error) {
	response := &NetworkListResponse{}
	listNetworks := network.NewListNetworksParams()
	resp, err := d.network.ListNetworks(listNetworks, d.auth)
	if err != nil {
		return response, err
	}
	response.Networks = resp.Payload
	return response, nil
}

// NetworkFind returns all networks that match given properties
func (d *Driver) NetworkFind(nfr *NetworkFindRequest) (*NetworkListResponse, error) {
	if nfr == nil {
		return d.NetworkList()
	}

	response := &NetworkListResponse{}
	var err error
	var resp *network.FindNetworksOK

	findNetworks := network.NewFindNetworksParams()
	req := &models.V1NetworkFindRequest{
		ID:                  nfr.ID,
		Name:                nfr.Name,
		Partitionid:         nfr.PartitionID,
		Projectid:           nfr.ProjectID,
		Prefixes:            nfr.Prefixes,
		Destinationprefixes: nfr.DestinationPrefixes,
		Nat:                 nfr.Nat,
		Privatesuper:        nfr.PrivateSuper,
		Underlay:            nfr.Underlay,
		Vrf:                 nfr.Vrf,
		Parentnetworkid:     nfr.ParentNetworkID,
	}
	findNetworks.SetBody(req)

	resp, err = d.network.FindNetworks(findNetworks, d.auth)
	if err != nil {
		return response, err
	}
	response.Networks = resp.Payload

	return response, nil
}

// NetworkCreate creates a new network
func (d *Driver) NetworkCreate(ncr *NetworkCreateRequest) (*NetworkDetailResponse, error) {
	response := &NetworkDetailResponse{}
	createNetwork := network.NewCreateNetworkParams()

	createRequest := &models.V1NetworkCreateRequest{
		ID:                  ncr.ID,
		Description:         ncr.Description,
		Name:                ncr.Name,
		Nat:                 &ncr.Nat,
		Partitionid:         ncr.Partitionid,
		Prefixes:            ncr.Prefixes,
		Destinationprefixes: ncr.Destinationprefixes,
		Vrf:                 ncr.Vrf,
		Vrfshared:           ncr.VrfShared,
		Privatesuper:        &ncr.PrivateSuper,
		Projectid:           ncr.Projectid,
		Underlay:            &ncr.Underlay,
	}
	createNetwork.SetBody(createRequest)
	resp, err := d.network.CreateNetwork(createNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// NetworkAcquire creates a new network
func (d *Driver) NetworkAcquire(ncr *NetworkAcquireRequest) (*NetworkDetailResponse, error) {
	response := &NetworkDetailResponse{}
	acquireNetwork := network.NewAcquireChildNetworkParams()

	acquireRequest := &models.V1NetworkAcquireRequest{
		Description: ncr.Description,
		Name:        ncr.Name,
		Partitionid: ncr.PartitionID,
		Projectid:   ncr.ProjectID,
		Labels:      ncr.Labels,
	}
	acquireNetwork.SetBody(acquireRequest)
	resp, err := d.network.AcquireChildNetwork(acquireNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// NetworkRelease creates a new network
func (d *Driver) NetworkRelease(id string) (*NetworkDetailResponse, error) {
	response := &NetworkDetailResponse{}
	releaseNetwork := network.NewReleaseChildNetworkParams()

	releaseNetwork.ID = id
	resp, err := d.network.ReleaseChildNetwork(releaseNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// NetworkUpdate creates a new network
func (d *Driver) NetworkUpdate(ncr *NetworkCreateRequest) (*NetworkDetailResponse, error) {
	response := &NetworkDetailResponse{}
	updateNetwork := network.NewUpdateNetworkParams()

	updateRequest := &models.V1NetworkUpdateRequest{
		ID:          ncr.ID,
		Description: ncr.Description,
		Name:        ncr.Name,
		Prefixes:    ncr.Prefixes,
	}
	updateNetwork.SetBody(updateRequest)
	resp, err := d.network.UpdateNetwork(updateNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// NetworkAddPrefix adds a prefix to a network
func (d *Driver) NetworkAddPrefix(nur *NetworkUpdateRequest) (*NetworkDetailResponse, error) {
	old, err := d.NetworkGet(nur.Networkid)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch network: %s to update:%v", nur.Networkid, err)
	}
	oldNetwork := old.Network
	newPrefixes := append(oldNetwork.Prefixes, nur.Prefix)

	response := &NetworkDetailResponse{}
	updateNetwork := network.NewUpdateNetworkParams()
	updateRequest := &models.V1NetworkUpdateRequest{
		ID:       &nur.Networkid,
		Prefixes: newPrefixes,
	}
	updateNetwork.SetBody(updateRequest)
	resp, err := d.network.UpdateNetwork(updateNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// NetworkRemovePrefix removes a prefix from a network
func (d *Driver) NetworkRemovePrefix(nur *NetworkUpdateRequest) (*NetworkDetailResponse, error) {
	old, err := d.NetworkGet(nur.Networkid)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch network: %s to update:%v", nur.Networkid, err)
	}
	oldNetwork := old.Network
	var newPrefixes []string
	for _, p := range oldNetwork.Prefixes {
		if p == nur.Prefix {
			continue
		}
		newPrefixes = append(newPrefixes, p)
	}

	response := &NetworkDetailResponse{}
	updateNetwork := network.NewUpdateNetworkParams()
	updateRequest := &models.V1NetworkUpdateRequest{
		ID:       &nur.Networkid,
		Prefixes: newPrefixes,
	}
	updateNetwork.SetBody(updateRequest)
	resp, err := d.network.UpdateNetwork(updateNetwork, d.auth)
	if err != nil {
		return response, err
	}
	response.Network = resp.Payload
	return response, nil
}

// IPGet gets a given IP
func (d *Driver) IPGet(ipaddress string) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	findIP := ip.NewFindIPParams()
	findIP.ID = ipaddress
	resp, err := d.ip.FindIP(findIP, d.auth)
	if err != nil {
		return response, err
	}
	response.IP = resp.Payload
	return response, nil
}

// IPUpdate updates an IP
func (d *Driver) IPUpdate(iur *IPUpdateRequest) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	updateIP := ip.NewUpdateIPParams()

	updateRequest := &models.V1IPUpdateRequest{
		Ipaddress:   &iur.IPAddress,
		Description: iur.Description,
		Name:        iur.Name,
		Iptype:      &iur.Type,
		Tags:        iur.Tags,
	}
	updateIP.SetBody(updateRequest)
	resp, err := d.ip.UpdateIP(updateIP, d.auth)
	if err != nil {
		return response, err
	}
	response.IP = resp.Payload
	return response, nil
}

// IPAssociate updates an IP and associates it with a cluster
func (d *Driver) IPAssociate(iur *IPAssociateRequest) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	updateIP := ip.NewUpdateIPParams()

	detail, err := d.IPGet(iur.IPAddress)
	if err != nil {
		return nil, err
	}

	ip := detail.IP
	tags := append(ip.Tags, fmt.Sprintf("%s=%s", TagIPClusterID, iur.ClusterID))
	tags = append(tags, iur.Tags...)
	updateRequest := &models.V1IPUpdateRequest{
		Ipaddress: ip.Ipaddress,
		Tags:      tags,
	}
	updateIP.SetBody(updateRequest)
	resp, err := d.ip.UpdateIP(updateIP, d.auth)
	if err != nil {
		return response, err
	}
	response.IP = resp.Payload
	return response, nil
}

// IPDeassociate updates an IP and deassociates it from a cluster
func (d *Driver) IPDeassociate(iur *IPDeassociateRequest) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	updateIP := ip.NewUpdateIPParams()

	detail, err := d.IPGet(iur.IPAddress)
	if err != nil {
		return nil, err
	}

	ip := detail.IP
	ct := fmt.Sprintf("%s=%s", TagIPClusterID, iur.ClusterID)
	tagsToRemove := map[string]interface{}{ct: nil}
	for _, t := range iur.Tags {
		tagsToRemove[t] = nil
	}

	newTags := []string{}
	for _, t := range ip.Tags {
		if _, ok := tagsToRemove[t]; ok {
			continue
		}
		newTags = append(newTags, t)
	}

	updateRequest := &models.V1IPUpdateRequest{
		Ipaddress: ip.Ipaddress,
		Tags:      newTags,
	}
	updateIP.SetBody(updateRequest)
	resp, err := d.ip.UpdateIP(updateIP, d.auth)
	if err != nil {
		return response, err
	}
	response.IP = resp.Payload
	return response, nil
}

// IPList lists all IPs
func (d *Driver) IPList() (*IPListResponse, error) {
	response := &IPListResponse{}
	listIPs := ip.NewListIpsParams()
	resp, err := d.ip.ListIps(listIPs, d.auth)
	if err != nil {
		return response, err
	}
	response.IPs = resp.Payload
	return response, nil
}

// IPFind returns all ips that match given properties
func (d *Driver) IPFind(ifr *IPFindRequest) (*IPListResponse, error) {
	if ifr == nil {
		return d.IPList()
	}

	response := &IPListResponse{}
	var err error
	var resp *ip.FindIpsOK

	findIPs := ip.NewFindIpsParams()
	req := &models.V1IPFindRequest{
		Ipaddress:     ifr.IPAddress,
		Projectid:     ifr.ProjectID,
		Networkprefix: ifr.ParentPrefixCidr,
		Networkid:     ifr.NetworkID,
		Machineid:     ifr.MachineID,
		Clusterid:     ifr.ClusterID,
		Type:          ifr.Type,
		Tags:          ifr.Tags,
	}
	findIPs.SetBody(req)

	resp, err = d.ip.FindIps(findIPs, d.auth)
	if err != nil {
		return response, err
	}
	response.IPs = resp.Payload

	return response, nil
}

// IPAcquire acquires an IP in a network for a project
func (d *Driver) IPAcquire(iar *IPAcquireRequest) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	acquireIPRequest := &models.V1IPAllocateRequest{
		Description: iar.Description,
		Name:        iar.Name,
		Networkid:   &iar.Networkid,
		Projectid:   &iar.Projectid,
		Machineid:   iar.Machineid,
		Clusterid:   iar.Clusterid,
		Iptype:      &iar.Type,
		Tags:        iar.Tags,
	}
	if iar.SpecificIP == "" {
		acquireIP := ip.NewAllocateIPParams()
		acquireIP.SetBody(acquireIPRequest)
		resp, err := d.ip.AllocateIP(acquireIP, d.auth)
		if err != nil {
			return response, err
		}
		response.IP = resp.Payload
	} else {
		acquireIP := ip.NewAllocateSpecificIPParams()
		acquireIP.IP = iar.SpecificIP
		acquireIP.SetBody(acquireIPRequest)
		resp, err := d.ip.AllocateSpecificIP(acquireIP, d.auth)
		if err != nil {
			return response, err
		}
		response.IP = resp.Payload
	}
	return response, nil
}

// IPDelete releases an IP
func (d *Driver) IPDelete(id string) (*IPDetailResponse, error) {
	response := &IPDetailResponse{}
	deleteIP := ip.NewReleaseIPParams()
	deleteIP.ID = id
	resp, err := d.ip.ReleaseIP(deleteIP, d.auth)
	if err != nil {
		return response, err
	}
	response.IP = resp.Payload
	return response, nil
}
