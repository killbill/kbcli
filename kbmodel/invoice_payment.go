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

// InvoicePayment invoice payment
//
// swagger:model InvoicePayment
type InvoicePayment struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// auth amount
	AuthAmount float64 `json:"authAmount,omitempty"`

	// captured amount
	CapturedAmount float64 `json:"capturedAmount,omitempty"`

	// credited amount
	CreditedAmount float64 `json:"creditedAmount,omitempty"`

	// currency
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency InvoicePaymentCurrencyEnum `json:"currency,omitempty"`

	// payment attempts
	PaymentAttempts []*PaymentAttempt `json:"paymentAttempts"`

	// payment external key
	PaymentExternalKey string `json:"paymentExternalKey,omitempty"`

	// payment Id
	// Format: uuid
	PaymentID strfmt.UUID `json:"paymentId,omitempty"`

	// payment method Id
	// Format: uuid
	PaymentMethodID strfmt.UUID `json:"paymentMethodId,omitempty"`

	// payment number
	PaymentNumber string `json:"paymentNumber,omitempty"`

	// purchased amount
	PurchasedAmount float64 `json:"purchasedAmount,omitempty"`

	// refunded amount
	RefundedAmount float64 `json:"refundedAmount,omitempty"`

	// target invoice Id
	// Format: uuid
	TargetInvoiceID strfmt.UUID `json:"targetInvoiceId,omitempty"`

	// transactions
	Transactions []*PaymentTransaction `json:"transactions"`
}

// Validate validates this invoice payment
func (m *InvoicePayment) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePaymentAttempts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetInvoiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoicePayment) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateAuditLogs(formats strfmt.Registry) error {
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

var invoicePaymentTypeCurrencyPropEnum []interface{}

func init() {
	var res []InvoicePaymentCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoicePaymentTypeCurrencyPropEnum = append(invoicePaymentTypeCurrencyPropEnum, v)
	}
}

type InvoicePaymentCurrencyEnum string

const (

	// InvoicePaymentCurrencyAED captures enum value "AED"
	InvoicePaymentCurrencyAED InvoicePaymentCurrencyEnum = "AED"

	// InvoicePaymentCurrencyAFN captures enum value "AFN"
	InvoicePaymentCurrencyAFN InvoicePaymentCurrencyEnum = "AFN"

	// InvoicePaymentCurrencyALL captures enum value "ALL"
	InvoicePaymentCurrencyALL InvoicePaymentCurrencyEnum = "ALL"

	// InvoicePaymentCurrencyAMD captures enum value "AMD"
	InvoicePaymentCurrencyAMD InvoicePaymentCurrencyEnum = "AMD"

	// InvoicePaymentCurrencyANG captures enum value "ANG"
	InvoicePaymentCurrencyANG InvoicePaymentCurrencyEnum = "ANG"

	// InvoicePaymentCurrencyAOA captures enum value "AOA"
	InvoicePaymentCurrencyAOA InvoicePaymentCurrencyEnum = "AOA"

	// InvoicePaymentCurrencyARS captures enum value "ARS"
	InvoicePaymentCurrencyARS InvoicePaymentCurrencyEnum = "ARS"

	// InvoicePaymentCurrencyAUD captures enum value "AUD"
	InvoicePaymentCurrencyAUD InvoicePaymentCurrencyEnum = "AUD"

	// InvoicePaymentCurrencyAWG captures enum value "AWG"
	InvoicePaymentCurrencyAWG InvoicePaymentCurrencyEnum = "AWG"

	// InvoicePaymentCurrencyAZN captures enum value "AZN"
	InvoicePaymentCurrencyAZN InvoicePaymentCurrencyEnum = "AZN"

	// InvoicePaymentCurrencyBAM captures enum value "BAM"
	InvoicePaymentCurrencyBAM InvoicePaymentCurrencyEnum = "BAM"

	// InvoicePaymentCurrencyBBD captures enum value "BBD"
	InvoicePaymentCurrencyBBD InvoicePaymentCurrencyEnum = "BBD"

	// InvoicePaymentCurrencyBDT captures enum value "BDT"
	InvoicePaymentCurrencyBDT InvoicePaymentCurrencyEnum = "BDT"

	// InvoicePaymentCurrencyBGN captures enum value "BGN"
	InvoicePaymentCurrencyBGN InvoicePaymentCurrencyEnum = "BGN"

	// InvoicePaymentCurrencyBHD captures enum value "BHD"
	InvoicePaymentCurrencyBHD InvoicePaymentCurrencyEnum = "BHD"

	// InvoicePaymentCurrencyBIF captures enum value "BIF"
	InvoicePaymentCurrencyBIF InvoicePaymentCurrencyEnum = "BIF"

	// InvoicePaymentCurrencyBMD captures enum value "BMD"
	InvoicePaymentCurrencyBMD InvoicePaymentCurrencyEnum = "BMD"

	// InvoicePaymentCurrencyBND captures enum value "BND"
	InvoicePaymentCurrencyBND InvoicePaymentCurrencyEnum = "BND"

	// InvoicePaymentCurrencyBOB captures enum value "BOB"
	InvoicePaymentCurrencyBOB InvoicePaymentCurrencyEnum = "BOB"

	// InvoicePaymentCurrencyBRL captures enum value "BRL"
	InvoicePaymentCurrencyBRL InvoicePaymentCurrencyEnum = "BRL"

	// InvoicePaymentCurrencyBSD captures enum value "BSD"
	InvoicePaymentCurrencyBSD InvoicePaymentCurrencyEnum = "BSD"

	// InvoicePaymentCurrencyBTN captures enum value "BTN"
	InvoicePaymentCurrencyBTN InvoicePaymentCurrencyEnum = "BTN"

	// InvoicePaymentCurrencyBWP captures enum value "BWP"
	InvoicePaymentCurrencyBWP InvoicePaymentCurrencyEnum = "BWP"

	// InvoicePaymentCurrencyBYR captures enum value "BYR"
	InvoicePaymentCurrencyBYR InvoicePaymentCurrencyEnum = "BYR"

	// InvoicePaymentCurrencyBZD captures enum value "BZD"
	InvoicePaymentCurrencyBZD InvoicePaymentCurrencyEnum = "BZD"

	// InvoicePaymentCurrencyCAD captures enum value "CAD"
	InvoicePaymentCurrencyCAD InvoicePaymentCurrencyEnum = "CAD"

	// InvoicePaymentCurrencyCDF captures enum value "CDF"
	InvoicePaymentCurrencyCDF InvoicePaymentCurrencyEnum = "CDF"

	// InvoicePaymentCurrencyCHF captures enum value "CHF"
	InvoicePaymentCurrencyCHF InvoicePaymentCurrencyEnum = "CHF"

	// InvoicePaymentCurrencyCLP captures enum value "CLP"
	InvoicePaymentCurrencyCLP InvoicePaymentCurrencyEnum = "CLP"

	// InvoicePaymentCurrencyCNY captures enum value "CNY"
	InvoicePaymentCurrencyCNY InvoicePaymentCurrencyEnum = "CNY"

	// InvoicePaymentCurrencyCOP captures enum value "COP"
	InvoicePaymentCurrencyCOP InvoicePaymentCurrencyEnum = "COP"

	// InvoicePaymentCurrencyCRC captures enum value "CRC"
	InvoicePaymentCurrencyCRC InvoicePaymentCurrencyEnum = "CRC"

	// InvoicePaymentCurrencyCUC captures enum value "CUC"
	InvoicePaymentCurrencyCUC InvoicePaymentCurrencyEnum = "CUC"

	// InvoicePaymentCurrencyCUP captures enum value "CUP"
	InvoicePaymentCurrencyCUP InvoicePaymentCurrencyEnum = "CUP"

	// InvoicePaymentCurrencyCVE captures enum value "CVE"
	InvoicePaymentCurrencyCVE InvoicePaymentCurrencyEnum = "CVE"

	// InvoicePaymentCurrencyCZK captures enum value "CZK"
	InvoicePaymentCurrencyCZK InvoicePaymentCurrencyEnum = "CZK"

	// InvoicePaymentCurrencyDJF captures enum value "DJF"
	InvoicePaymentCurrencyDJF InvoicePaymentCurrencyEnum = "DJF"

	// InvoicePaymentCurrencyDKK captures enum value "DKK"
	InvoicePaymentCurrencyDKK InvoicePaymentCurrencyEnum = "DKK"

	// InvoicePaymentCurrencyDOP captures enum value "DOP"
	InvoicePaymentCurrencyDOP InvoicePaymentCurrencyEnum = "DOP"

	// InvoicePaymentCurrencyDZD captures enum value "DZD"
	InvoicePaymentCurrencyDZD InvoicePaymentCurrencyEnum = "DZD"

	// InvoicePaymentCurrencyEGP captures enum value "EGP"
	InvoicePaymentCurrencyEGP InvoicePaymentCurrencyEnum = "EGP"

	// InvoicePaymentCurrencyERN captures enum value "ERN"
	InvoicePaymentCurrencyERN InvoicePaymentCurrencyEnum = "ERN"

	// InvoicePaymentCurrencyETB captures enum value "ETB"
	InvoicePaymentCurrencyETB InvoicePaymentCurrencyEnum = "ETB"

	// InvoicePaymentCurrencyEUR captures enum value "EUR"
	InvoicePaymentCurrencyEUR InvoicePaymentCurrencyEnum = "EUR"

	// InvoicePaymentCurrencyFJD captures enum value "FJD"
	InvoicePaymentCurrencyFJD InvoicePaymentCurrencyEnum = "FJD"

	// InvoicePaymentCurrencyFKP captures enum value "FKP"
	InvoicePaymentCurrencyFKP InvoicePaymentCurrencyEnum = "FKP"

	// InvoicePaymentCurrencyGBP captures enum value "GBP"
	InvoicePaymentCurrencyGBP InvoicePaymentCurrencyEnum = "GBP"

	// InvoicePaymentCurrencyGEL captures enum value "GEL"
	InvoicePaymentCurrencyGEL InvoicePaymentCurrencyEnum = "GEL"

	// InvoicePaymentCurrencyGGP captures enum value "GGP"
	InvoicePaymentCurrencyGGP InvoicePaymentCurrencyEnum = "GGP"

	// InvoicePaymentCurrencyGHS captures enum value "GHS"
	InvoicePaymentCurrencyGHS InvoicePaymentCurrencyEnum = "GHS"

	// InvoicePaymentCurrencyGIP captures enum value "GIP"
	InvoicePaymentCurrencyGIP InvoicePaymentCurrencyEnum = "GIP"

	// InvoicePaymentCurrencyGMD captures enum value "GMD"
	InvoicePaymentCurrencyGMD InvoicePaymentCurrencyEnum = "GMD"

	// InvoicePaymentCurrencyGNF captures enum value "GNF"
	InvoicePaymentCurrencyGNF InvoicePaymentCurrencyEnum = "GNF"

	// InvoicePaymentCurrencyGTQ captures enum value "GTQ"
	InvoicePaymentCurrencyGTQ InvoicePaymentCurrencyEnum = "GTQ"

	// InvoicePaymentCurrencyGYD captures enum value "GYD"
	InvoicePaymentCurrencyGYD InvoicePaymentCurrencyEnum = "GYD"

	// InvoicePaymentCurrencyHKD captures enum value "HKD"
	InvoicePaymentCurrencyHKD InvoicePaymentCurrencyEnum = "HKD"

	// InvoicePaymentCurrencyHNL captures enum value "HNL"
	InvoicePaymentCurrencyHNL InvoicePaymentCurrencyEnum = "HNL"

	// InvoicePaymentCurrencyHRK captures enum value "HRK"
	InvoicePaymentCurrencyHRK InvoicePaymentCurrencyEnum = "HRK"

	// InvoicePaymentCurrencyHTG captures enum value "HTG"
	InvoicePaymentCurrencyHTG InvoicePaymentCurrencyEnum = "HTG"

	// InvoicePaymentCurrencyHUF captures enum value "HUF"
	InvoicePaymentCurrencyHUF InvoicePaymentCurrencyEnum = "HUF"

	// InvoicePaymentCurrencyIDR captures enum value "IDR"
	InvoicePaymentCurrencyIDR InvoicePaymentCurrencyEnum = "IDR"

	// InvoicePaymentCurrencyILS captures enum value "ILS"
	InvoicePaymentCurrencyILS InvoicePaymentCurrencyEnum = "ILS"

	// InvoicePaymentCurrencyIMP captures enum value "IMP"
	InvoicePaymentCurrencyIMP InvoicePaymentCurrencyEnum = "IMP"

	// InvoicePaymentCurrencyINR captures enum value "INR"
	InvoicePaymentCurrencyINR InvoicePaymentCurrencyEnum = "INR"

	// InvoicePaymentCurrencyIQD captures enum value "IQD"
	InvoicePaymentCurrencyIQD InvoicePaymentCurrencyEnum = "IQD"

	// InvoicePaymentCurrencyIRR captures enum value "IRR"
	InvoicePaymentCurrencyIRR InvoicePaymentCurrencyEnum = "IRR"

	// InvoicePaymentCurrencyISK captures enum value "ISK"
	InvoicePaymentCurrencyISK InvoicePaymentCurrencyEnum = "ISK"

	// InvoicePaymentCurrencyJEP captures enum value "JEP"
	InvoicePaymentCurrencyJEP InvoicePaymentCurrencyEnum = "JEP"

	// InvoicePaymentCurrencyJMD captures enum value "JMD"
	InvoicePaymentCurrencyJMD InvoicePaymentCurrencyEnum = "JMD"

	// InvoicePaymentCurrencyJOD captures enum value "JOD"
	InvoicePaymentCurrencyJOD InvoicePaymentCurrencyEnum = "JOD"

	// InvoicePaymentCurrencyJPY captures enum value "JPY"
	InvoicePaymentCurrencyJPY InvoicePaymentCurrencyEnum = "JPY"

	// InvoicePaymentCurrencyKES captures enum value "KES"
	InvoicePaymentCurrencyKES InvoicePaymentCurrencyEnum = "KES"

	// InvoicePaymentCurrencyKGS captures enum value "KGS"
	InvoicePaymentCurrencyKGS InvoicePaymentCurrencyEnum = "KGS"

	// InvoicePaymentCurrencyKHR captures enum value "KHR"
	InvoicePaymentCurrencyKHR InvoicePaymentCurrencyEnum = "KHR"

	// InvoicePaymentCurrencyKMF captures enum value "KMF"
	InvoicePaymentCurrencyKMF InvoicePaymentCurrencyEnum = "KMF"

	// InvoicePaymentCurrencyKPW captures enum value "KPW"
	InvoicePaymentCurrencyKPW InvoicePaymentCurrencyEnum = "KPW"

	// InvoicePaymentCurrencyKRW captures enum value "KRW"
	InvoicePaymentCurrencyKRW InvoicePaymentCurrencyEnum = "KRW"

	// InvoicePaymentCurrencyKWD captures enum value "KWD"
	InvoicePaymentCurrencyKWD InvoicePaymentCurrencyEnum = "KWD"

	// InvoicePaymentCurrencyKYD captures enum value "KYD"
	InvoicePaymentCurrencyKYD InvoicePaymentCurrencyEnum = "KYD"

	// InvoicePaymentCurrencyKZT captures enum value "KZT"
	InvoicePaymentCurrencyKZT InvoicePaymentCurrencyEnum = "KZT"

	// InvoicePaymentCurrencyLAK captures enum value "LAK"
	InvoicePaymentCurrencyLAK InvoicePaymentCurrencyEnum = "LAK"

	// InvoicePaymentCurrencyLBP captures enum value "LBP"
	InvoicePaymentCurrencyLBP InvoicePaymentCurrencyEnum = "LBP"

	// InvoicePaymentCurrencyLKR captures enum value "LKR"
	InvoicePaymentCurrencyLKR InvoicePaymentCurrencyEnum = "LKR"

	// InvoicePaymentCurrencyLRD captures enum value "LRD"
	InvoicePaymentCurrencyLRD InvoicePaymentCurrencyEnum = "LRD"

	// InvoicePaymentCurrencyLSL captures enum value "LSL"
	InvoicePaymentCurrencyLSL InvoicePaymentCurrencyEnum = "LSL"

	// InvoicePaymentCurrencyLTL captures enum value "LTL"
	InvoicePaymentCurrencyLTL InvoicePaymentCurrencyEnum = "LTL"

	// InvoicePaymentCurrencyLVL captures enum value "LVL"
	InvoicePaymentCurrencyLVL InvoicePaymentCurrencyEnum = "LVL"

	// InvoicePaymentCurrencyLYD captures enum value "LYD"
	InvoicePaymentCurrencyLYD InvoicePaymentCurrencyEnum = "LYD"

	// InvoicePaymentCurrencyMAD captures enum value "MAD"
	InvoicePaymentCurrencyMAD InvoicePaymentCurrencyEnum = "MAD"

	// InvoicePaymentCurrencyMDL captures enum value "MDL"
	InvoicePaymentCurrencyMDL InvoicePaymentCurrencyEnum = "MDL"

	// InvoicePaymentCurrencyMGA captures enum value "MGA"
	InvoicePaymentCurrencyMGA InvoicePaymentCurrencyEnum = "MGA"

	// InvoicePaymentCurrencyMKD captures enum value "MKD"
	InvoicePaymentCurrencyMKD InvoicePaymentCurrencyEnum = "MKD"

	// InvoicePaymentCurrencyMMK captures enum value "MMK"
	InvoicePaymentCurrencyMMK InvoicePaymentCurrencyEnum = "MMK"

	// InvoicePaymentCurrencyMNT captures enum value "MNT"
	InvoicePaymentCurrencyMNT InvoicePaymentCurrencyEnum = "MNT"

	// InvoicePaymentCurrencyMOP captures enum value "MOP"
	InvoicePaymentCurrencyMOP InvoicePaymentCurrencyEnum = "MOP"

	// InvoicePaymentCurrencyMRO captures enum value "MRO"
	InvoicePaymentCurrencyMRO InvoicePaymentCurrencyEnum = "MRO"

	// InvoicePaymentCurrencyMUR captures enum value "MUR"
	InvoicePaymentCurrencyMUR InvoicePaymentCurrencyEnum = "MUR"

	// InvoicePaymentCurrencyMVR captures enum value "MVR"
	InvoicePaymentCurrencyMVR InvoicePaymentCurrencyEnum = "MVR"

	// InvoicePaymentCurrencyMWK captures enum value "MWK"
	InvoicePaymentCurrencyMWK InvoicePaymentCurrencyEnum = "MWK"

	// InvoicePaymentCurrencyMXN captures enum value "MXN"
	InvoicePaymentCurrencyMXN InvoicePaymentCurrencyEnum = "MXN"

	// InvoicePaymentCurrencyMYR captures enum value "MYR"
	InvoicePaymentCurrencyMYR InvoicePaymentCurrencyEnum = "MYR"

	// InvoicePaymentCurrencyMZN captures enum value "MZN"
	InvoicePaymentCurrencyMZN InvoicePaymentCurrencyEnum = "MZN"

	// InvoicePaymentCurrencyNAD captures enum value "NAD"
	InvoicePaymentCurrencyNAD InvoicePaymentCurrencyEnum = "NAD"

	// InvoicePaymentCurrencyNGN captures enum value "NGN"
	InvoicePaymentCurrencyNGN InvoicePaymentCurrencyEnum = "NGN"

	// InvoicePaymentCurrencyNIO captures enum value "NIO"
	InvoicePaymentCurrencyNIO InvoicePaymentCurrencyEnum = "NIO"

	// InvoicePaymentCurrencyNOK captures enum value "NOK"
	InvoicePaymentCurrencyNOK InvoicePaymentCurrencyEnum = "NOK"

	// InvoicePaymentCurrencyNPR captures enum value "NPR"
	InvoicePaymentCurrencyNPR InvoicePaymentCurrencyEnum = "NPR"

	// InvoicePaymentCurrencyNZD captures enum value "NZD"
	InvoicePaymentCurrencyNZD InvoicePaymentCurrencyEnum = "NZD"

	// InvoicePaymentCurrencyOMR captures enum value "OMR"
	InvoicePaymentCurrencyOMR InvoicePaymentCurrencyEnum = "OMR"

	// InvoicePaymentCurrencyPAB captures enum value "PAB"
	InvoicePaymentCurrencyPAB InvoicePaymentCurrencyEnum = "PAB"

	// InvoicePaymentCurrencyPEN captures enum value "PEN"
	InvoicePaymentCurrencyPEN InvoicePaymentCurrencyEnum = "PEN"

	// InvoicePaymentCurrencyPGK captures enum value "PGK"
	InvoicePaymentCurrencyPGK InvoicePaymentCurrencyEnum = "PGK"

	// InvoicePaymentCurrencyPHP captures enum value "PHP"
	InvoicePaymentCurrencyPHP InvoicePaymentCurrencyEnum = "PHP"

	// InvoicePaymentCurrencyPKR captures enum value "PKR"
	InvoicePaymentCurrencyPKR InvoicePaymentCurrencyEnum = "PKR"

	// InvoicePaymentCurrencyPLN captures enum value "PLN"
	InvoicePaymentCurrencyPLN InvoicePaymentCurrencyEnum = "PLN"

	// InvoicePaymentCurrencyPYG captures enum value "PYG"
	InvoicePaymentCurrencyPYG InvoicePaymentCurrencyEnum = "PYG"

	// InvoicePaymentCurrencyQAR captures enum value "QAR"
	InvoicePaymentCurrencyQAR InvoicePaymentCurrencyEnum = "QAR"

	// InvoicePaymentCurrencyRON captures enum value "RON"
	InvoicePaymentCurrencyRON InvoicePaymentCurrencyEnum = "RON"

	// InvoicePaymentCurrencyRSD captures enum value "RSD"
	InvoicePaymentCurrencyRSD InvoicePaymentCurrencyEnum = "RSD"

	// InvoicePaymentCurrencyRUB captures enum value "RUB"
	InvoicePaymentCurrencyRUB InvoicePaymentCurrencyEnum = "RUB"

	// InvoicePaymentCurrencyRWF captures enum value "RWF"
	InvoicePaymentCurrencyRWF InvoicePaymentCurrencyEnum = "RWF"

	// InvoicePaymentCurrencySAR captures enum value "SAR"
	InvoicePaymentCurrencySAR InvoicePaymentCurrencyEnum = "SAR"

	// InvoicePaymentCurrencySBD captures enum value "SBD"
	InvoicePaymentCurrencySBD InvoicePaymentCurrencyEnum = "SBD"

	// InvoicePaymentCurrencySCR captures enum value "SCR"
	InvoicePaymentCurrencySCR InvoicePaymentCurrencyEnum = "SCR"

	// InvoicePaymentCurrencySDG captures enum value "SDG"
	InvoicePaymentCurrencySDG InvoicePaymentCurrencyEnum = "SDG"

	// InvoicePaymentCurrencySEK captures enum value "SEK"
	InvoicePaymentCurrencySEK InvoicePaymentCurrencyEnum = "SEK"

	// InvoicePaymentCurrencySGD captures enum value "SGD"
	InvoicePaymentCurrencySGD InvoicePaymentCurrencyEnum = "SGD"

	// InvoicePaymentCurrencySHP captures enum value "SHP"
	InvoicePaymentCurrencySHP InvoicePaymentCurrencyEnum = "SHP"

	// InvoicePaymentCurrencySLL captures enum value "SLL"
	InvoicePaymentCurrencySLL InvoicePaymentCurrencyEnum = "SLL"

	// InvoicePaymentCurrencySOS captures enum value "SOS"
	InvoicePaymentCurrencySOS InvoicePaymentCurrencyEnum = "SOS"

	// InvoicePaymentCurrencySPL captures enum value "SPL"
	InvoicePaymentCurrencySPL InvoicePaymentCurrencyEnum = "SPL"

	// InvoicePaymentCurrencySRD captures enum value "SRD"
	InvoicePaymentCurrencySRD InvoicePaymentCurrencyEnum = "SRD"

	// InvoicePaymentCurrencySTD captures enum value "STD"
	InvoicePaymentCurrencySTD InvoicePaymentCurrencyEnum = "STD"

	// InvoicePaymentCurrencySVC captures enum value "SVC"
	InvoicePaymentCurrencySVC InvoicePaymentCurrencyEnum = "SVC"

	// InvoicePaymentCurrencySYP captures enum value "SYP"
	InvoicePaymentCurrencySYP InvoicePaymentCurrencyEnum = "SYP"

	// InvoicePaymentCurrencySZL captures enum value "SZL"
	InvoicePaymentCurrencySZL InvoicePaymentCurrencyEnum = "SZL"

	// InvoicePaymentCurrencyTHB captures enum value "THB"
	InvoicePaymentCurrencyTHB InvoicePaymentCurrencyEnum = "THB"

	// InvoicePaymentCurrencyTJS captures enum value "TJS"
	InvoicePaymentCurrencyTJS InvoicePaymentCurrencyEnum = "TJS"

	// InvoicePaymentCurrencyTMT captures enum value "TMT"
	InvoicePaymentCurrencyTMT InvoicePaymentCurrencyEnum = "TMT"

	// InvoicePaymentCurrencyTND captures enum value "TND"
	InvoicePaymentCurrencyTND InvoicePaymentCurrencyEnum = "TND"

	// InvoicePaymentCurrencyTOP captures enum value "TOP"
	InvoicePaymentCurrencyTOP InvoicePaymentCurrencyEnum = "TOP"

	// InvoicePaymentCurrencyTRY captures enum value "TRY"
	InvoicePaymentCurrencyTRY InvoicePaymentCurrencyEnum = "TRY"

	// InvoicePaymentCurrencyTTD captures enum value "TTD"
	InvoicePaymentCurrencyTTD InvoicePaymentCurrencyEnum = "TTD"

	// InvoicePaymentCurrencyTVD captures enum value "TVD"
	InvoicePaymentCurrencyTVD InvoicePaymentCurrencyEnum = "TVD"

	// InvoicePaymentCurrencyTWD captures enum value "TWD"
	InvoicePaymentCurrencyTWD InvoicePaymentCurrencyEnum = "TWD"

	// InvoicePaymentCurrencyTZS captures enum value "TZS"
	InvoicePaymentCurrencyTZS InvoicePaymentCurrencyEnum = "TZS"

	// InvoicePaymentCurrencyUAH captures enum value "UAH"
	InvoicePaymentCurrencyUAH InvoicePaymentCurrencyEnum = "UAH"

	// InvoicePaymentCurrencyUGX captures enum value "UGX"
	InvoicePaymentCurrencyUGX InvoicePaymentCurrencyEnum = "UGX"

	// InvoicePaymentCurrencyUSD captures enum value "USD"
	InvoicePaymentCurrencyUSD InvoicePaymentCurrencyEnum = "USD"

	// InvoicePaymentCurrencyUYU captures enum value "UYU"
	InvoicePaymentCurrencyUYU InvoicePaymentCurrencyEnum = "UYU"

	// InvoicePaymentCurrencyUZS captures enum value "UZS"
	InvoicePaymentCurrencyUZS InvoicePaymentCurrencyEnum = "UZS"

	// InvoicePaymentCurrencyVEF captures enum value "VEF"
	InvoicePaymentCurrencyVEF InvoicePaymentCurrencyEnum = "VEF"

	// InvoicePaymentCurrencyVND captures enum value "VND"
	InvoicePaymentCurrencyVND InvoicePaymentCurrencyEnum = "VND"

	// InvoicePaymentCurrencyVUV captures enum value "VUV"
	InvoicePaymentCurrencyVUV InvoicePaymentCurrencyEnum = "VUV"

	// InvoicePaymentCurrencyWST captures enum value "WST"
	InvoicePaymentCurrencyWST InvoicePaymentCurrencyEnum = "WST"

	// InvoicePaymentCurrencyXAF captures enum value "XAF"
	InvoicePaymentCurrencyXAF InvoicePaymentCurrencyEnum = "XAF"

	// InvoicePaymentCurrencyXCD captures enum value "XCD"
	InvoicePaymentCurrencyXCD InvoicePaymentCurrencyEnum = "XCD"

	// InvoicePaymentCurrencyXDR captures enum value "XDR"
	InvoicePaymentCurrencyXDR InvoicePaymentCurrencyEnum = "XDR"

	// InvoicePaymentCurrencyXOF captures enum value "XOF"
	InvoicePaymentCurrencyXOF InvoicePaymentCurrencyEnum = "XOF"

	// InvoicePaymentCurrencyXPF captures enum value "XPF"
	InvoicePaymentCurrencyXPF InvoicePaymentCurrencyEnum = "XPF"

	// InvoicePaymentCurrencyYER captures enum value "YER"
	InvoicePaymentCurrencyYER InvoicePaymentCurrencyEnum = "YER"

	// InvoicePaymentCurrencyZAR captures enum value "ZAR"
	InvoicePaymentCurrencyZAR InvoicePaymentCurrencyEnum = "ZAR"

	// InvoicePaymentCurrencyZMW captures enum value "ZMW"
	InvoicePaymentCurrencyZMW InvoicePaymentCurrencyEnum = "ZMW"

	// InvoicePaymentCurrencyZWD captures enum value "ZWD"
	InvoicePaymentCurrencyZWD InvoicePaymentCurrencyEnum = "ZWD"

	// InvoicePaymentCurrencyBTC captures enum value "BTC"
	InvoicePaymentCurrencyBTC InvoicePaymentCurrencyEnum = "BTC"
)

var InvoicePaymentCurrencyEnumValues = []string{
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

func (e InvoicePaymentCurrencyEnum) IsValid() bool {
	for _, v := range InvoicePaymentCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *InvoicePayment) validateCurrencyEnum(path, location string, value InvoicePaymentCurrencyEnum) error {
	if err := validate.EnumCase(path, location, value, invoicePaymentTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoicePayment) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validatePaymentAttempts(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentAttempts) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentAttempts); i++ {
		if swag.IsZero(m.PaymentAttempts[i]) { // not required
			continue
		}

		if m.PaymentAttempts[i] != nil {
			if err := m.PaymentAttempts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paymentAttempts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paymentAttempts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoicePayment) validatePaymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentId", "body", "uuid", m.PaymentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validatePaymentMethodID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentMethodId", "body", "uuid", m.PaymentMethodID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateTargetInvoiceID(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetInvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("targetInvoiceId", "body", "uuid", m.TargetInvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateTransactions(formats strfmt.Registry) error {
	if swag.IsZero(m.Transactions) { // not required
		return nil
	}

	for i := 0; i < len(m.Transactions); i++ {
		if swag.IsZero(m.Transactions[i]) { // not required
			continue
		}

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this invoice payment based on the context it is used
func (m *InvoicePayment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaymentAttempts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoicePayment) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InvoicePayment) contextValidatePaymentAttempts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PaymentAttempts); i++ {

		if m.PaymentAttempts[i] != nil {
			if err := m.PaymentAttempts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paymentAttempts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paymentAttempts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoicePayment) contextValidateTransactions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Transactions); i++ {

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoicePayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoicePayment) UnmarshalBinary(b []byte) error {
	var res InvoicePayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
