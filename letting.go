package nqb

import "bytes"

type LettingPath interface {
	HavingPath
	Letting(aliases ...*Alias) HavingPath
}

type defaultLettingPath struct {
	*defaultHavingPath
}

func newDefaultLettingPath(parent Path) *defaultLettingPath {
	return &defaultLettingPath{newDefaultHavingPath(parent)}
}

func (p *defaultLettingPath) Letting(aliases ...*Alias) HavingPath {
	p.setElement(&lettingElement{aliases})
	return newDefaultHavingPath(p)
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
