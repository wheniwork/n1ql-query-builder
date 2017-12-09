package nqb

import (
	"bytes"
	"errors"
)

const (
	and = "AND"
	or  = "OR"
	not = "NOT"
)

// And builds an AND condition
//  AND evaluates to TRUE only if all conditions are TRUE.
func And(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, and, cond...)
	})
}

// Or builds an OR condition
//  OR evaluates to TRUE if one of the conditions is TRUE.
func Or(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, or, cond...)
	})
}

// Not builds a NOT condition
//  NOT evaluates to TRUE if the expression does not match the condition.
func Not(cond BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, not, cond)
	})
}

func buildCondition(buf *bytes.Buffer, predicate string, cond ...BuildFunc) error {
	if len(predicate) == 0 {
		return errors.New("nqb: predicate cannot be empty")
	}

	for i, c := range cond {
		if i > 0 || predicate == not {
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
