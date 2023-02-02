package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// UserService implements awx Users apis.
type UserService struct {
	CrudImpl[User]
}

// ListUsersResponse represents `ListUsers` endpoint response.
type ListUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

type ListUsersEntitlementsResponse struct {
	Pagination
	Results []*ApplyRole `json:"results"`
}

func (u *UserService) ListUserRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListUsersEntitlementsResponse, error) {
	result := new(ListUsersEntitlementsResponse)
	endpoint := fmt.Sprintf("%s%d/roles/", usersAPIEndpoint, id)
	resp, err := u.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

func (u *UserService) UpdateUserRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	result := new(interface{})
	endpoint := fmt.Sprintf("%s%d/roles/", usersAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
