package kbcli
import (
	"time"
	"net/http"
	"strings"
)


type Response struct {
	Result       interface{}

	// Time when HTTP request was sent
	timestamp    time.Time
	// HTTP status for executed request
	status       int
	// Response object from http package
	httpResponse *http.Response
	// Body of server's response (JSON or otherwise)
	body         []byte
}


// Timestamp returns the time when HTTP request was sent.
func (r *Response) Timestamp() time.Time {
	return r.timestamp
}

// RawText returns the body of the server's response as raw text.
func (r *Response) RawText() string {
	return strings.TrimSpace(string(r.body))
}

// Status returns the HTTP status for the executed request, or 0 if request has
// not yet been sent.
func (r *Response) Status() int {
	return r.status
}

// HttpResponse returns the underlying Response object from http package.
func (r *Response) HttpResponse() *http.Response {
	return r.httpResponse
}


