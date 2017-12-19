package nqb

type WhereClause interface {
	GroupByClause

	// Where adds a WHERE clause
	Where(expression interface{}) GroupByClause
}

type defaultWhereClause struct {
	*defaultGroupByClause
}

func newDefaultWhereClause(parent Statement) *defaultWhereClause {
	return &defaultWhereClause{newDefaultGroupByClause(parent)}
}

func (c *defaultWhereClause) Where(expression interface{}) GroupByClause {
	switch expression.(type) {
	case *Expression:
		c.setElement(&whereElement{expression.(*Expression)})
	default:
		c.setElement(&whereElement{X(expression)})
	}

	return newDefaultGroupByClause(c)
}

type whereElement struct {
	expression *Expression
}

func (e *whereElement) export() string {
	return "WHERE " + e.expression.String()
}
