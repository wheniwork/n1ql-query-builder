package nqb

import "bytes"

type LettingClause interface {
	HavingClause
	Letting(aliases ...*Alias) HavingClause
}

type defaultLettingClause struct {
	*defaultHavingClause
}

func newDefaultLettingClause(parent Statement) *defaultLettingClause {
	return &defaultLettingClause{newDefaultHavingClause(parent)}
}

func (c *defaultLettingClause) Letting(aliases ...*Alias) HavingClause {
	c.setElement(&lettingElement{aliases})
	return newDefaultHavingClause(c)
}

type lettingElement struct {
	aliases []*Alias
}

func (e *lettingElement) export() string {
	buf := bytes.NewBufferString("LETTING ")

	for i, alias := range e.aliases {
		buf.WriteString(alias.String())

		// todo improve?
		if i < len(e.aliases)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
