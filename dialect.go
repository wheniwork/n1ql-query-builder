package nqb

import (
	"strings"
)

func EscapeIdentifier(s string) string {
	return escapeIdentifier(s, "`")
}

func escapeIdentifier(s, quote string) string {
	part := strings.SplitN(s, ".", 2)

	if len(part) == 2 {
		return escapeIdentifier(part[0], quote) + "." + escapeIdentifier(part[1], quote)
	}

	return quote + s + quote
}
