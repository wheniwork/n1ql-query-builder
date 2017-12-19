package nqb

type FromClause interface {
	LetClause

	From(from interface{}) AsKeyword
}

type defaultFromClause struct {
	*defaultLetClause
}

func newDefaultFromClause(parent Statement) *defaultFromClause {
	return &defaultFromClause{newDefaultLetClause(parent)}
}

func (p *defaultFromClause) From(from interface{}) AsKeyword {
	p.setElement(&fromElement{toString(from)})
	return newDefaultAsKeyword(p)
}

type fromElement struct {
	from string
}

func (e *fromElement) export() string {
	return "FROM " + e.from
}
