# N1QL Query Builder
**Fluent Couchbase N1QL Query Builder for Go**

[![Build Status](https://travis-ci.org/wheniwork/n1ql-query-builder.svg?branch=master)](https://travis-ci.org/wheniwork/n1ql-query-builder)
[![GoDoc](https://godoc.org/github.com/wheniwork/n1ql-query-builder?status.svg)](https://godoc.org/github.com/wheniwork/n1ql-query-builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/wheniwork/n1ql-query-builder)](https://goreportcard.com/report/github.com/wheniwork/n1ql-query-builder)
[![Coverage Status](https://coveralls.io/repos/github/wheniwork/n1ql-query-builder/badge.svg?branch=master)](https://coveralls.io/github/wheniwork/n1ql-query-builder?branch=master)
[![codecov](https://codecov.io/gh/wheniwork/n1ql-query-builder/branch/master/graph/badge.svg)](https://codecov.io/gh/wheniwork/n1ql-query-builder)

_The API, which is based on the query DSL from the [Couchbase Java SDK](https://developer.couchbase.com/documentation/server/current/sdk/java/n1ql-queries-with-sdk.html), is currently experimental and may change._

## Statements

### [`SELECT`](https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/select-syntax.html)

See the [godoc](https://godoc.org/github.com/wheniwork/n1ql-query-builder) or [the tests](select_test.go) for examples.

## Expressions

http://godoc.org/github.com/wheniwork/n1ql-query-builder/#Expression

## TODO

* More functions
    * [x] Aggregate
    * [ ] Array
    * [ ] Case
    * [ ] Collections
    * [ ] Comparison
    * [ ] Conditional
    * [ ] Date
    * [ ] JSON
    * [x] Meta
    * [x] Number
    * [ ] Object
    * [ ] Pattern Matching
    * [x] String
    * [ ] Type
* DML statements
* Integration tests (?)
* Open a pull request to make this part of the official Couchbase Go SDK
