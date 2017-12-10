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
	p.setElement(newAsAlement(alias))
	return newDefaultKeysPath(p)
}

type JoinType string

const (
	DefaultJoin JoinType = ""
	Inner       JoinType = "INNER"
	Left        JoinType = "LEFT"
	LeftOuter   JoinType = "LEFT OUTER"
)

type joinElement struct {
	joinType JoinType
	from     string
}

func newJoinElement(joinType JoinType, from string) *joinElement {
	return &joinElement{joinType, from}
}

func (e *joinElement) Export() string {
	buf := bytes.Buffer{}

	if e.joinType != DefaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("JOIN ")
	buf.WriteString(e.from)

	return buf.String()
}
