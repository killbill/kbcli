// Package kbcli is just a wrapper to pull in kbclient and kbmodel packages.
// Don't use this package directly. Use the kbmodel and kbclient/... packages.
package kbcli

import (
	_ "github.com/go-openapi/runtime"
	_ "github.com/killbill/kbcli/kbclient"
)
