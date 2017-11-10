package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateStmt(t *testing.T) {
	buf := &buffer{}
	builder := Update("keyspace").Set("a", 1).Where(Eq("b", 2))
	err := builder.Build()
	assert.NoError(t, err)

	assert.Equal(t, "UPDATE `keyspace` SET `a` = ? WHERE (`b` = ?)", buf.String())
	assert.Equal(t, []interface{}{1, 2}, buf.Value())
}

func BenchmarkUpdateValuesSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Update("keyspace").Set("a", 1).Build()
	}
}

func BenchmarkUpdateMapSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Update("keyspace").SetMap(map[string]interface{}{"a": 1, "b": 2}).Build()
	}
}
