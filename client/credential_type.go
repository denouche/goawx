package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CredentialTypeService interface {
	ListCredentialTypes(params map[string]string) ([]*CredentialType,
		*ListCredentialTypeResponse, error)
	CreateCredentialType(data map[string]interface{}, params map[string]string) (*CredentialType, error)
	GetCredentialTypeByID(id int, params map[string]string) (*CredentialType, error)
	UpdateCredentialTypeByID(id int, data map[string]interface{}, params map[string]string) (*CredentialType, error)
	DeleteCredentialTypeByID(id int, params map[string]string) error
}

type ListCredentialTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

const credentialTypesAPIEndpoint = "/api/v2/credential_types/"

func (cs *awx) ListCredentialTypes(params map[string]string) ([]*CredentialType,
	*ListCredentialTypeResponse, error) {
	result := new(ListCredentialTypeResponse)
	resp, err := cs.client.Requester.GetJSON(credentialTypesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (cs *awx) CreateCredentialType(data map[string]interface{}, params map[string]string) (*CredentialType, error) {
	result := new(CredentialType)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PostJSON(credentialTypesAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *awx) GetCredentialTypeByID(id int, params map[string]string) (*CredentialType, error) {
	result := new(CredentialType)
	endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)
	resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *awx) UpdateCredentialTypeByID(id int, data map[string]interface{}, params map[string]string) (*CredentialType, error) {
	result := new(CredentialType)
	endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PutJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *awx) DeleteCredentialTypeByID(id int, params map[string]string) error {
	endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)
	resp, err := cs.client.Requester.Delete(endpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}
