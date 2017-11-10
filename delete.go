package nqb

// deleteStmt builds `DELETE ...`
type deleteStmt struct {
	buf      *buffer
	keyspace string

	where []Builder
}

// Build builds `DELETE ...` in dialect
func (b *deleteStmt) Build() error {
	if b.keyspace == "" {
		return ErrKeyspaceNotSpecified
	}

	b.buf.WriteString("DELETE FROM ")
	b.buf.WriteString(EscapeIdentifier(b.keyspace))

	if len(b.where) > 0 {
		b.buf.WriteString(" WHERE ")
		err := And(b.where...).Build(b.buf)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteFrom creates a deleteStmt
func DeleteFrom(keyspace string) *deleteStmt {
	return &deleteStmt{
		buf:      &buffer{},
		keyspace: keyspace,
	}
}

// Where adds a where condition
func (b *deleteStmt) Where(query interface{}, value ...interface{}) *deleteStmt {
	switch query := query.(type) {
	case string:
		b.where = append(b.where, Expr(query, value...))
	case Builder:
		b.where = append(b.where, query)
	}
	return b
}
