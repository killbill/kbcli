package gen

type TenantAttributes struct {
	ApiKey      string `json:"apiKey"`
	ApiSecret   string `json:"apiSecret"`
	TenantId    string `json:"tenantId"`
	ExternalKey string `json:"externalKey"`
	auditLogs   interface{}
}
