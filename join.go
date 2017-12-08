package nqb

import (
	"bytes"
	"fmt"
)

type join struct {
	joinType joinType
	fromPath string
	alias    string
	onKeys   *OnKeysClause
	onKeyFor *onKeyForClause
}

type joinType string

const Left joinType = " LEFT "
const LeftOuter joinType = " LEFT OUTER "

type OnKeysClause struct {
	primary    bool
	expression string
}

func OnKeys(primary bool, expression string) OnKeysClause {
	return OnKeysClause{primary, expression}
}

type onKeyForClause struct {
	primary    bool
	rhsExpr    string
	lhsExprKey string
	forLhsExpr string
}

func OnKeysFor(primary bool, rhsExpression, lhsExpressionKey, forLhsExpression string) *onKeyForClause {
	return &onKeyForClause{primary, rhsExpression, lhsExpressionKey, forLhsExpression}
}

func (j *join) Build(buf *bytes.Buffer) {
	if len(j.joinType) > 0 {
		buf.WriteString(fmt.Sprintf(" %s", j.joinType))
	}

	buf.WriteString(fmt.Sprintf(" JOIN %s ", escapeIdentifiers(j.fromPath)))

	if len(j.alias) > 0 {
		buf.WriteString(fmt.Sprintf("AS %s ", escapeIdentifiers(j.alias)))
	}

	buf.WriteString("ON ")

	if j.onKeys != nil {
		if j.onKeys.primary {
			buf.WriteString("PRIMARY ")
		}

		buf.WriteString(fmt.Sprintf("KEYS %s", escapeIdentifiers(j.onKeys.expression)))
	}

	if j.onKeyFor != nil {
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
}
