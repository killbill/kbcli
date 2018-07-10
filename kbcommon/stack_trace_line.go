package kbcommon

// StackTraceLine represents a line of stack trace.
type StackTraceLine struct {
	// class name
	ClassName string `json:"className,omitempty"`

	// file name
	FileName string `json:"fileName,omitempty"`

	// is native method
	IsNativeMethod *bool `json:"isNativeMethod,omitempty"`

	// line number
	LineNumber int32 `json:"lineNumber,omitempty"`

	// method name
	MethodName string `json:"methodName,omitempty"`
}
