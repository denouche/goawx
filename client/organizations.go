package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// OrganizationsService implements awx organizations apis.
type OrganizationsService struct {
	CrudImpl[Organization]
}

// ListOrganizationsResponse represents `ListOrganizations` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organization `json:"results"`
}

// DisAssociateGalaxyCredentials remove Credentials form an awx job template
func (p *OrganizationsService) DisAssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("%s%d/galaxy_credentials/", organizationsAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields = []string{"id", "disassociate"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateGalaxyCredentials adding credentials to Organization.
func (p *OrganizationsService) AssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)

	endpoint := fmt.Sprintf("%s%d/galaxy_credentials/", organizationsAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
