package metalgo

import (
	v1 "github.com/metal-stack/masterdata-api/api/rest/v1"
	"github.com/metal-stack/metal-go/api/client/project"
	"github.com/metal-stack/metal-go/api/models"
)

// ProjectListResponse is the response of a ProjectList action
type ProjectListResponse struct {
	Project []*models.V1ProjectResponse
}

// ProjectGetResponse is the response of a ProjectGet action
type ProjectGetResponse struct {
	Project *models.V1ProjectResponse
}

// ProjectFindRequest is the find request struct
type ProjectFindRequest struct {
	ID     string
	Name   string
	Tenant string
}

// ProjectList return all projects
func (d *Driver) ProjectList() (*ProjectListResponse, error) {
	response := &ProjectListResponse{}
	listProjects := project.NewListProjectsParams()
	resp, err := d.project.ListProjects(listProjects, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// ProjectFind return projects by given findRequest
func (d *Driver) ProjectFind(pfr v1.ProjectFindRequest) (*ProjectListResponse, error) {
	response := &ProjectListResponse{}
	findProjects := project.NewFindProjectsParams()
	body := &models.V1ProjectFindRequest{}
	if pfr.Id != nil {
		body.ID = *pfr.Id
	}
	if pfr.Name != nil {
		body.Name = *pfr.Name
	}
	if pfr.TenantId != nil {
		body.TenantID = *pfr.TenantId
	}
	findProjects.Body = body
	resp, err := d.project.FindProjects(findProjects, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// ProjectGet return a Project
func (d *Driver) ProjectGet(projectID string) (*ProjectGetResponse, error) {
	response := &ProjectGetResponse{}
	getProject := project.NewFindProjectParams()
	getProject.ID = projectID
	resp, err := d.project.FindProject(getProject, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// ProjectCreate a new Project
func (d *Driver) ProjectCreate(pcr v1.ProjectCreateRequest) (*ProjectGetResponse, error) {
	response := &ProjectGetResponse{}
	createProject := project.NewCreateProjectParams()
	body := &models.V1ProjectCreateRequest{
		Description: pcr.Description,
		Meta: &models.V1Meta{
			ID:          pcr.Meta.Id,
			Kind:        pcr.Meta.Kind,
			Apiversion:  pcr.Meta.Apiversion,
			Annotations: pcr.Meta.Annotations,
			Labels:      pcr.Meta.Labels,
			Version:     pcr.Meta.Version,
		},
		Name:     pcr.Name,
		TenantID: pcr.TenantId,
		Quotas:   ToV1QuotaSet(pcr.Quotas),
	}
	createProject.Body = body
	resp, err := d.project.CreateProject(createProject, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// ProjectUpdate update a Project
func (d *Driver) ProjectUpdate(pur v1.ProjectUpdateRequest) (*ProjectGetResponse, error) {
	response := &ProjectGetResponse{}
	updateProject := project.NewUpdateProjectParams()
	body := &models.V1ProjectUpdateRequest{
		Description: pur.Description,
		Meta: &models.V1Meta{
			ID:          pur.Meta.Id,
			Kind:        pur.Meta.Kind,
			Apiversion:  pur.Meta.Apiversion,
			Annotations: pur.Meta.Annotations,
			Labels:      pur.Meta.Labels,
			Version:     pur.Meta.Version,
		},
		Name:     pur.Name,
		TenantID: pur.TenantId,
		Quotas:   ToV1QuotaSet(pur.Quotas),
	}
	updateProject.Body = body
	resp, err := d.project.UpdateProject(updateProject, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// ProjectDelete delete a Project
func (d *Driver) ProjectDelete(projectID string) (*ProjectGetResponse, error) {
	response := &ProjectGetResponse{}
	getProject := project.NewDeleteProjectParams()
	getProject.ID = projectID
	resp, err := d.project.DeleteProject(getProject, d.auth)
	if err != nil {
		return response, err
	}
	response.Project = resp.Payload
	return response, nil
}

// Helper

// ToV1QuotaSet convert a masterdata-api v1 QuotaSet to a swagger V1Quotaset
func ToV1QuotaSet(q *v1.QuotaSet) *models.V1QuotaSet {
	if q == nil {
		return nil
	}
	return &models.V1QuotaSet{
		Cluster: ToV1Quota(q.Cluster),
		Machine: ToV1Quota(q.Machine),
		IP:      ToV1Quota(q.Ip),
		Project: ToV1Quota(q.Project),
	}
}

// ToV1Quota convert a masterdata-api v1 Quota to a swagger V1Quota
func ToV1Quota(q *v1.Quota) *models.V1Quota {
	if q == nil {
		return nil
	}
	if q.Quota == nil {
		return nil
	}
	return &models.V1Quota{
		Quota: *q.Quota,
	}
}
