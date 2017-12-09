package nqb

import (
	"bytes"
)

type join struct {
	joinType joinType
	fromPath string
	alias    string
}

type lookupJoin struct {
	join
	onKeys OnKeysClause
}

type indexJoin struct {
	join
	onKeyFor OnKeyForClause
}

type joinType string

const Inner joinType = "INNER"
const Left joinType = "LEFT"
const LeftOuter joinType = "LEFT OUTER"

// OnKeysClause represents an ON KEYS clause used in lookup joins
type OnKeysClause struct {
	primary    bool
	expression string
}

// OnKeys creates a lookup join predicate
func OnKeys(primary bool, expression string) OnKeysClause {
	return OnKeysClause{primary, expression}
}

// OnKeyForClause represents an ON KEY FOR clause used in index joins
type OnKeyForClause struct {
	primary    bool
	rhsExpr    string //
	lhsExprKey string // attribute in rhs-expression referencing primary key for lhs-expression
	forLhsExpr string // keyspace or expression corresponding to the left hand side of JOIN
}

// OnKeyFor creates an index join predicate
//
// rhsExpression: keyspace or expression corresponding to the right hand side of JOIN
func OnKeyFor(primary bool, rhsExpression, lhsExpressionKey, forLhsExpression string) OnKeyForClause {
	return OnKeyForClause{primary, rhsExpression, lhsExpressionKey, forLhsExpression}
}

func (j *join) startClause(buf *bytes.Buffer) {
	if len(j.joinType) > 0 {
		buf.WriteString(" ")
		buf.WriteString(string(j.joinType))
	}

	buf.WriteString(" JOIN ")
	buf.WriteString(escapeIdentifiers(j.fromPath))
	buf.WriteString(" ")

	if len(j.alias) > 0 {
		buf.WriteString("AS ")
		buf.WriteString(escapeIdentifiers(j.alias))
		buf.WriteString(" ")
	}

	buf.WriteString("ON ")
}

func (j *lookupJoin) build(buf *bytes.Buffer) {
	j.startClause(buf)

	if j.onKeys.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString("KEYS ")
	buf.WriteString(escapeIdentifiers(j.onKeys.expression))
}

func (j *indexJoin) build(buf *bytes.Buffer) {
	j.startClause(buf)

	if j.onKeyFor.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString("KEY ")
	buf.WriteString(escapeIdentifiers(j.onKeyFor.rhsExpr))
	buf.WriteString(".")
	buf.WriteString(escapeIdentifiers(j.onKeyFor.lhsExprKey))
	buf.WriteString(" FOR ")
	buf.WriteString(escapeIdentifiers(j.onKeyFor.forLhsExpr))
}
