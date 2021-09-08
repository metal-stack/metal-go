package metalgo

import (
	"github.com/metal-stack/metal-go/client/operations"
	"github.com/metal-stack/metal-go/models"
)

// TenantListResponse is the response of a TenantList action
type TenantListResponse struct {
	Tenant []*models.V1TenantResponse
}

// TenantGetResponse is the response of a TenantGet action
type TenantGetResponse struct {
	Tenant *models.V1TenantResponse
}

// TenantFindRequest is the find request struct
type TenantFindRequest struct {
	ID     string
	Name   string
	Tenant string
}

// TenantList return all Tenants
func (d *Driver) TenantList() (*TenantListResponse, error) {
	response := &TenantListResponse{}
	listTenants := operations.NewListTenantsParams()
	resp, err := d.Client.ListTenants(listTenants, nil)
	if err != nil {
		return response, err
	}
	response.Tenant = resp.Payload
	return response, nil
}

// TenantGet return a Tenant
func (d *Driver) TenantGet(TenantID string) (*TenantGetResponse, error) {
	response := &TenantGetResponse{}
	getTenant := operations.NewGetTenantParams()
	getTenant.ID = TenantID
	resp, err := d.Client.GetTenant(getTenant, nil)
	if err != nil {
		return response, err
	}
	response.Tenant = resp.Payload
	return response, nil
}
