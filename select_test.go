package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectStmt(t *testing.T) {
	builder := Select("a", "b").
		From(Select("a").From("keyspace")).
		LeftJoin("keyspace2", "keyspace.a1 = keyspace.a2").
		Distinct().
		Where(Eq("c", 1)).
		GroupBy("d").
		Having(Eq("e", 2)).
		OrderAsc("f").
		Limit(3).
		Offset(4)
	err := builder.Build()
	assert.NoError(t, err)
	assert.Equal(t, "SELECT DISTINCT `a`, `b` FROM ? LEFT JOIN `keyspace2` ON keyspace.a1 = keyspace.a2 WHERE (`c` = ?) GROUP BY d HAVING (`e` = ?) ORDER BY f ASC LIMIT 3 OFFSET 4", builder.buf.String())
	// two functions cannot be compared
	assert.Equal(t, 3, len(builder.buf.String()))
}

func BenchmarkSelectSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Select("a", "b").From("keyspace").Where(Eq("c", 1)).OrderAsc("d").Build()
	}
}
