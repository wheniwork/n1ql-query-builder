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
	p.setElement(&asElement{alias})
	return newDefaultKeysPath(p)
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
