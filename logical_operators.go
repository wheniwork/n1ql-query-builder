package nqb

import "bytes"

// `AND` evaluates to TRUE only if all conditions are TRUE.
func And(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "AND", cond...)
	})
}

// `OR` evaluates to TRUE if one of the conditions is TRUE.
func Or(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "OR", cond...)
	})
}

// `NOT` evaluates to TRUE if the expression does not match the condition.
func Not(cond BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "NOT", cond)
	})
}
