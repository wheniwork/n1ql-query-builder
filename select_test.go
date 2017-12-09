package nqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

func TestSelectStatement_From(t *testing.T) {
	builder := Select().
		From("keyspace", nil, "")

	err := builder.Build()

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, "SELECT * FROM `keyspace`", query)
}

func TestSelectStatement_Distinct(t *testing.T) {
	builder := Select(ResultExpr("foo", "")).
		Distinct().
		From("keyspace", nil, "")

	err := builder.Build()

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, "SELECT DISTINCT `foo` FROM `keyspace`", query)
}

func TestSelectStatement_Where(t *testing.T) {
	builder := Select().
		From("keyspace", nil, "").
		Where(Eq("col", "1")).Where(Neq("loc", "loc"))

	err := builder.Build()

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, "SELECT * FROM `keyspace` WHERE (`col` = $1) AND (`loc` != $loc)", query)
}

func TestSelectStatement_LookupJoin(t *testing.T) {
	builder := Select(ResultExpr("baz.*", "bar")).
		From("foo", nil, "baz").
		LookupJoin(Inner, "foo", "bar", OnKeys(true, "baz.fooId")).
		Where(Eq("foo.type", "1")).
		Where(Eq("baz.type", "2")).
		Where(Eq("baz.fooId", "3"))

	err := builder.Build()

	expected := "SELECT `baz`.`*` AS `bar` FROM `foo` AS `baz` INNER JOIN `foo` AS `bar` ON PRIMARY KEYS `baz`.`fooId` WHERE (`foo`.`type` = $1) AND (`baz`.`type` = $2) AND (`baz`.`fooId` = $3)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_IndexJoin(t *testing.T) {
	builder := Select(ResultExpr("baz.*", "bar")).
		From("foo", nil, "baz").
		IndexJoin(Left, "foo", "bar", OnKeyFor(true, "baz", "fooId", "foo")).
		Where(Eq("foo.type", "1")).
		Where(Eq("baz.type", "2")).
		Where(Eq("baz.fooId", "3"))

	err := builder.Build()

	expected := "SELECT `baz`.`*` AS `bar` FROM `foo` AS `baz` LEFT JOIN `foo` AS `bar` ON PRIMARY KEY `baz`.`fooId` FOR `foo` WHERE (`foo`.`type` = $1) AND (`baz`.`type` = $2) AND (`baz`.`fooId` = $3)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_UseIndex(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		From("beer-sample", nil, "").
		UseIndex(IndexRef("beer_abv", GSI)).
		Where(Gt("abv", "1"))

	err := builder.Build()

	expected := "SELECT `name` AS `abv` FROM `beer-sample` USE INDEX (`beer_abv` USING GSI) WHERE (`abv` > $1)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Raw(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		Raw().
		From("beer-sample", nil, "").
		Where(Gt("abv", "1"))

	err := builder.Build()

	expected := "SELECT RAW `name` AS `abv` FROM `beer-sample` WHERE (`abv` > $1)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_OrderAsc(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		From("beer-sample", nil, "").
		Where(Gt("abv", "1")).
		OrderAsc("abv")

	err := builder.Build()

	expected := "SELECT `name` AS `abv` FROM `beer-sample` WHERE (`abv` > $1) ORDER BY `abv` ASC"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_OrderDesc(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		From("beer-sample", nil, "").
		Where(Gt("abv", "1")).
		OrderDesc("abv")

	err := builder.Build()

	expected := "SELECT `name` AS `abv` FROM `beer-sample` WHERE (`abv` > $1) ORDER BY `abv` DESC"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Limit(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		From("beer-sample", nil, "").
		Where(Gt("abv", "1")).
		Limit(1)

	err := builder.Build()

	expected := "SELECT `name` AS `abv` FROM `beer-sample` WHERE (`abv` > $1) LIMIT 1"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Offset(t *testing.T) {
	builder := Select(ResultExpr("name", "abv")).
		From("beer-sample", nil, "").
		Where(Gt("abv", "1")).
		Offset(1)

	err := builder.Build()

	expected := "SELECT `name` AS `abv` FROM `beer-sample` WHERE (`abv` > $1) OFFSET 1"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}
