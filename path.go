package nqb

import "bytes"

type Path interface {
	render() string
	setElement(element Element)
	String() string
}

type abstractPath struct {
	element Element
	parent  Path
}

func newAbstractPath(parent Path) *abstractPath {
	return &abstractPath{parent: parent}
}

func (p *abstractPath) render() string {
	buf := bytes.Buffer{}

	if p.parent != nil {
		buf.WriteString(p.parent.render())
		buf.WriteString(" ")
	}

	if p.element != nil {
		buf.WriteString(p.element.Export())
	}

	return buf.String()
}

func (p *abstractPath) setElement(element Element) {
	p.element = element
}

func (p *abstractPath) String() string {
	return p.render()
}
