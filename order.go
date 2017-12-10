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
	p.setElement(newOrderByElement(orderings...))
	return newDefaultLimitPath(p)
}

type orderByElement struct {
	sorts []Sort
}

func newOrderByElement(sorts ...Sort) *orderByElement {
	return &orderByElement{sorts}
}

func (e *orderByElement) Export() string {
	buf := bytes.NewBufferString("ORDER BY ")

	for i, sort := range e.sorts {
		buf.WriteString(sort.String())
		if i < len(e.sorts)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Sort struct {
	expression *Expression
	ordering   *Order
}

func newSort(expression *Expression, ordering *Order) *Sort {
	return &Sort{expression, ordering}
}

func DefaultSortExpr(expression *Expression) *Sort {
	return newSort(expression, nil)
}

func DefaultSort(expression string) *Sort {
	return DefaultSortExpr(X(expression))
}

func DescSortExpr(expression *Expression) *Sort {
	desc := DESC
	return newSort(expression, &desc)
}

func DescSort(expression string) *Sort {
	return DescSortExpr(X(expression))
}

func AscSortExpr(expression *Expression) *Sort {
	asc := ASC
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
