package nqb

import (
	"bytes"
)

type unionElement struct {
	all  bool
	with string
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
