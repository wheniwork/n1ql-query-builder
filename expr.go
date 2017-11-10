package nqb

type raw struct {
	Query string
	Value []interface{}
}

// Expr should be used when N1QL syntax is not supported
func Expr(query string, value ...interface{}) Builder {
	return &raw{Query: query, Value: value}
}

func (raw *raw) Build(buf *buffer) error {
	buf.WriteString(raw.Query)
	buf.WriteValue(raw.Value...)
	return nil
}
