package nqb

import (
	"bytes"
	"fmt"
)

type unionElement struct {
	all  bool
	with string
}

func newUnionElement(all bool) *unionElement {
	return &unionElement{all, ""}
}

func newUnionElementPath(all bool, with string) *unionElement {
	return &unionElement{all, with}
}

func newUnionElementPathStmt(all bool, with Statement) *unionElement {
	return &unionElement{all, fmt.Sprintf("%s", with)}
}

func (e *unionElement) export() string {
	buf := bytes.NewBufferString("UNION")

	if e.all {
		buf.WriteString(" ALL")
	}

	if e.with != "" {
		buf.WriteString(" ")
		buf.WriteString(e.with)
	}

	return buf.String()
}
