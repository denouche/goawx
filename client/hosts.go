package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// HostService implements awx Hosts apis.
type HostService struct {
	CrudImpl[Host]
}

// AssociateGroup implement the awx group association request
type AssociateGroup struct {
	ID        int  `json:"id"`
	Associate bool `json:"associate"`
}

// ListHostsResponse represents `ListHosts` endpoint response.
type ListHostsResponse struct {
	Pagination
	Results []*Host `json:"results"`
}

// AssociateGroup update an awx Host
func (h *HostService) AssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("%s%d/groups/", hostsAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociateGroup update an awx Host
func (h *HostService) DisAssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("%s%d/groups/", hostsAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
