package kblib

// supportedTagTargets - All supported object types
var supportedTagTargets = map[string]bool{
	"BUNDLE":             true,
	"INVOICE_ITEM":       true,
	"PAYMENT":            true,
	"TENANT":             true,
	"BLOCKING_STATES":    true,
	"INVOICE":            true,
	"TENANT_KVS":         true,
	"PAYMENT_ATTEMPT":    true,
	"SERVICE_BROADCAST":  true,
	"TAG_DEFINITION":     true,
	"PAYMENT_METHOD":     true,
	"SUBSCRIPTION":       true,
	"TRANSACTION":        true,
	"ACCOUNT_EMAIL":      true,
	"CUSTOM_FIELD":       true,
	"INVOICE_PAYMENT":    true,
	"SUBSCRIPTION_EVENT": true,
	"TAG":                true,
	"ACCOUNT":            true,
}

// GetAllSupportedTagTargets returns all supported targets for tags.
func GetAllSupportedTagTargets() []string {
	var result = make([]string, len(supportedTagTargets))
	var i = 0
	for k := range supportedTagTargets {
		result[i] = k
		i++
	}
	return result
}

// IsValidTagTarget returns true if the target is valid tag target
func IsValidTagTarget(tt string) bool {
	return supportedTagTargets[tt]
}
