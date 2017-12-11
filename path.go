package nqb

import (
	"bytes"
	"strings"
)

type Path interface {
	render() string
	setElement(element element)
	String() string
}

type abstractPath struct {
	element element
	parent  Path
}

func (p *abstractPath) render() string {
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

func (p *abstractPath) setElement(element element) {
	p.element = element
}

func (p *abstractPath) String() string {
	return strings.TrimSpace(p.render())
}
