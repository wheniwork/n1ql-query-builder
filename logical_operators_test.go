package nqb_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

var testFunc = func(buf *bytes.Buffer) error {
	return nil
}

func TestAnd(t *testing.T) {
	buf := &bytes.Buffer{}
	err := And(testFunc, testFunc).Build(buf)

	assert.NoError(t, err)
	assert.Equal(t, "() AND ()", buf.String())
}

func TestAnd_One_Cond(t *testing.T) {
	buf := &bytes.Buffer{}
	err := And(testFunc).Build(buf)

	assert.NoError(t, err)
	assert.Equal(t, "()", buf.String())
}

func TestOr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := Or(testFunc, testFunc).Build(buf)

	assert.NoError(t, err)
	assert.Equal(t, "() OR ()", buf.String())
}

func TestOr_One_Cond(t *testing.T) {
	buf := &bytes.Buffer{}
	err := Or(testFunc).Build(buf)

	assert.NoError(t, err)
	assert.Equal(t, "()", buf.String())
}

func TestNot(t *testing.T) {
	buf := &bytes.Buffer{}
	err := Not(testFunc).Build(buf)

	assert.NoError(t, err)
	assert.Equal(t, " NOT ()", buf.String())
}
