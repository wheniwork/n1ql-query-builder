package nqb

import (
	"bytes"
	"fmt"
)

type exceptElement struct {
	all  bool
	with string
}

func newExceptElement(all bool) *exceptElement {
	return &exceptElement{all, ""}
}

func newExceptElementPath(all bool, with string) *exceptElement {
	return &exceptElement{all, with}
}

func newExceptElementPathStmt(all bool, with Statement) *exceptElement {
	return &exceptElement{all, fmt.Sprintf("%s", with)}
}

func (e *exceptElement) export() string {
	buf := bytes.NewBufferString("EXCEPT")

	if e.all {
		buf.WriteString(" ALL")
	}

	if e.with != "" {
		buf.WriteString(" ")
		buf.WriteString(e.with)
	}

	return buf.String()
}
