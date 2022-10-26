package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// TeamService implements awx teams apis.
type TeamService struct {
	client *Client
}

// ListTeamsResponse represents `ListTeams` endpoint response.
type ListTeamsResponse struct {
	Pagination
	Results []*Team `json:"results"`
}

type ListTeamRoleEntitlementsResponse struct {
	Pagination
	Results []*ApplyRole `json:"results"`
}

type ListTeamUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

const teamsAPIEndpoint = "/api/v2/teams/"

// ListTeams shows list of awx teams.
func (p *TeamService) ListTeams(params map[string]string) ([]*Team, *ListTeamsResponse, error) {
	result := new(ListTeamsResponse)
	resp, err := p.client.Requester.GetJSON(teamsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (p *TeamService) ListTeamRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListTeamRoleEntitlementsResponse, error) {
	result := new(ListTeamRoleEntitlementsResponse)
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

func (p *TeamService) GetTeamUsers(id int, allPages bool, params map[string]string) ([]*User, *ListTeamUsersResponse, error) {
	endpoint := fmt.Sprintf("%s%d/users/", teamsAPIEndpoint, id)
	if allPages {
		users, err := p.getAllTeamUsersPages(endpoint, params)
		if err != nil {
			return nil, nil, err
		}
		return users, nil, nil
	} else {
		result := new(ListTeamUsersResponse)
		resp, err := p.client.Requester.GetJSON(endpoint, result, params)
		if err != nil {
			return nil, result, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, result, err
		}
		return result.Results, result, nil
	}
}

// GetTeamByID shows the details of a team.
func (p *TeamService) GetTeamByID(id int, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateTeam creates an awx team.
func (p *TeamService) CreateTeam(data map[string]interface{}, params map[string]string) (*Team, error) {
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Team)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if team exists and return proper error

	resp, err := p.client.Requester.PostJSON(teamsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTeam update an awx Team.
func (p *TeamService) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
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

func (p *TeamService) UpdateTeamRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	result := new(interface{})
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
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

// DeleteTeam delete an awx Team.
func (p *TeamService) DeleteTeam(id int) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
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
func (p *TeamService) getAllTeamUsersPages(firstURL string, params map[string]string) ([]*User, error) {
	results := make([]*User, 0)
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

		result := new(ListTeamUsersResponse)
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
