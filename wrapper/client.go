package killbill

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"strconv"
	"strings"
	"time"

	"github.com/killbill/kbcli/v2/kbclient"
	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbclient/admin"
	"github.com/killbill/kbcli/v2/kbclient/bundle"
	"github.com/killbill/kbcli/v2/kbclient/invoice"
	"github.com/killbill/kbcli/v2/kbclient/subscription"
	"github.com/killbill/kbcli/v2/kbclient/tenant"
	"github.com/killbill/kbcli/v2/kbcommon"
	"github.com/killbill/kbcli/v2/kbmodel"
)

const CreatedBy = "embs"
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
//
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

func (cli *RawClient) CreateTenant(ctx context.Context, apiKey string, apiSecret string, extKey string) (*strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.CrossTenantClient.Tenant.CreateTenant(ctx, &tenant.CreateTenantParams{
		Body: &kbmodel.Tenant{
			APIKey:      swag.String(apiKey),
			APISecret:   swag.String(apiSecret),
			ExternalKey: extKey,
		},
		// TODO Should it be set to false instead?
		UseGlobalDefault: swag.Bool(true),
		// Setting ProcessLocationHeader does not work so we cannot get out tenant back
		ProcessLocationHeader: false,
	})
	if err != nil {
		return nil, err
	}

	// This returns an empty ID
	return &res.Payload.TenantID, nil
}

func (cli *RawClient) GetTenant(ctx context.Context, apiKey string) (*strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.CrossTenantClient.Tenant.GetTenantByAPIKey(ctx, &tenant.GetTenantByAPIKeyParams{
		APIKey:                swag.String(apiKey),
		ProcessLocationHeader: false,
	})
	if err != nil {
		return nil, err
	}
	return &res.Payload.TenantID, nil
}

func (cli *RawClient) CreateAccount(ctx context.Context, extKey string, currency string, locale string, bcd int32) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	body := &kbmodel.Account{
		ExternalKey:   extKey,
		Currency:      kbmodel.AccountCurrencyEnum(currency),
		Locale:        locale,
		ReferenceTime: strfmt.DateTime(time.Now()),
	}

	if bcd > 0 {
		body.BillCycleDayLocal = bcd
	}

	return cli.CreateAccountFromBody(ctx, body)
}

func (cli *RawClient) CreateAccountFromBody(ctx context.Context, body *kbmodel.Account) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.CreateAccount(ctx, &account.CreateAccountParams{
		Body:                  body,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return "", err
	}
	return res.Payload.AccountID, nil
}

func (cli *RawClient) GetAccount(ctx context.Context, accountId strfmt.UUID, withBalance bool) (*kbmodel.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccount(ctx, &account.GetAccountParams{
		AccountID:                accountId,
		AccountWithBalanceAndCBA: &withBalance,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) GetAccountByKey(ctx context.Context, externalKey string) (*kbmodel.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccountByKey(ctx, &account.GetAccountByKeyParams{
		ExternalKey: externalKey,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) AddAccountCustomFields(ctx context.Context, accountId strfmt.UUID, fields map[string]string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	body := make([]*kbmodel.CustomField, len(fields))
	idx := 0
	for name, value := range fields {
		body[idx] = &kbmodel.CustomField{
			Name:       &name,
			ObjectID:   accountId,
			ObjectType: kbmodel.CustomFieldObjectTypeACCOUNT,
			Value:      &value,
		}
		idx++
	}

	_, err := cli.TenantClient.Account.CreateAccountCustomFields(ctx, &account.CreateAccountCustomFieldsParams{
		Body:                  body,
		AccountID:             accountId,
		ProcessLocationHeader: true,
	})
	return err
}

func (cli *RawClient) GetAccountTags(ctx context.Context, accountId strfmt.UUID) ([]*kbmodel.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccountTags(ctx, &account.GetAccountTagsParams{
		AccountID: accountId,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) AddAccountTags(ctx context.Context, accountId strfmt.UUID, tags ...string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := make([]strfmt.UUID, len(tags))
	for i := 0; i < len(tags); i++ {
		fmtTags[i] = strfmt.UUID(tags[i])
	}

	_, err := cli.TenantClient.Account.CreateAccountTags(ctx, &account.CreateAccountTagsParams{
		AccountID: accountId,
		Body:      fmtTags,
	})
	return err
}

func (cli *RawClient) DeleteAccountTags(ctx context.Context, accountId strfmt.UUID, tags ...string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := make([]strfmt.UUID, len(tags))
	for i := 0; i < len(tags); i++ {
		fmtTags[i] = strfmt.UUID(tags[i])
	}

	_, err := cli.TenantClient.Account.DeleteAccountTags(ctx, &account.DeleteAccountTagsParams{
		AccountID: accountId,
		TagDef:    fmtTags,
	})
	return err
}

func (cli *RawClient) GetSubscription(ctx context.Context, subId strfmt.UUID) (*kbmodel.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Subscription.GetSubscription(ctx, &subscription.GetSubscriptionParams{
		SubscriptionID: subId,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) GetBundle(ctx context.Context, bundleId strfmt.UUID) (*kbmodel.Bundle, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Bundle.GetBundle(ctx, &bundle.GetBundleParams{
		BundleID: bundleId,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) GetBundleByKey(ctx context.Context, extKey string) (*kbmodel.Bundle, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Bundle.GetBundleByKey(ctx, &bundle.GetBundleByKeyParams{
		ExternalKey: extKey,
	})
	if err != nil {
		return nil, err
	}
	if len(res.Payload) > 1 {
		return nil, fmt.Errorf("Unexpected number of bundles for extKey=%s: Got %d bundles", extKey, len(res.Payload))
	}
	return res.Payload[0], nil
}

func (cli *RawClient) GetSubscriptionByKey(ctx context.Context, extKey string) (*kbmodel.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Subscription.GetSubscriptionByKey(ctx, &subscription.GetSubscriptionByKeyParams{
		ExternalKey: extKey,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) TransferBundle(ctx context.Context, bundleId strfmt.UUID, destAccountId strfmt.UUID, requestedDate *strfmt.Date, subExtkeys map[string]string, follow bool) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	// Serialize subscription external keys, rely on undocumented mechanism until https://github.com/killbill/killbill/issues/1627 is fixed.
	props := make([]string, 0)
	for k, v := range subExtkeys {
		prop := fmt.Sprintf("KB_SUB_ID_%s=%s", k, v)
		props = append(props, prop)
	}

	billingPolicy := "IMMEDIATE"
	res, err := cli.TenantClient.Bundle.TransferBundle(ctx, &bundle.TransferBundleParams{
		BillingPolicy: &billingPolicy,
		Body: &kbmodel.Bundle{
			AccountID: &destAccountId,
		},
		PluginProperty:        props,
		BundleID:              bundleId,
		RequestedDate:         requestedDate,
		ProcessLocationHeader: follow,
	})
	if err != nil {
		return "", err
	}
	var result strfmt.UUID
	if follow {
		result = res.Payload.BundleID
	} else {
		location := res.HttpResponse.GetHeader("Location")
		parts := strings.Split(location, "/")
		result = strfmt.UUID(parts[len(parts)-1])
	}
	return result, nil
}

func (cli *RawClient) CreateSubscription(ctx context.Context, accountId strfmt.UUID, extKey string, entitlementDate *strfmt.Date, billingDate *strfmt.Date, planId string, bcd int, follow bool) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Subscription.CreateSubscription(ctx, &subscription.CreateSubscriptionParams{
		Body: &kbmodel.Subscription{
			AccountID:         accountId,
			PlanName:          swag.String(planId),
			BillCycleDayLocal: int32(bcd),
			ExternalKey:       extKey,
		},
		EntitlementDate:       entitlementDate,
		BillingDate:           billingDate,
		ProcessLocationHeader: follow,
		SkipResponse:          swag.Bool(!follow),
	})
	if err != nil {
		return "", err
	}

	var result strfmt.UUID
	if follow {
		result = res.Payload.SubscriptionID
	} else {
		location := res.HttpResponse.GetHeader("Location")
		parts := strings.Split(location, "/")
		result = strfmt.UUID(parts[len(parts)-1])
	}
	return result, nil
}

type PlanDescr struct {
	ExtKey string
	PlanId string
	Bcd    int
}

func (cli *RawClient) CreateSubscriptions(pctx context.Context, accountId strfmt.UUID, date *strfmt.Date, planDescs []PlanDescr, follow bool) ([]*kbmodel.Bundle, error) {

	prof := getProfValue(pctx)

	ctx, cancel := context.WithTimeout(pctx, cli.Timeout)
	defer cancel()

	body := make([]*kbmodel.BulkSubscriptionsBundle, 0)
	for _, desc := range planDescs {
		body = append(body, &kbmodel.BulkSubscriptionsBundle{
			BaseEntitlementAndAddOns: []*kbmodel.Subscription{
				{
					AccountID:         accountId,
					PlanName:          swag.String(desc.PlanId),
					BillCycleDayLocal: int32(desc.Bcd),
					ExternalKey:       desc.ExtKey,
				},
			},
		})
	}

	res, err := cli.TenantClient.Subscription.CreateSubscriptionsWithAddOns(ctx, &subscription.CreateSubscriptionsWithAddOnsParams{
		Body:                  body,
		EntitlementDate:       date,
		BillingDate:           date,
		WithProfilingInfo:     prof.profInfo,
		ProcessLocationHeader: follow,
		SkipResponse:          swag.Bool(!follow),
	})
	if err != nil {
		return nil, err
	}
	if prof.profInfo != nil {
		setProfRes(prof, res.HttpResponse)
	}
	return res.Payload, nil
}

func (cli *RawClient) PauseSubscription(ctx context.Context, subscriptionId strfmt.UUID, date *strfmt.Date) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Subscription.AddSubscriptionBlockingState(ctx, &subscription.AddSubscriptionBlockingStateParams{
		Body: &kbmodel.BlockingState{
			BlockedID:          subscriptionId,
			IsBlockBilling:     true,
			IsBlockChange:      false,
			IsBlockEntitlement: true,
			// TODO Abstract those ?
			Service:   "EMBS",
			StateName: "EMBS-PAUSE",
			Type:      kbmodel.BlockingStateTypeSUBSCRIPTION,
		},
		RequestedDate:  date,
		SubscriptionID: subscriptionId,
	})
	return err
}

func (cli *RawClient) ResumeSubscription(ctx context.Context, subscriptionId strfmt.UUID, date *strfmt.Date) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Subscription.AddSubscriptionBlockingState(ctx, &subscription.AddSubscriptionBlockingStateParams{
		Body: &kbmodel.BlockingState{
			BlockedID:          subscriptionId,
			IsBlockBilling:     false,
			IsBlockChange:      false,
			IsBlockEntitlement: false,
			Service:            "EMBS",
			StateName:          "EMBS-RESUME",
			Type:               kbmodel.BlockingStateTypeSUBSCRIPTION,
		},
		RequestedDate:  date,
		SubscriptionID: subscriptionId,
	})
	return err
}

func (cli *RawClient) CancelSubscription(ctx context.Context, subscriptionId strfmt.UUID, date *strfmt.Date) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	var policy *string
	if date == nil {
		policy = swag.String("END_OF_TERM")
	}

	_, err := cli.TenantClient.Subscription.CancelSubscriptionPlan(ctx, &subscription.CancelSubscriptionPlanParams{
		EntitlementPolicy:          policy,
		BillingPolicy:              policy,
		RequestedDate:              date,
		SubscriptionID:             subscriptionId,
		UseRequestedDateForBilling: swag.Bool(true),
	})
	return err
}

func (cli *RawClient) AddSubscriptionCustomFields(ctx context.Context, subscriptionId strfmt.UUID, fields map[string]string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	body := make([]*kbmodel.CustomField, len(fields))
	idx := 0
	for name, value := range fields {
		body[idx] = &kbmodel.CustomField{
			Name:       &name,
			ObjectID:   subscriptionId,
			ObjectType: kbmodel.CustomFieldObjectTypeSUBSCRIPTION,
			Value:      &value,
		}
		idx++
	}

	_, err := cli.TenantClient.Subscription.CreateSubscriptionCustomFields(ctx, &subscription.CreateSubscriptionCustomFieldsParams{
		Body:                  body,
		SubscriptionID:        subscriptionId,
		ProcessLocationHeader: true,
	})
	return err
}

func (cli *RawClient) TriggerDryRunCreateSubscription(ctx context.Context, accountId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionSTARTBILLING, "", "", product, priceList, effDt, targetDt)
}

func (cli *RawClient) TriggerDryRunChangePlan(ctx context.Context, accountId strfmt.UUID, bundleId strfmt.UUID, subscriptionId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionCHANGE, bundleId, subscriptionId, product, priceList, effDt, targetDt)
}

func (cli *RawClient) ChangePlan(ctx context.Context, subscriptionId strfmt.UUID, planName, billingPeriod, priceList string, requestedDate *strfmt.Date) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()
	enumValue := kbmodel.SubscriptionBillingPeriodEnum(billingPeriod)
	_, err := cli.TenantClient.Subscription.ChangeSubscriptionPlan(ctx, &subscription.ChangeSubscriptionPlanParams{
		SubscriptionID: subscriptionId,
		Body: &kbmodel.Subscription{
			PlanName:      &planName,
			BillingPeriod: &enumValue,
			PriceList:     &priceList,
		},
		RequestedDate: requestedDate,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cli *RawClient) UpdateSubscriptionBCD(ctx context.Context, subscriptionId strfmt.UUID, bcd int, requestedDate *strfmt.Date) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Subscription.UpdateSubscriptionBCD(ctx, &subscription.UpdateSubscriptionBCDParams{
		Body: &kbmodel.Subscription{
			BillCycleDayLocal: int32(bcd),
		},
		EffectiveFromDate: requestedDate,
		SubscriptionID:    subscriptionId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cli *RawClient) TriggerDryRunCancelSubscription(ctx context.Context, accountId strfmt.UUID, bundleId strfmt.UUID, subscriptionId strfmt.UUID, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	return cli.triggerDryRunAction(ctx, accountId, kbmodel.InvoiceDryRunDryRunActionSTOPBILLING, bundleId, subscriptionId, "", "", effDt, targetDt)
}

func (cli *RawClient) triggerDryRunAction(ctx context.Context, accountId strfmt.UUID, dryRunAction kbmodel.InvoiceDryRunDryRunActionEnum, bundleId strfmt.UUID, subscriptionId strfmt.UUID, product string, priceList string, effDt strfmt.Date, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	invOk, invNA, err := cli.TenantClient.Invoice.GenerateDryRunInvoice(ctx, &invoice.GenerateDryRunInvoiceParams{
		AccountID: accountId,
		Body: &kbmodel.InvoiceDryRun{
			DryRunAction:    dryRunAction,
			DryRunType:      kbmodel.InvoiceDryRunDryRunTypeSUBSCRIPTIONACTION,
			EffectiveDate:   effDt,
			BundleID:        bundleId,
			SubscriptionID:  subscriptionId,
			BillingPeriod:   kbmodel.InvoiceDryRunBillingPeriodMONTHLY,
			ProductCategory: kbmodel.InvoiceDryRunProductCategoryBASE,
			ProductName:     product,
			PriceListName:   priceList,
		},
		ProcessLocationHeader: true,
		TargetDate:            &targetDt,
	})
	if err != nil {
		return nil, err
	}

	if invNA != nil {
		return nil, nil
	} else {
		return invOk.Payload, nil
	}
}

// GetInvoice fetchs an invoice using the given invoiceID.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Invoice.GetInvoice(ctx, &invoice.GetInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetAllInvoices fetchs all invoice for the given account.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetAllInvoices(ctx context.Context, accountID strfmt.UUID) ([]*kbmodel.Invoice, error) {
	return cli.GetInvoicesBetweenDates(ctx, accountID, nil, nil)
}

// GetInvoicesBetweenDates returns all invoices belonging to an
// account, whose invoice date is in between the given arguments.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetInvoicesBetweenDates(ctx context.Context, accountID strfmt.UUID, from *strfmt.Date, to *strfmt.Date) ([]*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()
	b := true

	res, err := cli.TenantClient.Account.GetInvoicesForAccount(ctx, &account.GetInvoicesForAccountParams{
		AccountID:             accountID,
		StartDate:             from,
		EndDate:               to,
		IncludeVoidedInvoices: &b,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetMonthlyInvoice returns the invoice matching the given target
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetMonthlyInvoice(ctx context.Context, accountID strfmt.UUID, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	var from, to strfmt.Date

	res, err := cli.TenantClient.Account.GetInvoicesForAccount(ctx, &account.GetInvoicesForAccountParams{
		AccountID: accountID,
		StartDate: &from,
		EndDate:   &to,
	})
	if err != nil {
		return nil, err
	}
	// Kill Bill could return a successful status but without any invoices matching the query,
	// so we transform this as an error to make it easier on the consumer
	if len(res.Payload) == 0 {
		return nil, &account.GetInvoicesForAccountNotFound{
			HttpResponse: res.HttpResponse,
		}
	}
	return res.Payload[0], nil
}

// TriggerTargetInvoice creates a DRAFT invoice for the given account,
// using the given target date.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) TriggerTargetInvoice(ctx context.Context, accountID strfmt.UUID, targetDate strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	inv, err := cli.TenantClient.Invoice.CreateFutureInvoice(ctx, &invoice.CreateFutureInvoiceParams{
		AccountID:             accountID,
		TargetDate:            &targetDate,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return nil, err
	}

	return inv.Payload, nil
}

// VoidInvoice voids an invoice. If the invoice was already paid, voiding it fails.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) VoidInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Invoice.VoidInvoice(ctx, &invoice.VoidInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// CommitInvoice commits an invoice. If the current invoice status is different from DRAFT, it returns an error.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) CommitInvoice(ctx context.Context, invoiceID strfmt.UUID) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Invoice.CommitInvoice(ctx, &invoice.CommitInvoiceParams{
		InvoiceID: invoiceID,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// PayInvoice adds a full-amount payment to the given invoice.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) PayInvoice(ctx context.Context, invoiceID strfmt.UUID, accountID strfmt.UUID, amount float64, externalPayment bool) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, _, err := cli.TenantClient.Invoice.CreateInstantPayment(ctx, &invoice.CreateInstantPaymentParams{
		InvoiceID:       invoiceID,
		ExternalPayment: &externalPayment,
		Body: &kbmodel.InvoicePayment{
			AccountID:       accountID,
			PurchasedAmount: amount,
		},
	})

	return err
}

// IsInvoiceWrittenOff checks if the given invoice is written-off, returning true or false
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) IsInvoiceWrittenOff(ctx context.Context, invoiceID strfmt.UUID) (bool, string, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()
	reason := ""
	audit := "FULL"

	// Get the Invoice tags with the audit information
	res, err := cli.TenantClient.Invoice.GetInvoiceTags(ctx, &invoice.GetInvoiceTagsParams{
		InvoiceID: invoiceID,
		Audit:     &audit,
	})
	if err != nil {
		return false, reason, err
	}

	for _, tag := range res.Payload {
		if tag.TagDefinitionID == WRITTEN_OFF {
			// Gets tag definition audit logs and returns Comments field value from the latest one
			if len(tag.AuditLogs) >= 1 {
				reason = tag.AuditLogs[len(tag.AuditLogs)-1].Comments
			}
			return true, reason, nil
		}
	}

	return false, reason, nil
}

// WriteOffInvoice writes-off the given invoice.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) WriteOffInvoice(ctx context.Context, invoiceID strfmt.UUID, reason string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := []strfmt.UUID{WRITTEN_OFF}

	_, err := cli.TenantClient.Invoice.CreateInvoiceTags(ctx, &invoice.CreateInvoiceTagsParams{
		InvoiceID:        invoiceID,
		Body:             fmtTags,
		XKillbillComment: &reason,
	})

	return err
}

func (cli *RawClient) addPaymentMethod(ctx context.Context, accountId strfmt.UUID, isDefault bool) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Account.CreatePaymentMethod(ctx, &account.CreatePaymentMethodParams{
		AccountID: accountId,
		Body: &kbmodel.PaymentMethod{
			AccountID:  accountId,
			IsDefault:  isDefault,
			PluginName: "__EXTERNAL_PAYMENT__",
		},
		IsDefault: &isDefault,
	})
	return err
}

func (cli *RawClient) InvalidateTenantCache(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Admin.InvalidatesCacheByTenant(ctx, &admin.InvalidatesCacheByTenantParams{})
	return err
}
