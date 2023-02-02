package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// TeamService implements awx teams apis.
type TeamService struct {
	CrudImpl[Team]
}

// ListTeamsResponse represents `ListTeams` endpoint response.
type ListTeamsResponse struct {
	Pagination
	Results []*Team `json:"results"`
}

type ListTeamRolesResponse struct {
	Pagination
	Results []*ApplyRole `json:"results"`
}

type ListTeamObjectRolesResponse struct {
	Pagination
	Results []*ObjectRoles `json:"results"`
}

type ListTeamUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

func (t *TeamService) ListTeamRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListTeamRolesResponse, error) {
	result := new(ListTeamRolesResponse)
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

func (t *TeamService) GetTeamObjectRoles(id int, params map[string]string, pagination *PaginationRequest) ([]*ApplyRole, *ListTeamRolesResponse, error) {
	result := new(ListTeamRolesResponse)
	endpoint := fmt.Sprintf("%s%d/object_roles/", teamsAPIEndpoint, id)
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

func (t *TeamService) GetTeamUsers(id int, params map[string]string, pagination *PaginationRequest) ([]*User, *ListTeamUsersResponse, error) {
	endpoint := fmt.Sprintf("%s%d/users/", teamsAPIEndpoint, id)
	if *pagination.AllPages {
		users, err := t.getAllTeamUsersPages(endpoint, params)
		if err != nil {
			return nil, nil, err
		}
		return users, nil, nil
	} else {
		result := new(ListTeamUsersResponse)
		resp, err := t.client.Requester.GetJSON(endpoint, result, params)
		if err != nil {
			return nil, result, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, result, err
		}
		return result.Results, result, nil
	}
}

func (t *TeamService) GetTeamAccessList(id int, params map[string]string, pagination *PaginationRequest) ([]*User, *ListTeamUsersResponse, error) {
	endpoint := fmt.Sprintf("%s%d/access_list/", teamsAPIEndpoint, id)
	if *pagination.AllPages {
		users, err := t.getAllTeamUsersPages(endpoint, params)
		if err != nil {
			return nil, nil, err
		}
		return users, nil, nil
	} else {
		result := new(ListTeamUsersResponse)
		resp, err := t.client.Requester.GetJSON(endpoint, result, params)
		if err != nil {
			return nil, result, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, result, err
		}
		return result.Results, result, nil
	}
}

// AddTeamUser will add the user as member in destination team
func (t *TeamService) AddTeamUser(id int, data map[string]interface{}) error {
	endpoint := fmt.Sprintf("%s%d/users/", teamsAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields = []string{"id", "associate"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	if err := CheckResponse(resp); err != nil {
		return err
	}

	return nil
}

// RemoveTeamUser will remove the user from destination team without deleting the user
func (t *TeamService) RemoveTeamUser(id int, data map[string]interface{}) error {
	endpoint := fmt.Sprintf("%s%d/users/", teamsAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields = []string{"id", "disassociate"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	if err := CheckResponse(resp); err != nil {
		return err
	}

	return nil
}

func (t *TeamService) UpdateTeamRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	result := new(interface{})
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeam delete an awx Team.
func (t *TeamService) DeleteTeam(id int) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)

	resp, err := t.client.Requester.Delete(endpoint, result, nil)
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
func (t *TeamService) getAllTeamUsersPages(firstURL string, params map[string]string) ([]*User, error) {
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
		resp, err := t.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
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
