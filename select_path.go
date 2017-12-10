package nqb

type SelectPath interface {
	Path

	SelectExpr(expressions ...*Expression) FromPath

	Select(expressions ...string) FromPath

	SelectAllExpr(expressions ...*Expression) FromPath

	SelectAll(expressions ...string) FromPath

	SelectDistinctExpr(expressions ...*Expression) FromPath

	SelectDistinct(expressions ...string) FromPath

	SelectRawExpr(expression *Expression) FromPath

	SelectRaw(expression string) FromPath
}

type defaultSelectPath struct {
	*abstractPath
}

func newDefaultSelectPath(parent Path) *defaultSelectPath {
	return &defaultSelectPath{newAbstractPath(parent)}
}

func (p *defaultSelectPath) SelectExpr(expressions ...*Expression) FromPath {
	p.setElement(newSelectElement(DefaultSelect, expressions...))
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) Select(expressions ...string) FromPath {
	var converted []*Expression
	for _, expression := range expressions {
		converted = append(converted, X(expression))
	}
	return p.SelectExpr(converted...)
}

func (p *defaultSelectPath) SelectAllExpr(expressions ...*Expression) FromPath {
	p.setElement(newSelectElement(All, expressions...))
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectAll(expressions ...string) FromPath {
	var converted []*Expression
	for _, expression := range expressions {
		converted = append(converted, X(expression))
	}
	return p.SelectAllExpr(converted...)
}

func (p *defaultSelectPath) SelectDistinctExpr(expressions ...*Expression) FromPath {
	p.setElement(newSelectElement(Distinct, expressions...))
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectDistinct(expressions ...string) FromPath {
	var converted []*Expression
	for _, expression := range expressions {
		converted = append(converted, X(expression))
	}
	return p.SelectDistinctExpr(converted...)
}

func (p *defaultSelectPath) SelectRawExpr(expression *Expression) FromPath {
	p.setElement(newSelectElement(Raw, expression))
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectRaw(expression string) FromPath {
	return p.SelectRawExpr(X(expression))
}
