package killbill

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/killbill/kbcli/v2/kbclient/subscription"
	"github.com/killbill/kbcli/v2/kbmodel"
	"strings"
)

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
