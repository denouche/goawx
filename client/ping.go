package awx

// PingService implements awx ping apis.
type PingService interface {
	Ping() (*Ping, error)
}

const pingAPIEndpoint = "/api/v2/ping/"

// Ping do ping with awx servers.
func (p *awx) Ping() (*Ping, error) {
	result := new(Ping)
	resp, err := p.client.Requester.GetJSON(pingAPIEndpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
