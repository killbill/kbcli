package gen

import "encoding/json"

type TenantAttributes struct {
	ApiKey      string `json:"apiKey"`
	ApiSecret   string `json:"apiSecret"`
	TenantId    string `json:"tenantId"`
	ExternalKey string `json:"externalKey"`
	auditLogs   interface{}
}


func (data *TenantAttributes) FromJson(raw []byte) error {
	return json.Unmarshal(raw, data)
}