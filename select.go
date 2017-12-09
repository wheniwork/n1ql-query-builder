package nqb

import (
	"bytes"
	"fmt"
)

type selectStatement struct {
	buf *bytes.Buffer

	distinct bool
	raw      bool

	resultExpressions []*resultExpression

	keyspace string
	subquery *selectStatement
	alias    string

	keys    []string
	primary bool

	joins   []join
	nests   []nest
	unnests []unnest

	indexRefs []indexRef

	let     []BuildFunc
	where   []BuildFunc
	groupBy []BuildFunc
	letting []BuildFunc
	having  []BuildFunc
	orderBy []BuildFunc

	limit  int64
	offset int64
}

type resultExpression struct {
	pathOrExpression string
	alias            string
}

// ResultExpr creates a result expression comprised of a N1QL path or expression, and an optional alias
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/index.html#n1ql-lang-ref__N1QL_Expressions
func ResultExpr(pathOrExpression string, alias string) *resultExpression {
	return &resultExpression{
		pathOrExpression: pathOrExpression,
		alias:            alias,
	}
}

// Select initializes a `SELECT` statement
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/select-syntax.html
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/selectclause.html
func Select(resultExpression ...*resultExpression) *selectStatement {
	return &selectStatement{
		buf:               &bytes.Buffer{},
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
func (b *selectStatement) Raw() *selectStatement {
	b.raw = true
	return b
}

// From optionally specifies a keyspace or subquery, with or without an alias
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/from.html
func (b *selectStatement) From(keyspace string, subquery *selectStatement, alias string) *selectStatement {
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

// LookupJoin specifies a lookup join
func (b *selectStatement) LookupJoin(joinType joinType, fromPath string, alias string, onKeys OnKeysClause) *selectStatement {
	b.joins = append(b.joins, join{joinType, fromPath, alias, &onKeys, nil})
	return b
}

// IndexJoin specifies an index join
func (b *selectStatement) IndexJoin(
	joinType joinType, fromPath string, alias string, onKeys *OnKeysClause, onKeyFor *onKeyForClause,
) *selectStatement {
	b.joins = append(b.joins, join{joinType, fromPath, alias, onKeys, onKeyFor})
	return b
}

// Nest specifies a `NEST`
func (b *selectStatement) Nest(joinType joinType, fromPath string, alias *string, onKeys OnKeysClause) *selectStatement {
	b.nests = append(b.nests, nest{&joinType, fromPath, alias, onKeys})
	return b
}

// Unnest specifies an `UNNEST`
func (b *selectStatement) Unnest(joinType joinType, flatten bool, expression string, alias *string) *selectStatement {
	b.unnests = append(b.unnests, unnest{&joinType, flatten, expression, alias})
	return b
}

// UseIndex specifies index references (hints) for the `USE INDEX` clause
// https://developer.couchbase.com/documentation/server/5.0/n1ql/n1ql-language-reference/hints.html
func (b *selectStatement) UseIndex(indexRef ...indexRef) *selectStatement {
	b.indexRefs = indexRef
	return b
}

func let(alias, expression string) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		buf.WriteString(fmt.Sprintf(" (%s = %s) ", alias, expression))
		return nil
	})
}

// Let adds a `LET` clause
// https://developer.couchbase.com/documentation/server/5.0/n1ql/n1ql-language-reference/let.html
func (b *selectStatement) Let(alias, expression string) *selectStatement {
	b.let = append(b.let, let(alias, expression))
	return b
}

// Where adds a `WHERE` condition
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/where.html
func (b *selectStatement) Where(condition BuildFunc) *selectStatement {
	b.where = append(b.where, condition)
	return b
}

// GroupBy specifies columns for `GROUP BY`
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/groupby.html
func (b *selectStatement) GroupBy(col ...string) *selectStatement {
	for _, group := range col {
		b.groupBy = append(b.groupBy, func(buf *bytes.Buffer) error {
			buf.WriteString(group)
			return nil
		})
	}
	return b
}

// Letting adds a `LETTING` clause
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/groupby.html
func (b *selectStatement) Letting(alias, expression string) *selectStatement {
	b.letting = append(b.letting, let(alias, expression))
	return b
}

// Having specifies a `HAVING` condition
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/groupby.html
func (b *selectStatement) Having(condition BuildFunc) *selectStatement {
	b.having = append(b.having, condition)
	return b
}

// OrderAsc specifies an ascending `ORDER BY`
func (b *selectStatement) OrderAsc(col string) *selectStatement {
	b.orderBy = append(b.orderBy, order(col, asc))
	return b
}

// OrderDesc specifies a descending `ORDER BY`
func (b *selectStatement) OrderDesc(col string) *selectStatement {
	b.orderBy = append(b.orderBy, order(col, desc))
	return b
}

// Limit specifies a `LIMIT`
func (b *selectStatement) Limit(n uint64) *selectStatement {
	b.limit = int64(n)
	return b
}

// Offset specifies an `OFFSET`
func (b *selectStatement) Offset(n uint64) *selectStatement {
	b.offset = int64(n)
	return b
}

// Build builds a `SELECT` statement
func (b *selectStatement) Build() error {
	b.buf.WriteString("SELECT ")

	if b.distinct {
		b.buf.WriteString("DISTINCT ")
	}

	if b.raw {
		b.buf.WriteString(" RAW ")
	}

	if len(b.resultExpressions) == 0 {
		b.buf.WriteString("*")
	} else {
		for i, resultExpression := range b.resultExpressions {
			if i > 0 {
				b.buf.WriteString(", ")
			}

			b.buf.WriteString(escapeIdentifiers(resultExpression.pathOrExpression))

			if len(resultExpression.alias) > 0 {
				b.buf.WriteString(" AS ")
				b.buf.WriteString(escapeIdentifiers(resultExpression.alias))
			}
		}
	}

	if len(b.keyspace) > 0 {
		b.buf.WriteString(" FROM ")
		b.buf.WriteString(escapeIdentifiers(b.keyspace))

		if len(b.alias) > 0 {
			b.buf.WriteString(" AS ")
			b.buf.WriteString(escapeIdentifiers(b.alias))
		}
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

	if b.subquery != nil {
		b.buf.WriteString(" ( ")

		if err := b.subquery.Build(); err != nil {
			return err
		}

		b.buf.WriteString(b.subquery.String())
		b.buf.WriteString(" ) ")
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

	b.buildIndexRefs()

	if err := b.buildLet(); err != nil {
		return err
	}

	if err := b.buildWhere(); err != nil {
		return err
	}

	if err := b.buildGroupBy(); err != nil {
		return err
	}

	if err := b.buildOrderBy(); err != nil {
		return err
	}

	b.buildLimit()
	b.buildOffset()

	return nil
}

func (b *selectStatement) String() string {
	return b.buf.String()
}

func (b *selectStatement) buildIndexRefs() {
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
}

func (b *selectStatement) buildLet() error {
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
	return nil
}

func (b *selectStatement) buildWhere() error {
	if len(b.where) > 0 {
		b.buf.WriteString(" WHERE ")
		err := And(b.where...).Build(b.buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *selectStatement) buildGroupBy() error {
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
			for i, letting := range b.letting {
				if i > 0 {
					b.buf.WriteString(", ")
				}
				if err := letting.Build(b.buf); err != nil {
					return err
				}
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
	return nil
}

func (b *selectStatement) buildOrderBy() error {
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
	return nil
}

func (b *selectStatement) buildLimit() {
	if b.limit >= 0 {
		b.buf.WriteString(" LIMIT ")
		b.buf.WriteString(fmt.Sprint(b.limit))
	}
}

func (b *selectStatement) buildOffset() {
	if b.offset >= 0 {
		b.buf.WriteString(" OFFSET ")
		b.buf.WriteString(fmt.Sprint(b.offset))
	}
}
