package nqb

import "bytes"

type GroupByPath interface {
	SelectResultPath

	GroupByExpr(expressions ...*Expression) LettingPath

	GroupBy(identifiers ...string) LettingPath
}

type defaultGroupByPath struct {
	*defaultSelectResultPath
}

func newDefaultGroupByPath(parent Path) *defaultGroupByPath {
	return &defaultGroupByPath{newDefaultSelectResultPath(parent)}
}

func (p *defaultGroupByPath) GroupByExpr(expressions ...*Expression) LettingPath {
	p.setElement(&groupByElement{expressions})
	return newDefaultLettingPath(p)
}

func (p *defaultGroupByPath) GroupBy(identifiers ...string) LettingPath {
	var expressions []*Expression
	for _, identifier := range identifiers {
		expressions = append(expressions, X(identifier))
	}
	return p.GroupByExpr(expressions...)
}

type groupByElement struct {
	expressions []*Expression
}

func (e *groupByElement) export() string {
	n1ql := bytes.NewBufferString("GROUP BY ")

	for i, expression := range e.expressions {
		n1ql.WriteString(expression.String())

		// todo improve?
		if i < len(e.expressions)-1 {
			n1ql.WriteString(", ")
		}
	}

	return n1ql.String()
}