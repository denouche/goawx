package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

const inventorySourcesSchedulesAPIEndpoint = "/api/v2/inventory_sources/%d/schedules/"

// InventorySourcesSchedulesService implements awx inventory sources schedules apis.
type InventorySourcesSchedulesService struct {
	client *Client
}

// ListInventorySourcesSchedules shows a list of schedules for a given inventory_source
func (is *InventorySourcesSchedulesService) ListInventorySourcesSchedules(id int, params map[string]string) ([]*Schedule, *ListSchedulesResponse, error) {
	result := new(ListSchedulesResponse)
	resp, err := is.client.Requester.GetJSON(
		fmt.Sprintf(inventorySourcesSchedulesAPIEndpoint, id),
		result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateInventorySourcesSchedule will create a schedule for an existing inventory_source
func (is *InventorySourcesSchedulesService) CreateInventorySourcesSchedule(id int, data map[string]interface{}, params map[string]string) (*Schedule, error) {
	mandatoryFields = []string{"name", "rrule"}
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

	resp, err := is.client.Requester.PostJSON(
		fmt.Sprintf(inventorySourcesSchedulesAPIEndpoint, id),
		bytes.NewReader(payload), result, params,
	)

	log.Println("OK")

	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
