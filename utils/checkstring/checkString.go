package utils

import (
	"strings"
)

// CheckString function checks the validity of a string
func CheckString(toBeChecked string) (okay bool) {
	if strings.TrimSpace(toBeChecked) != "" {
		okay = true
	}
	return
}
