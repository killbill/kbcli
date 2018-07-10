package kbcommon

import "strings"

// ParseLocationHeader parses the given location header and returns
// the path (removes hostname, port)
func ParseLocationHeader(location string) string {
	if location == "" {
		return ""
	}
	prefixLen := 8 // https:// or http://
	if len(location) < prefixLen {
		return ""
	}
	// Remove hostname
	location = location[prefixLen:]
	idx := strings.Index(location, "/")
	if idx == -1 {
		return ""
	}
	return location[idx:]
}
