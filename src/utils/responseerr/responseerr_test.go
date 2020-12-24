package responseerr

import (
	"errors"
	"singulr/src/utils/errmessages"
	"testing"
)

func TestGetStrErr(t *testing.T) {

	var err = errors.New("A good error")

	if GetStrErr(err) != "A good error" {
		t.Errorf(errmessages.SimpleErrorMessage(
			"GetStrErr not working as expected",
		))
	}
	if GetStrErr(nil) != "" {
		t.Errorf(errmessages.SimpleErrorMessage(
			"Nil error not returned as an empty string",
		))
	}
}
