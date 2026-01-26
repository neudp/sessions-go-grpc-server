package app

import "errors"

var ErrInvalidArgument = errors.New("invalid argument")

type AppError struct {
	msg  string
	errs []error
}

func NewAppError(msg string, errs ...error) *AppError {
	return &AppError{msg, errs}
}

func (err *AppError) Message() string {
	return err.msg
}

func (err *AppError) Error() string {
	return err.msg + "\n" + errors.Join(err.errs...).Error()
}

func (err *AppError) Unwrap() []error {
	return err.errs
}
