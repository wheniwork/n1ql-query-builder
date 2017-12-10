package nqb

import "bytes"

type JoinPath interface {
	KeysPath
	As(alias string) KeysPath
}

type defaultJoinPath struct {
	*defaultKeysPath
}

func newDefaultJoinPath(parent Path) *defaultJoinPath {
	return &defaultJoinPath{newDefaultKeysPath(parent)}
}

func (p *defaultJoinPath) As(alias string) KeysPath {
	p.setElement(&asElement{alias})
	return newDefaultKeysPath(p)
}

type joinType string

const (
	defaultJoin joinType = ""
	inner       joinType = "INNER"
	left        joinType = "LEFT"
	leftOuter   joinType = "LEFT OUTER"
)

type joinElement struct {
	joinType joinType
	from     string
}

func (e *joinElement) export() string {
	buf := bytes.Buffer{}

	if e.joinType != defaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("JOIN ")
	buf.WriteString(e.from)

	return buf.String()
}
