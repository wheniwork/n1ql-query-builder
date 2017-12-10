package nqb

import "strconv"

type OffsetPath interface {
	Statement
	Path
	Offset(offset int) Statement
}

type defaultOffsetPath struct {
	*abstractPath
}

func newDefaultOffsetPath(parent Path) *defaultOffsetPath {
	return &defaultOffsetPath{newAbstractPath(parent)}
}

func (p *defaultOffsetPath) Offset(offset int) Statement {
	p.setElement(&offsetElement{offset})
	return p
}

type offsetElement struct {
	offset int
}

func (e *offsetElement) export() string {
	return "OFFSET " + strconv.Itoa(e.offset)
}
