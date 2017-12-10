package nqb

import "bytes"

type GroupByPath interface {
	SelectResultPath

	// GroupBy adds a GROUP BY clause.
	GroupBy(identifiers ...interface{}) LettingPath
}

type defaultGroupByPath struct {
	*defaultSelectResultPath
}

func newDefaultGroupByPath(parent Path) *defaultGroupByPath {
	return &defaultGroupByPath{newDefaultSelectResultPath(parent)}
}

func (p *defaultGroupByPath) GroupBy(identifiers ...interface{}) LettingPath {
	p.setElement(&groupByElement{toExpressions(identifiers)})
	return newDefaultLettingPath(p)
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
