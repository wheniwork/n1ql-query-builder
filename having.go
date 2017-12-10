package nqb

type HavingPath interface {
	SelectResultPath
	Having(condition *Expression) SelectResultPath
}

type defaultHavingPath struct {
	*defaultSelectResultPath
}

func newDefaultHavingPath(parent Path) *defaultHavingPath {
	return &defaultHavingPath{newDefaultSelectResultPath(parent)}
}

func (p *defaultHavingPath) Having(condition *Expression) SelectResultPath {
	p.setElement(newHavingElement(condition))
	return newDefaultSelectResultPath(p)
}

type havingElement struct {
	expression *Expression
}

func newHavingElement(expression *Expression) *havingElement {
	return &havingElement{expression}
}

func (e *havingElement) Export() string {
	return "HAVING " + e.expression.String()
}
