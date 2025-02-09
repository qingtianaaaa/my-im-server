package errs

import "my-im-server/sdk/tools/errs/stack"

const stackSkip = 4

func WrapErr(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	return nil
}

func WrapMsg(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	err = NewErrorWrapper(err, toString(msg, kv...))
	return stack.New(err, stackSkip)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return stack.New(err, stackSkip)
}
