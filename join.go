package nqb

import "fmt"

type join struct {
	joinType *joinType
	fromPath string
	alias    *string
	onKeys   *onKeysClause
	onKeyFor *onKeyForClause
}

type joinType string

const Left joinType = " LEFT "
const LeftOuter joinType = " LEFT OUTER "

type onKeysClause struct {
	primary    bool
	expression string
}

func OnKeys(primary bool, expression string) onKeysClause {
	return onKeysClause{primary, expression}
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

func (j *join) Build(buf *buffer) {
	if j.joinType != nil {
		buf.WriteString(fmt.Sprintf(" %s ", j.joinType))
	}

	buf.WriteString(fmt.Sprintf("JOIN %s ", j.fromPath))

	if j.alias != nil {
		buf.WriteString(fmt.Sprintf("AS %s ", j.alias))
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
