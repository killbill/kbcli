package kblib

import "github.com/go-openapi/strfmt"

// ParseKeyOrID parses given input string that is either UUID or external key.
// External keys may be prefixed with +.
func ParseKeyOrID(keyOrID string) (idOrKey string, isID bool) {
	if strfmt.IsUUID(keyOrID) {
		return keyOrID, true
	}
	key := keyOrID
	if key[0:1] == "+" {
		return key[1:], false
	}

	return key, false
}
