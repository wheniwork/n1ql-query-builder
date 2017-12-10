package nqb

import "bytes"

type NestPath interface {
	KeysPath
	As(alias string) KeysPath
}

type defaultNestPath struct {
	*defaultKeysPath
}

func newDefaultNestPath(parent Path) *defaultNestPath {
	return &defaultNestPath{newDefaultKeysPath(parent)}
}

func (p *defaultNestPath) As(alias string) KeysPath {
	p.setElement(newAsAlement(alias))
	return newDefaultKeysPath(p)
}

type nestElement struct {
	joinType JoinType
	from     string
}

func newNestElement(joinType JoinType, from string) *nestElement {
	return &nestElement{joinType, from}
}

func (e *nestElement) Export() string {
	buf := bytes.Buffer{}

	if e.joinType != DefaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("NEST ")
	buf.WriteString(e.from)

	return buf.String()
}
