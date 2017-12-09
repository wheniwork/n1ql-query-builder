package nqb

import (
	"bytes"
	"strings"
)

func escapeIdentifiers(s string) string {
	part := strings.SplitN(s, ".", 2)

	if len(part) == 2 {
		return escapeIdentifiers(part[0]) + "." + escapeIdentifiers(part[1])
	}

	buf := bytes.NewBufferString("`")
	buf.WriteString(s)
	buf.WriteString("`")

	return buf.String()
}
