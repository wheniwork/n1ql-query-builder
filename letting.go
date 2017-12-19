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

func (p *defaultLettingClause) Letting(aliases ...*Alias) HavingClause {
	p.setElement(&lettingElement{aliases})
	return newDefaultHavingClause(p)
}

type lettingElement struct {
	aliases []*Alias
}

func (e *lettingElement) export() string {
	n1ql := bytes.NewBufferString("LETTING ")

	for i, alias := range e.aliases {
		n1ql.WriteString(alias.String())

		// todo improve?
		if i < len(e.aliases)-1 {
			n1ql.WriteString(", ")
		}
	}

	return n1ql.String()
}
