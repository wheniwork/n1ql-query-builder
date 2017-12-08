package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectStatement_From(t *testing.T) {
	keyspace := "keyspace"
	builder := Select().
		From(&keyspace, nil, nil)

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT * FROM `keyspace`", builder.String())
}

func TestSelectStatement_Distinct(t *testing.T) {
	keyspace := "keyspace"
	builder := Select(ResultExpr("foo", nil)).
		Distinct().
		From(&keyspace, nil, nil)

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT DISTINCT `foo` FROM `keyspace`", builder.String())
}

func TestSelectStatement_Where(t *testing.T) {
	keyspace := "keyspace"
	builder := Select().
		From(&keyspace, nil, nil).
		Where(Eq("col", "1")).Where(Neq("loc", "loc"))

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT * FROM `keyspace` WHERE (`col` = $1) AND (`loc` != $loc)", builder.String())
}

//
//func BenchmarkSelectSQL(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Select("a", "b").From("keyspace").Where(Eq("c", 1)).OrderAsc("d").Build()
//	}
//}
