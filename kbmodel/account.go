// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
// swagger:model Account
type Account struct {

	// account balance
	AccountBalance float64 `json:"accountBalance,omitempty"`

	// account c b a
	AccountCBA float64 `json:"accountCBA,omitempty"`

	// account Id
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// address1
	Address1 string `json:"address1,omitempty"`

	// address2
	Address2 string `json:"address2,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bill cycle day local
	BillCycleDayLocal int32 `json:"billCycleDayLocal,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// currency
	Currency AccountCurrencyEnum `json:"currency,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`

	// first name length
	FirstNameLength int32 `json:"firstNameLength,omitempty"`

	// is migrated
	IsMigrated *bool `json:"isMigrated,omitempty"`

	// is notified for invoices
	IsNotifiedForInvoices *bool `json:"isNotifiedForInvoices,omitempty"`

	// is payment delegated to parent
	IsPaymentDelegatedToParent *bool `json:"isPaymentDelegatedToParent,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notes
	Notes string `json:"notes,omitempty"`

	// parent account Id
	ParentAccountID strfmt.UUID `json:"parentAccountId,omitempty"`

	// payment method Id
	PaymentMethodID strfmt.UUID `json:"paymentMethodId,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// postal code
	PostalCode string `json:"postalCode,omitempty"`

	// reference time
	ReferenceTime strfmt.DateTime `json:"referenceTime,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParentAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReferenceTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateAuditLogs(formats strfmt.Registry) error {

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
				}
				return err
			}

		}

	}

	return nil
}

var accountTypeCurrencyPropEnum []interface{}

func init() {
	var res []AccountCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeCurrencyPropEnum = append(accountTypeCurrencyPropEnum, v)
	}
}

type AccountCurrencyEnum string

const (

	// AccountCurrencyAED captures enum value "AED"
	AccountCurrencyAED AccountCurrencyEnum = "AED"

	// AccountCurrencyAFN captures enum value "AFN"
	AccountCurrencyAFN AccountCurrencyEnum = "AFN"

	// AccountCurrencyALL captures enum value "ALL"
	AccountCurrencyALL AccountCurrencyEnum = "ALL"

	// AccountCurrencyAMD captures enum value "AMD"
	AccountCurrencyAMD AccountCurrencyEnum = "AMD"

	// AccountCurrencyANG captures enum value "ANG"
	AccountCurrencyANG AccountCurrencyEnum = "ANG"

	// AccountCurrencyAOA captures enum value "AOA"
	AccountCurrencyAOA AccountCurrencyEnum = "AOA"

	// AccountCurrencyARS captures enum value "ARS"
	AccountCurrencyARS AccountCurrencyEnum = "ARS"

	// AccountCurrencyAUD captures enum value "AUD"
	AccountCurrencyAUD AccountCurrencyEnum = "AUD"

	// AccountCurrencyAWG captures enum value "AWG"
	AccountCurrencyAWG AccountCurrencyEnum = "AWG"

	// AccountCurrencyAZN captures enum value "AZN"
	AccountCurrencyAZN AccountCurrencyEnum = "AZN"

	// AccountCurrencyBAM captures enum value "BAM"
	AccountCurrencyBAM AccountCurrencyEnum = "BAM"

	// AccountCurrencyBBD captures enum value "BBD"
	AccountCurrencyBBD AccountCurrencyEnum = "BBD"

	// AccountCurrencyBDT captures enum value "BDT"
	AccountCurrencyBDT AccountCurrencyEnum = "BDT"

	// AccountCurrencyBGN captures enum value "BGN"
	AccountCurrencyBGN AccountCurrencyEnum = "BGN"

	// AccountCurrencyBHD captures enum value "BHD"
	AccountCurrencyBHD AccountCurrencyEnum = "BHD"

	// AccountCurrencyBIF captures enum value "BIF"
	AccountCurrencyBIF AccountCurrencyEnum = "BIF"

	// AccountCurrencyBMD captures enum value "BMD"
	AccountCurrencyBMD AccountCurrencyEnum = "BMD"

	// AccountCurrencyBND captures enum value "BND"
	AccountCurrencyBND AccountCurrencyEnum = "BND"

	// AccountCurrencyBOB captures enum value "BOB"
	AccountCurrencyBOB AccountCurrencyEnum = "BOB"

	// AccountCurrencyBRL captures enum value "BRL"
	AccountCurrencyBRL AccountCurrencyEnum = "BRL"

	// AccountCurrencyBSD captures enum value "BSD"
	AccountCurrencyBSD AccountCurrencyEnum = "BSD"

	// AccountCurrencyBTN captures enum value "BTN"
	AccountCurrencyBTN AccountCurrencyEnum = "BTN"

	// AccountCurrencyBWP captures enum value "BWP"
	AccountCurrencyBWP AccountCurrencyEnum = "BWP"

	// AccountCurrencyBYR captures enum value "BYR"
	AccountCurrencyBYR AccountCurrencyEnum = "BYR"

	// AccountCurrencyBZD captures enum value "BZD"
	AccountCurrencyBZD AccountCurrencyEnum = "BZD"

	// AccountCurrencyCAD captures enum value "CAD"
	AccountCurrencyCAD AccountCurrencyEnum = "CAD"

	// AccountCurrencyCDF captures enum value "CDF"
	AccountCurrencyCDF AccountCurrencyEnum = "CDF"

	// AccountCurrencyCHF captures enum value "CHF"
	AccountCurrencyCHF AccountCurrencyEnum = "CHF"

	// AccountCurrencyCLP captures enum value "CLP"
	AccountCurrencyCLP AccountCurrencyEnum = "CLP"

	// AccountCurrencyCNY captures enum value "CNY"
	AccountCurrencyCNY AccountCurrencyEnum = "CNY"

	// AccountCurrencyCOP captures enum value "COP"
	AccountCurrencyCOP AccountCurrencyEnum = "COP"

	// AccountCurrencyCRC captures enum value "CRC"
	AccountCurrencyCRC AccountCurrencyEnum = "CRC"

	// AccountCurrencyCUC captures enum value "CUC"
	AccountCurrencyCUC AccountCurrencyEnum = "CUC"

	// AccountCurrencyCUP captures enum value "CUP"
	AccountCurrencyCUP AccountCurrencyEnum = "CUP"

	// AccountCurrencyCVE captures enum value "CVE"
	AccountCurrencyCVE AccountCurrencyEnum = "CVE"

	// AccountCurrencyCZK captures enum value "CZK"
	AccountCurrencyCZK AccountCurrencyEnum = "CZK"

	// AccountCurrencyDJF captures enum value "DJF"
	AccountCurrencyDJF AccountCurrencyEnum = "DJF"

	// AccountCurrencyDKK captures enum value "DKK"
	AccountCurrencyDKK AccountCurrencyEnum = "DKK"

	// AccountCurrencyDOP captures enum value "DOP"
	AccountCurrencyDOP AccountCurrencyEnum = "DOP"

	// AccountCurrencyDZD captures enum value "DZD"
	AccountCurrencyDZD AccountCurrencyEnum = "DZD"

	// AccountCurrencyEGP captures enum value "EGP"
	AccountCurrencyEGP AccountCurrencyEnum = "EGP"

	// AccountCurrencyERN captures enum value "ERN"
	AccountCurrencyERN AccountCurrencyEnum = "ERN"

	// AccountCurrencyETB captures enum value "ETB"
	AccountCurrencyETB AccountCurrencyEnum = "ETB"

	// AccountCurrencyEUR captures enum value "EUR"
	AccountCurrencyEUR AccountCurrencyEnum = "EUR"

	// AccountCurrencyFJD captures enum value "FJD"
	AccountCurrencyFJD AccountCurrencyEnum = "FJD"

	// AccountCurrencyFKP captures enum value "FKP"
	AccountCurrencyFKP AccountCurrencyEnum = "FKP"

	// AccountCurrencyGBP captures enum value "GBP"
	AccountCurrencyGBP AccountCurrencyEnum = "GBP"

	// AccountCurrencyGEL captures enum value "GEL"
	AccountCurrencyGEL AccountCurrencyEnum = "GEL"

	// AccountCurrencyGGP captures enum value "GGP"
	AccountCurrencyGGP AccountCurrencyEnum = "GGP"

	// AccountCurrencyGHS captures enum value "GHS"
	AccountCurrencyGHS AccountCurrencyEnum = "GHS"

	// AccountCurrencyGIP captures enum value "GIP"
	AccountCurrencyGIP AccountCurrencyEnum = "GIP"

	// AccountCurrencyGMD captures enum value "GMD"
	AccountCurrencyGMD AccountCurrencyEnum = "GMD"

	// AccountCurrencyGNF captures enum value "GNF"
	AccountCurrencyGNF AccountCurrencyEnum = "GNF"

	// AccountCurrencyGTQ captures enum value "GTQ"
	AccountCurrencyGTQ AccountCurrencyEnum = "GTQ"

	// AccountCurrencyGYD captures enum value "GYD"
	AccountCurrencyGYD AccountCurrencyEnum = "GYD"

	// AccountCurrencyHKD captures enum value "HKD"
	AccountCurrencyHKD AccountCurrencyEnum = "HKD"

	// AccountCurrencyHNL captures enum value "HNL"
	AccountCurrencyHNL AccountCurrencyEnum = "HNL"

	// AccountCurrencyHRK captures enum value "HRK"
	AccountCurrencyHRK AccountCurrencyEnum = "HRK"

	// AccountCurrencyHTG captures enum value "HTG"
	AccountCurrencyHTG AccountCurrencyEnum = "HTG"

	// AccountCurrencyHUF captures enum value "HUF"
	AccountCurrencyHUF AccountCurrencyEnum = "HUF"

	// AccountCurrencyIDR captures enum value "IDR"
	AccountCurrencyIDR AccountCurrencyEnum = "IDR"

	// AccountCurrencyILS captures enum value "ILS"
	AccountCurrencyILS AccountCurrencyEnum = "ILS"

	// AccountCurrencyIMP captures enum value "IMP"
	AccountCurrencyIMP AccountCurrencyEnum = "IMP"

	// AccountCurrencyINR captures enum value "INR"
	AccountCurrencyINR AccountCurrencyEnum = "INR"

	// AccountCurrencyIQD captures enum value "IQD"
	AccountCurrencyIQD AccountCurrencyEnum = "IQD"

	// AccountCurrencyIRR captures enum value "IRR"
	AccountCurrencyIRR AccountCurrencyEnum = "IRR"

	// AccountCurrencyISK captures enum value "ISK"
	AccountCurrencyISK AccountCurrencyEnum = "ISK"

	// AccountCurrencyJEP captures enum value "JEP"
	AccountCurrencyJEP AccountCurrencyEnum = "JEP"

	// AccountCurrencyJMD captures enum value "JMD"
	AccountCurrencyJMD AccountCurrencyEnum = "JMD"

	// AccountCurrencyJOD captures enum value "JOD"
	AccountCurrencyJOD AccountCurrencyEnum = "JOD"

	// AccountCurrencyJPY captures enum value "JPY"
	AccountCurrencyJPY AccountCurrencyEnum = "JPY"

	// AccountCurrencyKES captures enum value "KES"
	AccountCurrencyKES AccountCurrencyEnum = "KES"

	// AccountCurrencyKGS captures enum value "KGS"
	AccountCurrencyKGS AccountCurrencyEnum = "KGS"

	// AccountCurrencyKHR captures enum value "KHR"
	AccountCurrencyKHR AccountCurrencyEnum = "KHR"

	// AccountCurrencyKMF captures enum value "KMF"
	AccountCurrencyKMF AccountCurrencyEnum = "KMF"

	// AccountCurrencyKPW captures enum value "KPW"
	AccountCurrencyKPW AccountCurrencyEnum = "KPW"

	// AccountCurrencyKRW captures enum value "KRW"
	AccountCurrencyKRW AccountCurrencyEnum = "KRW"

	// AccountCurrencyKWD captures enum value "KWD"
	AccountCurrencyKWD AccountCurrencyEnum = "KWD"

	// AccountCurrencyKYD captures enum value "KYD"
	AccountCurrencyKYD AccountCurrencyEnum = "KYD"

	// AccountCurrencyKZT captures enum value "KZT"
	AccountCurrencyKZT AccountCurrencyEnum = "KZT"

	// AccountCurrencyLAK captures enum value "LAK"
	AccountCurrencyLAK AccountCurrencyEnum = "LAK"

	// AccountCurrencyLBP captures enum value "LBP"
	AccountCurrencyLBP AccountCurrencyEnum = "LBP"

	// AccountCurrencyLKR captures enum value "LKR"
	AccountCurrencyLKR AccountCurrencyEnum = "LKR"

	// AccountCurrencyLRD captures enum value "LRD"
	AccountCurrencyLRD AccountCurrencyEnum = "LRD"

	// AccountCurrencyLSL captures enum value "LSL"
	AccountCurrencyLSL AccountCurrencyEnum = "LSL"

	// AccountCurrencyLTL captures enum value "LTL"
	AccountCurrencyLTL AccountCurrencyEnum = "LTL"

	// AccountCurrencyLVL captures enum value "LVL"
	AccountCurrencyLVL AccountCurrencyEnum = "LVL"

	// AccountCurrencyLYD captures enum value "LYD"
	AccountCurrencyLYD AccountCurrencyEnum = "LYD"

	// AccountCurrencyMAD captures enum value "MAD"
	AccountCurrencyMAD AccountCurrencyEnum = "MAD"

	// AccountCurrencyMDL captures enum value "MDL"
	AccountCurrencyMDL AccountCurrencyEnum = "MDL"

	// AccountCurrencyMGA captures enum value "MGA"
	AccountCurrencyMGA AccountCurrencyEnum = "MGA"

	// AccountCurrencyMKD captures enum value "MKD"
	AccountCurrencyMKD AccountCurrencyEnum = "MKD"

	// AccountCurrencyMMK captures enum value "MMK"
	AccountCurrencyMMK AccountCurrencyEnum = "MMK"

	// AccountCurrencyMNT captures enum value "MNT"
	AccountCurrencyMNT AccountCurrencyEnum = "MNT"

	// AccountCurrencyMOP captures enum value "MOP"
	AccountCurrencyMOP AccountCurrencyEnum = "MOP"

	// AccountCurrencyMRO captures enum value "MRO"
	AccountCurrencyMRO AccountCurrencyEnum = "MRO"

	// AccountCurrencyMUR captures enum value "MUR"
	AccountCurrencyMUR AccountCurrencyEnum = "MUR"

	// AccountCurrencyMVR captures enum value "MVR"
	AccountCurrencyMVR AccountCurrencyEnum = "MVR"

	// AccountCurrencyMWK captures enum value "MWK"
	AccountCurrencyMWK AccountCurrencyEnum = "MWK"

	// AccountCurrencyMXN captures enum value "MXN"
	AccountCurrencyMXN AccountCurrencyEnum = "MXN"

	// AccountCurrencyMYR captures enum value "MYR"
	AccountCurrencyMYR AccountCurrencyEnum = "MYR"

	// AccountCurrencyMZN captures enum value "MZN"
	AccountCurrencyMZN AccountCurrencyEnum = "MZN"

	// AccountCurrencyNAD captures enum value "NAD"
	AccountCurrencyNAD AccountCurrencyEnum = "NAD"

	// AccountCurrencyNGN captures enum value "NGN"
	AccountCurrencyNGN AccountCurrencyEnum = "NGN"

	// AccountCurrencyNIO captures enum value "NIO"
	AccountCurrencyNIO AccountCurrencyEnum = "NIO"

	// AccountCurrencyNOK captures enum value "NOK"
	AccountCurrencyNOK AccountCurrencyEnum = "NOK"

	// AccountCurrencyNPR captures enum value "NPR"
	AccountCurrencyNPR AccountCurrencyEnum = "NPR"

	// AccountCurrencyNZD captures enum value "NZD"
	AccountCurrencyNZD AccountCurrencyEnum = "NZD"

	// AccountCurrencyOMR captures enum value "OMR"
	AccountCurrencyOMR AccountCurrencyEnum = "OMR"

	// AccountCurrencyPAB captures enum value "PAB"
	AccountCurrencyPAB AccountCurrencyEnum = "PAB"

	// AccountCurrencyPEN captures enum value "PEN"
	AccountCurrencyPEN AccountCurrencyEnum = "PEN"

	// AccountCurrencyPGK captures enum value "PGK"
	AccountCurrencyPGK AccountCurrencyEnum = "PGK"

	// AccountCurrencyPHP captures enum value "PHP"
	AccountCurrencyPHP AccountCurrencyEnum = "PHP"

	// AccountCurrencyPKR captures enum value "PKR"
	AccountCurrencyPKR AccountCurrencyEnum = "PKR"

	// AccountCurrencyPLN captures enum value "PLN"
	AccountCurrencyPLN AccountCurrencyEnum = "PLN"

	// AccountCurrencyPYG captures enum value "PYG"
	AccountCurrencyPYG AccountCurrencyEnum = "PYG"

	// AccountCurrencyQAR captures enum value "QAR"
	AccountCurrencyQAR AccountCurrencyEnum = "QAR"

	// AccountCurrencyRON captures enum value "RON"
	AccountCurrencyRON AccountCurrencyEnum = "RON"

	// AccountCurrencyRSD captures enum value "RSD"
	AccountCurrencyRSD AccountCurrencyEnum = "RSD"

	// AccountCurrencyRUB captures enum value "RUB"
	AccountCurrencyRUB AccountCurrencyEnum = "RUB"

	// AccountCurrencyRWF captures enum value "RWF"
	AccountCurrencyRWF AccountCurrencyEnum = "RWF"

	// AccountCurrencySAR captures enum value "SAR"
	AccountCurrencySAR AccountCurrencyEnum = "SAR"

	// AccountCurrencySBD captures enum value "SBD"
	AccountCurrencySBD AccountCurrencyEnum = "SBD"

	// AccountCurrencySCR captures enum value "SCR"
	AccountCurrencySCR AccountCurrencyEnum = "SCR"

	// AccountCurrencySDG captures enum value "SDG"
	AccountCurrencySDG AccountCurrencyEnum = "SDG"

	// AccountCurrencySEK captures enum value "SEK"
	AccountCurrencySEK AccountCurrencyEnum = "SEK"

	// AccountCurrencySGD captures enum value "SGD"
	AccountCurrencySGD AccountCurrencyEnum = "SGD"

	// AccountCurrencySHP captures enum value "SHP"
	AccountCurrencySHP AccountCurrencyEnum = "SHP"

	// AccountCurrencySLL captures enum value "SLL"
	AccountCurrencySLL AccountCurrencyEnum = "SLL"

	// AccountCurrencySOS captures enum value "SOS"
	AccountCurrencySOS AccountCurrencyEnum = "SOS"

	// AccountCurrencySPL captures enum value "SPL"
	AccountCurrencySPL AccountCurrencyEnum = "SPL"

	// AccountCurrencySRD captures enum value "SRD"
	AccountCurrencySRD AccountCurrencyEnum = "SRD"

	// AccountCurrencySTD captures enum value "STD"
	AccountCurrencySTD AccountCurrencyEnum = "STD"

	// AccountCurrencySVC captures enum value "SVC"
	AccountCurrencySVC AccountCurrencyEnum = "SVC"

	// AccountCurrencySYP captures enum value "SYP"
	AccountCurrencySYP AccountCurrencyEnum = "SYP"

	// AccountCurrencySZL captures enum value "SZL"
	AccountCurrencySZL AccountCurrencyEnum = "SZL"

	// AccountCurrencyTHB captures enum value "THB"
	AccountCurrencyTHB AccountCurrencyEnum = "THB"

	// AccountCurrencyTJS captures enum value "TJS"
	AccountCurrencyTJS AccountCurrencyEnum = "TJS"

	// AccountCurrencyTMT captures enum value "TMT"
	AccountCurrencyTMT AccountCurrencyEnum = "TMT"

	// AccountCurrencyTND captures enum value "TND"
	AccountCurrencyTND AccountCurrencyEnum = "TND"

	// AccountCurrencyTOP captures enum value "TOP"
	AccountCurrencyTOP AccountCurrencyEnum = "TOP"

	// AccountCurrencyTRY captures enum value "TRY"
	AccountCurrencyTRY AccountCurrencyEnum = "TRY"

	// AccountCurrencyTTD captures enum value "TTD"
	AccountCurrencyTTD AccountCurrencyEnum = "TTD"

	// AccountCurrencyTVD captures enum value "TVD"
	AccountCurrencyTVD AccountCurrencyEnum = "TVD"

	// AccountCurrencyTWD captures enum value "TWD"
	AccountCurrencyTWD AccountCurrencyEnum = "TWD"

	// AccountCurrencyTZS captures enum value "TZS"
	AccountCurrencyTZS AccountCurrencyEnum = "TZS"

	// AccountCurrencyUAH captures enum value "UAH"
	AccountCurrencyUAH AccountCurrencyEnum = "UAH"

	// AccountCurrencyUGX captures enum value "UGX"
	AccountCurrencyUGX AccountCurrencyEnum = "UGX"

	// AccountCurrencyUSD captures enum value "USD"
	AccountCurrencyUSD AccountCurrencyEnum = "USD"

	// AccountCurrencyUYU captures enum value "UYU"
	AccountCurrencyUYU AccountCurrencyEnum = "UYU"

	// AccountCurrencyUZS captures enum value "UZS"
	AccountCurrencyUZS AccountCurrencyEnum = "UZS"

	// AccountCurrencyVEF captures enum value "VEF"
	AccountCurrencyVEF AccountCurrencyEnum = "VEF"

	// AccountCurrencyVND captures enum value "VND"
	AccountCurrencyVND AccountCurrencyEnum = "VND"

	// AccountCurrencyVUV captures enum value "VUV"
	AccountCurrencyVUV AccountCurrencyEnum = "VUV"

	// AccountCurrencyWST captures enum value "WST"
	AccountCurrencyWST AccountCurrencyEnum = "WST"

	// AccountCurrencyXAF captures enum value "XAF"
	AccountCurrencyXAF AccountCurrencyEnum = "XAF"

	// AccountCurrencyXCD captures enum value "XCD"
	AccountCurrencyXCD AccountCurrencyEnum = "XCD"

	// AccountCurrencyXDR captures enum value "XDR"
	AccountCurrencyXDR AccountCurrencyEnum = "XDR"

	// AccountCurrencyXOF captures enum value "XOF"
	AccountCurrencyXOF AccountCurrencyEnum = "XOF"

	// AccountCurrencyXPF captures enum value "XPF"
	AccountCurrencyXPF AccountCurrencyEnum = "XPF"

	// AccountCurrencyYER captures enum value "YER"
	AccountCurrencyYER AccountCurrencyEnum = "YER"

	// AccountCurrencyZAR captures enum value "ZAR"
	AccountCurrencyZAR AccountCurrencyEnum = "ZAR"

	// AccountCurrencyZMW captures enum value "ZMW"
	AccountCurrencyZMW AccountCurrencyEnum = "ZMW"

	// AccountCurrencyZWD captures enum value "ZWD"
	AccountCurrencyZWD AccountCurrencyEnum = "ZWD"

	// AccountCurrencyBTC captures enum value "BTC"
	AccountCurrencyBTC AccountCurrencyEnum = "BTC"
)

var AccountCurrencyEnumValues = []string{
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

func (e AccountCurrencyEnum) IsValid() bool {
	for _, v := range AccountCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Account) validateCurrencyEnum(path, location string, value AccountCurrencyEnum) error {
	if err := validate.Enum(path, location, value, accountTypeCurrencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateParentAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentAccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentAccountId", "body", "uuid", m.ParentAccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validatePaymentMethodID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentMethodId", "body", "uuid", m.PaymentMethodID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateReferenceTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceTime) { // not required
		return nil
	}

	if err := validate.FormatOf("referenceTime", "body", "date-time", m.ReferenceTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
