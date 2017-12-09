package nqb_test

import (
	"bytes"
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
	builder := Select(ResultPath("foo", "")).
		Distinct().
		From("keyspace", nil, "")

	err := builder.Build()

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, "SELECT DISTINCT `foo` FROM `keyspace`", query)
}

func TestSelectStatement_Raw(t *testing.T) {
	builder := Select(ResultPath("name", "abv")).
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

func TestSelectStatement_UseKeys(t *testing.T) {
	builder := Select().
		From("beer-sample", nil, "").
		UseKeys(true, "12345")

	err := builder.Build()

	expected := "SELECT * FROM `beer-sample` USE PRIMARY KEYS \"12345\""

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_UseKeys_Multiple(t *testing.T) {
	builder := Select().
		From("beer-sample", nil, "").
		UseKeys(true, "12345", "67890")

	err := builder.Build()

	expected := "SELECT * FROM `beer-sample` USE PRIMARY KEYS [ \"12345\", \"67890\" ]"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
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
	builder := Select(ResultPath("baz.*", "bar")).
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
	builder := Select(ResultPath("baz.*", "bar")).
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

func TestSelectStatement_Nest(t *testing.T) {
	builder := Select(ResultPath("baz.*", "bar")).
		From("foo", nil, "baz").
		Nest(Inner, "foo", "bar", OnKeys(true, "baz.fooId")).
		Where(Eq("foo.type", "1")).
		Where(Eq("baz.type", "2")).
		Where(Eq("baz.fooId", "3"))

	err := builder.Build()

	expected := "SELECT `baz`.`*` AS `bar` FROM `foo` AS `baz` INNER NEST `foo` AS `bar` ON PRIMARY KEYS `baz`.`fooId` WHERE (`foo`.`type` = $1) AND (`baz`.`type` = $2) AND (`baz`.`fooId` = $3)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Unnest(t *testing.T) {
	builder := Select(ResultPath("c.name", ""), ResultPath("a.*", "")).
		From("customer", nil, "c").
		Unnest(Inner, "c.address", "a")

	err := builder.Build()

	expected := "SELECT `c`.`name`, `a`.`*` FROM `customer` AS `c` INNER UNNEST `c`.`address` AS `a`"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_UseIndex(t *testing.T) {
	builder := Select(ResultPath("name", "abv")).
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

func TestSelectStatement_Let(t *testing.T) {
	builder := Select().Let("foo", "bar").Let("name", "abv")

	err := builder.Build()

	expected := "SELECT * LET (foo = bar), (name = abv)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_GroupBy(t *testing.T) {
	builder := Select(ResultPath("relation", ""), ResultExpr("COUNT(*)", "count")).
		From("tutorial", nil, "").GroupBy("relation")

	err := builder.Build()

	expected := "SELECT `relation`, COUNT(*) AS `count` FROM `tutorial` GROUP BY `relation`"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Letting(t *testing.T) {
	builder := Select(ResultPath("relation", ""), ResultExpr("COUNT(*)", "count")).
		From("tutorial", nil, "").GroupBy("relation").Letting("foo", "bar")

	err := builder.Build()

	expected := "SELECT `relation`, COUNT(*) AS `count` FROM `tutorial` GROUP BY `relation` LETTING (foo = bar)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_Having(t *testing.T) {
	builder := Select(ResultPath("relation", ""), ResultExpr("COUNT(*)", "count")).
		From("tutorial", nil, "").GroupBy("relation").Having(func(buf *bytes.Buffer) error {
		buf.WriteString("COUNT(*) > 1")
		return nil
	})

	err := builder.Build()

	expected := "SELECT `relation`, COUNT(*) AS `count` FROM `tutorial` GROUP BY `relation` HAVING (COUNT(*) > 1)"

	assert.NoError(t, err)

	query := builder.String()
	t.Log(query)

	assert.Equal(t, expected, query)
}

func TestSelectStatement_OrderAsc(t *testing.T) {
	builder := Select(ResultPath("name", "abv")).
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
	builder := Select(ResultPath("name", "abv")).
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
	builder := Select(ResultPath("name", "abv")).
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
	builder := Select(ResultPath("name", "abv")).
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
