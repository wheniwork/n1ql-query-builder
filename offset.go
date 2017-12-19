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

func (c *defaultOffsetClause) Offset(offset int) Statement {
	c.setElement(&offsetElement{offset})
	return c
}

type offsetElement struct {
	offset int
}

func (e *offsetElement) export() string {
	return "OFFSET " + strconv.Itoa(e.offset)
}
