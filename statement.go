package nqb

import (
	"bytes"
	"strings"
)

type Statement interface {
	render() string
	setElement(element element)
	String() string
}

type abstractStatement struct {
	element element
	parent  Statement
}

func (s *abstractStatement) render() string {
	buf := bytes.Buffer{}

	if s.parent != nil {
		buf.WriteString(s.parent.render())
		buf.WriteString(" ")
	}

	if s.element != nil {
		buf.WriteString(s.element.export())
	}

	return buf.String()
}

func (s *abstractStatement) setElement(element element) {
	s.element = element
}

func (s *abstractStatement) String() string {
	return strings.TrimSpace(s.render())
}
