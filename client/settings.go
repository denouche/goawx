package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// SettingService implements awx settings apis.
type SettingService interface {
	ListSettings(params map[string]string) ([]*SettingSummary, *ListSettingsResponse, error)
	GetSettingsBySlug(slug string, params map[string]string) (*Setting, error)
	UpdateSettings(slug string, data map[string]interface{}, params map[string]string) (*Setting, error)
	DeleteSettings(slug string) (*Setting, error)
}

// ListSettingsResponse represents `ListSettings` endpoint response.
type ListSettingsResponse struct {
	Pagination
	Results []*SettingSummary `json:"results"`
}

const settingsAPIEndpoint = "/api/v2/settings/"

// ListSettings shows list of awx settings.
func (p *awx) ListSettings(params map[string]string) ([]*SettingSummary, *ListSettingsResponse, error) {
	result := new(ListSettingsResponse)
	resp, err := p.client.Requester.GetJSON(settingsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// GetSettingById shows the details of a setting.
func (p *awx) GetSettingsBySlug(slug string, params map[string]string) (*Setting, error) {
	result := new(Setting)
	endpoint := fmt.Sprintf("%s%s/", settingsAPIEndpoint, slug)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateSetting update an awx Setting.
func (p *awx) UpdateSettings(slug string, data map[string]interface{}, params map[string]string) (*Setting, error) {
	result := new(Setting)
	endpoint := fmt.Sprintf("%s%s", settingsAPIEndpoint, slug)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteSetting delete an awx Setting.
func (p *awx) DeleteSettings(slug string) (*Setting, error) {
	result := new(Setting)
	endpoint := fmt.Sprintf("%s%s", settingsAPIEndpoint, slug)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
