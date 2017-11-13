package nqb

import (
	"strings"
	"fmt"
)

func escapeIdentifiers(s string) string {
	part := strings.SplitN(s, ".", 2)

	if len(part) == 2 {
		return escapeIdentifiers(part[0]) + "." + escapeIdentifiers(part[1])
	}

	return fmt.Sprintf("`%s`", s)
}
