package nqb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftOuterJoin(t *testing.T) {
	buf := &bytes.Buffer{}
	j := join{LeftOuter, "foo", "bar"}
	j.startClause(buf)

	startClause := buf.String()

	t.Log(startClause)
	assert.Equal(t, " LEFT OUTER JOIN `foo` AS `bar` ON ", startClause)
}
