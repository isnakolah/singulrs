package utils

import (
	"fmt"
	"testing"
)

type testpair struct {
	funcErrMessage, actualErrMessage string
}

var tests = []testpair{
	{
		ErrorMessage("Test Base Message", "Test Label", float64(100), float64(10)),
		fmt.Sprintf(
			"\n\n-> %s. \n-> For %q expected %g got %g\n\n", "Test Base Message", "Test Label", float64(100), float64(10),
		),
	},
	{
		ErrorMessage("This does work", "Just another test", float64(0), float64(-10)),
		fmt.Sprintf(
			"\n\n-> %s. \n-> For %q expected %g got %g\n\n", "This does work", "Just another test", float64(0), float64(-10),
		),
	},
	{
		ErrorMessage("Another one", "We are almost done", float64(10000008.2394710), float64(-10)),
		fmt.Sprintf(
			"\n\n-> %s. \n-> For %q expected %g got %g\n\n", "Another one", "We are almost done", float64(10000008.2394710), float64(-10),
		),
	},
}

func TestErrorMessages(t *testing.T) {
	for _, pair := range tests {
		if pair.funcErrMessage != pair.actualErrMessage {
			t.Error("Error message function not working")
		}
	}
}
