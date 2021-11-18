package utils

import (
	"bytes"
)

// JoinStrWithSep joins multi string with the separator
// @sep: such as: "-", "&", "_", and so on
// @originStr: the string to be joined
// return: (string, error), for example: ("a=1&b=2", nil)
func JoinStrWithSep(sep string, originStr ...string) string {
	size := len(originStr)
	if size == 0 {
		return ""
	}
	var buf bytes.Buffer
	var i int
	for i = 0; i < size-1; i++ {
		buf.WriteString(originStr[i])
		buf.WriteString(sep)
	}
	buf.WriteString(originStr[i])
	return buf.String()
}
