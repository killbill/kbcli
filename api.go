package kbcli

import (
	"github.com/killbill/kbcli/models/gen"
	"net/url"
	"net/http"
)

const TENANT_PATH string = "/tenants"



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


func CreateTenant(s *Session) (*gen.TenantAttributes, error) {

	var tenantInfo JsonDeserializer
	tenantInfo = &gen.TenantAttributes{ApiKey:s.ApiKey, ApiSecret:s.ApiSecret}
	resp, err := s.Post(TENANT_PATH, tenantInfo)
	if err != nil {
		s.log("Failed to post request for path", TENANT_PATH)
		s.log(err)
		return &gen.TenantAttributes{}, err
	}

	tenantInfo = &gen.TenantAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], tenantInfo)
		return resp.Result.(*gen.TenantAttributes), err
	} else {
		return &gen.TenantAttributes{}, err
	}
}

