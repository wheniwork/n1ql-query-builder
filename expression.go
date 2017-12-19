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
	value interface{}
}

// X creates an arbitrary expression from the given value.
//
// No quoting or escaping will be done on the input.
// In addition, it is not checked if the given value is an actual valid (N1QL syntax wise) expression.
func X(value interface{}) *Expression {
	switch value.(type) {
	case *Expression:
		return value.(*Expression)
	case bool:
		if value == true {
			return trueExpr
		}
		return falseExpr
	default:
		return &Expression{value}
	}
}

// Sub creates an expression from a given sub-Statement, wrapping it in parentheses.
func Sub(statement Statement) *Expression {
	return &Expression{"(" + statement.(fmt.Stringer).String() + ")"}
}

// Par wraps an Expression in parentheses.
func Par(expression *Expression) *Expression {
	return infix(expression.String(), "(", ")")
}

// Path constructs a path ("a.b.c") from Expressions or values.
// Strings are considered identifiers (so they won't be quoted).
func Path(pathComponents ...interface{}) *Expression {
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

// P is an alias for the Path function
func P(pathComponents ...interface{}) *Expression {
	return Path(pathComponents...)
}

// I constructs an identifier or list of identifiers escaped using back-quotes (`).
//
// Useful for example for identifiers that contains a dash like "beer-sample".
// Multiple identifiers are returned as a list of escaped identifiers separated by ", ".
func I(identifiers ...string) *Expression {
	return &Expression{wrapWith('`', identifiers...)}
}

// S constructs an identifier or list of identifiers which will be quoted as strings (with "").
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
func (e *Expression) And(right interface{}) *Expression {
	return infix("AND", e.String(), toString(right))
}

// Or OR-combines two expressions.
func (e *Expression) Or(right interface{}) *Expression {
	return infix("OR", e.String(), toString(right))
}

// Eq combines two expressions with the equals operator ("=").
func (e *Expression) Eq(right interface{}) *Expression {
	return infix("=", e.String(), toString(right))
}

// Ne combines two expressions with the not equals operator ("!=").
func (e *Expression) Ne(right interface{}) *Expression {
	return infix("!=", e.String(), toString(right))
}

// Gt combines two expressions with the greater than operator (">").
func (e *Expression) Gt(right interface{}) *Expression {
	return infix(">", e.String(), toString(right))
}

// Lt combines two expressions with the less than operator ("<").
func (e *Expression) Lt(right interface{}) *Expression {
	return infix("<", e.String(), toString(right))
}

// Gte combines two expressions with the greater or equals than operator (">=").
func (e *Expression) Gte(right interface{}) *Expression {
	return infix(">=", e.String(), toString(right))
}

// Concat combines two expressions with the concatenation operator ("||").
func (e *Expression) Concat(right interface{}) *Expression {
	return infix("||", e.String(), toString(right))
}

// Lte combines two expressions with the less or equals than operator ("<=").
func (e *Expression) Lte(right interface{}) *Expression {
	return infix("<=", e.String(), toString(right))
}

// IsValued appends an "IS VALUED" to the expression.
func (e *Expression) IsValued() *Expression {
	return postfix("IS VALUED", e.String())
}

// IsNotValued appends an "IS NOT VALUED" to the expression.
func (e *Expression) IsNotValued() *Expression {
	return postfix("IS NOT VALUED", e.String())
}

// IsNull appends an "IS NULL" to the expression.
func (e *Expression) IsNull() *Expression {
	return postfix("IS NULL", e.String())
}

// IsNotNull appends an "IS NOT NULL" to the expression.
func (e *Expression) IsNotNull() *Expression {
	return postfix("IS NOT NULL", e.String())
}

// IsMissing appends an "IS MISSING" to the expression.
func (e *Expression) IsMissing() *Expression {
	return postfix("IS MISSING", e.String())
}

// IsNotMissing appends an "IS NOT MISSING" to the expression.
func (e *Expression) IsNotMissing() *Expression {
	return postfix("IS NOT MISSING", e.String())
}

// Between adds a BETWEEN clause between the current and the given expression.
func (e *Expression) Between(right interface{}) *Expression {
	return infix("BETWEEN", e.String(), toString(right))
}

// NotBetween adds a NOT BETWEEN clause between the current and the given expression.
func (e *Expression) NotBetween(right interface{}) *Expression {
	return infix("NOT BETWEEN", e.String(), toString(right))
}

// Like adds a LIKE clause between the current and the given expression.
func (e *Expression) Like(right interface{}) *Expression {
	return infix("LIKE", e.String(), toString(right))
}

// NotLike adds a NOT LIKE clause between the current and the given expression.
func (e *Expression) NotLike(right interface{}) *Expression {
	return infix("NOT LIKE", e.String(), toString(right))
}

// Exists prefixes the current expression with the EXISTS clause.
func (e *Expression) Exists() *Expression {
	return prefix("EXISTS", e.String())
}

// In adds a IN clause between the current and the given expression.
func (e *Expression) In(right interface{}) *Expression {
	return infix("IN", e.String(), toString(right))
}

// NotIn adds a NOT IN clause between the current and the given expression.
func (e *Expression) NotIn(right interface{}) *Expression {
	return infix("NOT IN", e.String(), toString(right))
}

// As Adds a AS clause between the current and the given expression. Often used to alias an identifier.
func (e *Expression) As(alias interface{}) *Expression {
	return infix("AS", e.String(), toString(alias))
}

// Add establishes arithmetic addition between current and given expression.
func (e *Expression) Add(value interface{}) *Expression {
	return infix("+", e.String(), toString(value))
}

// Subtract establishes arithmetic v between current and given expression.
func (e *Expression) Subtract(value interface{}) *Expression {
	return infix("-", e.String(), toString(value))
}

// Multiply establishes arithmetic multiplication between current and given expression.
func (e *Expression) Multiply(value interface{}) *Expression {
	return infix("*", e.String(), toString(value))
}

// Divide establishes arithmetic division between current and given expression.
func (e *Expression) Divide(value interface{}) *Expression {
	return infix("/", e.String(), toString(value))
}

// Get an attribute of an object using the given value as attribute name.
func (e *Expression) Get(expression interface{}) *Expression {
	switch expression.(type) {
	case *Expression:
		return e.Get(expression.(*Expression).String())
	default:
		return &Expression{Path(e.String(), X(expression))}
	}
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
// Separates multiple arguments with a ", "
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
	switch e.value.(type) {
	case string:
		return e.value.(string)
	default:
		return fmt.Sprint(e.value)
	}
}

func toExpressions(values ...interface{}) []*Expression {
	var converted []*Expression
	for _, value := range values {
		switch value.(type) {
		case *Expression:
			converted = append(converted, value.(*Expression))
		default:
			converted = append(converted, X(value))
		}
	}
	return converted
}

func toString(value interface{}) string {
	switch value.(type) {
	case *Expression:
		return value.(*Expression).String()
	case string:
		return value.(string)
	default:
		return fmt.Sprint(value)
	}
}
