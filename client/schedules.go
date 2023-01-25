package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// SchedulesService implements awx projects apis.
type SchedulesService interface {
	ListSchedule(params map[string]string) ([]*Schedule, *ListSchedulesResponse, error)
	GetScheduleByID(id int, params map[string]string) (*Schedule, error)
	CreateSchedule(data map[string]interface{}, params map[string]string) (*Schedule, error)
	UpdateSchedule(id int, data map[string]interface{}, params map[string]string) (*Schedule, error)
	DeleteSchedule(id int) (*Schedule, error)
}

// ListSchedulesResponse represents `List` endpoint response.
type ListSchedulesResponse struct {
	Pagination
	Results []*Schedule `json:"results"`
}

const schedulesAPIEndpoint = "/api/v2/schedules/"

func (s *awx) ListSchedule(params map[string]string) ([]*Schedule, *ListSchedulesResponse, error) {
	result := new(ListSchedulesResponse)
	resp, err := s.client.Requester.GetJSON(schedulesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// GetByID shows the details of a schedule.
func (s *awx) GetScheduleByID(id int, params map[string]string) (*Schedule, error) {
	result := new(Schedule)
	endpoint := fmt.Sprintf("%s%d/", schedulesAPIEndpoint, id)
	resp, err := s.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Create creates an awx schedule.
func (s *awx) CreateSchedule(data map[string]interface{}, params map[string]string) (*Schedule, error) {
	mandatoryFields = []string{"name", "rrule", "unified_job_template"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Schedule)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PostJSON(schedulesAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Update update an awx schedule.
func (s *awx) UpdateSchedule(id int, data map[string]interface{}, params map[string]string) (*Schedule, error) {
	result := new(Schedule)
	endpoint := fmt.Sprintf("%s%d", schedulesAPIEndpoint, id)
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

// Delete delete an awx schedule.
func (s *awx) DeleteSchedule(id int) (*Schedule, error) {
	result := new(Schedule)
	endpoint := fmt.Sprintf("%s%d", schedulesAPIEndpoint, id)

	resp, err := s.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
