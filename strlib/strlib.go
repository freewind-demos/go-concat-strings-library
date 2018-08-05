package strlib

import (
	"bytes"
)

func ConcatStrs(strs ...string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(strs); i++ {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}

