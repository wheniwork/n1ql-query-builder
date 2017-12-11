package nqb

import (
	"bytes"
)

type exceptElement struct {
	all  bool
	with string
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
