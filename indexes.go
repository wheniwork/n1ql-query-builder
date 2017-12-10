package nqb

type IndexType string

const (
	GSI  IndexType = "GSI"
	View IndexType = "View"
)

type indexReference struct {
	indexReference string
}

func newIndexReference(representation string) *indexReference {
	return &indexReference{representation}
}

func (i *indexReference) String() string {
	return i.indexReference
}

func IndexRef(indexName string) *indexReference {
	return IndexRefType(indexName, "")
}

func IndexRefType(indexName string, indexType IndexType) *indexReference {
	if len(indexType) > 0 {
		return newIndexReference("`" + indexName + "`")
	}

	return newIndexReference("`" + indexName + "` USING " + string(indexType))
}
