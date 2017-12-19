package nqb

type SelectClause interface {
	Statement

	Select(expressions ...interface{}) FromClause

	SelectAll(expressions ...interface{}) FromClause

	SelectDistinct(expressions ...interface{}) FromClause

	SelectRaw(expression interface{}) FromClause
}

type defaultSelectClause struct {
	*abstractStatement
}

func newDefaultSelectClause(parent Statement) *defaultSelectClause {
	return &defaultSelectClause{&abstractStatement{parent: parent}}
}

func (p *defaultSelectClause) Select(expressions ...interface{}) FromClause {
	p.setElement(&selectElement{defaultSelect, toExpressions(expressions...)})
	return newDefaultFromClause(p)
}

func (p *defaultSelectClause) SelectAll(expressions ...interface{}) FromClause {
	p.setElement(&selectElement{all, toExpressions(expressions...)})
	return newDefaultFromClause(p)
}

func (p *defaultSelectClause) SelectDistinct(expressions ...interface{}) FromClause {
	p.setElement(&selectElement{distinct, toExpressions(expressions...)})
	return newDefaultFromClause(p)
}

func (p *defaultSelectClause) SelectRaw(expression interface{}) FromClause {
	p.setElement(&selectElement{raw, []*Expression{X(expression)}})
	return newDefaultFromClause(p)
}
