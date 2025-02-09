package errs

import (
	"bytes"
	"fmt"
)

func toString(s string, kv ...any) string {
	if len(kv) == 0 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteString(s)
	for i := 0; i < len(kv); i += 2 {
		if buf.Len() > 0 {
			buf.WriteString(", ")
		}
		key := fmt.Sprintf("%v", kv[i])
		buf.WriteString(key)
		buf.WriteString("=")
		if i+1 < len(kv) {
			buf.WriteString(fmt.Sprintf("%v", kv[i+1]))
		} else {
			buf.WriteString("MISSING VALUE")
		}
	}
	return buf.String()
}
