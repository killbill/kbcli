package kbcli

import (
	"github.com/killbill/kbcli/models/gen"
	"net/url"
	"net/http"
	"strings"
)

const EMPTY_STRING = ""

const SLASH string = "/"
const TENANT_PATH string = "/tenants"
const ACCOUNT_PATH string = "/accounts"
const PAYMENT_PATH string = "/payments"

const PAYMENT_METHOD_PATH string = "/paymentMethods"

const COMBO_PATH = "/combo"

const QUERY_EXTERNAL_KEY = "externalKey"
const QUERY_AUDIT = "audit"
const QUERY_PAYMENT_METHOD_IS_DEFAULT = "isDefault"

func CreateSession(killbillIp, killbillPort, userName, password, apiKey, apiSecret, createdBy  string, debugLog bool) (s *Session) {

	headers := make(http.Header)
	headers.Set("X-Killbill-CreatedBy", createdBy)

	return &Session{
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

	attr := &gen.TenantAttributes{ApiKey:s.ApiKey, ApiSecret:s.ApiSecret}
	resp, err := s.Post(TENANT_PATH, attr, nil)
	if err != nil {
		s.log("Failed to post request for path", TENANT_PATH)
		s.log(err)
		return &gen.TenantAttributes{}, err
	}

	var deserializer JsonDeserializer = &gen.TenantAttributes{}
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

func CreateAccount(s *Session, attr *gen.AccountAttributes, params *QueryParams) (*gen.AccountAttributes, error) {
	resp, err := s.Post(ACCOUNT_PATH, attr, params)
	if err != nil {
		s.log("Failed to post request for path", ACCOUNT_PATH)
		s.log(err)
		return &gen.AccountAttributes{}, err
	}

	var deserializer JsonDeserializer = &gen.AccountAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], deserializer, nil)
		return resp.Result.(*gen.AccountAttributes), err
	} else {
		return &gen.AccountAttributes{}, err
	}
}

func GetAccount(s *Session, accountId string, params *QueryParams) (*gen.AccountAttributes, error) {
	var deserializer JsonDeserializer = &gen.AccountAttributes{}
	var url string
	if accountId != EMPTY_STRING {
		resourceParts := []string{ACCOUNT_PATH, accountId}
		url = strings.Join(resourceParts, SLASH)
	} else {
		url = ACCOUNT_PATH
	}
	resp, err := s.Get(url, deserializer, params)
	return resp.Result.(*gen.AccountAttributes), err
}


//
//                            PAYMENT
//
func CreatePaymentMethod(s *Session, attr *gen.PaymentMethodAttributes, params *QueryParams) (*gen.PaymentMethodAttributes, error) {

	resourceParts := []string{ACCOUNT_PATH, SLASH, attr.AccountId, PAYMENT_METHOD_PATH}
	url := strings.Join(resourceParts, "")
	resp, err := s.Post(url, attr, params)
	if err != nil {
		s.log("Failed to post request for path:", url)
		s.log(err)
		return &gen.PaymentMethodAttributes{}, err
	}

	var deserializer JsonDeserializer = &gen.PaymentMethodAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], deserializer, nil)
		return resp.Result.(*gen.PaymentMethodAttributes), err
	} else {
		return &gen.PaymentMethodAttributes{}, err
	}
}


func CreatePaymentAuthorization(s *Session, accountId string, attr *gen.PaymentTransactionAttributes, params *QueryParams) (*gen.PaymentAttributes, error) {
	resourceParts := []string{ACCOUNT_PATH, SLASH, accountId, PAYMENT_PATH}
	url := strings.Join(resourceParts, "")
	return createPaymentTransaction(s, url, attr, params)
}

func CreatePaymentPurchase(s *Session, accountId string, attr *gen.PaymentTransactionAttributes, params *QueryParams) (*gen.PaymentAttributes, error) {
	resourceParts := []string{ACCOUNT_PATH, SLASH, accountId, PAYMENT_PATH}
	url := strings.Join(resourceParts, "")
	return createPaymentTransaction(s, url, attr, params)
}

func CreatePaymentCredit(s *Session, accountId string, attr *gen.PaymentTransactionAttributes, params *QueryParams) (*gen.PaymentAttributes, error) {
	resourceParts := []string{ACCOUNT_PATH, SLASH, accountId, PAYMENT_PATH}
	url := strings.Join(resourceParts, "")
	return createPaymentTransaction(s, url, attr, params)
}

func CreatePaymentCapture(s *Session, accountId string, attr *gen.PaymentTransactionAttributes, params *QueryParams) (*gen.PaymentAttributes, error) {
	resourceParts := []string{PAYMENT_PATH, SLASH, attr.PaymentId}
	url := strings.Join(resourceParts, "")
	return createPaymentTransaction(s, url, attr, params)
}

func CreateComboAuthorization(s *Session, attr interface{}, params *QueryParams) (*gen.PaymentAttributes, error) {
	resourceParts := []string{PAYMENT_PATH, COMBO_PATH}
	url := strings.Join(resourceParts, "")
	return createPaymentTransaction(s, url, attr, params)
}


func createPaymentTransaction(s *Session, url string, attr interface{}, params *QueryParams) (*gen.PaymentAttributes, error) {

	resp, err := s.Post(url, attr, nil)
	if err != nil {
		s.log("Failed to post request for path:", url)
		s.log(err)
		return &gen.PaymentAttributes{}, err
	}

	var deserializer JsonDeserializer = &gen.PaymentAttributes{}
	if value, ok := resp.HttpResponse().Header["Location"]; ok {
		resp, err = s.Get(value[0], deserializer, nil)
		return resp.Result.(*gen.PaymentAttributes), err
	} else {
		return &gen.PaymentAttributes{}, err
	}
}



