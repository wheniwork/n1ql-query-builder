package nqb

import "bytes"

type UnnestPath interface {
	LetPath
	As(alias string) LetPath
}

type defaultUnnestPath struct {
	*defaultLetPath
}

func newDefaultUnnestPath(parent Path) *defaultUnnestPath {
	return &defaultUnnestPath{newDefaultLetPath(parent)}
}

func (p *defaultUnnestPath) As(alias string) LetPath {
	p.setElement(&asElement{alias})
	return newDefaultLetPath(p)
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
