package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapeIdentifiers(t *testing.T) {
	assert.Equal(t, "`foo`.`bar`.`baz`", escapeIdentifiers("foo.bar.baz"))
}
