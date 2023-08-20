package errutil

import (
	"runtime"
	"strconv"
	"strings"
)

type NewStackErrorOption struct {
	// frame = frame[skip:endSkip]
	skip    uint
	endSkip int
	depth   uint
	prefix  string
}

func newDefaultOption() *NewStackErrorOption {
	return &NewStackErrorOption{
		skip:    1,
		endSkip: 4,
		depth:   32,
		prefix:  "stack:\n",
	}
}

type WithOptionFunc func(opt *NewStackErrorOption)

// WithSkipOption skips the caller stack frame.
func WithSkipOption(skip uint) WithOptionFunc {
	return func(opt *NewStackErrorOption) {
		opt.skip = skip
	}
}

// WithDepthOption defines the max depth of stack frame.
func WithDepthOption(depth uint) WithOptionFunc {
	return func(opt *NewStackErrorOption) {
		opt.depth = depth
	}
}

// StackError is an error with caller stack.
type StackError struct {
	stack string
}

var _ error = (*StackError)(nil)

// NewStackError returns an error with caller stack.
func NewStackError(opts ...WithOptionFunc) *StackError {
	opt := newDefaultOption()
	for _, with := range opts {
		with(opt)
	}

	return &StackError{
		stack: stack(opt),
	}
}

func (e *StackError) Error() string {
	return e.stack
}

func stack(opt *NewStackErrorOption) string {
	pc := make([]uintptr, opt.depth)
	n := runtime.Callers(int(opt.skip)+2, pc)
	if n > opt.endSkip {
		n -= opt.endSkip
	}
	frames := runtime.CallersFrames(pc[:n])

	var builder strings.Builder
	builder.WriteString(opt.prefix)
	for {
		frame, more := frames.Next()

		builder.WriteString(frame.Function)
		builder.WriteByte('\n')
		builder.WriteByte('\t')
		builder.WriteString(frame.File)
		builder.WriteByte(':')
		builder.WriteString(strconv.Itoa(frame.Line))
		builder.WriteByte('\n')

		if !more {
			break
		}
	}

	return builder.String()
}
