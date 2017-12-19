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

func (c *defaultSelectClause) Select(expressions ...interface{}) FromClause {
	c.setElement(&selectElement{defaultSelect, toExpressions(expressions...)})
	return newDefaultFromClause(c)
}

func (c *defaultSelectClause) SelectAll(expressions ...interface{}) FromClause {
	c.setElement(&selectElement{all, toExpressions(expressions...)})
	return newDefaultFromClause(c)
}

func (c *defaultSelectClause) SelectDistinct(expressions ...interface{}) FromClause {
	c.setElement(&selectElement{distinct, toExpressions(expressions...)})
	return newDefaultFromClause(c)
}

func (c *defaultSelectClause) SelectRaw(expression interface{}) FromClause {
	c.setElement(&selectElement{raw, []*Expression{X(expression)}})
	return newDefaultFromClause(c)
}
