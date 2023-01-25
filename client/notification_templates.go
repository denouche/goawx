package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// NotificationTemplatesService implements awx projects apis.
type NotificationTemplatesService interface {
	List(params map[string]string) ([]*NotificationTemplate, *ListNotificationTemplatesResponse, error)
	GetByID(id int, params map[string]string) (*NotificationTemplate, error)
	Create(data map[string]interface{}, params map[string]string) (*NotificationTemplate, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*NotificationTemplate, error)
	Delete(id int) (*NotificationTemplate, error)
}

// ListNotificationTemplatesResponse represents `List` endpoint response.
type ListNotificationTemplatesResponse struct {
	Pagination
	Results []*NotificationTemplate `json:"results"`
}

const notificationTemplatesAPIEndpoint = "/api/v2/notification_templates/"

func (s *awx) List(params map[string]string) ([]*NotificationTemplate, *ListNotificationTemplatesResponse, error) {
	result := new(ListNotificationTemplatesResponse)
	resp, err := s.client.Requester.GetJSON(notificationTemplatesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// GetByID shows the details of a notification_template.
func (s *awx) GetByID(id int, params map[string]string) (*NotificationTemplate, error) {
	result := new(NotificationTemplate)
	endpoint := fmt.Sprintf("%s%d/", notificationTemplatesAPIEndpoint, id)
	resp, err := s.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Create creates an awx notification_template.
func (s *awx) Create(data map[string]interface{}, params map[string]string) (*NotificationTemplate, error) {
	mandatoryFields = []string{"name", "organization", "notification_type"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(NotificationTemplate)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PostJSON(notificationTemplatesAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Update update an awx notification_template.
func (s *awx) Update(id int, data map[string]interface{}, params map[string]string) (*NotificationTemplate, error) {
	result := new(NotificationTemplate)
	endpoint := fmt.Sprintf("%s%d", notificationTemplatesAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Delete delete an awx notification_template.
func (s *awx) Delete(id int) (*NotificationTemplate, error) {
	result := new(NotificationTemplate)
	endpoint := fmt.Sprintf("%s%d", notificationTemplatesAPIEndpoint, id)

	resp, err := s.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
