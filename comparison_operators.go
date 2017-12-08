package nqb

import (
	"bytes"
	"fmt"
)

func buildComparison(buf *bytes.Buffer, operator string, column string, value *string) error {
	buf.WriteString(escapeIdentifiers(column))
	buf.WriteString(" ")
	buf.WriteString(operator)

	if value != nil {
		buf.WriteString(" ")
		buf.WriteString(*value)
	}

	return nil
}

func queryPlaceholder(placeholder string) *string {
	p := fmt.Sprintf("$%s", placeholder)
	return &p
}

// Eq builds an equal to (`=`) comparison.
func Eq(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "=", column, queryPlaceholder(placeholder))
	})
}

// Neq builds a not equal to (`!=`) comparison.
func Neq(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "!=", column, queryPlaceholder(placeholder))
	})
}

// Gt builds a greater than (`>`) comparison.
func Gt(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, ">", column, queryPlaceholder(placeholder))
	})
}

// Gte builds a greater than or equal to (`>=`) comparison.
func Gte(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, ">=", column, queryPlaceholder(placeholder))
	})
}

// Lt builds a less than (`<`) comparison.
func Lt(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "<", column, queryPlaceholder(placeholder))
	})
}

// Lte builds a less than or equal to (`<=`) comparison.
func Lte(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "<=", column, queryPlaceholder(placeholder))
	})
}

func betweenPlaceholders(placeholder1, placeholder2 string) *string {
	p := fmt.Sprintf("$%s AND $%s", placeholder1, placeholder2)
	return &p
}

// Between builds a `BETWEEN` comparison.
//
// Search criteria for a query where the value is between two values,
// including the end values specified in the range.
// Values can be numbers, text, or dates.
func Between(column, placeholder1, placeholder2 string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "BETWEEN", column, betweenPlaceholders(placeholder1, placeholder2))
	})
}

// NotBetween builds a `NOT BETWEEN` comparison.
//
// Search criteria for a query where the value is outside the range of two values,
// including the end values specified in the range.
// Values can be numbers, text, or dates.
func NotBetween(column, placeholder1, placeholder2 string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "NOT BETWEEN", column, betweenPlaceholders(placeholder1, placeholder2))
	})
}

// Like builds a `LIKE` comparison.
//
// Match string with a wildcard expression.
// Use % for zero or more wildcards and _ to match any character at this place in a string.
//
// The wildcard characters can be escaped by preceding them with a backslash (\).
// Backslash itself can also be escaped by preceding it with another backslash.
func Like(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "LIKE", column, queryPlaceholder(placeholder))
	})
}

// NotLike builds a `NOT LIKE` comparison.
// Inverse of LIKE. Return TRUE if string is not similar to given string.
func NotLike(column, placeholder string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "NOT LIKE", column, queryPlaceholder(placeholder))
	})
}

// IsNull builds an `IS NULL` comparison.
// Field has value of NULL.
func IsNull(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS NULL", column, nil)
	})
}

// IsNotNull builds an `IS NOT NULL` comparison.
// Field has value or is missing.
func IsNotNull(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS NOT NULL", column, nil)
	})
}

// IsMissing builds an `IS MISSING` comparison.
// No value for field found.
func IsMissing(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS MISSING", column, nil)
	})
}

// IsNotMissing builds an `IS NOT MISSING` comparison.
// Value for field found or value is NULL.
func IsNotMissing(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS NOT MISSING", column, nil)
	})
}

// IsValued builds an `IS VALUED` comparison.
// Value for field found. Value is neither missing nor NULL.
func IsValued(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS VALUED", column, nil)
	})
}

// IsNotValued builds an `IS NOT VALUED` comparison.
// Value for field not found. Value is NULL.
func IsNotValued(column string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		return buildComparison(buf, "IS NOT VALUED", column, nil)
	})
}
