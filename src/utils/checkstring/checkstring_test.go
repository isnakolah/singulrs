package checkstring

import (
	"singulr/src/utils/errmessages"
	"testing"
)

var tests = map[string]bool{
	"     ":          false,
	"Good":           true,
	"12354262      ": true,
	"   we 13446   ": true,
}

func TestValidString(t *testing.T) {
	for inputedString, expected := range tests {
		if CheckString(inputedString) != expected {
			t.Errorf(errmessages.SimpleErrorMessage(
				"CheckString function not working as expected.",
			))
		}
	}
}
