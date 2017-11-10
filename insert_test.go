package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type insertTest struct {
	A int
	C string `db:"b"`
}

func TestInsertStmt(t *testing.T) {
	buf := &buffer{}
	builder := InsertInto("keyspace").Columns("a", "b").Values(1, "one").Record(&insertTest{
		A: 2,
		C: "two",
	})
	err := builder.Build()
	assert.NoError(t, err)
	assert.Equal(t, "INSERT INTO `keyspace` (`a`,`b`) VALUES (?,?), (?,?)", buf.String())
	assert.Equal(t, []interface{}{1, "one", 2, "two"}, buf.Value())
}

func BenchmarkInsertValuesSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInto("keyspace").Columns("a", "b").Values(1, "one").Build()
	}
}

func BenchmarkInsertRecordSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInto("keyspace").Columns("a", "b").Record(&insertTest{
			A: 2,
			C: "two",
		}).Build()
	}
}
