package nqb

import "bytes"

type NestClause interface {
	KeysClauses
	As(alias string) KeysClauses
}

type defaultNestClause struct {
	*defaultKeysClauses
}

func newDefaultNestClause(parent Statement) *defaultNestClause {
	return &defaultNestClause{newDefaultKeysClauses(parent)}
}

func (c *defaultNestClause) As(alias string) KeysClauses {
	c.setElement(&asElement{alias})
	return newDefaultKeysClauses(c)
}

type nestElement struct {
	joinType joinType
	from     string
}

func (e *nestElement) export() string {
	buf := bytes.Buffer{}

	if e.joinType != defaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("NEST ")
	buf.WriteString(e.from)

	return buf.String()
}
