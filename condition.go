package nqb

import (
	"bytes"
	"errors"
)

func buildCondition(buf *bytes.Buffer, predicate string, cond ...BuildFunc) error {
	if len(predicate) == 0 {
		return errors.New("nqb: predicate cannot be empty")
	}

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
