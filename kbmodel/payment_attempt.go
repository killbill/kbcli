// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentAttempt payment attempt
//
// swagger:model PaymentAttempt
type PaymentAttempt struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// Transaction amount, required except for void operations
	Amount float64 `json:"amount,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// Amount currency (account currency unless specified)
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency PaymentAttemptCurrencyEnum `json:"currency,omitempty"`

	// effective date
	// Format: date-time
	EffectiveDate strfmt.DateTime `json:"effectiveDate,omitempty"`

	// payment external key
	PaymentExternalKey string `json:"paymentExternalKey,omitempty"`

	// payment method Id
	// Format: uuid
	PaymentMethodID strfmt.UUID `json:"paymentMethodId,omitempty"`

	// plugin name
	PluginName string `json:"pluginName,omitempty"`

	// plugin properties
	PluginProperties []*PluginProperty `json:"pluginProperties"`

	// state name
	StateName string `json:"stateName,omitempty"`

	// transaction external key
	TransactionExternalKey string `json:"transactionExternalKey,omitempty"`

	// transaction Id
	// Format: uuid
	TransactionID strfmt.UUID `json:"transactionId,omitempty"`

	// transaction type
	// Enum: [AUTHORIZE CAPTURE CHARGEBACK CREDIT PURCHASE REFUND VOID]
	TransactionType PaymentAttemptTransactionTypeEnum `json:"transactionType,omitempty"`
}

// Validate validates this payment attempt
func (m *PaymentAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttempt) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttempt) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var paymentAttemptTypeCurrencyPropEnum []interface{}

func init() {
	var res []PaymentAttemptCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAttemptTypeCurrencyPropEnum = append(paymentAttemptTypeCurrencyPropEnum, v)
	}
}

type PaymentAttemptCurrencyEnum string

const (

	// PaymentAttemptCurrencyAED captures enum value "AED"
	PaymentAttemptCurrencyAED PaymentAttemptCurrencyEnum = "AED"

	// PaymentAttemptCurrencyAFN captures enum value "AFN"
	PaymentAttemptCurrencyAFN PaymentAttemptCurrencyEnum = "AFN"

	// PaymentAttemptCurrencyALL captures enum value "ALL"
	PaymentAttemptCurrencyALL PaymentAttemptCurrencyEnum = "ALL"

	// PaymentAttemptCurrencyAMD captures enum value "AMD"
	PaymentAttemptCurrencyAMD PaymentAttemptCurrencyEnum = "AMD"

	// PaymentAttemptCurrencyANG captures enum value "ANG"
	PaymentAttemptCurrencyANG PaymentAttemptCurrencyEnum = "ANG"

	// PaymentAttemptCurrencyAOA captures enum value "AOA"
	PaymentAttemptCurrencyAOA PaymentAttemptCurrencyEnum = "AOA"

	// PaymentAttemptCurrencyARS captures enum value "ARS"
	PaymentAttemptCurrencyARS PaymentAttemptCurrencyEnum = "ARS"

	// PaymentAttemptCurrencyAUD captures enum value "AUD"
	PaymentAttemptCurrencyAUD PaymentAttemptCurrencyEnum = "AUD"

	// PaymentAttemptCurrencyAWG captures enum value "AWG"
	PaymentAttemptCurrencyAWG PaymentAttemptCurrencyEnum = "AWG"

	// PaymentAttemptCurrencyAZN captures enum value "AZN"
	PaymentAttemptCurrencyAZN PaymentAttemptCurrencyEnum = "AZN"

	// PaymentAttemptCurrencyBAM captures enum value "BAM"
	PaymentAttemptCurrencyBAM PaymentAttemptCurrencyEnum = "BAM"

	// PaymentAttemptCurrencyBBD captures enum value "BBD"
	PaymentAttemptCurrencyBBD PaymentAttemptCurrencyEnum = "BBD"

	// PaymentAttemptCurrencyBDT captures enum value "BDT"
	PaymentAttemptCurrencyBDT PaymentAttemptCurrencyEnum = "BDT"

	// PaymentAttemptCurrencyBGN captures enum value "BGN"
	PaymentAttemptCurrencyBGN PaymentAttemptCurrencyEnum = "BGN"

	// PaymentAttemptCurrencyBHD captures enum value "BHD"
	PaymentAttemptCurrencyBHD PaymentAttemptCurrencyEnum = "BHD"

	// PaymentAttemptCurrencyBIF captures enum value "BIF"
	PaymentAttemptCurrencyBIF PaymentAttemptCurrencyEnum = "BIF"

	// PaymentAttemptCurrencyBMD captures enum value "BMD"
	PaymentAttemptCurrencyBMD PaymentAttemptCurrencyEnum = "BMD"

	// PaymentAttemptCurrencyBND captures enum value "BND"
	PaymentAttemptCurrencyBND PaymentAttemptCurrencyEnum = "BND"

	// PaymentAttemptCurrencyBOB captures enum value "BOB"
	PaymentAttemptCurrencyBOB PaymentAttemptCurrencyEnum = "BOB"

	// PaymentAttemptCurrencyBRL captures enum value "BRL"
	PaymentAttemptCurrencyBRL PaymentAttemptCurrencyEnum = "BRL"

	// PaymentAttemptCurrencyBSD captures enum value "BSD"
	PaymentAttemptCurrencyBSD PaymentAttemptCurrencyEnum = "BSD"

	// PaymentAttemptCurrencyBTN captures enum value "BTN"
	PaymentAttemptCurrencyBTN PaymentAttemptCurrencyEnum = "BTN"

	// PaymentAttemptCurrencyBWP captures enum value "BWP"
	PaymentAttemptCurrencyBWP PaymentAttemptCurrencyEnum = "BWP"

	// PaymentAttemptCurrencyBYR captures enum value "BYR"
	PaymentAttemptCurrencyBYR PaymentAttemptCurrencyEnum = "BYR"

	// PaymentAttemptCurrencyBZD captures enum value "BZD"
	PaymentAttemptCurrencyBZD PaymentAttemptCurrencyEnum = "BZD"

	// PaymentAttemptCurrencyCAD captures enum value "CAD"
	PaymentAttemptCurrencyCAD PaymentAttemptCurrencyEnum = "CAD"

	// PaymentAttemptCurrencyCDF captures enum value "CDF"
	PaymentAttemptCurrencyCDF PaymentAttemptCurrencyEnum = "CDF"

	// PaymentAttemptCurrencyCHF captures enum value "CHF"
	PaymentAttemptCurrencyCHF PaymentAttemptCurrencyEnum = "CHF"

	// PaymentAttemptCurrencyCLP captures enum value "CLP"
	PaymentAttemptCurrencyCLP PaymentAttemptCurrencyEnum = "CLP"

	// PaymentAttemptCurrencyCNY captures enum value "CNY"
	PaymentAttemptCurrencyCNY PaymentAttemptCurrencyEnum = "CNY"

	// PaymentAttemptCurrencyCOP captures enum value "COP"
	PaymentAttemptCurrencyCOP PaymentAttemptCurrencyEnum = "COP"

	// PaymentAttemptCurrencyCRC captures enum value "CRC"
	PaymentAttemptCurrencyCRC PaymentAttemptCurrencyEnum = "CRC"

	// PaymentAttemptCurrencyCUC captures enum value "CUC"
	PaymentAttemptCurrencyCUC PaymentAttemptCurrencyEnum = "CUC"

	// PaymentAttemptCurrencyCUP captures enum value "CUP"
	PaymentAttemptCurrencyCUP PaymentAttemptCurrencyEnum = "CUP"

	// PaymentAttemptCurrencyCVE captures enum value "CVE"
	PaymentAttemptCurrencyCVE PaymentAttemptCurrencyEnum = "CVE"

	// PaymentAttemptCurrencyCZK captures enum value "CZK"
	PaymentAttemptCurrencyCZK PaymentAttemptCurrencyEnum = "CZK"

	// PaymentAttemptCurrencyDJF captures enum value "DJF"
	PaymentAttemptCurrencyDJF PaymentAttemptCurrencyEnum = "DJF"

	// PaymentAttemptCurrencyDKK captures enum value "DKK"
	PaymentAttemptCurrencyDKK PaymentAttemptCurrencyEnum = "DKK"

	// PaymentAttemptCurrencyDOP captures enum value "DOP"
	PaymentAttemptCurrencyDOP PaymentAttemptCurrencyEnum = "DOP"

	// PaymentAttemptCurrencyDZD captures enum value "DZD"
	PaymentAttemptCurrencyDZD PaymentAttemptCurrencyEnum = "DZD"

	// PaymentAttemptCurrencyEGP captures enum value "EGP"
	PaymentAttemptCurrencyEGP PaymentAttemptCurrencyEnum = "EGP"

	// PaymentAttemptCurrencyERN captures enum value "ERN"
	PaymentAttemptCurrencyERN PaymentAttemptCurrencyEnum = "ERN"

	// PaymentAttemptCurrencyETB captures enum value "ETB"
	PaymentAttemptCurrencyETB PaymentAttemptCurrencyEnum = "ETB"

	// PaymentAttemptCurrencyEUR captures enum value "EUR"
	PaymentAttemptCurrencyEUR PaymentAttemptCurrencyEnum = "EUR"

	// PaymentAttemptCurrencyFJD captures enum value "FJD"
	PaymentAttemptCurrencyFJD PaymentAttemptCurrencyEnum = "FJD"

	// PaymentAttemptCurrencyFKP captures enum value "FKP"
	PaymentAttemptCurrencyFKP PaymentAttemptCurrencyEnum = "FKP"

	// PaymentAttemptCurrencyGBP captures enum value "GBP"
	PaymentAttemptCurrencyGBP PaymentAttemptCurrencyEnum = "GBP"

	// PaymentAttemptCurrencyGEL captures enum value "GEL"
	PaymentAttemptCurrencyGEL PaymentAttemptCurrencyEnum = "GEL"

	// PaymentAttemptCurrencyGGP captures enum value "GGP"
	PaymentAttemptCurrencyGGP PaymentAttemptCurrencyEnum = "GGP"

	// PaymentAttemptCurrencyGHS captures enum value "GHS"
	PaymentAttemptCurrencyGHS PaymentAttemptCurrencyEnum = "GHS"

	// PaymentAttemptCurrencyGIP captures enum value "GIP"
	PaymentAttemptCurrencyGIP PaymentAttemptCurrencyEnum = "GIP"

	// PaymentAttemptCurrencyGMD captures enum value "GMD"
	PaymentAttemptCurrencyGMD PaymentAttemptCurrencyEnum = "GMD"

	// PaymentAttemptCurrencyGNF captures enum value "GNF"
	PaymentAttemptCurrencyGNF PaymentAttemptCurrencyEnum = "GNF"

	// PaymentAttemptCurrencyGTQ captures enum value "GTQ"
	PaymentAttemptCurrencyGTQ PaymentAttemptCurrencyEnum = "GTQ"

	// PaymentAttemptCurrencyGYD captures enum value "GYD"
	PaymentAttemptCurrencyGYD PaymentAttemptCurrencyEnum = "GYD"

	// PaymentAttemptCurrencyHKD captures enum value "HKD"
	PaymentAttemptCurrencyHKD PaymentAttemptCurrencyEnum = "HKD"

	// PaymentAttemptCurrencyHNL captures enum value "HNL"
	PaymentAttemptCurrencyHNL PaymentAttemptCurrencyEnum = "HNL"

	// PaymentAttemptCurrencyHRK captures enum value "HRK"
	PaymentAttemptCurrencyHRK PaymentAttemptCurrencyEnum = "HRK"

	// PaymentAttemptCurrencyHTG captures enum value "HTG"
	PaymentAttemptCurrencyHTG PaymentAttemptCurrencyEnum = "HTG"

	// PaymentAttemptCurrencyHUF captures enum value "HUF"
	PaymentAttemptCurrencyHUF PaymentAttemptCurrencyEnum = "HUF"

	// PaymentAttemptCurrencyIDR captures enum value "IDR"
	PaymentAttemptCurrencyIDR PaymentAttemptCurrencyEnum = "IDR"

	// PaymentAttemptCurrencyILS captures enum value "ILS"
	PaymentAttemptCurrencyILS PaymentAttemptCurrencyEnum = "ILS"

	// PaymentAttemptCurrencyIMP captures enum value "IMP"
	PaymentAttemptCurrencyIMP PaymentAttemptCurrencyEnum = "IMP"

	// PaymentAttemptCurrencyINR captures enum value "INR"
	PaymentAttemptCurrencyINR PaymentAttemptCurrencyEnum = "INR"

	// PaymentAttemptCurrencyIQD captures enum value "IQD"
	PaymentAttemptCurrencyIQD PaymentAttemptCurrencyEnum = "IQD"

	// PaymentAttemptCurrencyIRR captures enum value "IRR"
	PaymentAttemptCurrencyIRR PaymentAttemptCurrencyEnum = "IRR"

	// PaymentAttemptCurrencyISK captures enum value "ISK"
	PaymentAttemptCurrencyISK PaymentAttemptCurrencyEnum = "ISK"

	// PaymentAttemptCurrencyJEP captures enum value "JEP"
	PaymentAttemptCurrencyJEP PaymentAttemptCurrencyEnum = "JEP"

	// PaymentAttemptCurrencyJMD captures enum value "JMD"
	PaymentAttemptCurrencyJMD PaymentAttemptCurrencyEnum = "JMD"

	// PaymentAttemptCurrencyJOD captures enum value "JOD"
	PaymentAttemptCurrencyJOD PaymentAttemptCurrencyEnum = "JOD"

	// PaymentAttemptCurrencyJPY captures enum value "JPY"
	PaymentAttemptCurrencyJPY PaymentAttemptCurrencyEnum = "JPY"

	// PaymentAttemptCurrencyKES captures enum value "KES"
	PaymentAttemptCurrencyKES PaymentAttemptCurrencyEnum = "KES"

	// PaymentAttemptCurrencyKGS captures enum value "KGS"
	PaymentAttemptCurrencyKGS PaymentAttemptCurrencyEnum = "KGS"

	// PaymentAttemptCurrencyKHR captures enum value "KHR"
	PaymentAttemptCurrencyKHR PaymentAttemptCurrencyEnum = "KHR"

	// PaymentAttemptCurrencyKMF captures enum value "KMF"
	PaymentAttemptCurrencyKMF PaymentAttemptCurrencyEnum = "KMF"

	// PaymentAttemptCurrencyKPW captures enum value "KPW"
	PaymentAttemptCurrencyKPW PaymentAttemptCurrencyEnum = "KPW"

	// PaymentAttemptCurrencyKRW captures enum value "KRW"
	PaymentAttemptCurrencyKRW PaymentAttemptCurrencyEnum = "KRW"

	// PaymentAttemptCurrencyKWD captures enum value "KWD"
	PaymentAttemptCurrencyKWD PaymentAttemptCurrencyEnum = "KWD"

	// PaymentAttemptCurrencyKYD captures enum value "KYD"
	PaymentAttemptCurrencyKYD PaymentAttemptCurrencyEnum = "KYD"

	// PaymentAttemptCurrencyKZT captures enum value "KZT"
	PaymentAttemptCurrencyKZT PaymentAttemptCurrencyEnum = "KZT"

	// PaymentAttemptCurrencyLAK captures enum value "LAK"
	PaymentAttemptCurrencyLAK PaymentAttemptCurrencyEnum = "LAK"

	// PaymentAttemptCurrencyLBP captures enum value "LBP"
	PaymentAttemptCurrencyLBP PaymentAttemptCurrencyEnum = "LBP"

	// PaymentAttemptCurrencyLKR captures enum value "LKR"
	PaymentAttemptCurrencyLKR PaymentAttemptCurrencyEnum = "LKR"

	// PaymentAttemptCurrencyLRD captures enum value "LRD"
	PaymentAttemptCurrencyLRD PaymentAttemptCurrencyEnum = "LRD"

	// PaymentAttemptCurrencyLSL captures enum value "LSL"
	PaymentAttemptCurrencyLSL PaymentAttemptCurrencyEnum = "LSL"

	// PaymentAttemptCurrencyLTL captures enum value "LTL"
	PaymentAttemptCurrencyLTL PaymentAttemptCurrencyEnum = "LTL"

	// PaymentAttemptCurrencyLVL captures enum value "LVL"
	PaymentAttemptCurrencyLVL PaymentAttemptCurrencyEnum = "LVL"

	// PaymentAttemptCurrencyLYD captures enum value "LYD"
	PaymentAttemptCurrencyLYD PaymentAttemptCurrencyEnum = "LYD"

	// PaymentAttemptCurrencyMAD captures enum value "MAD"
	PaymentAttemptCurrencyMAD PaymentAttemptCurrencyEnum = "MAD"

	// PaymentAttemptCurrencyMDL captures enum value "MDL"
	PaymentAttemptCurrencyMDL PaymentAttemptCurrencyEnum = "MDL"

	// PaymentAttemptCurrencyMGA captures enum value "MGA"
	PaymentAttemptCurrencyMGA PaymentAttemptCurrencyEnum = "MGA"

	// PaymentAttemptCurrencyMKD captures enum value "MKD"
	PaymentAttemptCurrencyMKD PaymentAttemptCurrencyEnum = "MKD"

	// PaymentAttemptCurrencyMMK captures enum value "MMK"
	PaymentAttemptCurrencyMMK PaymentAttemptCurrencyEnum = "MMK"

	// PaymentAttemptCurrencyMNT captures enum value "MNT"
	PaymentAttemptCurrencyMNT PaymentAttemptCurrencyEnum = "MNT"

	// PaymentAttemptCurrencyMOP captures enum value "MOP"
	PaymentAttemptCurrencyMOP PaymentAttemptCurrencyEnum = "MOP"

	// PaymentAttemptCurrencyMRO captures enum value "MRO"
	PaymentAttemptCurrencyMRO PaymentAttemptCurrencyEnum = "MRO"

	// PaymentAttemptCurrencyMUR captures enum value "MUR"
	PaymentAttemptCurrencyMUR PaymentAttemptCurrencyEnum = "MUR"

	// PaymentAttemptCurrencyMVR captures enum value "MVR"
	PaymentAttemptCurrencyMVR PaymentAttemptCurrencyEnum = "MVR"

	// PaymentAttemptCurrencyMWK captures enum value "MWK"
	PaymentAttemptCurrencyMWK PaymentAttemptCurrencyEnum = "MWK"

	// PaymentAttemptCurrencyMXN captures enum value "MXN"
	PaymentAttemptCurrencyMXN PaymentAttemptCurrencyEnum = "MXN"

	// PaymentAttemptCurrencyMYR captures enum value "MYR"
	PaymentAttemptCurrencyMYR PaymentAttemptCurrencyEnum = "MYR"

	// PaymentAttemptCurrencyMZN captures enum value "MZN"
	PaymentAttemptCurrencyMZN PaymentAttemptCurrencyEnum = "MZN"

	// PaymentAttemptCurrencyNAD captures enum value "NAD"
	PaymentAttemptCurrencyNAD PaymentAttemptCurrencyEnum = "NAD"

	// PaymentAttemptCurrencyNGN captures enum value "NGN"
	PaymentAttemptCurrencyNGN PaymentAttemptCurrencyEnum = "NGN"

	// PaymentAttemptCurrencyNIO captures enum value "NIO"
	PaymentAttemptCurrencyNIO PaymentAttemptCurrencyEnum = "NIO"

	// PaymentAttemptCurrencyNOK captures enum value "NOK"
	PaymentAttemptCurrencyNOK PaymentAttemptCurrencyEnum = "NOK"

	// PaymentAttemptCurrencyNPR captures enum value "NPR"
	PaymentAttemptCurrencyNPR PaymentAttemptCurrencyEnum = "NPR"

	// PaymentAttemptCurrencyNZD captures enum value "NZD"
	PaymentAttemptCurrencyNZD PaymentAttemptCurrencyEnum = "NZD"

	// PaymentAttemptCurrencyOMR captures enum value "OMR"
	PaymentAttemptCurrencyOMR PaymentAttemptCurrencyEnum = "OMR"

	// PaymentAttemptCurrencyPAB captures enum value "PAB"
	PaymentAttemptCurrencyPAB PaymentAttemptCurrencyEnum = "PAB"

	// PaymentAttemptCurrencyPEN captures enum value "PEN"
	PaymentAttemptCurrencyPEN PaymentAttemptCurrencyEnum = "PEN"

	// PaymentAttemptCurrencyPGK captures enum value "PGK"
	PaymentAttemptCurrencyPGK PaymentAttemptCurrencyEnum = "PGK"

	// PaymentAttemptCurrencyPHP captures enum value "PHP"
	PaymentAttemptCurrencyPHP PaymentAttemptCurrencyEnum = "PHP"

	// PaymentAttemptCurrencyPKR captures enum value "PKR"
	PaymentAttemptCurrencyPKR PaymentAttemptCurrencyEnum = "PKR"

	// PaymentAttemptCurrencyPLN captures enum value "PLN"
	PaymentAttemptCurrencyPLN PaymentAttemptCurrencyEnum = "PLN"

	// PaymentAttemptCurrencyPYG captures enum value "PYG"
	PaymentAttemptCurrencyPYG PaymentAttemptCurrencyEnum = "PYG"

	// PaymentAttemptCurrencyQAR captures enum value "QAR"
	PaymentAttemptCurrencyQAR PaymentAttemptCurrencyEnum = "QAR"

	// PaymentAttemptCurrencyRON captures enum value "RON"
	PaymentAttemptCurrencyRON PaymentAttemptCurrencyEnum = "RON"

	// PaymentAttemptCurrencyRSD captures enum value "RSD"
	PaymentAttemptCurrencyRSD PaymentAttemptCurrencyEnum = "RSD"

	// PaymentAttemptCurrencyRUB captures enum value "RUB"
	PaymentAttemptCurrencyRUB PaymentAttemptCurrencyEnum = "RUB"

	// PaymentAttemptCurrencyRWF captures enum value "RWF"
	PaymentAttemptCurrencyRWF PaymentAttemptCurrencyEnum = "RWF"

	// PaymentAttemptCurrencySAR captures enum value "SAR"
	PaymentAttemptCurrencySAR PaymentAttemptCurrencyEnum = "SAR"

	// PaymentAttemptCurrencySBD captures enum value "SBD"
	PaymentAttemptCurrencySBD PaymentAttemptCurrencyEnum = "SBD"

	// PaymentAttemptCurrencySCR captures enum value "SCR"
	PaymentAttemptCurrencySCR PaymentAttemptCurrencyEnum = "SCR"

	// PaymentAttemptCurrencySDG captures enum value "SDG"
	PaymentAttemptCurrencySDG PaymentAttemptCurrencyEnum = "SDG"

	// PaymentAttemptCurrencySEK captures enum value "SEK"
	PaymentAttemptCurrencySEK PaymentAttemptCurrencyEnum = "SEK"

	// PaymentAttemptCurrencySGD captures enum value "SGD"
	PaymentAttemptCurrencySGD PaymentAttemptCurrencyEnum = "SGD"

	// PaymentAttemptCurrencySHP captures enum value "SHP"
	PaymentAttemptCurrencySHP PaymentAttemptCurrencyEnum = "SHP"

	// PaymentAttemptCurrencySLL captures enum value "SLL"
	PaymentAttemptCurrencySLL PaymentAttemptCurrencyEnum = "SLL"

	// PaymentAttemptCurrencySOS captures enum value "SOS"
	PaymentAttemptCurrencySOS PaymentAttemptCurrencyEnum = "SOS"

	// PaymentAttemptCurrencySPL captures enum value "SPL"
	PaymentAttemptCurrencySPL PaymentAttemptCurrencyEnum = "SPL"

	// PaymentAttemptCurrencySRD captures enum value "SRD"
	PaymentAttemptCurrencySRD PaymentAttemptCurrencyEnum = "SRD"

	// PaymentAttemptCurrencySTD captures enum value "STD"
	PaymentAttemptCurrencySTD PaymentAttemptCurrencyEnum = "STD"

	// PaymentAttemptCurrencySVC captures enum value "SVC"
	PaymentAttemptCurrencySVC PaymentAttemptCurrencyEnum = "SVC"

	// PaymentAttemptCurrencySYP captures enum value "SYP"
	PaymentAttemptCurrencySYP PaymentAttemptCurrencyEnum = "SYP"

	// PaymentAttemptCurrencySZL captures enum value "SZL"
	PaymentAttemptCurrencySZL PaymentAttemptCurrencyEnum = "SZL"

	// PaymentAttemptCurrencyTHB captures enum value "THB"
	PaymentAttemptCurrencyTHB PaymentAttemptCurrencyEnum = "THB"

	// PaymentAttemptCurrencyTJS captures enum value "TJS"
	PaymentAttemptCurrencyTJS PaymentAttemptCurrencyEnum = "TJS"

	// PaymentAttemptCurrencyTMT captures enum value "TMT"
	PaymentAttemptCurrencyTMT PaymentAttemptCurrencyEnum = "TMT"

	// PaymentAttemptCurrencyTND captures enum value "TND"
	PaymentAttemptCurrencyTND PaymentAttemptCurrencyEnum = "TND"

	// PaymentAttemptCurrencyTOP captures enum value "TOP"
	PaymentAttemptCurrencyTOP PaymentAttemptCurrencyEnum = "TOP"

	// PaymentAttemptCurrencyTRY captures enum value "TRY"
	PaymentAttemptCurrencyTRY PaymentAttemptCurrencyEnum = "TRY"

	// PaymentAttemptCurrencyTTD captures enum value "TTD"
	PaymentAttemptCurrencyTTD PaymentAttemptCurrencyEnum = "TTD"

	// PaymentAttemptCurrencyTVD captures enum value "TVD"
	PaymentAttemptCurrencyTVD PaymentAttemptCurrencyEnum = "TVD"

	// PaymentAttemptCurrencyTWD captures enum value "TWD"
	PaymentAttemptCurrencyTWD PaymentAttemptCurrencyEnum = "TWD"

	// PaymentAttemptCurrencyTZS captures enum value "TZS"
	PaymentAttemptCurrencyTZS PaymentAttemptCurrencyEnum = "TZS"

	// PaymentAttemptCurrencyUAH captures enum value "UAH"
	PaymentAttemptCurrencyUAH PaymentAttemptCurrencyEnum = "UAH"

	// PaymentAttemptCurrencyUGX captures enum value "UGX"
	PaymentAttemptCurrencyUGX PaymentAttemptCurrencyEnum = "UGX"

	// PaymentAttemptCurrencyUSD captures enum value "USD"
	PaymentAttemptCurrencyUSD PaymentAttemptCurrencyEnum = "USD"

	// PaymentAttemptCurrencyUYU captures enum value "UYU"
	PaymentAttemptCurrencyUYU PaymentAttemptCurrencyEnum = "UYU"

	// PaymentAttemptCurrencyUZS captures enum value "UZS"
	PaymentAttemptCurrencyUZS PaymentAttemptCurrencyEnum = "UZS"

	// PaymentAttemptCurrencyVEF captures enum value "VEF"
	PaymentAttemptCurrencyVEF PaymentAttemptCurrencyEnum = "VEF"

	// PaymentAttemptCurrencyVND captures enum value "VND"
	PaymentAttemptCurrencyVND PaymentAttemptCurrencyEnum = "VND"

	// PaymentAttemptCurrencyVUV captures enum value "VUV"
	PaymentAttemptCurrencyVUV PaymentAttemptCurrencyEnum = "VUV"

	// PaymentAttemptCurrencyWST captures enum value "WST"
	PaymentAttemptCurrencyWST PaymentAttemptCurrencyEnum = "WST"

	// PaymentAttemptCurrencyXAF captures enum value "XAF"
	PaymentAttemptCurrencyXAF PaymentAttemptCurrencyEnum = "XAF"

	// PaymentAttemptCurrencyXCD captures enum value "XCD"
	PaymentAttemptCurrencyXCD PaymentAttemptCurrencyEnum = "XCD"

	// PaymentAttemptCurrencyXDR captures enum value "XDR"
	PaymentAttemptCurrencyXDR PaymentAttemptCurrencyEnum = "XDR"

	// PaymentAttemptCurrencyXOF captures enum value "XOF"
	PaymentAttemptCurrencyXOF PaymentAttemptCurrencyEnum = "XOF"

	// PaymentAttemptCurrencyXPF captures enum value "XPF"
	PaymentAttemptCurrencyXPF PaymentAttemptCurrencyEnum = "XPF"

	// PaymentAttemptCurrencyYER captures enum value "YER"
	PaymentAttemptCurrencyYER PaymentAttemptCurrencyEnum = "YER"

	// PaymentAttemptCurrencyZAR captures enum value "ZAR"
	PaymentAttemptCurrencyZAR PaymentAttemptCurrencyEnum = "ZAR"

	// PaymentAttemptCurrencyZMW captures enum value "ZMW"
	PaymentAttemptCurrencyZMW PaymentAttemptCurrencyEnum = "ZMW"

	// PaymentAttemptCurrencyZWD captures enum value "ZWD"
	PaymentAttemptCurrencyZWD PaymentAttemptCurrencyEnum = "ZWD"

	// PaymentAttemptCurrencyBTC captures enum value "BTC"
	PaymentAttemptCurrencyBTC PaymentAttemptCurrencyEnum = "BTC"
)

var PaymentAttemptCurrencyEnumValues = []string{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GGP",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"IMP",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JEP",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SLL",
	"SOS",
	"SPL",
	"SRD",
	"STD",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TVD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XDR",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZWD",
	"BTC",
}

func (e PaymentAttemptCurrencyEnum) IsValid() bool {
	for _, v := range PaymentAttemptCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *PaymentAttempt) validateCurrencyEnum(path, location string, value PaymentAttemptCurrencyEnum) error {
	if err := validate.EnumCase(path, location, value, paymentAttemptTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAttempt) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttempt) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttempt) validatePaymentMethodID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentMethodId", "body", "uuid", m.PaymentMethodID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttempt) validatePluginProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.PluginProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.PluginProperties); i++ {
		if swag.IsZero(m.PluginProperties[i]) { // not required
			continue
		}

		if m.PluginProperties[i] != nil {
			if err := m.PluginProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaymentAttempt) validateTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionID) { // not required
		return nil
	}

	if err := validate.FormatOf("transactionId", "body", "uuid", m.TransactionID.String(), formats); err != nil {
		return err
	}

	return nil
}

var paymentAttemptTypeTransactionTypePropEnum []interface{}

func init() {
	var res []PaymentAttemptTransactionTypeEnum
	if err := json.Unmarshal([]byte(`["AUTHORIZE","CAPTURE","CHARGEBACK","CREDIT","PURCHASE","REFUND","VOID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAttemptTypeTransactionTypePropEnum = append(paymentAttemptTypeTransactionTypePropEnum, v)
	}
}

type PaymentAttemptTransactionTypeEnum string

const (

	// PaymentAttemptTransactionTypeAUTHORIZE captures enum value "AUTHORIZE"
	PaymentAttemptTransactionTypeAUTHORIZE PaymentAttemptTransactionTypeEnum = "AUTHORIZE"

	// PaymentAttemptTransactionTypeCAPTURE captures enum value "CAPTURE"
	PaymentAttemptTransactionTypeCAPTURE PaymentAttemptTransactionTypeEnum = "CAPTURE"

	// PaymentAttemptTransactionTypeCHARGEBACK captures enum value "CHARGEBACK"
	PaymentAttemptTransactionTypeCHARGEBACK PaymentAttemptTransactionTypeEnum = "CHARGEBACK"

	// PaymentAttemptTransactionTypeCREDIT captures enum value "CREDIT"
	PaymentAttemptTransactionTypeCREDIT PaymentAttemptTransactionTypeEnum = "CREDIT"

	// PaymentAttemptTransactionTypePURCHASE captures enum value "PURCHASE"
	PaymentAttemptTransactionTypePURCHASE PaymentAttemptTransactionTypeEnum = "PURCHASE"

	// PaymentAttemptTransactionTypeREFUND captures enum value "REFUND"
	PaymentAttemptTransactionTypeREFUND PaymentAttemptTransactionTypeEnum = "REFUND"

	// PaymentAttemptTransactionTypeVOID captures enum value "VOID"
	PaymentAttemptTransactionTypeVOID PaymentAttemptTransactionTypeEnum = "VOID"
)

var PaymentAttemptTransactionTypeEnumValues = []string{
	"AUTHORIZE",
	"CAPTURE",
	"CHARGEBACK",
	"CREDIT",
	"PURCHASE",
	"REFUND",
	"VOID",
}

func (e PaymentAttemptTransactionTypeEnum) IsValid() bool {
	for _, v := range PaymentAttemptTransactionTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *PaymentAttempt) validateTransactionTypeEnum(path, location string, value PaymentAttemptTransactionTypeEnum) error {
	if err := validate.EnumCase(path, location, value, paymentAttemptTypeTransactionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAttempt) validateTransactionType(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransactionTypeEnum("transactionType", "body", m.TransactionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payment attempt based on the context it is used
func (m *PaymentAttempt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePluginProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttempt) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaymentAttempt) contextValidatePluginProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PluginProperties); i++ {

		if m.PluginProperties[i] != nil {
			if err := m.PluginProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttempt) UnmarshalBinary(b []byte) error {
	var res PaymentAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
