package nqb

import (
	"bytes"
)

type nest struct {
	joinType joinType
	fromPath string
	alias    *string
	onKeys   OnKeysClause
}

// Build builds a NEST clause
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/from.html#story-h2-6
func (n *nest) build(buf *bytes.Buffer) {
	if len(n.joinType) > 0 {
		buf.WriteString(" ")
		buf.WriteString(string(n.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("NEST ")
	buf.WriteString(escapeIdentifiers(n.fromPath))
	buf.WriteString(" ")

	if n.alias != nil {
		buf.WriteString("AS ")
		buf.WriteString(escapeIdentifiers(*n.alias))
		buf.WriteString(" ")
	}

	buf.WriteString("ON ")

	if n.onKeys.primary {
		buf.WriteString("PRIMARY ")
	}

	buf.WriteString("KEYS ")
	buf.WriteString(escapeIdentifiers(n.onKeys.expression))
}

type unnest struct {
	joinType   joinType
	flatten    bool
	expression string
	alias      *string
}

// Build builds an UNNEST clause
// https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/from.html#story-h2-5
func (u *unnest) build(buf *bytes.Buffer) {
	if len(u.joinType) > 0 {
		buf.WriteString(" ")
		buf.WriteString(string(u.joinType))
		buf.WriteString(" ")
	}

	if u.flatten {
		buf.WriteString("FLATTEN ")
	} else {
		buf.WriteString("UNNEST ")
	}

	buf.WriteString(escapeIdentifiers(u.expression))

	if u.alias != nil {
		buf.WriteString("AS ")
		buf.WriteString(escapeIdentifiers(*u.alias))
		buf.WriteString(" ")
	}
}
