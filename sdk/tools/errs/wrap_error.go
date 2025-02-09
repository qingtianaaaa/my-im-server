package errs

import "errors"

type ErrorWrapper interface {
	Is(err error) bool
	Wrap() error
	UnWrap() error
	WrapMsg(msg string, kv ...any) error
	error
}

type errorWrapper struct {
	error
	msg string
}

func (e *errorWrapper) Is(err error) bool {
	if err == nil {
		return false
	}
	var t *errorWrapper
	ok := errors.As(err, &t)
	return ok && t.msg == e.msg
}

func (e *errorWrapper) Wrap() error {
	return Wrap(e)
}

func (e *errorWrapper) Error() string {
	return e.msg
}

func (e *errorWrapper) UnWrap() error {
	return e.error
}

func (e *errorWrapper) WrapMsg(msg string, kv ...any) error {
	return WrapMsg(e, msg, kv...)
}

func NewErrorWrapper(err error, msg string) ErrorWrapper {
	return &errorWrapper{
		error: err,
		msg:   msg,
	}
}
