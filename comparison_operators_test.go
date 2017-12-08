package nqb_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

func TestEq(t *testing.T) {
	buf := bytes.Buffer{}
	Eq("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` = $1", buf.String())
}

func TestNeq(t *testing.T) {
	buf := bytes.Buffer{}
	Neq("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` != $1", buf.String())
}

func TestGt(t *testing.T) {
	buf := bytes.Buffer{}
	Gt("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` > $1", buf.String())
}

func TestGte(t *testing.T) {
	buf := bytes.Buffer{}
	Gte("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` >= $1", buf.String())
}

func TestLt(t *testing.T) {
	buf := bytes.Buffer{}
	Lt("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` < $1", buf.String())
}

func TestLte(t *testing.T) {
	buf := bytes.Buffer{}
	Lte("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` <= $1", buf.String())
}

func TestBetween(t *testing.T) {
	buf := bytes.Buffer{}
	Between("foo", "1", "2").Build(&buf)
	assert.Equal(t, "`foo` BETWEEN $1 AND $2", buf.String())
}

func TestNotBetween(t *testing.T) {
	buf := bytes.Buffer{}
	NotBetween("foo", "1", "2").Build(&buf)
	assert.Equal(t, "`foo` NOT BETWEEN $1 AND $2", buf.String())
}

func TestLike(t *testing.T) {
	buf := bytes.Buffer{}
	Like("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` LIKE $1", buf.String())
}

func TestNotLike(t *testing.T) {
	buf := bytes.Buffer{}
	NotLike("foo", "1").Build(&buf)
	assert.Equal(t, "`foo` NOT LIKE $1", buf.String())
}

func TestIsNull(t *testing.T) {
	buf := bytes.Buffer{}
	IsNull("foo").Build(&buf)
	assert.Equal(t, "`foo` IS NULL", buf.String())
}

func TestIsNotNull(t *testing.T) {
	buf := bytes.Buffer{}
	IsNotNull("foo").Build(&buf)
	assert.Equal(t, "`foo` IS NOT NULL", buf.String())
}

func TestIsMissing(t *testing.T) {
	buf := bytes.Buffer{}
	IsMissing("foo").Build(&buf)
	assert.Equal(t, "`foo` IS MISSING", buf.String())
}

func TestIsNotMissing(t *testing.T) {
	buf := bytes.Buffer{}
	IsNotMissing("foo").Build(&buf)
	assert.Equal(t, "`foo` IS NOT MISSING", buf.String())
}

func TestIsValued(t *testing.T) {
	buf := bytes.Buffer{}
	IsValued("foo").Build(&buf)
	assert.Equal(t, "`foo` IS VALUED", buf.String())
}

func TestIsNotValued(t *testing.T) {
	buf := bytes.Buffer{}
	IsNotValued("foo").Build(&buf)
	assert.Equal(t, "`foo` IS NOT VALUED", buf.String())
}
