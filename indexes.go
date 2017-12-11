package nqb

type IndexType string

const (
	GSI  IndexType = "GSI"
	View IndexType = "VIEW"
)

type indexReference struct {
	representation string
}

func (i *indexReference) String() string {
	return i.representation
}

func IndexRef(indexName string) *indexReference {
	return IndexRefType(indexName, "")
}

func IndexRefType(indexName string, indexType IndexType) *indexReference {
	if len(indexType) == 0 {
		return &indexReference{"`" + indexName + "`"}
	}

	return &indexReference{"`" + indexName + "` USING " + string(indexType)}
}
