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

func (p *abstractStatement) render() string {
	buf := bytes.Buffer{}

	if p.parent != nil {
		buf.WriteString(p.parent.render())
		buf.WriteString(" ")
	}

	if p.element != nil {
		buf.WriteString(p.element.export())
	}

	return buf.String()
}

func (p *abstractStatement) setElement(element element) {
	p.element = element
}

func (p *abstractStatement) String() string {
	return strings.TrimSpace(p.render())
}
