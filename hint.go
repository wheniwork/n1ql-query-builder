package nqb

import "bytes"

type UseIndexClause interface {
	KeysClauses

	UseIndexRef(indexes ...*indexReference) KeysClauses
	UseIndex(indexes ...string) KeysClauses
}

type defaultUseIndexClause struct {
	*defaultKeysClauses
}

func newDefaultHintClause(parent Statement) *defaultUseIndexClause {
	return &defaultUseIndexClause{newDefaultKeysClauses(parent)}
}

func (p *defaultUseIndexClause) UseIndexRef(indexes ...*indexReference) KeysClauses {
	p.setElement(&hintIndexElement{indexes})
	return newDefaultKeysClauses(p)
}

func (p *defaultUseIndexClause) UseIndex(indexes ...string) KeysClauses {
	var indexRefs []*indexReference
	for _, index := range indexes {
		indexRef := IndexRef(index)
		indexRefs = append(indexRefs, indexRef)
	}
	return p.UseIndexRef(indexRefs...)
}

type hintIndexElement struct {
	indexReferences []*indexReference
}

func (e *hintIndexElement) export() string {
	if e.indexReferences == nil || len(e.indexReferences) < 1 {
		return ""
	}

	n1ql := bytes.NewBufferString("USE INDEX (")

	for i, indexReference := range e.indexReferences {
		if i > 0 {
			n1ql.WriteString(",")
		}
		n1ql.WriteString(indexReference.String())
	}

	n1ql.WriteString(")")

	return n1ql.String()
}
