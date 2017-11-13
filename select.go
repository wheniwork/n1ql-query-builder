package nqb

import "fmt"

type selectStatement struct {
	buf *buffer

	distinct bool

	resultExpressions []resultExpression

	raw *rawExpression

	keyspace *string
	subquery *selectStatement
	alias *string

	keys    []string
	primary bool

	joins   []join
	nests   []nest
	unnests []unnest

	indexRefs []indexRef

	let     []BuildFunc
	where   []Builder
	groupBy []BuildFunc
	letting []Builder
	having  []Builder
	orderBy []BuildFunc

	limit  int64
	offset int64
}

type resultExpression struct {
	pathOrExpression string
	alias            *string
}

func ResultExpr(pathOrExpression string, alias *string) *resultExpression {
	return &resultExpression{
		pathOrExpression: pathOrExpression,
		alias:            alias,
	}
}

type rawExpression struct {
	element    bool
	expression string
	alias *string
}

// Select creates a selectStatement
func Select(resultExpression ...resultExpression) *selectStatement {
	return &selectStatement{
		buf:               &buffer{},
		resultExpressions: resultExpression,
		limit:             -1,
		offset:            -1,
	}
}

// Distinct adds `DISTINCT`
func (b *selectStatement) Distinct() *selectStatement {
	b.distinct = true
	return b
}

// Raw adds `RAW`
func (b *selectStatement) Raw(element bool, expression string, alias *string) *selectStatement {
	b.raw = &rawExpression{element, expression, alias}
	return b
}

// From specifies keyspace
func (b *selectStatement) From(keyspace *string, subquery *selectStatement, alias *string) *selectStatement {
	b.keyspace = keyspace
	b.subquery = subquery
	b.alias = alias
	return b
}

// UseKeys specifies keys to use
func (b *selectStatement) UseKeys(primary bool, expression ...string) *selectStatement {
	b.keys = expression
	b.primary = primary
	return b
}

func (b *selectStatement) LookupJoin(joinType joinType, fromPath string, alias *string, onKeys onKeysClause) *selectStatement {
	b.joins = append(b.joins, join{&joinType, fromPath, alias, &onKeys, nil})
	return b
}

func (b *selectStatement) IndexJoin(
	joinType joinType, fromPath string, alias *string, onKeys *onKeysClause, onKeyFor *onKeyForClause,
) *selectStatement {
	b.joins = append(b.joins, join{&joinType, fromPath, alias, onKeys, onKeyFor})
	return b
}

func (b *selectStatement) Nest(joinType joinType, fromPath string, alias *string, onKeys onKeysClause) *selectStatement {
	b.nests = append(b.nests, nest{&joinType, fromPath, alias, onKeys})
	return b
}

func (b *selectStatement) Unnest(joinType joinType, flatten bool, expression string, alias *string) *selectStatement {
	b.unnests = append(b.unnests, unnest{&joinType, flatten, expression, alias})
	return b
}

func (b *selectStatement) UseIndex(indexRef ...indexRef) *selectStatement {
	b.indexRefs = indexRef
	return b
}

func (b *selectStatement) Let(alias, expression string) *selectStatement {
	b.let = append(b.let, func(buf *buffer) error {
		buf.WriteString(fmt.Sprintf(" (%s = %s) ", alias, expression))
	})
	return b
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

// GroupBy specifies resultExpressions for grouping
func (b *selectStatement) GroupBy(col ...string) *selectStatement {
	for _, group := range col {
		b.groupBy = append(b.groupBy, func(buf *buffer) error {
			buf.WriteString(group)
			return nil
		})
	}
	return b
}

// OrderBy specifies resultExpressions for ordering
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

// Build builds `SELECT ...`
func (b *selectStatement) Build() error {
	if len(b.resultExpressions) == 0 {
		return ErrColumnNotSpecified
	}

	b.buf.WriteString("SELECT ")

	if b.distinct {
		b.buf.WriteString("DISTINCT ")
	}

	for i, resultExpression := range b.resultExpressions {
		if i > 0 {
			b.buf.WriteString(", ")
		}
		b.buf.WriteString(escapeIdentifiers(resultExpression.pathOrExpression))

		if resultExpression.alias != nil {
			b.buf.WriteString(escapeIdentifiers(*resultExpression.alias))
		}
	}

	if b.raw != nil {
		if b.raw.element {
			b.buf.WriteString(" ELEMENT ")
		} else {
			b.buf.WriteString(" RAW ")
		}

		b.buf.WriteString(b.raw.expression)

		if b.raw.alias != nil {
			b.buf.WriteString(fmt.Sprintf(" AS %s", escapeIdentifiers(*b.raw.alias)))
		}
	}

	if b.keyspace != nil {
		b.buf.WriteString(" FROM ")
		b.buf.WriteString(escapeIdentifiers(*b.keyspace))

		if b.alias != nil {
			b.buf.WriteString(fmt.Sprintf(" %s ", escapeIdentifiers(*b.alias)))
		}

		if len(b.joins) > 0 {
			for _, join := range b.joins {
				join.Build(b.buf)
			}
		}

		if len(b.nests) > 0 {
			for _, nest := range b.nests {
				nest.Build(b.buf)
			}
		}

		if len(b.unnests) > 0 {
			for _, unnest := range b.unnests {
				unnest.Build(b.buf)
			}
		}
	}

	if b.subquery != nil {
		b.buf.WriteString(" ( ")

		if err := b.subquery.Build(); err != nil {
			return err
		}

		b.buf.WriteString(b.subquery.String())
		b.buf.WriteString(" ) ")
	}

	if len(b.keys) > 0 {
		b.buf.WriteString(" USE ")

		if b.primary {
			b.buf.WriteString("PRIMARY ")
		}

		b.buf.WriteString("KEYS ")

		if len(b.keys) == 1 {
			b.buf.WriteString(fmt.Sprintf(`"%s"`, escapeIdentifiers(b.keys[0])))
		} else {
			b.buf.WriteString("[ ")
			for i, key := range b.keys {
				if i > 0 {
					b.buf.WriteString(", ")
				}
				b.buf.WriteString(fmt.Sprintf(`"%s"`, escapeIdentifiers(key)))
			}
			b.buf.WriteString(" ]")
		}
	}

	if len(b.indexRefs) > 0 {
		b.buf.WriteString(" USE INDEX (")

		for i, indexRef := range b.indexRefs {
			if i > 0 {
				b.buf.WriteString(", ")
			}

			b.buf.WriteString(escapeIdentifiers(indexRef.name))

			if indexRef.using != nil {
				b.buf.WriteString(fmt.Sprintf(" USING %s", *indexRef.using))
			}
		}

		b.buf.WriteString(")")
	}

	if len(b.let) > 0 {
		b.buf.WriteString(" LET ")
		for i, let := range b.let {
			if i > 0 {
				b.buf.WriteString(", ")
			}
			if err := let.Build(b.buf); err != nil {
				return err
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
			if err := group.Build(b.buf); err != nil {
				return err
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

		//todo union, intersect, except
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
