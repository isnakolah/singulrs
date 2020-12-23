package utils

import "fmt"

// Err struct is blueprint for the error messages
type Err struct {
	baseMessage, label string
	expected, got      float64
}

// ErrorMessage function creates a new error message
func ErrorMessage(baseMessage, label string, expected, got float64) string {
	message := Err{baseMessage, label, expected, got}
	return fmt.Sprintf(
		"\n\n-> %s. \n-> For %q expected %g got %g\n\n", message.baseMessage, message.label, message.expected, message.got,
	)
}

// SimpleErrorMessage function creates a simple error message
func SimpleErrorMessage(message string) string {
	return fmt.Sprintf("\n\n-> %s.\n\n", message)
}
