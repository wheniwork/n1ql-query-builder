package nqb

import (
	"bytes"
	"fmt"
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

const Left joinType = " LEFT "
const LeftOuter joinType = " LEFT OUTER "

type OnKeysClause struct {
	primary    bool
	expression string
}

// OnKeys creates a lookup join predicate
func OnKeys(primary bool, expression string) OnKeysClause {
	return OnKeysClause{primary, expression}
}

type OnKeyForClause struct {
	primary    bool
	rhsExpr    string
	lhsExprKey string
	forLhsExpr string
}

// OnKeysFor creates an index join predicate
func OnKeysFor(primary bool, rhsExpression, lhsExpressionKey, forLhsExpression string) OnKeyForClause {
	return OnKeyForClause{primary, rhsExpression, lhsExpressionKey, forLhsExpression}
}

func (j *join) startClause(buf *bytes.Buffer) {
	if len(j.joinType) > 0 {
		buf.WriteString(fmt.Sprintf(" %s", j.joinType))
	}

	buf.WriteString(fmt.Sprintf(" JOIN %s ", escapeIdentifiers(j.fromPath)))

	if len(j.alias) > 0 {
		buf.WriteString(fmt.Sprintf("AS %s ", escapeIdentifiers(j.alias)))
	}

	buf.WriteString("ON ")
}

func (j *lookupJoin) build(buf *bytes.Buffer) {
	j.startClause(buf)

	if j.onKeys.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString(fmt.Sprintf("KEYS %s", escapeIdentifiers(j.onKeys.expression)))
}

func (j *indexJoin) build(buf *bytes.Buffer) {
	j.startClause(buf)

	if j.onKeyFor.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString(
		fmt.Sprintf(
			"KEY %s.%s FOR %s",
			escapeIdentifiers(j.onKeyFor.rhsExpr),
			escapeIdentifiers(j.onKeyFor.lhsExprKey),
			escapeIdentifiers(j.onKeyFor.forLhsExpr),
		),
	)
}
