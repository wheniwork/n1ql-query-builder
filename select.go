package nqb

import "fmt"

type selectStatement struct {
	buf *buffer

	distinct bool
	raw      bool
	element  bool

	resultExpression []ResultExpression
	keyspace         *string
	alias            *string
	joinKeyspaces    []BuildFunc

	where   []Builder
	groupBy []BuildFunc
	letting []Builder
	having  []Builder
	orderBy []BuildFunc

	limit  int64
	offset int64
}

type ResultExpression struct {
	PathOrExpression string
	Alias            *string
}

// Select creates a selectStatement
func Select(resultExpression ...ResultExpression) *selectStatement {
	return &selectStatement{
		buf:              &buffer{},
		resultExpression: resultExpression,
		limit:            -1,
		offset:           -1,
	}
}

// From specifies keyspace
func (b *selectStatement) From(keyspace string, alias *string) *selectStatement {
	b.keyspace = &keyspace
	b.alias = alias
	return b
}

// Distinct adds `DISTINCT`
func (b *selectStatement) Distinct() *selectStatement {
	b.distinct = true
	return b
}

// Raw adds `RAW`
func (b *selectStatement) Raw() *selectStatement {
	b.raw = true
	return b
}

// Element adds `ELEMENT`
func (b *selectStatement) Element() *selectStatement {
	b.element = true
	return b
}

// Join joins keyspace on condition
func (b *selectStatement) Join(keyspace, on interface{}) *selectStatement {
	b.joinKeyspaces = append(b.joinKeyspaces, join(inner, keyspace, on))
	return b
}

func (b *selectStatement) LeftJoin(keyspace, on interface{}) *selectStatement {
	b.joinKeyspaces = append(b.joinKeyspaces, join(left, keyspace, on))
	return b
}

func (b *selectStatement) RightJoin(keyspace, on interface{}) *selectStatement {
	b.joinKeyspaces = append(b.joinKeyspaces, join(right, keyspace, on))
	return b
}

func (b *selectStatement) FullJoin(keyspace, on interface{}) *selectStatement {
	b.joinKeyspaces = append(b.joinKeyspaces, join(full, keyspace, on))
	return b
}

// As creates alias for select statement
// deprecated
func (b *selectStatement) As(alias string) *selectStatement {
	return as(b, alias).(*selectStatement)
}

// Where adds a where condition
func (b *selectStatement) Where(query interface{}, value ...interface{}) *selectStatement {
	switch query := query.(type) {
	case string:
		b.where = append(b.where, Expr(query, value...))
	case Builder:
		b.where = append(b.where, query)
	}
	return b
}

// Letting adds a letting clause
func (b *selectStatement) Letting(query interface{}, value ...interface{}) *selectStatement {
	switch query := query.(type) {
	case string:
		b.letting = append(b.letting, Expr(query, value...))
	case Builder:
		b.letting = append(b.letting, query)
	}
	return b
}

// Having adds a having condition
func (b *selectStatement) Having(query interface{}, value ...interface{}) *selectStatement {
	switch query := query.(type) {
	case string:
		b.having = append(b.having, Expr(query, value...))
	case Builder:
		b.having = append(b.having, query)
	}
	return b
}

// GroupBy specifies resultExpression for grouping
func (b *selectStatement) GroupBy(col ...string) *selectStatement {
	for _, group := range col {
		b.groupBy = append(b.groupBy, func(buf *buffer) error {
			buf.WriteString(group)
			return nil
		})
	}
	return b
}

// OrderBy specifies resultExpression for ordering
func (b *selectStatement) OrderAsc(col string) *selectStatement {
	b.orderBy = append(b.orderBy, order(col, asc))
	return b
}

func (b *selectStatement) OrderDesc(col string) *selectStatement {
	b.orderBy = append(b.orderBy, order(col, desc))
	return b
}

// Limit adds limit
func (b *selectStatement) Limit(n uint64) *selectStatement {
	b.limit = int64(n)
	return b
}

// Offset adds offset
func (b *selectStatement) Offset(n uint64) *selectStatement {
	b.offset = int64(n)
	return b
}

// Build builds `SELECT ...` in dialect
func (b *selectStatement) Build() error {
	if len(b.resultExpression) == 0 {
		return ErrColumnNotSpecified
	}

	b.buf.WriteString("SELECT ")

	if b.distinct {
		b.buf.WriteString("DISTINCT ")
	}


	for i, col := range b.resultExpression {
		if i > 0 {
			b.buf.WriteString(", ")
		}
		b.buf.WriteString(EscapeIdentifier(col))
	}

	if b.raw {
		b.buf.WriteString(" RAW ")
	}

	if b.element {
		b.buf.WriteString(" ELEMENT ")
	}

	if b.keyspace != nil {
		b.buf.WriteString(" FROM ")
		b.buf.WriteString(*b.keyspace)

		if b.alias != nil {
			b.buf.WriteString(fmt.Sprintf(" %s ", b.alias))
		}

		if len(b.joinKeyspaces) > 0 {
			for _, join := range b.joinKeyspaces {
				err := join.Build(b.buf)
				if err != nil {
					return err
				}
			}
		}
	}

	if len(b.where) > 0 {
		b.buf.WriteString(" WHERE ")
		err := And(b.where...).Build(b.buf)
		if err != nil {
			return err
		}
	}

	if len(b.groupBy) > 0 {
		b.buf.WriteString(" GROUP BY ")
		for i, group := range b.groupBy {
			if i > 0 {
				b.buf.WriteString(", ")
			}
			err := group.Build(b.buf)
			if err != nil {
				return err
			}
		}
	}

	if len(b.letting) > 0 {
		b.buf.WriteString(" LETTING ")
		err := And(b.letting...).Build(b.buf)
		if err != nil {
			return err
		}
	}

	if len(b.having) > 0 {
		b.buf.WriteString(" HAVING ")
		err := And(b.having...).Build(b.buf)
		if err != nil {
			return err
		}
	}

	if len(b.orderBy) > 0 {
		b.buf.WriteString(" ORDER BY ")
		for i, order := range b.orderBy {
			if i > 0 {
				b.buf.WriteString(", ")
			}
			err := order.Build(b.buf)
			if err != nil {
				return err
			}
		}
	}

	if b.limit >= 0 {
		b.buf.WriteString(" LIMIT ")
		b.buf.WriteString(fmt.Sprint(b.limit))
	}

	if b.offset >= 0 {
		b.buf.WriteString(" OFFSET ")
		b.buf.WriteString(fmt.Sprint(b.offset))
	}
	return nil
}

func (b *selectStatement) String() string {
	return b.buf.String()
}
