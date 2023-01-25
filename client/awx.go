package awx

import (
	"fmt"
	"net/http"
)

// This variable is mandatory and to be populated for creating services API
var mandatoryFields []string

type AWX interface {
	PingService
	ApplicationService
	ExecutionEnvironmentsService
	InventoriesService
	JobService
	JobTemplateService
	JobTemplateNotificationTemplatesService
	ProjectService
	ProjectUpdatesService
	UserService
	GroupService
	HostService
	CredentialsService
	CredentialTypeService
	CredentialInputSourceService
	InventorySourcesService
	InventoryGroupService
	InstanceGroupsService
	NotificationTemplatesService
	OrganizationsService
	SchedulesService
	SettingService
	TeamService
	WorkflowJobTemplateScheduleService
	WorkflowJobTemplateService
	WorkflowJobTemplateNodeService
	WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNotificationTemplatesService
}

// AWX represents awx api endpoints with services, and using
// client to communicate with awx server.
type awx struct {
	client *Client
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
func NewAWX(baseURL, userName, passwd string, client *http.Client) (AWX, error) {
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
	_, err := newAWX.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

// NewAWXToken creates an AWX handler with token support.
func NewAWXToken(baseURL, token string, client *http.Client) (AWX, error) {
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
	_, err := newAWX.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

func newAWX(c *Client) AWX {
	return &awx{
		client: c,
	}
}
