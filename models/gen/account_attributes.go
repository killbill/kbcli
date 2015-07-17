package gen

import "encoding/json"


type AccountAttributes struct {
	AccountId      string `json:"accountId"`
	Name string `json:"name"`
	ExternalKey   string `json:"externalKey"`
	Email    string `json:"email"`
	Currency string `json:"currency"`
	auditLogs   interface{}
}

func (data *AccountAttributes) FromJson(raw []byte) error {
	return json.Unmarshal(raw, data)
}