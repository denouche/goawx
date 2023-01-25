package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// InstanceGroupsService implements awx execution environments apis.
type InstanceGroupsService interface {
	ListInstanceGroups(params map[string]string) ([]*InstanceGroup, *ListInstanceGroupsResponse, error)
	GetInstanceGroupByID(id int, params map[string]string) (*InstanceGroup, error)
	CreateInstanceGroup(data map[string]interface{}, params map[string]string) (*InstanceGroup, error)
	UpdateInstanceGroup(id int, data map[string]interface{}, params map[string]string) (*InstanceGroup, error)
	DeleteInstanceGroup(id int) (*InstanceGroup, error)
}

// ListInstanceGroupsResponse represents `ListInstanceGroups` endpoint response.
type ListInstanceGroupsResponse struct {
	Pagination
	Results []*InstanceGroup `json:"results"`
}

const InstanceGroupsAPIEndpoint = "/api/v2/instance_groups/"

// ListInstanceGroups shows list of awx execution environments.
func (p *awx) ListInstanceGroups(params map[string]string) ([]*InstanceGroup, *ListInstanceGroupsResponse, error) {
	result := new(ListInstanceGroupsResponse)
	resp, err := p.client.Requester.GetJSON(InstanceGroupsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// GetInstanceGroupByID shows the details of a InstanceGroup.
func (p *awx) GetInstanceGroupByID(id int, params map[string]string) (*InstanceGroup, error) {
	result := new(InstanceGroup)
	endpoint := fmt.Sprintf("%s%d/", InstanceGroupsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateInstanceGroup creates an awx InstanceGroup.
func (p *awx) CreateInstanceGroup(data map[string]interface{}, params map[string]string) (*InstanceGroup, error) {
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(InstanceGroup)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(InstanceGroupsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateInstanceGroup update an awx InstanceGroup.
func (p *awx) UpdateInstanceGroup(id int, data map[string]interface{}, params map[string]string) (*InstanceGroup, error) {
	result := new(InstanceGroup)
	endpoint := fmt.Sprintf("%s%d", InstanceGroupsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteInstanceGroup delete an awx InstanceGroup.
func (p *awx) DeleteInstanceGroup(id int) (*InstanceGroup, error) {
	result := new(InstanceGroup)
	endpoint := fmt.Sprintf("%s%d", InstanceGroupsAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
