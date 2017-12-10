package nqb

import (
	"bytes"
	"fmt"
)

type intersectElement struct {
	all  bool
	with string
}

func newIntersectElement(all bool) *intersectElement {
	return &intersectElement{all, ""}
}

func newIntersectElementPath(all bool, with string) *intersectElement {
	return &intersectElement{all, with}
}

func newIntersectElementPathStmt(all bool, with Statement) *intersectElement {
	return &intersectElement{all, fmt.Sprintf("%s", with)}
}

func (e *intersectElement) export() string {
	buf := bytes.NewBufferString("INTERSECT")

	if e.all {
		buf.WriteString(" ALL")
	}

	if e.with != "" {
		buf.WriteString(" ")
		buf.WriteString(e.with)
	}

	return buf.String()
}
