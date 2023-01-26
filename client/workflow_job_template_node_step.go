package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// WorkflowJobTemplateNodeStepService implements awx job template nodes apis.
type WorkflowJobTemplateNodeStepService interface {
	ListWorkflowJobTemplateSuccessNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error)
	CreateWorkflowJobTemplateSuccessNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error)
	ListWorkflowJobTemplateFailureNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error)
	CreateWorkflowJobTemplateFailureNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error)
	ListWorkflowJobTemplateAlwaysNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error)
	CreateWorkflowJobTemplateAlwaysNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error)
}

func (jt *awx) ListWorkflowJobTemplateSuccessNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	return jt.listWorkflowJobTemplateNodeSteps(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/success_nodes/"), params)
}
func (jt *awx) CreateWorkflowJobTemplateSuccessNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	return jt.createWorkflowJobTemplateNodeStep(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/success_nodes/"), data, params)

}
func (jt *awx) ListWorkflowJobTemplateFailureNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	return jt.listWorkflowJobTemplateNodeSteps(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/failure_nodes/"), params)

}
func (jt *awx) CreateWorkflowJobTemplateFailureNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	return jt.createWorkflowJobTemplateNodeStep(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/failure_nodes/"), data, params)

}
func (jt *awx) ListWorkflowJobTemplateAlwaysNodeSteps(id int, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	return jt.listWorkflowJobTemplateNodeSteps(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/always_nodes/"), params)

}
func (jt *awx) CreateWorkflowJobTemplateAlwaysNodeStep(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	return jt.createWorkflowJobTemplateNodeStep(id, fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/always_nodes/"), data, params)

}

// ListWorkflowJobTemplateNodeSteps shows a list of job templates nodes.
func (jt *awx) listWorkflowJobTemplateNodeSteps(id int, endpoint string, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	workflowJobTemplateNodesActionEndpoint := fmt.Sprintf(endpoint, id)
	return fetchWorkflowJobTemplateNode(jt.client, params, workflowJobTemplateNodesActionEndpoint)
}

func fetchWorkflowJobTemplateNode(client *Client, params map[string]string, workflowJobTemplateNodesActionEndpoint string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	result := new(ListWorkflowJobTemplateNodesResponse)
	resp, err := client.Requester.GetJSON(workflowJobTemplateNodesActionEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func createWorkflowJobTemplateNode(client *Client, data map[string]interface{}, params map[string]string, workflowJobTemplateNodesActionEndpoint string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"unified_job_template", "identifier"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := client.Requester.PostJSON(workflowJobTemplateNodesActionEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	log.Printf("Created ID %v", result.ID)
	return result, nil
}

// CreateWorkflowJobTemplateNodeStep will be create a template node for a existing node
func (jt *awx) createWorkflowJobTemplateNodeStep(id int, endpoint string, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	workflowJobTemplateNodesActionEndpoint := fmt.Sprintf(endpoint, id)
	return createWorkflowJobTemplateNode(jt.client, data, params, workflowJobTemplateNodesActionEndpoint)
}
