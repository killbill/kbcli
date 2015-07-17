package kbcli

import (
	"runtime"
	"path/filepath"
	"encoding/json"
	"fmt"
)

func pretty(v interface{}) string {
	// Get source file and line
	_, file, line, _ := runtime.Caller(1)

	// Get relative filename
	filename := filepath.Base(file)

	// Convert to JSON
	b, _ := json.MarshalIndent(v, "", "\t")

	// Make that all pretty together
	return fmt.Sprintf("%s:%d: \n%s\n", filename, line, string(b))
}

