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

func newDefaultUseIndexClause(parent Statement) *defaultUseIndexClause {
	return &defaultUseIndexClause{newDefaultKeysClauses(parent)}
}

func (c *defaultUseIndexClause) UseIndexRef(indexes ...*indexReference) KeysClauses {
	c.setElement(&useIndexElement{indexes})
	return newDefaultKeysClauses(c)
}

func (c *defaultUseIndexClause) UseIndex(indexes ...string) KeysClauses {
	var indexRefs []*indexReference
	for _, index := range indexes {
		indexRef := IndexRef(index)
		indexRefs = append(indexRefs, indexRef)
	}
	return c.UseIndexRef(indexRefs...)
}

type useIndexElement struct {
	indexReferences []*indexReference
}

func (e *useIndexElement) export() string {
	if e.indexReferences == nil || len(e.indexReferences) < 1 {
		return ""
	}

	buf := bytes.NewBufferString("USE INDEX (")

	for i, indexReference := range e.indexReferences {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(indexReference.String())
	}

	buf.WriteString(")")

	return buf.String()
}
