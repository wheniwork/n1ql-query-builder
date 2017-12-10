package nqb

import "bytes"

type selectType string

const (
	DefaultSelect selectType = ""
	All           selectType = "ALL"
	Distinct      selectType = "DISTINCT"
	Raw           selectType = "RAW"
)

type selectElement struct {
	selectType  selectType
	expressions []*Expression
}

func newSelectElement(selectType selectType, expressions ...*Expression) *selectElement {
	return &selectElement{selectType, expressions}
}

func (s *selectElement) Export() string {
	buf := bytes.NewBufferString("SELECT ")
	if s.selectType != DefaultSelect {
		buf.WriteString(string(s.selectType))
		buf.WriteString(" ")
	}

	for i, expression := range s.expressions {
		buf.WriteString(expression.String())
		if i < len(s.expressions)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
