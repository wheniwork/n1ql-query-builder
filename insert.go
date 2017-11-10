package nqb

import (
	"bytes"
	"reflect"
)

// insertStmt builds `INSERT INTO ...`
type insertStmt struct {
	buf          *buffer
	keyspace     string
	columns      []string
	values       [][]interface{}
	returnColumn []string
}

// Build builds `INSERT INTO ...` in dialect
func (b *insertStmt) Build() error {
	if b.keyspace == "" {
		return ErrKeyspaceNotSpecified
	}

	if len(b.columns) == 0 {
		return ErrColumnNotSpecified
	}

	b.buf.WriteString("INSERT INTO ")
	b.buf.WriteString(EscapeIdentifier(b.keyspace))

	placeholderBuf := new(bytes.Buffer)
	placeholderBuf.WriteString("(")
	b.buf.WriteString(" (")
	for i, col := range b.columns {
		if i > 0 {
			b.buf.WriteString(",")
			placeholderBuf.WriteString(",")
		}
		b.buf.WriteString(EscapeIdentifier(col))
		placeholderBuf.WriteString(placeholder)
	}
	b.buf.WriteString(") VALUES ")
	placeholderBuf.WriteString(")")
	placeholderStr := placeholderBuf.String()

	for i, tuple := range b.values {
		if i > 0 {
			b.buf.WriteString(", ")
		}
		b.buf.WriteString(placeholderStr)

		b.buf.WriteValue(tuple...)
	}

	if len(b.returnColumn) > 0 {
		b.buf.WriteString(" RETURNING ")
		for i, col := range b.returnColumn {
			if i > 0 {
				b.buf.WriteString(",")
			}
			b.buf.WriteString(EscapeIdentifier(col))
		}
	}

	return nil
}

// InsertInto creates an insertStmt
func InsertInto(keyspace string) *insertStmt {
	return &insertStmt{
		buf:      &buffer{},
		keyspace: keyspace,
	}
}

// Columns adds columns
func (b *insertStmt) Columns(column ...string) *insertStmt {
	b.columns = column
	return b
}

// Values adds a tuple for columns
func (b *insertStmt) Values(value ...interface{}) *insertStmt {
	b.values = append(b.values, value)
	return b
}

// Record adds a tuple for columns from a struct
func (b *insertStmt) Record(structValue interface{}) *insertStmt {
	v := reflect.Indirect(reflect.ValueOf(structValue))

	if v.Kind() == reflect.Struct {
		var value []interface{}
		m := structMap(v)
		for _, key := range b.columns {
			if val, ok := m[key]; ok {
				value = append(value, val.Interface())
			} else {
				value = append(value, nil)
			}
		}
		b.Values(value...)
	}
	return b
}

func (b *insertStmt) Returning(column ...string) *insertStmt {
	b.returnColumn = column
	return b
}
