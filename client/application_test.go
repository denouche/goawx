package awx

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type AuthenticatorMock struct {
}

func (am *AuthenticatorMock) addAuthenticationHeaders(*http.Request) {
	return
}
func Test_awx_CreateApplication(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.CreateApplication(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteApplication(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.DeleteApplication(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetApplicationByID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("{\"name\":\"toto\"}"))
	}))
	defer func() {
		srv.Close()
	}()
	srvErr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte("{\"name\":\"toto\"}"))
	}))
	defer func() {
		srvErr.Close()
	}()
	srvErr100 := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusFound)
		res.Write([]byte("{\"name\":\"toto\"}"))
	}))
	defer func() {
		srvErr100.Close()
	}()
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *Application
		wantErr   bool
		wantError error
	}{
		{
			name: "Application found",
			fields: fields{
				client: &Client{
					BaseURL: srv.URL,
					Requester: &Requester{
						Base:          srv.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srv.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:    &Application{Name: "toto"},
			wantErr: false,
		},
		{
			name: "Error while requesting",
			fields: fields{
				client: &Client{
					BaseURL: srvErr.URL,
					Requester: &Requester{
						Base:          srvErr.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvErr.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Errors:\n- name: []"),
		},
		{
			name: "Response status is not 2XX",
			fields: fields{
				client: &Client{
					BaseURL: srvErr100.URL,
					Requester: &Requester{
						Base:          srvErr100.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvErr100.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Errors:\n- name: []"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.GetApplicationByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ListApplication(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Application
		want1   *ListApplicationResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, got1, err := c.ListApplication(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListApplication() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListApplication() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_UpdateApplication(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.UpdateApplication(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIRequest_SetHeader(t *testing.T) {
	type fields struct {
		Method   string
		Endpoint string
		Payload  io.Reader
		Headers  http.Header
		Suffix   string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *APIRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &APIRequest{
				Method:   tt.fields.Method,
				Endpoint: tt.fields.Endpoint,
				Payload:  tt.fields.Payload,
				Headers:  tt.fields.Headers,
				Suffix:   tt.fields.Suffix,
			}
			if got := ar.SetHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicAuth_addAuthenticationHeaders(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := &BasicAuth{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			ba.addAuthenticationHeaders(tt.args.r)
		})
	}
}

func TestCheckResponse(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("CheckResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAPIRequest(t *testing.T) {
	type args struct {
		method   string
		endpoint string
		payload  io.Reader
	}
	tests := []struct {
		name string
		args args
		want *APIRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIRequest(tt.args.method, tt.args.endpoint, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAWX(t *testing.T) {
	type args struct {
		baseURL  string
		userName string
		passwd   string
		client   *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    AWX
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAWX(tt.args.baseURL, tt.args.userName, tt.args.passwd, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWX() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAWXToken(t *testing.T) {
	type args struct {
		baseURL string
		token   string
		client  *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    AWX
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAWXToken(tt.args.baseURL, tt.args.token, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWXToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWXToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_Delete(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.Delete(tt.args.endpoint, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_Do(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		ar             *APIRequest
		responseStruct interface{}
		options        []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.Do(tt.args.ar, tt.args.responseStruct, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_Get(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.Get(tt.args.endpoint, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_GetJSON(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		responseStruct interface{}
		query          map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.GetJSON(tt.args.endpoint, tt.args.responseStruct, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_PatchJSON(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		payload        io.Reader
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.PatchJSON(tt.args.endpoint, tt.args.payload, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatchJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatchJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_Post(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		payload        io.Reader
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.Post(tt.args.endpoint, tt.args.payload, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_PostJSON(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		payload        io.Reader
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.PostJSON(tt.args.endpoint, tt.args.payload, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_PutJSON(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		endpoint       string
		payload        io.Reader
		responseStruct interface{}
		querystring    map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.PutJSON(tt.args.endpoint, tt.args.payload, tt.args.responseStruct, tt.args.querystring)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_ReadJSONResponse(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		response       *http.Response
		responseStruct interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.ReadJSONResponse(tt.args.response, tt.args.responseStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJSONResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadJSONResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequester_ReadRawResponse(t *testing.T) {
	type fields struct {
		Base          string
		Authenticator Authenticator
		Client        *http.Client
	}
	type args struct {
		response       *http.Response
		responseStruct interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				Base:          tt.fields.Base,
				Authenticator: tt.fields.Authenticator,
				Client:        tt.fields.Client,
			}
			got, err := r.ReadRawResponse(tt.args.response, tt.args.responseStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRawResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRawResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenAuth_addAuthenticationHeaders(t *testing.T) {
	type fields struct {
		Token string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := &TokenAuth{
				Token: tt.fields.Token,
			}
			ta.addAuthenticationHeaders(tt.args.r)
		})
	}
}

func TestValidateParams(t *testing.T) {
	type args struct {
		data            map[string]interface{}
		mandatoryFields []string
	}
	tests := []struct {
		name         string
		args         args
		wantNotfound []string
		wantStatus   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNotfound, gotStatus := ValidateParams(tt.args.data, tt.args.mandatoryFields)
			if !reflect.DeepEqual(gotNotfound, tt.wantNotfound) {
				t.Errorf("ValidateParams() gotNotfound = %v, want %v", gotNotfound, tt.wantNotfound)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ValidateParams() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func Test_awx_AddTeamUser(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id   int
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			if err := t.AddTeamUser(tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t1.Errorf("AddTeamUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_awx_AssociateCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.AssociateCredentials(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateGalaxyCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.AssociateGalaxyCredentials(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateGalaxyCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateGalaxyCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.AssociateGroup(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateJobTemplateNotificationTemplatesError(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.AssociateJobTemplateNotificationTemplatesError(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateJobTemplateNotificationTemplatesError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateJobTemplateNotificationTemplatesError() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateJobTemplateNotificationTemplatesStarted(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.AssociateJobTemplateNotificationTemplatesStarted(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateJobTemplateNotificationTemplatesStarted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateJobTemplateNotificationTemplatesStarted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateJobTemplateNotificationTemplatesSuccess(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.AssociateJobTemplateNotificationTemplatesSuccess(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateJobTemplateNotificationTemplatesSuccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateJobTemplateNotificationTemplatesSuccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateWorkflowJobTemplateNotificationTemplatesApprovals(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.AssociateWorkflowJobTemplateNotificationTemplatesApprovals(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesApprovals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesApprovals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateWorkflowJobTemplateNotificationTemplatesError(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.AssociateWorkflowJobTemplateNotificationTemplatesError(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesError() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateWorkflowJobTemplateNotificationTemplatesStarted(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.AssociateWorkflowJobTemplateNotificationTemplatesStarted(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesStarted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesStarted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_AssociateWorkflowJobTemplateNotificationTemplatesSuccess(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.AssociateWorkflowJobTemplateNotificationTemplatesSuccess(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesSuccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssociateWorkflowJobTemplateNotificationTemplatesSuccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CancelJob(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CancelJobResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &awx{
				client: tt.fields.client,
			}
			got, err := j.CancelJob(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateApplication1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.CreateApplication(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateCredentialInputSource(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialInputSource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.CreateCredentialInputSource(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredentialInputSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCredentialInputSource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateCredentialType(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.CreateCredentialType(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredentialType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCredentialType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Credential
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.CreateCredentials(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateExecutionEnvironment(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateExecutionEnvironment(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &awx{
				client: tt.fields.client,
			}
			got, err := g.CreateGroup(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateHost(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.CreateHost(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateInstanceGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InstanceGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateInstanceGroup(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInstanceGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstanceGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateInventory(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.CreateInventory(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateInventorySource(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InventorySource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.CreateInventorySource(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInventorySource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInventorySource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateJobTemplate(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateNotificationTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.CreateNotificationTemplate(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNotificationTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNotificationTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateOrganization(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateOrganization(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrganization() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateProject(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateProject(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateSchedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.CreateSchedule(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateTeam(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.CreateTeam(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreateTeam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateUser(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, err := u.CreateUser(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplate(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplateAlwaysNodeStep(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplateAlwaysNodeStep(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplateAlwaysNodeStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplateAlwaysNodeStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplateFailureNodeStep(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplateFailureNodeStep(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplateFailureNodeStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplateFailureNodeStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplateNode(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplateNode(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplateNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplateSchedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplateSchedule(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplateSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplateSchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateWorkflowJobTemplateSuccessNodeStep(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.CreateWorkflowJobTemplateSuccessNodeStep(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkflowJobTemplateSuccessNodeStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkflowJobTemplateSuccessNodeStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteApplication1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.DeleteApplication(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteCredentialInputSourceByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			if err := cs.DeleteCredentialInputSourceByID(tt.args.id, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCredentialInputSourceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_awx_DeleteCredentialTypeByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			if err := cs.DeleteCredentialTypeByID(tt.args.id, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCredentialTypeByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_awx_DeleteCredentialsByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			if err := cs.DeleteCredentialsByID(tt.args.id, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCredentialsByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_awx_DeleteExecutionEnvironment(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteExecutionEnvironment(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &awx{
				client: tt.fields.client,
			}
			got, err := g.DeleteGroup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteHost(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.DeleteHost(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteInstanceGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InstanceGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteInstanceGroup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInstanceGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInstanceGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteInventory(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.DeleteInventory(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteInventorySource(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InventorySource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.DeleteInventorySource(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInventorySource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInventorySource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DeleteJobTemplate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteNotificationTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DeleteNotificationTemplate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteNotificationTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteNotificationTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteOrganization(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteOrganization(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteOrganization() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteProject(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteProject(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteSchedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DeleteSchedule(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteSettings(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Setting
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteSettings(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSettings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteTeam(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.DeleteTeam(tt.args.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("DeleteTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DeleteTeam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteUser(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, err := u.DeleteUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteWorkflowJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DeleteWorkflowJobTemplate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWorkflowJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteWorkflowJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteWorkflowJobTemplateNode(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DeleteWorkflowJobTemplateNode(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWorkflowJobTemplateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteWorkflowJobTemplateNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisAssociateCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DisAssociateCredentials(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisAssociateCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisAssociateCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisAssociateGalaxyCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DisAssociateGalaxyCredentials(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisAssociateGalaxyCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisAssociateGalaxyCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisAssociateGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.DisAssociateGroup(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisAssociateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisAssociateGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateJobTemplateNotificationTemplatesError(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DisassociateJobTemplateNotificationTemplatesError(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesError() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateJobTemplateNotificationTemplatesStarted(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DisassociateJobTemplateNotificationTemplatesStarted(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesStarted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesStarted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateJobTemplateNotificationTemplatesSuccess(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.DisassociateJobTemplateNotificationTemplatesSuccess(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesSuccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateJobTemplateNotificationTemplatesSuccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateWorkflowJobTemplateNotificationTemplatesApprovals(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DisassociateWorkflowJobTemplateNotificationTemplatesApprovals(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesApprovals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesApprovals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateWorkflowJobTemplateNotificationTemplatesError(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DisassociateWorkflowJobTemplateNotificationTemplatesError(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesError() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateWorkflowJobTemplateNotificationTemplatesStarted(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DisassociateWorkflowJobTemplateNotificationTemplatesStarted(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesStarted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesStarted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DisassociateWorkflowJobTemplateNotificationTemplatesSuccess(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.DisassociateWorkflowJobTemplateNotificationTemplatesSuccess(tt.args.jobTemplateID, tt.args.notificationTemplateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetApplicationByID1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.GetApplicationByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetCredentialInputSourceByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialInputSource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.GetCredentialInputSourceByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialInputSourceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialInputSourceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetCredentialTypeByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.GetCredentialTypeByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialTypeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetCredentialsByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Credential
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.GetCredentialsByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialsByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetExecutionEnvironmentByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetExecutionEnvironmentByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExecutionEnvironmentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExecutionEnvironmentByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetGroupByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &awx{
				client: tt.fields.client,
			}
			got, err := g.GetGroupByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetHostByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.GetHostByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHostByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetHostSummaries(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []HostSummary
		want1   *HostSummariesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &awx{
				client: tt.fields.client,
			}
			got, got1, err := j.GetHostSummaries(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHostSummaries() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetHostSummaries() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_GetInstanceGroupByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InstanceGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetInstanceGroupByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInstanceGroupByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstanceGroupByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetInventory(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.GetInventory(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetInventoryByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.GetInventoryByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventoryByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetInventorySource(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InventorySource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.GetInventorySource(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventorySource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventorySource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetInventorySourceByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InventorySource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.GetInventorySourceByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventorySourceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventorySourceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetJob(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &awx{
				client: tt.fields.client,
			}
			got, err := j.GetJob(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetJobEvents(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []JobEvent
		want1   *JobEventsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &awx{
				client: tt.fields.client,
			}
			got, got1, err := j.GetJobEvents(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJobEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJobEvents() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetJobEvents() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_GetJobTemplateByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.GetJobTemplateByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJobTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJobTemplateByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetNotificationTemplateByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.GetNotificationTemplateByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotificationTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotificationTemplateByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetOrganizationsByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetOrganizationsByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrganizationsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrganizationsByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetProjectByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetProjectByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetScheduleByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.GetScheduleByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetScheduleByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScheduleByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetSettingsBySlug(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		slug   string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Setting
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetSettingsBySlug(tt.args.slug, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSettingsBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSettingsBySlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetTeamAccessList(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id         int
		params     map[string]string
		pagination *PaginationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*User
		want1   *ListTeamUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, got1, err := t.GetTeamAccessList(tt.args.id, tt.args.params, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTeamAccessList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTeamAccessList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("GetTeamAccessList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_GetTeamByID(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.GetTeamByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTeamByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTeamByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetTeamObjectRoles(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id         int
		params     map[string]string
		pagination *PaginationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ApplyRole
		want1   *ListTeamRolesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, got1, err := t.GetTeamObjectRoles(tt.args.id, tt.args.params, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTeamObjectRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTeamObjectRoles() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("GetTeamObjectRoles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_GetTeamUsers(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id         int
		params     map[string]string
		pagination *PaginationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*User
		want1   *ListTeamUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, got1, err := t.GetTeamUsers(tt.args.id, tt.args.params, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTeamUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTeamUsers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("GetTeamUsers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_GetUserByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, err := u.GetUserByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetWorkflowJobTemplateByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.GetWorkflowJobTemplateByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkflowJobTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkflowJobTemplateByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_GetWorkflowJobTemplateNodeByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.GetWorkflowJobTemplateNodeByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkflowJobTemplateNodeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkflowJobTemplateNodeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_LaunchJob(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobLaunch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.LaunchJob(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("LaunchJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LaunchJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_LaunchWorkflow(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobLaunch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.LaunchWorkflow(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("LaunchWorkflow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LaunchWorkflow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ListApplication1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Application
		want1   *ListApplicationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, got1, err := c.ListApplication(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListApplication() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListApplication() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListCredentialInputSources(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*CredentialInputSource
		want1   *ListCredentialInputSourceResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, got1, err := cs.ListCredentialInputSources(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCredentialInputSources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCredentialInputSources() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListCredentialInputSources() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListCredentialTypes(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*CredentialType
		want1   *ListCredentialTypeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, got1, err := cs.ListCredentialTypes(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCredentialTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCredentialTypes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListCredentialTypes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListCredentials(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Credential
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.ListCredentials(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ListExecutionEnvironments(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ExecutionEnvironment
		want1   *ListExecutionEnvironmentsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, got1, err := p.ListExecutionEnvironments(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListExecutionEnvironments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListExecutionEnvironments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListExecutionEnvironments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListGroups(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Group
		want1   *ListGroupsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &awx{
				client: tt.fields.client,
			}
			got, got1, err := g.ListGroups(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListGroups() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListGroups() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListHosts(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Host
		want1   *ListHostsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, got1, err := h.ListHosts(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListHosts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListHosts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListInstanceGroups(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*InstanceGroup
		want1   *ListInstanceGroupsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, got1, err := p.ListInstanceGroups(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListInstanceGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInstanceGroups() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListInstanceGroups() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListInventories(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Inventory
		want1   *ListInventoriesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, got1, err := i.ListInventories(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListInventories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInventories() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListInventories() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListInventoryGroups(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Group
		want1   *ListGroupsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, got1, err := i.ListInventoryGroups(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListInventoryGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInventoryGroups() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListInventoryGroups() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListInventorySources(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*InventorySource
		want1   *ListInventorySourcesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, got1, err := i.ListInventorySources(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListInventorySources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInventorySources() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListInventorySources() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListJobTemplates(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*JobTemplate
		want1   *ListJobTemplatesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListJobTemplates(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListJobTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListJobTemplates() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListJobTemplates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListNotificationTemplates(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*NotificationTemplate
		want1   *ListNotificationTemplatesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, got1, err := s.ListNotificationTemplates(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListNotificationTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNotificationTemplates() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListNotificationTemplates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListOrganizations(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.ListOrganizations(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListOrganizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOrganizations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ListProjects(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Project
		want1   *ListProjectsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, got1, err := p.ListProjects(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListProjects() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListProjects() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListSchedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Schedule
		want1   *ListSchedulesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, got1, err := s.ListSchedule(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSchedule() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListSchedule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListSettings(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*SettingSummary
		want1   *ListSettingsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, got1, err := p.ListSettings(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSettings() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListSettings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListTeamRoleEntitlements(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ApplyRole
		want1   *ListTeamRolesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, got1, err := t.ListTeamRoleEntitlements(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ListTeamRoleEntitlements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ListTeamRoleEntitlements() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("ListTeamRoleEntitlements() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListTeams(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Team
		want1   *ListTeamsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, got1, err := t.ListTeams(tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ListTeams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ListTeams() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("ListTeams() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListUserRoleEntitlements(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ApplyRole
		want1   *ListUsersEntitlementsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, got1, err := u.ListUserRoleEntitlements(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUserRoleEntitlements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUserRoleEntitlements() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListUserRoleEntitlements() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListUsers(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*User
		want1   *ListUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, got1, err := u.ListUsers(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUsers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListUsers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplateAlwaysNodeSteps(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplateAlwaysNodeSteps(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplateAlwaysNodeSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplateAlwaysNodeSteps() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplateAlwaysNodeSteps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplateFailureNodeSteps(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplateFailureNodeSteps(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplateFailureNodeSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplateFailureNodeSteps() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplateFailureNodeSteps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplateNodes(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplateNodes(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplateNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplateNodes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplateNodes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplateSchedules(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Schedule
		want1   *ListSchedulesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplateSchedules(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplateSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplateSchedules() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplateSchedules() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplateSuccessNodeSteps(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplateSuccessNodeSteps(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplateSuccessNodeSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplateSuccessNodeSteps() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplateSuccessNodeSteps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_ListWorkflowJobTemplates(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplate
		want1   *ListWorkflowJobTemplatesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.ListWorkflowJobTemplates(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWorkflowJobTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListWorkflowJobTemplates() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListWorkflowJobTemplates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_awx_Ping(t *testing.T) {
	type fields struct {
		client *Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Ping
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ProjectUpdateCancel(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ProjectUpdateCancel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.ProjectUpdateCancel(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectUpdateCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectUpdateCancel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ProjectUpdateGet(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.ProjectUpdateGet(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectUpdateGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectUpdateGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_RelaunchJob(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobLaunch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &awx{
				client: tt.fields.client,
			}
			got, err := j.RelaunchJob(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelaunchJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelaunchJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_RemoveTeamUser(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id   int
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			if err := t.RemoveTeamUser(tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t1.Errorf("RemoveTeamUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_awx_UpdateApplication1(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.UpdateApplication(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateCredentialInputSourceByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialInputSource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.UpdateCredentialInputSourceByID(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCredentialInputSourceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCredentialInputSourceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateCredentialTypeByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CredentialType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.UpdateCredentialTypeByID(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCredentialTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCredentialTypeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateCredentialsByID(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Credential
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.UpdateCredentialsByID(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCredentialsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCredentialsByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateExecutionEnvironment(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateExecutionEnvironment(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &awx{
				client: tt.fields.client,
			}
			got, err := g.UpdateGroup(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateHost(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Host
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.UpdateHost(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateInstanceGroup(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InstanceGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateInstanceGroup(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateInstanceGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateInstanceGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateInventory(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.UpdateInventory(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateInventorySource(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *InventorySource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &awx{
				client: tt.fields.client,
			}
			got, err := i.UpdateInventorySource(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateInventorySource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateInventorySource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.UpdateJobTemplate(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateNotificationTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.UpdateNotificationTemplate(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateNotificationTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateNotificationTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateOrganization(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateOrganization(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrganization() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateProject(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateProject(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateSchedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.UpdateSchedule(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateSettings(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		slug   string
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Setting
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateSettings(tt.args.slug, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSettings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateTeam(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.UpdateTeam(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdateTeam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateTeamRoleEntitlement(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.UpdateTeamRoleEntitlement(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateTeamRoleEntitlement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdateTeamRoleEntitlement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateUser(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, err := u.UpdateUser(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateUserRoleEntitlement(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &awx{
				client: tt.fields.client,
			}
			got, err := u.UpdateUserRoleEntitlement(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserRoleEntitlement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUserRoleEntitlement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateWorkflowJobTemplate(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.UpdateWorkflowJobTemplate(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWorkflowJobTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateWorkflowJobTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateWorkflowJobTemplateNode(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.UpdateWorkflowJobTemplateNode(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWorkflowJobTemplateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateWorkflowJobTemplateNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_associateJobTemplateNotificationTemplatesForType(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
		typ                    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.associateJobTemplateNotificationTemplatesForType(tt.args.jobTemplateID, tt.args.notificationTemplateID, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("associateJobTemplateNotificationTemplatesForType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("associateJobTemplateNotificationTemplatesForType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_associateWorkflowJobTemplateNotificationTemplatesForType(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
		typ                    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.associateWorkflowJobTemplateNotificationTemplatesForType(tt.args.jobTemplateID, tt.args.notificationTemplateID, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("associateWorkflowJobTemplateNotificationTemplatesForType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("associateWorkflowJobTemplateNotificationTemplatesForType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_createWorkflowJobTemplateNodeStep(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id       int
		endpoint string
		data     map[string]interface{}
		params   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.createWorkflowJobTemplateNodeStep(tt.args.id, tt.args.endpoint, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("createWorkflowJobTemplateNodeStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createWorkflowJobTemplateNodeStep() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_disassociateJobTemplateNotificationTemplatesForType(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
		typ                    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, err := jt.disassociateJobTemplateNotificationTemplatesForType(tt.args.jobTemplateID, tt.args.notificationTemplateID, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("disassociateJobTemplateNotificationTemplatesForType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("disassociateJobTemplateNotificationTemplatesForType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_disassociateWorkflowJobTemplateNotificationTemplatesForType(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		jobTemplateID          int
		notificationTemplateID int
		typ                    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NotificationTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &awx{
				client: tt.fields.client,
			}
			got, err := s.disassociateWorkflowJobTemplateNotificationTemplatesForType(tt.args.jobTemplateID, tt.args.notificationTemplateID, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("disassociateWorkflowJobTemplateNotificationTemplatesForType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("disassociateWorkflowJobTemplateNotificationTemplatesForType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_getAllPages(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		firstURL string
		params   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Credential
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &awx{
				client: tt.fields.client,
			}
			got, err := cs.getAllPages(tt.args.firstURL, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllPages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_getAllTeamUsersPages(t1 *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		firstURL string
		params   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &awx{
				client: tt.fields.client,
			}
			got, err := t.getAllTeamUsersPages(tt.args.firstURL, tt.args.params)
			if (err != nil) != tt.wantErr {
				t1.Errorf("getAllTeamUsersPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("getAllTeamUsersPages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_getOrganizationAllPages(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		firstURL string
		params   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Organization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.getOrganizationAllPages(tt.args.firstURL, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrganizationAllPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrganizationAllPages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_listWorkflowJobTemplateNodeSteps(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id       int
		endpoint string
		params   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := &awx{
				client: tt.fields.client,
			}
			got, got1, err := jt.listWorkflowJobTemplateNodeSteps(tt.args.id, tt.args.endpoint, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("listWorkflowJobTemplateNodeSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listWorkflowJobTemplateNodeSteps() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("listWorkflowJobTemplateNodeSteps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_createWorkflowJobTemplateNode(t *testing.T) {
	type args struct {
		client                                 *Client
		data                                   map[string]interface{}
		params                                 map[string]string
		workflowJobTemplateNodesActionEndpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    *WorkflowJobTemplateNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createWorkflowJobTemplateNode(tt.args.client, tt.args.data, tt.args.params, tt.args.workflowJobTemplateNodesActionEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("createWorkflowJobTemplateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createWorkflowJobTemplateNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchWorkflowJobTemplateNode(t *testing.T) {
	type args struct {
		client                                 *Client
		params                                 map[string]string
		workflowJobTemplateNodesActionEndpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    []*WorkflowJobTemplateNode
		want1   *ListWorkflowJobTemplateNodesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := fetchWorkflowJobTemplateNode(tt.args.client, tt.args.params, tt.args.workflowJobTemplateNodesActionEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchWorkflowJobTemplateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchWorkflowJobTemplateNode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("fetchWorkflowJobTemplateNode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_newAWX(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		args args
		want AWX
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAWX(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAWX() = %v, want %v", got, tt.want)
			}
		})
	}
}
