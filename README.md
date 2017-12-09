# N1QL Query Builder
**Fluent Couchbase N1QL Query Builder for Go**

[![Build Status](https://travis-ci.org/wheniwork/n1ql-query-builder.svg?branch=master)](https://travis-ci.org/wheniwork/n1ql-query-builder)
[![GoDoc](https://godoc.org/github.com/wheniwork/n1ql-query-builder?status.svg)](https://godoc.org/github.com/wheniwork/n1ql-query-builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/wheniwork/n1ql-query-builder)](https://goreportcard.com/report/github.com/wheniwork/n1ql-query-builder)
[![Coverage Status](https://coveralls.io/repos/github/wheniwork/n1ql-query-builder/badge.svg?branch=master)](https://coveralls.io/github/wheniwork/n1ql-query-builder?branch=master)
[![codecov](https://codecov.io/gh/wheniwork/n1ql-query-builder/branch/master/graph/badge.svg)](https://codecov.io/gh/wheniwork/n1ql-query-builder)

_The API is currently experimental and may change._

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
	qb := nqb.Select(nqb.ResultPath("baz.*", "bar")).
    		From("foo", nil, "baz").
    		LookupJoin(nqb.Inner, "foo", "bar", nqb.OnKeys(false, "baz.fooId")).
    		Where(nqb.Eq("foo.type", "1")).
    		Where(nqb.Eq("baz.type", "2")).
    		Where(nqb.Eq("baz.fooId", "3"))
    
    if err := qb.Build(); err != nil {
    	panic(err)
    }
    
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
	qb := Select(ResultPath("baz.*", "bar")).
        From("foo", nil, "baz").
        LookupJoin(Inner, "foo", "bar", OnKeys(false, "baz.fooId")).
        Where(Eq("foo.type", "1")).
        Where(Eq("baz.type", "2")).
        Where(Eq("baz.fooId", "3"))
    
    if err := qb.Build(); err != nil {
    	panic(err)
    }
    
    query := gocb.NewN1qlQuery(qb.String())
    
    // execute query, specifying parameters accordingly...
}
```

## TODO

* DML statements
* Integration tests (?)
