# N1QL Query Builder
**Fluent Couchbase N1QL Query Builder for Go**

[![Build Status](https://travis-ci.org/wheniwork/n1ql-query-builder.svg?branch=master)](https://travis-ci.org/wheniwork/n1ql-query-builder)
[![GoDoc](https://godoc.org/github.com/wheniwork/n1ql-query-builder?status.svg)](https://godoc.org/github.com/wheniwork/n1ql-query-builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/wheniwork/n1ql-query-builder)](https://goreportcard.com/report/github.com/wheniwork/n1ql-query-builder)
[![Coverage Status](https://coveralls.io/repos/github/wheniwork/n1ql-query-builder/badge.svg?branch=master)](https://coveralls.io/github/wheniwork/n1ql-query-builder?branch=master)
[![codecov](https://codecov.io/gh/wheniwork/n1ql-query-builder/branch/master/graph/badge.svg)](https://codecov.io/gh/wheniwork/n1ql-query-builder)

_The API, which is based on the query DSL from the [Couchbase Java SDK](), is currently experimental and may change._

## Implemented Statements

### [`SELECT`](https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/select-syntax.html)

Example usage:
```go
package main

import (
	"github.com/wheniwork/n1ql-query-builder"
	"gopkg.in/couchbase/gocb.v1"
)

func main() {
	// fixme
    query := gocb.NewN1qlQuery(qb.String())
    
    // execute query, specifying parameters accordingly...
}
```

Example using ["dot" import declaration](https://golang.org/ref/spec#Import_declarations):
```go
package main

import (
	. "github.com/wheniwork/n1ql-query-builder"
	"gopkg.in/couchbase/gocb.v1"
)

func main() {
	//fixme
    
    query := gocb.NewN1qlQuery(qb.String())
    
    // execute query, specifying parameters accordingly...
}
```

## TODO

* DML statements
* Integration tests (?)
