package nqb

type indexType string

const View indexType = "VIEW"
const GSI indexType = "GSI"

type indexRef struct {
	name  string
	using *indexType
}

func IndexRef(name string, using *indexType) *indexRef {
	return &indexRef{name, using}
}
