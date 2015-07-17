package kbcli

import (
	"github.com/killbill/kbcli/models/gen"
	"net/url"
	"net/http"
)

const TENANT_PATH string = "/tenants"
const ACCOUNT_PATH string = "/accounts"

const QUERY_EXTERNAL_KEY = "externalKey";


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

//
//                            TENANT
//


func CreateTenant(s *Session) (*gen.TenantAttributes, error) {

	var deserializer JsonDeserializer
	deserializer = &gen.TenantAttributes{ApiKey:s.ApiKey, ApiSecret:s.ApiSecret}
	resp, err := s.Post(TENANT_PATH, deserializer)
	if err != nil {
		s.log("Failed to post request for path", TENANT_PATH)
		s.log(err)
		return &gen.TenantAttributes{}, err
	}

	deserializer = &gen.TenantAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], deserializer, nil)
		return resp.Result.(*gen.TenantAttributes), err
	} else {
		return &gen.TenantAttributes{}, err
	}
}

//
//                            ACCOUNT
//

func CreateAccount(s *Session, attr *gen.AccountAttributes) (*gen.AccountAttributes, error) {

	var deserializer JsonDeserializer
	deserializer = attr
	resp, err := s.Post(ACCOUNT_PATH, deserializer)
	if err != nil {
		s.log("Failed to post request for path", ACCOUNT_PATH)
		s.log(err)
		return &gen.AccountAttributes{}, err
	}

	deserializer = &gen.AccountAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], deserializer, nil)
		return resp.Result.(*gen.AccountAttributes), err
	} else {
		return &gen.AccountAttributes{}, err
	}
}

func GetAccountByKey(s *Session, key string) (*gen.AccountAttributes, error) {
	var deserializer JsonDeserializer
	deserializer = &gen.AccountAttributes{}
	params := make(QueryParams)
	params[QUERY_EXTERNAL_KEY] = key
	resp, err := s.Get(ACCOUNT_PATH, deserializer, &params)
	return resp.Result.(*gen.AccountAttributes), err
}

