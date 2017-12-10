package nqb

import (
	"bytes"
	"fmt"
)

type Expression struct {
	value interface{}
}

func newExpression(value interface{}) *Expression {
	return &Expression{value}
}

func X(value interface{}) *Expression {
	return newExpression(value)
}

func S(strings ...string) *Expression {
	return newExpression(wrapWith('"', strings...))
}

func (e *Expression) String() string {
	return fmt.Sprintf("%s", e.value)
}

func wrapWith(wrapper byte, input ...string) string {
	escaped := bytes.Buffer{}

	for _, i := range input {
		escaped.WriteString(", ")
		escaped.WriteByte(wrapper)
		escaped.WriteString(i)
		escaped.WriteByte(wrapper)
	}

	str := escaped.String()

	if len(str) > 2 {
		return str[2:]
	}

	return str
}
