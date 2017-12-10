package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapWith(t *testing.T) {
	str := wrapWith('`', "foo")

	t.Log(str)
	assert.Equal(t, "`foo`", str)
}

func TestWrapWith_MultipleArgs(t *testing.T) {
	str := wrapWith('`', "foo", "bar", "baz")

	t.Log(str)
	assert.Equal(t, "`foo`, `bar`, `baz`", str)
}
