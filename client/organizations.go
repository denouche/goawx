package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// OrganizationsService implements awx organizations apis.
type OrganizationsService interface {
	ListOrganizations(params map[string]string) ([]*Organization, error)
	GetOrganizationsByID(id int, params map[string]string) (*Organization, error)
	CreateOrganization(data map[string]interface{}, params map[string]string) (*Organization, error)
	UpdateOrganization(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
	DeleteOrganization(id int) (*Organization, error)
	DisAssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
	AssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
}

// ListOrganizationsResponse represents `ListOrganizations` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organization `json:"results"`
}

const organizationsAPIEndpoint = "/api/v2/organizations/"

// ListOrganizations shows list of awx organizations.
func (p *awx) ListOrganizations(params map[string]string) ([]*Organization, error) {
	results, err := p.getOrganizationAllPages(organizationsAPIEndpoint, params)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetOrganizationsByID shows the details of a Organization.
func (p *awx) GetOrganizationsByID(id int, params map[string]string) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("%s%d/", organizationsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateOrganization creates an awx Organization.
func (p *awx) CreateOrganization(data map[string]interface{}, params map[string]string) (*Organization, error) {
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Organization)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(organizationsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateOrganization update an awx Organization.
func (p *awx) UpdateOrganization(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("%s%d", organizationsAPIEndpoint, id)
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

// DeleteOrganization delete an awx Organization.
func (p *awx) DeleteOrganization(id int) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("%s%d", organizationsAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociateGalaxyCredentials remove Credentials form an awx job template
func (p *awx) DisAssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
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
func (p *awx) AssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
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

// Must be replaced by a generic function
// But upgrade to version go 1.18 before
func (p *awx) getOrganizationAllPages(firstURL string, params map[string]string) ([]*Organization, error) {
	results := make([]*Organization, 0)
	nextURL := firstURL
	for {
		nextURLParsed, err := url.Parse(nextURL)
		if err != nil {
			return nil, err
		}

		nextURLQueryParams := make(map[string]string)
		for paramName, paramValues := range nextURLParsed.Query() {
			if len(paramValues) > 0 {
				nextURLQueryParams[paramName] = paramValues[0]
			}
		}

		for paramName, paramValue := range params {
			nextURLQueryParams[paramName] = paramValue
		}

		result := new(ListOrganizationsResponse)
		resp, err := p.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		results = append(results, result.Results...)

		if result.Next == nil || result.Next.(string) == "" {
			break
		}
		nextURL = result.Next.(string)
	}
	return results, nil
}
