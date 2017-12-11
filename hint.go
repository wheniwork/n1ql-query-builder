package nqb

import "bytes"

type HintPath interface {
	KeysPath

	UseIndexRef(indexes ...*indexReference) KeysPath
	UseIndex(indexes ...string) KeysPath
}

type defaultHintPath struct {
	*defaultKeysPath
}

func newDefaultHintPath(parent Path) *defaultHintPath {
	return &defaultHintPath{newDefaultKeysPath(parent)}
}

func (p *defaultHintPath) UseIndexRef(indexes ...*indexReference) KeysPath {
	p.setElement(&hintIndexElement{indexes})
	return newDefaultKeysPath(p)
}

func (p *defaultHintPath) UseIndex(indexes ...string) KeysPath {
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
