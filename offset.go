package nqb

import "strconv"

type OffsetClause interface {
	Statement
	Offset(offset int) Statement
}

type defaultOffsetClause struct {
	*abstractStatement
}

func newDefaultOffsetClause(parent Statement) *defaultOffsetClause {
	return &defaultOffsetClause{&abstractStatement{parent: parent}}
}

func (p *defaultOffsetClause) Offset(offset int) Statement {
	p.setElement(&offsetElement{offset})
	return p
}

type offsetElement struct {
	offset int
}

func (e *offsetElement) export() string {
	return "OFFSET " + strconv.Itoa(e.offset)
}
