package nqb

const (
	placeholder = "?"
)

// identifier is a type of string
type I string

func (i I) Build(buf *buffer) error {
	buf.WriteString(EscapeIdentifier(string(i)))
	return nil
}

// As creates an alias for expr. e.g. SELECT `a1` AS `a2`
func (i I) As(alias string) Builder {
	return as(i, alias)
}

func as(expr interface{}, alias string) Builder {
	return BuildFunc(func(buf *buffer) error {
		buf.WriteString(placeholder)
		buf.WriteValue(expr)
		buf.WriteString(" AS ")
		buf.WriteString(EscapeIdentifier(alias))
		return nil
	})
}
