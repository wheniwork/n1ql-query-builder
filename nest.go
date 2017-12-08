package nqb

import (
	"bytes"
	"fmt"
)

type nest struct {
	joinType *joinType
	fromPath string
	alias    *string
	onKeys   OnKeysClause
}

func (n *nest) Build(buf *bytes.Buffer) {
	if n.joinType != nil {
		buf.WriteString(fmt.Sprintf(" %s ", *n.joinType))
	}

	buf.WriteString(fmt.Sprintf("NEST %s ", escapeIdentifiers(n.fromPath)))

	if n.alias != nil {
		buf.WriteString(fmt.Sprintf("AS %s ", escapeIdentifiers(*n.alias)))
	}

	buf.WriteString("ON ")

	if n.onKeys.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString(fmt.Sprintf("KEYS %s", escapeIdentifiers(n.onKeys.expression)))
}

type unnest struct {
	joinType   *joinType
	flatten    bool
	expression string
	alias      *string
}

func (u *unnest) Build(buf *bytes.Buffer) {
	if u.joinType != nil {
		buf.WriteString(fmt.Sprintf(" %s ", *u.joinType))
	}

	if u.flatten {
		buf.WriteString("FLATTEN ")
	} else {
		buf.WriteString("UNNEST ")
	}

	buf.WriteString(escapeIdentifiers(u.expression))

	if u.alias != nil {
		buf.WriteString(fmt.Sprintf("AS %s ", escapeIdentifiers(*u.alias)))
	}
}
