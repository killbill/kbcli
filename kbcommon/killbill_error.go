// Package kbcommon implements base library for go client.
package kbcommon

import (
	"fmt"
	"strings"
)

// KillbillError - default error message when an operation fails.
type KillbillError struct {
	// HTTPCode - http return code
	HTTPCode int

	// cause class name
	CauseClassName string `json:"causeClassName,omitempty"`

	// cause message
	CauseMessage string `json:"causeMessage,omitempty"`

	// class name
	ClassName string `json:"className,omitempty"`

	// code
	Code int `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// stack trace
	StackTrace []*StackTraceLine `json:"stackTrace"`
}

// NewKillbillError creates new error
func NewKillbillError(httpCode int) *KillbillError {
	return &KillbillError{
		HTTPCode: httpCode,
	}
}

// Error returns error message
func (k *KillbillError) Error() string {
	return fmt.Sprintf("[%d] %s", k.HTTPCode, k.Message)
}

// FormatStackTrace returns formatted stack trace.
func (k *KillbillError) FormatStackTrace() string {
	var result []string
	for _, l := range k.StackTrace {
		line := fmt.Sprintf("%s.%s(%s):%d", l.ClassName, l.MethodName, l.FileName, l.LineNumber)
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

// FormatFull formats all the fields of this error.
func (k *KillbillError) FormatFull() string {
	return fmt.Sprintf("httpcode = %d, class = %s, cause = %s, cause_msg = %s, code = %d, desc = %s", k.HTTPCode, k.ClassName, k.CauseClassName, k.CauseMessage, k.Code, k.Message)
}
