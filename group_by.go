package nqb

import "bytes"

type GroupByClause interface {
	SelectResult

	// GroupBy adds a GROUP BY clause.
	GroupBy(identifiers ...interface{}) LettingClause
}

type defaultGroupByClause struct {
	*defaultSelectResult
}

func newDefaultGroupByClause(parent Statement) *defaultGroupByClause {
	return &defaultGroupByClause{newDefaultSelectResult(parent)}
}

func (p *defaultGroupByClause) GroupBy(identifiers ...interface{}) LettingClause {
	p.setElement(&groupByElement{toExpressions(identifiers...)})
	return newDefaultLettingClause(p)
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
