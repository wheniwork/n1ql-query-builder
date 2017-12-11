package nqb

import "bytes"

type OrderByPath interface {
	LimitPath
	OrderBy(orderings ...*sort) LimitPath
}

type defaultOrderByPath struct {
	*defaultLimitPath
}

func newDefaultOrderByPath(parent Path) *defaultOrderByPath {
	return &defaultOrderByPath{newDefaultLimitPath(parent)}
}

func (p *defaultOrderByPath) OrderBy(orderings ...*sort) LimitPath {
	p.setElement(&orderByElement{orderings})
	return newDefaultLimitPath(p)
}

type orderByElement struct {
	sorts []*sort
}

func (e *orderByElement) export() string {
	buf := bytes.NewBufferString("ORDER BY ")

	for i, sort := range e.sorts {
		buf.WriteString(sort.String())

		// todo improve
		if i < len(e.sorts)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}

type order string

const (
	asc  order = "ASC"
	desc order = "DESC"
)

type sort struct {
	expression interface{}
	ordering   *order
}

// DefaultSort won't specify an order in the resulting expression.
func DefaultSort(expression interface{}) *sort {
	return &sort{expression, nil}
}

// Desc specifies descending order in the resulting expression.
func Desc(expression interface{}) *sort {
	desc := desc
	return &sort{expression, &desc}
}

// Asc specifies ascending order in the resulting expression.
func Asc(expression interface{}) *sort {
	asc := asc
	return &sort{expression, &asc}
}

func (s *sort) String() string {
	expr := toString(s.expression)

	if s.ordering != nil {
		return expr + " " + string(*s.ordering)
	}

	return expr
}
