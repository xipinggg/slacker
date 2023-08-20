package errutil

import (
	"fmt"
)

// WithStack with the stack.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	stackErr := NewStackError(WithSkipOption(2))
	return fmt.Errorf("%w\n%w", err, stackErr)
}

// WithMessage with the message.
func WithMessage(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}

	msg = fmt.Sprintf(msg, args...)
	return fmt.Errorf("%s\n%w", msg, err)
}

// Wrap with the message and the stack.
func Wrap(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}

	msg = fmt.Sprintf(msg, args...)
	stackErr := NewStackError(WithSkipOption(2))
	return fmt.Errorf("%s\n%w\n%w", msg, err, stackErr)
}
