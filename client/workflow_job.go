package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Enum of job statuses.
const (
	WorkflowJobStatusNew        = "new"
	WorkflowJobStatusPending    = "pending"
	WorkflowJobStatusWaiting    = "waiting"
	WorkflowJobStatusRunning    = "running"
	WorkflowJobStatusSuccessful = "successful"
	WorkflowJobStatusFailed     = "failed"
	WorkflowJobStatusError      = "error"
	WorkflowJobStatusCanceled   = "canceled"
)

// WorkflowJobService implements awx job apis.
type WorkflowJobService struct {
	client *Client
}

// JobEventsResponse represents `JobEvents` endpoint response.
type WokflowJobEventsResponse struct {
	Pagination
	Results []JobEvent `json:"results"`
}

// CancelJobResponse represents `CancelJob` endpoint response.
type CancelWorkflowJobResponse struct {
	Detail string `json:"detail"`
}

const WorkflowJobAPIEndpoint = "/api/v2/workflow_jobs/"

// GetWorkflowJob shows the details of a job.
func (j *WorkflowJobService) GetWorkflowJob(id int, params map[string]string) (*Job, error) {
	result := new(Job)
	endpoint := fmt.Sprintf("%s%d/", WorkflowJobAPIEndpoint, id)
	resp, err := j.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CancelJob cancels a job.
func (j *WorkflowJobService) CancelWorkflowJob(id int, data map[string]interface{}, params map[string]string) (*CancelJobResponse, error) {
	result := new(CancelJobResponse)
	endpoint := fmt.Sprintf("%s%d/cancel/", WorkflowJobAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := j.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// RelaunchJob relaunch a job.
func (j *WorkflowJobService) RelaunchWorkflowJob(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("%s%d/relaunch/", WorkflowJobAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := j.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
