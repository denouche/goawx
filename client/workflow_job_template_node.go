package awx

// WorkflowJobTemplateNodeService implements awx job template node apis.
type WorkflowJobTemplateNodeService struct {
	CrudImpl[WorkflowJobTemplateNode]
}

// ListWorkflowJobTemplateNodesResponse represents `ListWorkflowJobTemplateNodes` endpoint response.
type ListWorkflowJobTemplateNodesResponse struct {
	Pagination
	Results []*WorkflowJobTemplateNode `json:"results"`
}

const workflowJobTemplateNodeAPIEndpoint = "/api/v2/workflow_job_template_nodes/"
