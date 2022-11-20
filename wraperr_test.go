package wraperr_test

import (
	"errors"
	"testing"

	"github.com/Nolions/wraperr"
)

func TestNew(t *testing.T) {
	statusCode := 404
	code := 40400
	message := "Not found"
	previous := errors.New("Previous error")

	err := wraperr.New(statusCode, code, message, previous)

	if statusCode != err.StatusCode {
		t.Error("Status code not equal")
	}
	if code != err.Code {
		t.Error("Code not equal")
	}
	if message != err.Message {
		t.Error("Message not equal")
	}
	if previous != err.Previous {
		t.Error("Previous error not equal")
	}
}
