package nqb

type SelectPath interface {
	Path

	Select(expressions ...interface{}) FromPath

	SelectAll(expressions ...interface{}) FromPath

	SelectDistinct(expressions ...interface{}) FromPath

	SelectRaw(expression interface{}) FromPath
}

type defaultSelectPath struct {
	*abstractPath
}

func newDefaultSelectPath(parent Path) *defaultSelectPath {
	return &defaultSelectPath{newAbstractPath(parent)}
}

func (p *defaultSelectPath) Select(expressions ...interface{}) FromPath {
	p.setElement(&selectElement{defaultSelect, toExpressions(expressions)})
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectAll(expressions ...interface{}) FromPath {
	p.setElement(&selectElement{all, toExpressions(expressions)})
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectDistinct(expressions ...interface{}) FromPath {
	p.setElement(&selectElement{distinct, toExpressions(expressions)})
	return newDefaultFromPath(p)
}

func (p *defaultSelectPath) SelectRaw(expression interface{}) FromPath {
	switch expression.(type) {
	case *Expression:
		p.setElement(&selectElement{raw, []*Expression{expression.(*Expression)}})
	default:
		p.setElement(&selectElement{raw, []*Expression{X(expression)}})
	}

	return newDefaultFromPath(p)
}
