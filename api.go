package kbcli

import (
	"github.com/killbill/kbcli/models/gen"
	"net/url"
	"net/http"
)


func CreateSession(killbillIp, killbillPort, userName, password, apiKey, apiSecret, createdBy  string, debugLog bool) (s *Session) {
	headers := make(http.Header)
	headers.Set("X-Killbill-CreatedBy", createdBy)
	return &Session {
		KillbillIp: killbillIp,
		KillbillPort: killbillPort,
		Userinfo: url.UserPassword(userName, password),
		ApiKey:apiKey,
		ApiSecret:apiSecret,
		Header: &headers,
		Log: debugLog,
	}
}


func CreateTenant(s *Session) (*Response, error) {
	tenantInfo := &gen.TenantAttributes{ApiKey:s.ApiKey, ApiSecret:s.ApiSecret}
	resp, err := s.Post("/tenants", tenantInfo)
	s.log("*********** resp HttpResponse headers", resp.HttpResponse().Header)
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		return s.Get(value[0], tenantInfo)
	} else {
		return resp, err
	}
}

