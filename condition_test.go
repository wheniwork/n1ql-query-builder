package nqb_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

func TestCondition(t *testing.T) {
	for _, test := range []struct {
		cond  BuildFunc
		query string
	}{
		{
			cond:  Eq("col", "1"),
			query: "`col` = $1",
		},
		{
			cond:  Neq("col", "1"),
			query: "`col` != $1",
		},
		{
			cond:  Gt("col", "1"),
			query: "`col` > $1",
		},
		{
			cond:  Gte("col", "1"),
			query: "`col` >= $1",
		},
		{
			cond:  Lt("col", "1"),
			query: "`col` < $1",
		},
		{
			cond:  Lte("col", "1"),
			query: "`col` <= $1",
		},
		{
			cond:  And(Lt("a", "1"), Or(Gt("b", "2"), Neq("c", "3"))),
			query: "(`a` < $1) AND ((`b` > $2) OR (`c` != $3))",
		},
	} {
		buf := &bytes.Buffer{}
		err := test.cond.Build(buf)
		assert.NoError(t, err)
		assert.Equal(t, test.query, buf.String())
	}
}
