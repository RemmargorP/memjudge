package utils

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	matched, err := ValidateEmail("abs@asd.asd")
	if !matched {
		t.Error("Not matched: abs@asd.asd")
	}
	if err != nil {
		t.Error(err)
	}
	matched, err = ValidateEmail("abs@asdasd")
	if matched {
		t.Error("Matched: abs@asdasd")
	}
	if err != nil {
		t.Error(err)
	}
	matched, err = ValidateEmail("absaasd.asd")
	if matched {
		t.Error("Matched: absaasd.asd")
	}
	if err != nil {
		t.Error(err)
	}
}
