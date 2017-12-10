package nqb

type WherePath interface {
	GroupByPath

	WhereExpr(expression *Expression) GroupByPath

	Where(expression string) GroupByPath
}

type defaultWherePath struct {
	*defaultGroupByPath
}

func newDefaultWherePath(parent Path) *defaultWherePath {
	return &defaultWherePath{newDefaultGroupByPath(parent)}
}

func (p *defaultWherePath) WhereExpr(expression *Expression) GroupByPath {
	p.setElement(&whereElement{expression})
	return newDefaultGroupByPath(p)
}

func (p *defaultWherePath) Where(expression string) GroupByPath {
	return p.WhereExpr(X(expression))
}

type whereElement struct {
	expression *Expression
}

func (e *whereElement) export() string {
	return "WHERE " + e.expression.String()
}