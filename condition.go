package nqb

import "reflect"

func buildCond(buf *buffer, pred string, cond ...Builder) error {
	for i, c := range cond {
		if i > 0 {
			buf.WriteString(" ")
			buf.WriteString(pred)
			buf.WriteString(" ")
		}
		buf.WriteString("(")
		err := c.Build()
		if err != nil {
			return err
		}
		buf.WriteString(")")
	}
	return nil
}

// And creates AND from a list of conditions
func And(cond ...Builder) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCond(buf, "AND", cond...)
	})
}

// Or creates OR from a list of conditions
func Or(cond ...Builder) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCond(buf, "OR", cond...)
	})
}

func buildCmp(buf *buffer, pred string, column string, value interface{}) error {
	buf.WriteString(escapeIdentifiers(column))
	buf.WriteString(" ")
	buf.WriteString(pred)
	buf.WriteString(" ")
	buf.WriteString(placeholder)

	buf.WriteValue(value)
	return nil
}

// Eq is `=`.
// When values is nil, it will be translated to `IS NULL`.
// When values is a slice, it will be translated to `IN`.
// Otherwise it will be translated to `=`.
func Eq(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		if value == nil {
			buf.WriteString(escapeIdentifiers(column))
			buf.WriteString(" IS NULL")
			return nil
		}
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Slice {
			if v.Len() == 0 {
				buf.WriteString("false") //todo check this
				return nil
			}
			return buildCmp(buf, "IN", column, value)
		}
		return buildCmp(buf, "=", column, value)
	})
}

// Neq is `!=`.
// When values is nil, it will be translated to `IS NOT NULL`.
// When values is a slice, it will be translated to `NOT IN`.
// Otherwise it will be translated to `!=`.
func Neq(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		if value == nil {
			buf.WriteString(escapeIdentifiers(column))
			buf.WriteString(" IS NOT NULL")
			return nil
		}
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Slice {
			if v.Len() == 0 {
				buf.WriteString("true") //todo check this
				return nil
			}
			return buildCmp(buf, "NOT IN", column, value)
		}
		return buildCmp(buf, "!=", column, value)
	})
}

// Gt is `>`.
func Gt(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCmp(buf, ">", column, value)
	})
}

// Gte is '>='.
func Gte(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCmp(buf, ">=", column, value)
	})
}

// Lt is '<'.
func Lt(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCmp(buf, "<", column, value)
	})
}

// Lte is `<=`.
func Lte(column string, value interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		return buildCmp(buf, "<=", column, value)
	})
}
