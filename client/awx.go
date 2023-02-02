package awx

import (
	"fmt"
	"net/http"
)

// This variable is mandatory and to be populated for creating services API
var mandatoryFields []string

// AWX represents awx api endpoints with services, and using
// client to communicate with awx server.
type AWX struct {
	client *Client

	ApplicationService                              Crud[Application]
	ExecutionEnvironmentsService                    Crud[ExecutionEnvironment]
	PingService                                     *PingService
	InventoriesService                              Crud[Inventory]
	JobService                                      *JobService
	JobTemplateService                              *JobTemplateService
	JobTemplateNotificationTemplatesService         *JobTemplateNotificationTemplatesService
	ProjectService                                  Crud[Project]
	ProjectUpdatesService                           *ProjectUpdatesService
	UserService                                     *UserService
	GroupService                                    Crud[Groups]
	HostService                                     Crud[Host]
	CredentialsService                              Crud[Credential]
	CredentialTypeService                           Crud[CredentialType]
	CredentialInputSourceService                    Crud[CredentialInputSource]
	InventorySourcesService                         Crud[InventorySource]
	InventoryGroupService                           *InventoryGroupService
	InstanceGroupsService                           Crud[InstanceGroup]
	NotificationTemplatesService                    Crud[NotificationTemplate]
	OrganizationsService                            *OrganizationsService
	ScheduleService                                 Crud[Schedule]
	SettingService                                  *SettingService
	TeamService                                     *TeamService
	WorkflowJobTemplateScheduleService              *WorkflowJobTemplateScheduleService
	WorkflowJobTemplateService                      *WorkflowJobTemplateService
	WorkflowJobTemplateNodeService                  *WorkflowJobTemplateNodeService
	WorkflowJobTemplateNodeAlwaysService            *WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNodeFailureService           *WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNodeSuccessService           *WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNotificationTemplatesService *WorkflowJobTemplateNotificationTemplatesService
}

// Client implement http client.
type Client struct {
	BaseURL   string
	Requester *Requester
}

// CheckResponse do http response check, and return err if not in [200, 300).
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
}

// ValidateParams is to validate the input to use the services.
func ValidateParams(data map[string]interface{}, mandatoryFields []string) (notfound []string, status bool) {
	status = true
	for _, key := range mandatoryFields {
		_, exists := data[key]

		if !exists {
			notfound = append(notfound, key)
			status = false
		}
	}
	return notfound, status
}

// NewAWX news an awx handler with basic auth support, you could customize the http
// transport by passing custom client.
func NewAWX(baseURL, userName, passwd string, client *http.Client) (*AWX, error) {
	r := &Requester{Base: baseURL, Authenticator: &BasicAuth{Username: userName, Password: passwd}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
		Requester: r,
	}

	newAWX := newAWX(awxClient)

	// test the connection and return and error if there's an issue
	_, err := newAWX.PingService.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

// NewAWXToken creates an AWX handler with token support.
func NewAWXToken(baseURL, token string, client *http.Client) (*AWX, error) {
	r := &Requester{Base: baseURL, Authenticator: &TokenAuth{Token: token}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
		Requester: r,
	}

	newAWX := newAWX(awxClient)

	// test the connection and return and error if there's an issue
	_, err := newAWX.PingService.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

func newAWX(c *Client) *AWX {
	return &AWX{
		client: c,

		ApplicationService: CrudImpl[Application]{
			client:   c,
			endpoint: applicationAPIEndpoint,
		},
		ExecutionEnvironmentsService: CrudImpl[ExecutionEnvironment]{
			endpoint: executionEnvironmentsAPIEndpoint,
			client:   c,
		},
		PingService: &PingService{
			client: c,
		},
		InventoriesService: CrudImpl[Inventory]{
			endpoint: inventoriesAPIEndpoint,
			client:   c,
		},
		JobService: &JobService{
			CrudImpl: CrudImpl[Job]{
				client:   c,
				endpoint: jobAPIEndpoint,
			},
		},
		JobTemplateService: &JobTemplateService{
			CrudImpl: CrudImpl[JobTemplate]{
				client:   c,
				endpoint: jobTemplateAPIEndpoint,
			},
		},
		JobTemplateNotificationTemplatesService: &JobTemplateNotificationTemplatesService{
			client: c,
		},
		ProjectService: CrudImpl[Project]{
			endpoint: projectsAPIEndpoint,
			client:   c,
		},
		ProjectUpdatesService: &ProjectUpdatesService{
			client: c,
		},
		UserService: &UserService{
			CrudImpl: CrudImpl[User]{
				client:   c,
				endpoint: usersAPIEndpoint,
			},
		},
		GroupService: CrudImpl[Groups]{
			endpoint: groupsAPIEndpoint,
			client:   c,
		},
		HostService: CrudImpl[Host]{
			client: c,
		},
		CredentialsService: CrudImpl[Credential]{
			endpoint: credentialsAPIEndpoint,
			client:   c,
		},
		CredentialTypeService: CrudImpl[CredentialType]{
			endpoint: credentialTypesAPIEndpoint,
			client:   c,
		},
		CredentialInputSourceService: CrudImpl[CredentialInputSource]{
			endpoint: credentialInputSourceAPIEndpoint,
			client:   c,
		},
		InventorySourcesService: CrudImpl[InventorySource]{
			endpoint: inventorySourcesAPIEndpoint,
			client:   c,
		},
		InventoryGroupService: &InventoryGroupService{
			client: c,
		},
		InstanceGroupsService: CrudImpl[InstanceGroup]{
			endpoint: instanceGroupsAPIEndpoint,
			client:   c,
		},
		NotificationTemplatesService: CrudImpl[NotificationTemplate]{
			endpoint: notificationTemplatesAPIEndpoint,
			client:   c,
		},
		OrganizationsService: &OrganizationsService{
			CrudImpl: CrudImpl[Organization]{
				client:   c,
				endpoint: organizationsAPIEndpoint,
			},
		},
		ScheduleService: CrudImpl[Schedule]{
			client:   c,
			endpoint: schedulesAPIEndpoint,
		},
		SettingService: &SettingService{
			client: c,
		},
		TeamService: &TeamService{
			CrudImpl: CrudImpl[Team]{
				client:   c,
				endpoint: teamsAPIEndpoint,
			},
		},
		WorkflowJobTemplateScheduleService: &WorkflowJobTemplateScheduleService{
			client: c,
		},
		WorkflowJobTemplateService: &WorkflowJobTemplateService{
			CrudImpl: CrudImpl[WorkflowJobTemplate]{
				client:   c,
				endpoint: workflowJobTemplateAPIEndpoint,
			},
		},
		WorkflowJobTemplateNodeService: &WorkflowJobTemplateNodeService{
			CrudImpl: CrudImpl[WorkflowJobTemplateNode]{
				client:   c,
				endpoint: workflowJobTemplateNodeAPIEndpoint,
			},
		},
		WorkflowJobTemplateNodeSuccessService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/success_nodes/"),
			client:   c,
		},
		WorkflowJobTemplateNodeFailureService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/failure_nodes/"),
			client:   c,
		},
		WorkflowJobTemplateNodeAlwaysService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/always_nodes/"),
			client:   c,
		},
		WorkflowJobTemplateNotificationTemplatesService: &WorkflowJobTemplateNotificationTemplatesService{
			client: c,
		},
	}
}
