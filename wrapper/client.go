package killbill

import (
	"encoding/base64"
	"fmt"
	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"strconv"
	"strings"
	"time"

	"github.com/killbill/kbcli/v3/kbclient"
	"github.com/killbill/kbcli/v3/kbcommon"
)

const CreatedBy = "go-client"
const Comment = ""

const DebugHttpRequests = false

var NullDate = strfmt.Date(time.Time{})

type KillbillPolicy string

type Logger interface {
	LogError(err error, format string, a ...interface{})
}

const (
	NONE KillbillPolicy = ""
	IMM  KillbillPolicy = "IMMEDIATE"
	EOT  KillbillPolicy = "END_OF_TERM"
	SOT  KillbillPolicy = "START_OF_TERM"
)

// Kill Bill Control Tag
const (
	AUTO_PAY_OFF               = "00000000-0000-0000-0000-000000000001"
	AUTO_INVOICING_OFF         = "00000000-0000-0000-0000-000000000002"
	OVERDUE_ENFORCEMENT_OFF    = "00000000-0000-0000-0000-000000000003"
	WRITTEN_OFF                = "00000000-0000-0000-0000-000000000004"
	MANUAL_PAY                 = "00000000-0000-0000-0000-000000000005"
	TEST                       = "00000000-0000-0000-0000-000000000006"
	PARTNER                    = "00000000-0000-0000-0000-000000000007"
	AUTO_INVOICING_DRAFT       = "00000000-0000-0000-0000-000000000008"
	AUTO_INVOICING_REUSE_DRAFT = "00000000-0000-0000-0000-000000000009"
)

// Kill Bill Error Codes
const (
	SUB_CREATE_BP_EXISTS           = 1015
	ACCOUNT_ALREADY_EXISTS         = 3000
	ACCOUNT_DOES_NOT_EXIST_FOR_KEY = 3003
)

func NewKBClient(conf KillbillConfig) *RawClient {
	trp := NewClientTransport(conf.GetUrl())
	return &RawClient{
		Trp:               trp,
		ApiKey:            conf.GetApiKey(),
		ApiSecret:         conf.GetApiSecret(),
		CrossTenantClient: NewKillbillClient(trp, "", "", conf.GetUsername(), conf.GetPassword()),
		TenantClient:      NewKillbillClient(trp, conf.GetApiKey(), conf.GetApiSecret(), conf.GetUsername(), conf.GetPassword()),
		Timeout:           conf.GetTimeout(),
	}
}

type RawClient struct {
	// Low level transport
	Trp *transport.Runtime
	// Cross tenant operations (e.g Create tenant)
	CrossTenantClient *kbclient.KillBill
	// Per-tenant client
	ApiKey       string
	ApiSecret    string
	TenantClient *kbclient.KillBill
	// Default timeout
	Timeout time.Duration
}

// Technically this could be shared across clients if needs to be
func NewClientTransport(url string) *transport.Runtime {

	// Expect hostname:port
	parts := strings.Split(url, ":")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Unexpected KB URL %s ", url))
	}
	kbPort, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("Unexpected KB port %s ", parts[1]))
	}

	scheme := "http"
	if kbPort == 443 {
		scheme = "https"
	}

	trp := transport.New(url, "", []string{scheme})
	// Add text/xml producer which is not handled by openapi runtime.
	trp.Producers["text/xml"] = runtime.TextProducer()
	// Set this to true to dump http messages
	trp.Debug = DebugHttpRequests
	return trp
}

// Because of Kill Bill not always correctly returning the right error code on duplicate entry
// we need to add some dirty band-aid and specify 'ignrDup', 'logger' args to this function
// See https://github.com/killbill/killbill/issues/1344
func IsCriticalError(err error, kbCodeIgnore int, ignrDup bool, logger Logger) error {
	ke, ok := err.(*kbcommon.KillbillError)
	if ok {
		if ke.Code == kbCodeIgnore {
			logger.LogError(err, "Duplicate entry error ignored...")
			return nil
		}
	}
	// Hack
	if err != nil && ignrDup {
		//
		// It seems like we can get either:
		// Case1: "[500] java.sql.SQLIntegrityConstraintViolationException: (conn=81) Duplicate entry '37144e24-d5b4-419b-944f-4e42fd56b59e-916' for key 'accounts.accounts_external_key'"
		// Case 2: "[500] org.postgresql.util.PSQLException: ERROR: duplicate key value violates unique constraint "accounts_external_key Detail: Key (external_key, tenant_record_id)=(d585e7be-cad6-4743-95a0-6b5cf21df5c1, 1) already exists."

		checkErrFn := func(err string) bool {
			return (strings.Contains(err, "Duplicate entry") && strings.Contains(err, "for key")) /* Case 1*/ ||
				strings.Contains(err, "duplicate key value violates unique constraint") /* Case 2*/
		}
		if checkErrFn(err.Error()) {
			logger.LogError(err, "Duplicate entry (hack) error ignored...")
			return nil
		}
	}

	// Add debug for tracking potential test flakiness
	if err != nil {
		logger.LogError(err, "killbill#isCriticalError err=%+v, err#string =%s, kbCodeIgnore=%d, ignrDup=%t", err, err.Error(), kbCodeIgnore, ignrDup)
	}
	return err
}

func CreateAuthInfo(apiKey, apiSecret, user, pwd string) runtime.ClientAuthInfoWriter {
	authWriter := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		encoded := base64.StdEncoding.EncodeToString([]byte(user + ":" + pwd))
		if err := r.SetHeaderParam("Authorization", "Basic "+encoded); err != nil {
			return err
		}
		if apiKey != "" {
			if err := r.SetHeaderParam("X-KillBill-ApiKey", apiKey); err != nil {
				return err
			}
		}
		if apiSecret != "" {
			if err := r.SetHeaderParam("X-KillBill-ApiSecret", apiSecret); err != nil {
				return err
			}
		}
		return nil
	})
	return authWriter
}

func NewKillbillClient(trp *transport.Runtime, apiKey, apiSecret, user, pwd string) *kbclient.KillBill {
	authWriter := CreateAuthInfo(apiKey, apiSecret, user, pwd)
	client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})

	createdBy := CreatedBy
	comment := Comment
	reason := ""

	client.SetDefaults(kbclient.KillbillDefaults{
		CreatedBy: &createdBy,
		Comment:   &comment,
		Reason:    &reason,
	})
	return client
}
