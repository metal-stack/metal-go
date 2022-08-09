package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/tenant"
	"github.com/metal-stack/metal-go/api/models"
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
	listTenants := tenant.NewListTenantsParams()
	resp, err := d.Tenant().ListTenants(listTenants, nil)
	if err != nil {
		return response, err
	}
	response.Tenant = resp.Payload
	return response, nil
}

// TenantGet return a Tenant
func (d *Driver) TenantGet(tenantID string) (*TenantGetResponse, error) {
	response := &TenantGetResponse{}
	getTenant := tenant.NewGetTenantParams()
	getTenant.ID = tenantID
	resp, err := d.Tenant().GetTenant(getTenant, nil)
	if err != nil {
		return response, err
	}
	response.Tenant = resp.Payload
	return response, nil
}
