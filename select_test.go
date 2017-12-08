package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectStatement_From(t *testing.T) {
	builder := Select().
		From("keyspace", nil, "")

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT * FROM `keyspace`", builder.String())
}

func TestSelectStatement_Distinct(t *testing.T) {
	builder := Select(ResultExpr("foo", "")).
		Distinct().
		From("keyspace", nil, "")

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT DISTINCT `foo` FROM `keyspace`", builder.String())
}

func TestSelectStatement_Where(t *testing.T) {
	builder := Select().
		From("keyspace", nil, "").
		Where(Eq("col", "1")).Where(Neq("loc", "loc"))

	err := builder.Build()

	assert.NoError(t, err)
	assert.Equal(t, "SELECT * FROM `keyspace` WHERE (`col` = $1) AND (`loc` != $loc)", builder.String())
}

func TestSelectStatement_LookupJoin(t *testing.T) {
	builder := Select(ResultExpr("baz.*", "bar")).
		From("foo", nil, "baz").
		LookupJoin("", "foo", "bar", OnKeys(false, "baz.fooId")).
		Where(Eq("foo.type", "1")).
		Where(Eq("baz.type", "2")).
		Where(Eq("baz.fooId", "3"))

	err := builder.Build()

	expected := "SELECT `baz`.`*` AS `bar` FROM `foo` AS `baz` JOIN `foo` AS `bar` ON KEYS `baz`.`fooId` WHERE (`foo`.`type` = $1) AND (`baz`.`type` = $2) AND (`baz`.`fooId` = $3)"

	assert.NoError(t, err)
	assert.Equal(t, expected, builder.String())
}

//
//func BenchmarkSelectSQL(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Select("a", "b").From("keyspace").Where(Eq("c", 1)).OrderAsc("d").Build()
//	}
//}
