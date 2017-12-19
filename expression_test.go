package nqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

func TestP(t *testing.T) {
	path := Path("foo", "bar", "baz")
	str := path.String()

	t.Log(str)
	assert.Equal(t, "foo.bar.baz", str)
}
