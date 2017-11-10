package nqb

import "fmt"

type selectStmt struct {
	buf *buffer

	distinct bool
	raw      bool

	columns       []string
	keyspace      *string
	alias         *string
	joinKeyspaces []BuildFunc

	where   []Builder
	groupBy []BuildFunc
	letting []Builder
	having  []Builder
	orderBy []BuildFunc

	limit  int64
	offset int64
}

// Select creates a selectStmt
func Select(column ...string) *selectStmt {
	return &selectStmt{
		buf:     &buffer{},
		columns: column,
		limit:   -1,
		offset:  -1,
	}
}

// From specifies keyspace
func (b *selectStmt) From(keyspace string, alias *string) *selectStmt {
	b.keyspace = &keyspace
	b.alias = alias
	return b
}

// Distinct adds `DISTINCT`
func (b *selectStmt) Distinct() *selectStmt {
	b.distinct = true
	return b
}

// Raw adds `RAW`
func (b *selectStmt) Raw() *selectStmt {
	b.raw = true
	return b
}

// Join joins keyspace on condition
func (b *selectStmt) Join(keyspace, on interface{}) *selectStmt {
	b.joinKeyspaces = append(b.joinKeyspaces, join(inner, keyspace, on))
	return b
}

func (b *selectStmt) LeftJoin(keyspace, on interface{}) *selectStmt {
	b.joinKeyspaces = append(b.joinKeyspaces, join(left, keyspace, on))
	return b
}

func (b *selectStmt) RightJoin(keyspace, on interface{}) *selectStmt {
	b.joinKeyspaces = append(b.joinKeyspaces, join(right, keyspace, on))
	return b
}

func (b *selectStmt) FullJoin(keyspace, on interface{}) *selectStmt {
	b.joinKeyspaces = append(b.joinKeyspaces, join(full, keyspace, on))
	return b
}

// As creates alias for select statement
// deprecated
func (b *selectStmt) As(alias string) *selectStmt {
	return as(b, alias).(*selectStmt)
}

// Where adds a where condition
func (b *selectStmt) Where(query interface{}, value ...interface{}) *selectStmt {
	switch query := query.(type) {
	case string:
		b.where = append(b.where, Expr(query, value...))
	case Builder:
		b.where = append(b.where, query)
	}
	return b
}

// Letting adds a letting clause
func (b *selectStmt) Letting(query interface{}, value ...interface{}) *selectStmt {
	switch query := query.(type) {
	case string:
		b.letting = append(b.letting, Expr(query, value...))
	case Builder:
		b.letting = append(b.letting, query)
	}
	return b
}

// Having adds a having condition
func (b *selectStmt) Having(query interface{}, value ...interface{}) *selectStmt {
	switch query := query.(type) {
	case string:
		b.having = append(b.having, Expr(query, value...))
	case Builder:
		b.having = append(b.having, query)
	}
	return b
}

// GroupBy specifies columns for grouping
func (b *selectStmt) GroupBy(col ...string) *selectStmt {
	for _, group := range col {
		b.groupBy = append(b.groupBy, func(buf *buffer) error {
			buf.WriteString(group)
			return nil
		})
	}
	return b
}

// OrderBy specifies columns for ordering
func (b *selectStmt) OrderAsc(col string) *selectStmt {
	b.orderBy = append(b.orderBy, order(col, asc))
	return b
}

func (b *selectStmt) OrderDesc(col string) *selectStmt {
	b.orderBy = append(b.orderBy, order(col, desc))
	return b
}

// Limit adds limit
func (b *selectStmt) Limit(n uint64) *selectStmt {
	b.limit = int64(n)
	return b
}

// Offset adds offset
func (b *selectStmt) Offset(n uint64) *selectStmt {
	b.offset = int64(n)
	return b
}

// Build builds `SELECT ...` in dialect
func (b *selectStmt) Build() error {
	if len(b.columns) == 0 {
		return ErrColumnNotSpecified
	}

	b.buf.WriteString("SELECT ")

	if b.distinct {
		b.buf.WriteString("DISTINCT ")
	}

	if b.raw {
		b.buf.WriteString("RAW ")
	}

	for i, col := range b.columns {
		if i > 0 {
			b.buf.WriteString(", ")
		}
		b.buf.WriteString(EscapeIdentifier(col))
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

func (b *selectStmt) String() string {
	return b.buf.String()
}
