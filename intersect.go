package nqb

import (
	"bytes"
)

type intersectElement struct {
	all  bool
	with string
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
