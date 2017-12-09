package nqb

type indexType string

// View is used to specify a VIEW index type
const View indexType = "VIEW"

// GSI is used to specify a GSI index type
const GSI indexType = "GSI"

type indexRef struct {
	name  string
	using indexType
}

// IndexRef creates an index reference (hint) for the `USE INDEX` clause
// https://developer.couchbase.com/documentation/server/5.0/n1ql/n1ql-language-reference/hints.html
func IndexRef(name string, using indexType) *indexRef {
	return &indexRef{name, using}
}
