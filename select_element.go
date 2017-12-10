package nqb

import "bytes"

type selectType string

const (
	defaultSelect selectType = ""
	all           selectType = "ALL"
	distinct      selectType = "DISTINCT"
	raw           selectType = "RAW"
)

type selectElement struct {
	selectType  selectType
	expressions []*Expression
}

func (s *selectElement) export() string {
	buf := bytes.NewBufferString("SELECT ")
	if s.selectType != defaultSelect {
		buf.WriteString(string(s.selectType))
		buf.WriteString(" ")
	}

	for i, expression := range s.expressions {
		buf.WriteString(expression.String())

		// todo improve?
		if i < len(s.expressions)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
