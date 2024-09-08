package baseerror

import (
	"errors"
	"testing"
)

func TestBaseError_NewError(t *testing.T) {
	want := "test error"
	err := NewError(want)
	if value := err.Error(); value != want {
		t.Errorf("value mismatch. Recieved: %v, Expected: %v", value, want)
	}
}

func TestBaseError_WrapError(t *testing.T) {
	want := "test error"

	err := errors.New(want)
	wrapped := WrapError(err)

	if value := wrapped.Error(); value != want {
		t.Errorf("value mismatch. Recieved: %v, Expected: %v", value, want)
	}
}

func TestBaseError_Wrap(t *testing.T) {
	err := NewError("test error")
	wrapped := err.Wrap("wrapped error")

	want := "test error: wrapped error"
	if value := wrapped.Error(); value != want {
		t.Errorf("value mismatch. Recieved: %v, Expected: %v", value, want)
	}
}
