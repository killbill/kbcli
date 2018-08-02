package kblib

// BoolPtr return bool pointer from bool value
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr return string pointer from string value
func StringPtr(s string) *string {
	return &s
}

// Int64Ptr returns int64 pointer from int64 value
func Int64Ptr(i int64) *int64 {
	return &i
}

// IntPtr returns int pointer from int value
func IntPtr(i int) *int {
	return &i
}
