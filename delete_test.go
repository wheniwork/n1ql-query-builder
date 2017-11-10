package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt(t *testing.T) {
	buf := &buffer{}
	builder := DeleteFrom("keyspace").Where(Eq("a", 1))
	err := builder.Build()
	assert.NoError(t, err)
	assert.Equal(t, "DELETE FROM `keyspace` WHERE (`a` = ?)", buf.String())
	assert.Equal(t, []interface{}{1}, buf.Value())
}

func BenchmarkDeleteSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeleteFrom("keyspace").Where(Eq("a", 1)).Build()
	}
}
