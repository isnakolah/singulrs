package utils

import "fmt"

// ErrorMessage function creates a new error message
func ErrorMessage(baseMessage, label string, expected, got float64) string {
	return fmt.Sprintf("\n\n-> %s. \n-> For %q expected %g got %g\n\n", baseMessage, label, expected, got)
}
