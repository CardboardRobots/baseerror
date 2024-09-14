package baseerror

import (
	"errors"
	"fmt"
)

type BaseError struct {
	error
}

var (
	ErrAlreadyExists      = NewError("already exists")
	ErrFailedPrecondition = NewError("failed precondition")
	ErrInvalidArgument    = NewError("invalid argument")
	ErrNotFound           = NewError("not found")
	ErrOutOfRange         = NewError("out of range")
	ErrPermissionDenied   = NewError("permission denied")
	ErrUnauthenticated    = NewError("unauthenticated")
)

func NewError(text string) BaseError {
	return BaseError{
		error: errors.New(text),
	}
}

func WrapError(err error) BaseError {
	return BaseError{
		error: err,
	}
}

func (be BaseError) Unwrap() error {
	return be.error
}

func (be BaseError) Wrap(text string) BaseError {
	return BaseError{
		error: fmt.Errorf("%w: %v", be, text),
	}
}

func (be BaseError) WrapError(err error) BaseError {
	return BaseError{
		error: errors.Join(be, err),
	}
}
