package nqb

type HavingClause interface {
	SelectResult
	Having(condition *Expression) SelectResult
}

type defaultHavingClause struct {
	*defaultSelectResult
}

func newDefaultHavingClause(parent Statement) *defaultHavingClause {
	return &defaultHavingClause{newDefaultSelectResult(parent)}
}

func (p *defaultHavingClause) Having(condition *Expression) SelectResult {
	p.setElement(&havingElement{condition})
	return newDefaultSelectResult(p)
}

type havingElement struct {
	expression *Expression
}

func (e *havingElement) export() string {
	return "HAVING " + e.expression.String()
}
