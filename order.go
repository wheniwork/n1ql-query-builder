package nqb

import "bytes"

type OrderByPath interface {
	LimitPath
	OrderBy(orderings ...Sort) LimitPath
}

type defaultOrderByPath struct {
	*defaultLimitPath
}

func newDefaultOrderByPath(parent Path) *defaultOrderByPath {
	return &defaultOrderByPath{newDefaultLimitPath(parent)}
}

func (p *defaultOrderByPath) OrderBy(orderings ...Sort) LimitPath {
	p.setElement(&orderByElement{orderings})
	return newDefaultLimitPath(p)
}

type orderByElement struct {
	sorts []Sort
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

type Sort struct {
	expression *Expression
	ordering   *order
}

func newSort(expression *Expression, ordering *order) *Sort {
	return &Sort{expression, ordering}
}

func DefaultSortExpr(expression *Expression) *Sort {
	return newSort(expression, nil)
}

func DefaultSort(expression string) *Sort {
	return DefaultSortExpr(X(expression))
}

func DescSortExpr(expression *Expression) *Sort {
	desc := desc
	return newSort(expression, &desc)
}

func DescSort(expression string) *Sort {
	return DescSortExpr(X(expression))
}

func AscSortExpr(expression *Expression) *Sort {
	asc := asc
	return newSort(expression, &asc)
}

func AscSort(expression string) *Sort {
	return AscSortExpr(X(expression))
}

func (s *Sort) String() string {
	if s.ordering != nil {
		return s.expression.String() + " " + string(*s.ordering)
	}

	return s.expression.String()
}
