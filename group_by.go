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

func (c *defaultGroupByClause) GroupBy(identifiers ...interface{}) LettingClause {
	c.setElement(&groupByElement{toExpressions(identifiers...)})
	return newDefaultLettingClause(c)
}

type groupByElement struct {
	expressions []*Expression
}

func (e *groupByElement) export() string {
	buf := bytes.NewBufferString("GROUP BY ")

	for i, expression := range e.expressions {
		buf.WriteString(expression.String())

		// todo improve?
		if i < len(e.expressions)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
