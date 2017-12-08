package nqb

import (
	"bytes"
)

func buildCondition(buf *bytes.Buffer, predicate string, cond ...BuildFunc) error {
	for i, c := range cond {
		if i > 0 {
			buf.WriteString(" ")
			buf.WriteString(predicate)
			buf.WriteString(" ")
		}

		buf.WriteString("(")

		err := c(buf)

		if err != nil {
			return err
		}

		buf.WriteString(")")
	}

	return nil
}
