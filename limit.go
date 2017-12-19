package nqb

import "strconv"

type LimitClause interface {
	OffsetClause
	Limit(limit int) OffsetClause
}

type defaultLimitClause struct {
	*defaultOffsetClause
}

func newDefaultLimitClause(parent Statement) *defaultLimitClause {
	return &defaultLimitClause{newDefaultOffsetClause(parent)}
}

func (c *defaultLimitClause) Limit(limit int) OffsetClause {
	c.setElement(&limitElement{limit})
	return newDefaultOffsetClause(c)
}

type limitElement struct {
	limit int
}

func (e *limitElement) export() string {
	return "LIMIT " + strconv.Itoa(e.limit)
}
