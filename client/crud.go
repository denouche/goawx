package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type Crud[T any] interface {
	GetById(id int, params map[string]string) (*T, error)
	Create(data map[string]interface{}, params map[string]string) (*T, error)
	ListAll(params map[string]string) ([]*T, error)
	ListPartial(params map[string]string) ([]*T, *ListResponse[T], error)
	DeleteById(id int, params map[string]string) error
	UpdateById(id int, data map[string]interface{}, params map[string]string) (*T, error)
}

type CrudImpl[T any] struct {
	endpoint string
	client   *Client
}

type ListResponse[T any] struct {
	Pagination
	Results []*T `json:"results"`
}

func (crud CrudImpl[T]) GetById(id int, params map[string]string) (*T, error) {
	result := new(T)
	endpoint := fmt.Sprintf("%s%d", crud.endpoint, id)
	resp, err := crud.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (crud CrudImpl[T]) Create(data map[string]interface{}, params map[string]string) (*T, error) {
	result := new(T)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := crud.client.Requester.PostJSON(crud.endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (crud CrudImpl[T]) ListAll(params map[string]string) ([]*T, error) {
	results, err := crud.getAllPages(crud.endpoint, params)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (crud CrudImpl[T]) DeleteById(id int, params map[string]string) error {
	endpoint := fmt.Sprintf("%s%d", crud.endpoint, id)
	resp, err := crud.client.Requester.Delete(endpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

func (crud CrudImpl[T]) UpdateById(id int, data map[string]interface{}, params map[string]string) (*T, error) {
	result := new(T)
	endpoint := fmt.Sprintf("%s%d", crud.endpoint, id)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := crud.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (crud CrudImpl[T]) getAllPages(firstURL string, params map[string]string) ([]*T, error) {
	results := make([]*T, 0)
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

		result := new(ListResponse[T])
		resp, err := crud.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
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

func (crud CrudImpl[T]) ListPartial(params map[string]string) ([]*T, *ListResponse[T], error) {
	result := new(ListResponse[T])
	resp, err := crud.client.Requester.GetJSON(crud.endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
