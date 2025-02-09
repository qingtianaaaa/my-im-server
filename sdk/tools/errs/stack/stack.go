package stack

import (
	"path"
	"runtime"
	"strconv"
	"strings"
)

func New(err error, skip int) error {
	return &stackError{
		err:   err,
		stack: Callers(skip),
	}
}

type stackError struct {
	err   error
	stack []uintptr
}

func Callers(skip int) []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}

func (e *stackError) Error() string {
	if len(e.stack) == 0 {
		return e.err.Error()
	}
	var sb strings.Builder
	sb.WriteString("Error: ")
	sb.WriteString(e.err.Error())
	sb.WriteString(" |")
	for _, pc := range e.stack {
		fn := runtime.FuncForPC(pc - 1)
		if fn == nil {
			continue
		}
		name := path.Base(fn.Name())
		if strings.HasPrefix(name, "runtime.") {
			break
		}
		file, line := fn.FileLine(pc)
		sb.WriteString(" -> ")
		sb.WriteString(name)
		sb.WriteString("() ")
		sb.WriteString(file)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(line))
	}
	return sb.String()
}
