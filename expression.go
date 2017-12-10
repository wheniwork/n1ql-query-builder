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

// An identifier or list of identifiers escaped using back-quotes `.
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

// TRUE returns an expression representing boolean TRUE.
func TRUE() *Expression {
	return trueExpr
}

// FALSE returns an expression representing boolean FALSE.
func FALSE() *Expression {
	return falseExpr
}

// NULL returns an expression representing NULL.
func NULL() *Expression {
	return nullExpr
}

// MISSING returns an expression representing MISSING.
func MISSING() *Expression {
	return missingExpr
}

// Not negates the given expression by prefixing a NOT.
func (e *Expression) Not() *Expression {
	return prefix("NOT", e.String())
}

// And AND-combines two expressions.
func (e *Expression) And(right *Expression) *Expression {
	return infix("AND", e.String(), right.String())
}

// Or OR-combines two expressions.
func (e *Expression) Or(right *Expression) *Expression {
	return infix("OR", e.String(), right.String())
}

// Eq combines two expressions with the equals operator ("=").
func (e *Expression) Eq(right *Expression) *Expression {
	return infix("=", e.String(), right.String())
}

// Ne combines two expressions with the not equals operator ("!=").
func (e *Expression) Ne(right *Expression) *Expression {
	return infix("!=", e.String(), right.String())
}

// Gt combines two expressions with the greater than operator (">").
func (e *Expression) Gt(right *Expression) *Expression {
	return infix(">", e.String(), right.String())
}

// Lt combines two expressions with the less than operator ("<").
func (e *Expression) Lt(right *Expression) *Expression {
	return infix("<", e.String(), right.String())
}

// Gte combines two expressions with the greater or equals than operator (">=").
func (e *Expression) Gte(right *Expression) *Expression {
	return infix(">=", e.String(), right.String())
}

// Concat combines two expressions with the concatenation operator ("||").
func (e *Expression) Concat(right *Expression) *Expression {
	return infix("||", e.String(), right.String())
}

// Lte combines two expressions with the less or equals than operator ("<=").
func (e *Expression) Lte(right *Expression) *Expression {
	return infix("<=", e.String(), right.String())
}

//todo more operators

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
