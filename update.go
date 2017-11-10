package nqb

type updateStmt struct {
	buf      *buffer
	keyspace string
	values   map[string]interface{}

	where []Builder
}

func (b *updateStmt) Build() error {
	if b.keyspace == "" {
		return ErrKeyspaceNotSpecified
	}

	if len(b.values) == 0 {
		return ErrColumnNotSpecified
	}

	b.buf.WriteString("UPDATE ")
	b.buf.WriteString(EscapeIdentifier(b.keyspace))
	b.buf.WriteString(" SET ")

	i := 0
	for col, v := range b.values {
		if i > 0 {
			b.buf.WriteString(", ")
		}
		b.buf.WriteString(EscapeIdentifier(col))
		b.buf.WriteString(" = ")
		b.buf.WriteString(placeholder)

		b.buf.WriteValue(v)
		i++
	}

	if len(b.where) > 0 {
		b.buf.WriteString(" WHERE ")
		err := And(b.where...).Build(b.buf)
		if err != nil {
			return err
		}
	}
	return nil
}

// Update creates an updateStmt
func Update(keyspace string) *updateStmt {
	return &updateStmt{
		buf:      &buffer{},
		keyspace: keyspace,
		values:   make(map[string]interface{}),
	}
}

// Where adds a where condition
func (b *updateStmt) Where(query interface{}, value ...interface{}) *updateStmt {
	switch query := query.(type) {
	case string:
		b.where = append(b.where, Expr(query, value...))
	case Builder:
		b.where = append(b.where, query)
	}
	return b
}

// Set specifies a key-values pair
func (b *updateStmt) Set(column string, value interface{}) *updateStmt {
	b.values[column] = value
	return b
}

// SetMap specifies a list of key-values pair
func (b *updateStmt) SetMap(m map[string]interface{}) *updateStmt {
	for col, val := range m {
		b.Set(col, val)
	}
	return b
}
