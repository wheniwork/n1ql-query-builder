package nqb

import (
	"bytes"
	"fmt"
)

var (
	nullExpr    = &Expression{"NULL"}
	trueExpr    = &Expression{"TRUE"}
	falseExpr   = &Expression{"FALSE"}
	missingExpr = &Expression{"MISSING"}
	emptyExpr   = &Expression{""}
)

// Expression represents a N1QL Expression.
type Expression struct {
	value string
}

// Creates an arbitrary expression from the given value.
//
// No quoting or escaping will be done on the input.
// In addition, it is not checked if the given value is an actual valid (N1QL syntax wise) expression.
func X(value interface{}) *Expression {
	switch value.(type) {
	case bool:
		if value == true {
			return trueExpr
		}
		return falseExpr
	}
	return &Expression{fmt.Sprint(value)}
}

// Sub creates an expression from a given sub-Statement, wrapping it in parenthesis.
func Sub(statement Statement) *Expression {
	return &Expression{"(" + statement.(fmt.Stringer).String() + ")"}
}

// Par wraps an Expression in parenthesis.
func Par(expression *Expression) *Expression {
	return infix(expression.String(), "(", ")")
}

// Construct a path ("a.b.c") from Expressions or values. Strings are considered identifiers (so they won't be quoted).
func P(pathComponents ...interface{}) *Expression {
	if pathComponents == nil || len(pathComponents) == 0 {
		return emptyExpr
	}

	path := bytes.Buffer{}

	for i, p := range pathComponents {
		if i > 0 {
			path.WriteString(".")
		}

		switch p.(type) {
		case *Expression:
			path.WriteString(p.(*Expression).String())
		default:
			path.WriteString(fmt.Sprint(p))
		}
	}

	return &Expression{path.String()}
}

// An identifier or list of identifiers escaped using backquotes `.
//
// Useful for example for identifiers that contains a dash like "beer-sample".
// Multiple identifiers are returned as a list of escaped identifiers separated by ", ".
func I(identifiers ...string) *Expression {
	return &Expression{wrapWith('`', identifiers...)}
}

// An identifier or list of identifiers which will be quoted as strings (with "").
func S(strings ...string) *Expression {
	return &Expression{wrapWith('"', strings...)}
}

// Helper method to prefix a string.
func prefix(prefix, right string) *Expression {
	return &Expression{prefix + " " + right}
}

// Helper method to infix a string.
func infix(infix, left, right string) *Expression {
	return &Expression{left + " " + infix + " " + right}
}

// Helper method to postfix a string.
func postfix(postfix, left string) *Expression {
	return &Expression{left + " " + postfix}
}

// Helper method to wrap variadic arguments with the given character.
func wrapWith(wrapper byte, input ...string) string {
	escaped := bytes.Buffer{}

	for n, i := range input {
		if n > 0 {
			escaped.WriteString(", ")
		}

		escaped.WriteByte(wrapper)
		escaped.WriteString(i)
		escaped.WriteByte(wrapper)
	}

	return escaped.String()
}

func (e *Expression) String() string {
	return e.value
}
