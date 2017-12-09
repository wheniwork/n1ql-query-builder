package nqb

import "bytes"

// And builds an AND condition
//  AND evaluates to TRUE only if all conditions are TRUE.
func And(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "AND", cond...)
	})
}

// Or builds an OR condition
//  OR evaluates to TRUE if one of the conditions is TRUE.
func Or(cond ...BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "OR", cond...)
	})
}

// Not builds a NOT condition
//  NOT evaluates to TRUE if the expression does not match the condition.
func Not(cond BuildFunc) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildCondition(buf, "NOT", cond)
	})
}
