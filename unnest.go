package nqb

import "bytes"

type UnnestClause interface {
	LetClause
	As(alias string) LetClause
}

type defaultUnnestClause struct {
	*defaultLetClause
}

func newDefaultUnnestClause(parent Statement) *defaultUnnestClause {
	return &defaultUnnestClause{newDefaultLetClause(parent)}
}

func (p *defaultUnnestClause) As(alias string) LetClause {
	p.setElement(&asElement{alias})
	return newDefaultLetClause(p)
}

type unnestElement struct {
	joinType joinType
	path     string
}

func newUnnestElement(joinType joinType, path string) *unnestElement {
	return &unnestElement{joinType, path}
}

func (e *unnestElement) export() string {
	buf := bytes.Buffer{}

	if e.joinType != defaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("UNNEST ")
	buf.WriteString(e.path)

	return buf.String()
}
