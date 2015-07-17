package kbcli

import "net/http"

type QueryParams map[string]string

type Request struct {
	Url         string
	// GET, PUT, POST,..
	Method      string
	// Headers
	Header      *http.Header
	// Query QueryParams
	QueryParams *QueryParams
	// Body
	Body        interface{}
}

