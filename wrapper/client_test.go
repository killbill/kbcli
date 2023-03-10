package killbill

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/killbill/kbcli/v3/kbmodel"
	"os"
	"testing"
	"time"
)

const DefaultUrl string = "127.0.0.1:8080"
const DefaultUser string = "admin"
const DefaultPwd string = "password"
const DefaultToSec int64 = 5

// Max calls to make for async operations each delayed by SleepTime
const SleepTime time.Duration = 1 * time.Second
const MaxAttempts int = 5

func randStr() string {
	res, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return res.String()
}

type testInfo struct {
	cli  *RawClient
	t    *testing.T
	conf Config
}

func setupTestWithDefault(t *testing.T) *testInfo {
	return setupTest(randStr(), DefaultUrl, DefaultUser, DefaultPwd, DefaultToSec, t)
}

func setupTest(key string, url string, user string, pwd string, toSec int64, t *testing.T) *testInfo {
	conf := Config{
		Url:        url,
		Username:   user,
		Password:   pwd,
		ApiKey:     key,
		ApiSecret:  key,
		TimeoutSec: toSec,
	}

	cli := NewKBClient(&conf)
	return &testInfo{
		cli:  cli,
		conf: conf,
		t:    t,
	}
}

func (ti *testInfo) createTenant() {
	ctx := context.Background()
	_, err := ti.cli.CreateTenant(ctx, ti.conf.ApiKey, ti.conf.ApiSecret, ti.conf.ApiKey)
	if err != nil {
		ti.t.Fatalf("Failed to create tenant, err=%s", err)
	}
}

func (ti *testInfo) uploadCatalog(catalogName string) {

	curPath, err := os.Getwd()
	if err != nil {
		ti.t.Fatalf("Failed get current directory err=%s", err)
	}

	ctx := context.Background()
	catalogPath := fmt.Sprintf("%s/resource/%s", curPath, catalogName)
	err = ti.cli.UploadCatalogXML(ctx, catalogPath)
	if err != nil {
		ti.t.Fatalf("Failed to upload catalog from resource %s, err=%s", catalogPath, err)
	}
}

func (ti *testInfo) createAccount(accountKey string) strfmt.UUID {
	ctx := context.Background()
	accountId, err := ti.cli.CreateAccount(ctx, accountKey, "USD", "en_US", 0)
	if err != nil {
		ti.t.Fatalf("Failed create account err=%s", err)
	}
	//fmt.Printf("accountId = %s\n", accountId)
	return accountId
}

func (ti *testInfo) createSubscription(accountId strfmt.UUID, subKey string, effDt *strfmt.Date, planName string) strfmt.UUID {

	ctx := context.Background()
	subId, err := ti.cli.CreateSubscription(ctx, accountId, subKey, effDt, effDt, planName, 0, true)
	if err != nil {
		ti.t.Fatalf("Failed create subscription err=%s", err)
	}
	//fmt.Printf("subId = %s\n", subId)
	return subId
}

func (ti *testInfo) getAllInvoices(accountId strfmt.UUID) []*kbmodel.Invoice {

	ctx := context.Background()
	invs, err := ti.cli.GetAllInvoices(ctx, accountId)
	if err != nil {
		ti.t.Fatalf("Failed retrieve invoices err=%s", err)
	}
	return invs
}

func (ti *testInfo) fetchInvoicesUntil(accountId strfmt.UUID, cond func([]*kbmodel.Invoice) bool) bool {
	for i := 0; i < MaxAttempts; i++ {
		invs := ti.getAllInvoices(accountId)
		if cond(invs) {
			return true
		}
		time.Sleep(SleepTime)
	}
	return false
}

func TestSubscriptionScenario(t *testing.T) {

	test := setupTestWithDefault(t)

	test.createTenant()
	test.uploadCatalog("catalog.xml")

	accountKey := randStr()
	accountId := test.createAccount(accountKey)

	invs := test.getAllInvoices(accountId)
	if len(invs) != 0 {
		test.t.Fatalf("Expected no invoice, but got %d", len(invs))
	}

	subKey := randStr()
	effDt := strfmt.Date(time.Now().UTC())
	test.createSubscription(accountId, subKey, &effDt, "standard-monthly")

	res := test.fetchInvoicesUntil(accountId, func(input []*kbmodel.Invoice) bool {
		return len(input) == 1
	})
	if !res {
		test.t.Fatalf("Failed to fetch expected invoices after %d attempts", MaxAttempts)
	}
}
