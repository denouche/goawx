package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const workflowJobTemplateNotificationTemplatesAPIEndpoint = "/api/v2/workflow_job_templates/%d/notification_templates_%s/"

// WorkflowJobTemplateNotificationTemplatesService implements awx job template nodes apis.
type WorkflowJobTemplateNotificationTemplatesService interface {
	AssociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	AssociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	AssociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	AssociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	DisassociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	DisassociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	DisassociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
	DisassociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error)
}

func (s *awx) associateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID int, notificationTemplateID int, typ string) (*NotificationTemplate, error) {
	result := new(NotificationTemplate)

	data := map[string]interface{}{
		"id": notificationTemplateID,
	}

	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	endpoint := fmt.Sprintf(workflowJobTemplateNotificationTemplatesAPIEndpoint, jobTemplateID, typ)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateWorkflowJobTemplateNotificationTemplatesError will associate an error notification_template for a job_template
func (s *awx) AssociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.associateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "error")
}

// AssociateWorkflowJobTemplateNotificationTemplatesSuccess will associate a success notification_template for a job_template
func (s *awx) AssociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.associateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "success")
}

// AssociateWorkflowJobTemplateNotificationTemplatesStarted will associate a started notification_template for a job_template
func (s *awx) AssociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.associateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "started")
}

// AssociateWorkflowJobTemplateNotificationTemplatesApprovals will associate an approval notification_template for a job_template
func (s *awx) AssociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.associateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "approvals")
}

func (s *awx) disassociateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID int, notificationTemplateID int, typ string) (*NotificationTemplate, error) {
	result := new(NotificationTemplate)

	data := map[string]interface{}{
		"id":           notificationTemplateID,
		"disassociate": true,
	}
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	endpoint := fmt.Sprintf(workflowJobTemplateNotificationTemplatesAPIEndpoint, jobTemplateID, typ)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisassociateWorkflowJobTemplateNotificationTemplatesError will disassociate an error notification_template for a job_template
func (s *awx) DisassociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.disassociateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "error")
}

// DisassociateWorkflowJobTemplateNotificationTemplatesSuccess will disassociate a success notification_template for a job_template
func (s *awx) DisassociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.disassociateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "success")
}

// DisassociateWorkflowJobTemplateNotificationTemplatesStarted will disassociate a started notification_template for a job_template
func (s *awx) DisassociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.disassociateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "started")
}

// DisassociateWorkflowJobTemplateNotificationTemplatesApprovals will disassociate an approval notification_template for a job_template
func (s *awx) DisassociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*NotificationTemplate, error) {
	return s.disassociateWorkflowJobTemplateNotificationTemplatesForType(jobTemplateID, notificationTemplateID, "approvals")
}
