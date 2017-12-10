package nqb

type WherePath interface {
	GroupByPath

	// Where adds a WHERE clause
	Where(expression interface{}) GroupByPath
}

type defaultWherePath struct {
	*defaultGroupByPath
}

func newDefaultWherePath(parent Path) *defaultWherePath {
	return &defaultWherePath{newDefaultGroupByPath(parent)}
}

func (p *defaultWherePath) Where(expression interface{}) GroupByPath {
	switch expression.(type) {
	case *Expression:
		p.setElement(&whereElement{expression.(*Expression)})
	default:
		p.setElement(&whereElement{X(expression)})
	}

	return newDefaultGroupByPath(p)
}

type whereElement struct {
	expression *Expression
}

func (e *whereElement) export() string {
	return "WHERE " + e.expression.String()
}
