package nqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// ====================================
// General Select-From Tests (select-from-clause)
// ====================================
//

func TestGroupBy(t *testing.T) {
	path := newDefaultGroupByPath(nil).GroupBy(X("relation"))
	assert.Equal(t, "GROUP BY relation", path.String())
}

func TestGroupByWithHaving(t *testing.T) {
	path := newDefaultGroupByPath(nil).GroupBy(X("relation")).Having(X("count(*) > 1"))
	assert.Equal(t, "GROUP BY relation HAVING count(*) > 1", path.String())
}

func TestGroupByWithLetting(t *testing.T) {
	path := newDefaultGroupByPath(nil).GroupBy(X("relation")).Letting(NewAliasExpr("foo", X("bar")))
	assert.Equal(t, "GROUP BY relation LETTING foo = bar", path.String())
}

func TestGroupByWithLettingAndHaving(t *testing.T) {
	path := newDefaultGroupByPath(nil).
		GroupBy(X("relation")).
		Letting(NewAliasExpr("foo", X("bar")), NewAliasExpr("hello", S("world"))).
		Having(X("count(*) > 1"))
	assert.Equal(t, "GROUP BY relation LETTING foo = bar, hello = \"world\" HAVING count(*) > 1", path.String())
}

func TestWhere(t *testing.T) {
	path := newDefaultWherePath(nil).Where(X("age").Gt(X("20")))
	assert.Equal(t, "WHERE age > 20", path.String())

	path = newDefaultWherePath(nil).Where("age > 20")
	assert.Equal(t, "WHERE age > 20", path.String())
}

func TestWhereWithGroupBy(t *testing.T) {
	path := newDefaultWherePath(nil).Where(X("age > 20")).GroupBy(X("age"))
	assert.Equal(t, "WHERE age > 20 GROUP BY age", path.String())
}

func TestWhereWithGroupByAndHaving(t *testing.T) {
	path := newDefaultWherePath(nil).Where(X("age > 20")).GroupBy(X("age")).Having(X("count(*) > 10"))
	assert.Equal(t, "WHERE age > 20 GROUP BY age HAVING count(*) > 10", path.String())
}

func TestLet(t *testing.T) {
	path := newDefaultLetPath(nil).Let(NewAliasExpr("count", X("COUNT(*)")))
	assert.Equal(t, "LET count = COUNT(*)", path.String())

	path = newDefaultLetPath(nil).Let(NewAliasExpr("a", X("x > 5")), NewAliasExpr("b", S("foobar")))
	assert.Equal(t, "LET a = x > 5, b = \"foobar\"", path.String())
}

func TestLetWithWhere(t *testing.T) {
	path := newDefaultLetPath(nil).
		Let(NewAliasExpr("a", X("x > 5")), NewAliasExpr("b", S("foobar"))).
		Where(X("foo").Eq(S("bar")))
	assert.Equal(t, "LET a = x > 5, b = \"foobar\" WHERE foo = \"bar\"", path.String())
}

func TestJoins(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	pathString := newDefaultLetPath(nil).Join(sToken).String()
	pathExpression := newDefaultLetPath(nil).Join(eToken).String()
	assert.Equal(t, "JOIN a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftJoin(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftJoin(eToken).String()
	assert.Equal(t, "LEFT JOIN a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).InnerJoin(sToken).String()
	pathExpression = newDefaultLetPath(nil).InnerJoin(eToken).String()
	assert.Equal(t, "INNER JOIN a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftOuterJoin(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftOuterJoin(eToken).String()
	assert.Equal(t, "LEFT OUTER JOIN a", pathString)
	assert.Equal(t, pathString, pathExpression)
}

func TestNests(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	pathString := newDefaultLetPath(nil).Nest(sToken).String()
	pathExpression := newDefaultLetPath(nil).Nest(eToken).String()
	assert.Equal(t, "NEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftNest(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftNest(eToken).String()
	assert.Equal(t, "LEFT NEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).InnerNest(sToken).String()
	pathExpression = newDefaultLetPath(nil).InnerNest(eToken).String()
	assert.Equal(t, "INNER NEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftOuterNest(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftOuterNest(eToken).String()
	assert.Equal(t, "LEFT OUTER NEST a", pathString)
	assert.Equal(t, pathString, pathExpression)
}

func TestUnNests(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	pathString := newDefaultLetPath(nil).Unnest(sToken).String()
	pathExpression := newDefaultLetPath(nil).Unnest(eToken).String()
	assert.Equal(t, "UNNEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftUnnest(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftUnnest(eToken).String()
	assert.Equal(t, "LEFT UNNEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).InnerUnnest(sToken).String()
	pathExpression = newDefaultLetPath(nil).InnerUnnest(eToken).String()
	assert.Equal(t, "INNER UNNEST a", pathString)
	assert.Equal(t, pathString, pathExpression)

	pathString = newDefaultLetPath(nil).LeftOuterUnnest(sToken).String()
	pathExpression = newDefaultLetPath(nil).LeftOuterUnnest(eToken).String()
	assert.Equal(t, "LEFT OUTER UNNEST a", pathString)
	assert.Equal(t, pathString, pathExpression)
}

//
// ====================================
// General Select Tests (select-clause)
// ====================================
//

func TestSelect(t *testing.T) {
	statement := newDefaultSelectPath(nil).Select(X("firstname"), X("lastname"))
	assert.Equal(t, "SELECT firstname, lastname", statement.String())

	statement = newDefaultSelectPath(nil).SelectAll(X("firstname"))
	assert.Equal(t, "SELECT ALL firstname", statement.String())
}

func TestSelectWithUnion(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		Union().
		Select(X("a"), X("b"))
	assert.Equal(t, "SELECT firstname, lastname UNION SELECT a, b", statement.String())
}

func TestSelectWithUnionAll(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		UnionAll().
		Select(X("a"), X("b"))

	assert.Equal(t, "SELECT firstname, lastname UNION ALL SELECT a, b", statement.String())
}

func TestSelectWithIntersect(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo"))).
		Intersect().
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"INTERSECT " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement.String())
}

func TestSelectWithIntersectAll(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo"))).
		IntersectAll().
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"INTERSECT ALL " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement.String())
}

func TestSelectWithExcept(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo"))).
		Except().
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"EXCEPT " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement.String())
}

func TestSelectWithExceptAll(t *testing.T) {
	statement := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo"))).
		ExceptAll().
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"EXCEPT ALL " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement.String())
}

func TestSelectChainedWithUnion(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"UNION " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.UnionPath(statement2).String())
}

func TestSelectChainedWithUnionAll(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"UNION ALL " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.UnionAllPath(statement2).String())
}

func TestSelectChainedWithIntersect(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"INTERSECT " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.IntersectPath(statement2).String())
}

func TestSelectChainedWithIntersectAll(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"INTERSECT ALL " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.IntersectAllPath(statement2).String())
}

func TestSelectChainedWithExcept(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"EXCEPT " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.ExceptPath(statement2).String())
}

func TestSelectChainedWithExceptAll(t *testing.T) {
	statement1 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectPath(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("bar")))

	expected := "SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"foo\" " +
		"EXCEPT ALL " +
		"SELECT firstname, lastname " +
		"FROM foo " +
		"WHERE lastname = \"bar\""

	assert.Equal(t, expected, statement1.ExceptAllPath(statement2).String())
}

func TestOrderBy(t *testing.T) {
	statement := newDefaultOrderByPath(nil).OrderBy(Asc("firstname"))
	assert.Equal(t, "ORDER BY firstname ASC", statement.String())

	statement = newDefaultOrderByPath(nil).OrderBy(Asc("firstname"), Desc("lastname"))
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC", statement.String())

	statement = newDefaultOrderByPath(nil).OrderBy(DefaultSort("firstname"), DefaultSort(X("lastname")))
	assert.Equal(t, "ORDER BY firstname, lastname", statement.String())
}

func TestOrderByWithLimit(t *testing.T) {
	statement := newDefaultOrderByPath(nil).OrderBy(Asc("firstname")).Limit(5)
	assert.Equal(t, "ORDER BY firstname ASC LIMIT 5", statement.String())
}

func TestOrderByWithLimitAndOffset(t *testing.T) {
	statement := newDefaultOrderByPath(nil).
		OrderBy(Asc("firstname"), Desc("lastname")).
		Limit(5).
		Offset(10)
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC LIMIT 5 OFFSET 10", statement.String())
}

func TestOrderByWithOffset(t *testing.T) {
	statement := newDefaultOrderByPath(nil).
		OrderBy(Asc("firstname"), Desc("lastname")).
		Offset(3)
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC OFFSET 3", statement.String())
}

func TestOffset(t *testing.T) {
	statement := newDefaultOffsetPath(nil).Offset(3)
	assert.Equal(t, "OFFSET 3", statement.String())
}

func TestLimitWithOffset(t *testing.T) {
	statement := newDefaultLimitPath(nil).Limit(4).Offset(3)
	assert.Equal(t, "LIMIT 4 OFFSET 3", statement.String())
}

func TestHintIndexPathSingle(t *testing.T) {
	hint1 := newDefaultHintPath(nil).UseIndexRef(IndexRef("test"))
	hint2 := newDefaultHintPath(nil).UseIndex("test")

	assert.Equal(t, "USE INDEX (`test`)", hint1.String())
	assert.Equal(t, hint1.String(), hint2.String())

	typedHint1 := newDefaultHintPath(nil).UseIndexRef(IndexRefType("test", GSI))
	typedHint2 := newDefaultHintPath(nil).UseIndexRef(IndexRefType("test", View))

	assert.Equal(t, "USE INDEX (`test` USING GSI)", typedHint1.String())
	assert.Equal(t, "USE INDEX (`test` USING VIEW)", typedHint2.String())
}

func TestHintIndexPathMultiple(t *testing.T) {
	hint1 := newDefaultHintPath(nil).UseIndexRef(IndexRef("test"), IndexRef("test2"))
	hint2 := newDefaultHintPath(nil).UseIndex("test", "test2")

	assert.Equal(t, "USE INDEX (`test`,`test2`)", hint1.String())
	assert.Equal(t, hint1.String(), hint2.String())

	typedHint1 := newDefaultHintPath(nil).UseIndexRef(
		IndexRefType("test", GSI),
		IndexRefType("test", View))

	assert.Equal(t, "USE INDEX (`test` USING GSI,`test` USING VIEW)", typedHint1.String())
}

//
// ====================================
// From Tests (from-clause)
// ====================================
//

func TestSimpleFrom(t *testing.T) {
	statement := newDefaultFromPath(nil).From("default")
	assert.Equal(t, "FROM default", statement.String())

	statement2 := newDefaultFromPath(nil).From("beer-sample").As("b")
	assert.Equal(t, "FROM beer-sample AS b", statement2.String())
}

func TestFromWithKeys(t *testing.T) {
	statement := newDefaultFromPath(nil).
		From("beer-sample").
		As("b").
		UseKeys("a.id")
	assert.Equal(t, "FROM beer-sample AS b USE KEYS a.id", statement.String())

	statement = newDefaultFromPath(nil).
		From("beer-sample").
		As("b").
		UseKeysValues("my-brewery")
	assert.Equal(t, "FROM beer-sample AS b USE KEYS \"my-brewery\"", statement.String())

	//statement = newDefaultFromPath(nil).
	//	From("beer-sample").
	//	UseKeys(JsonArray.From("key1", "key2")) //fixme
	//assert.Equal(t, "FROM beer-sample USE KEYS [\"key1\",\"key2\"]", statement.String())

	statement = newDefaultFromPath(nil).
		From("beer-sample").
		UseKeysValues("key1", "key2")
	assert.Equal(t, "FROM beer-sample USE KEYS [\"key1\",\"key2\"]", statement.String())
}

func TestUnNest(t *testing.T) {
	statement := newDefaultFromPath(nil).
		From("tutorial").As("contact").
		Unnest("contact.children").
		Where(X("contact.fname").Eq(S("Dave")))
	assert.Equal(t, "FROM tutorial AS contact UNNEST contact.children WHERE contact.fname = \"Dave\"",
		statement.String())

	statement = newDefaultFromPath(nil).
		From("default").
		LeftOuterUnnest("foo.bar").
		LeftUnnest("bar.baz").
		InnerUnnest("x.y")
	assert.Equal(t, "FROM default LEFT OUTER UNNEST foo.bar LEFT UNNEST bar.baz INNER UNNEST x.y", statement.String())
}

func TestNest(t *testing.T) {
	statement := newDefaultFromPath(nil).
		From("users_with_orders").As("user").
		Nest("orders_with_users").As("orders")
	assert.Equal(t, "FROM users_with_orders AS user NEST orders_with_users AS orders", statement.String())

	statement = newDefaultFromPath(nil).
		From("default").
		LeftOuterNest("foo.bar").
		LeftNest("bar.baz").
		InnerNest("x.y")
	assert.Equal(t, "FROM default LEFT OUTER NEST foo.bar LEFT NEST bar.baz INNER NEST x.y", statement.String())
}

//func TestNestWithKeys(t *testing.T) {
//	statement := newDefaultFromPath(nil).
//		From("users_with_orders").As("user").
//		Nest("orders_with_users").As("orders").
//		OnKeys(X(JsonArray.From("key1", "key2"))) //fixme
//	assert.Equal(t, "FROM users_with_orders AS user NEST orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
//		statement.String())
//}

func TestJoin(t *testing.T) {
	statement := newDefaultFromPath(nil).
		From("users_with_orders").As("user").
		Join("orders_with_users").As("orders")
	assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders", statement.String())

	statement = newDefaultFromPath(nil).
		From("default").
		LeftOuterJoin("foo.bar").
		LeftJoin("bar.baz").
		InnerJoin("x.y")
	assert.Equal(t, "FROM default LEFT OUTER JOIN foo.bar LEFT JOIN bar.baz INNER JOIN x.y", statement.String())
}

func TestJoinWithKeys(t *testing.T) {
	//statement := newDefaultFromPath(nil).
	//	From("users_with_orders").As("user").
	//	Join("orders_with_users").As("orders").
	//	OnKeys(X(JsonArray.From("key1", "key2"))) //fixme
	//assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
	//	statement.String())

	//statement = newDefaultFromPath(nil).
	//	From("users_with_orders").As("user").
	//	Join("orders_with_users").As("orders").
	//	OnKeys(JsonArray.From("key1", "key2")) //fixme
	//assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
	//	statement.String())

	statement := newDefaultFromPath(nil).
		From("users_with_orders").As("user").
		Join("orders_with_users").As("orders").
		OnKeys("orders.id")
	assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS orders.id",
		statement.String())
}

func TestJoinWithEscapedNamespace(t *testing.T) {
	statement := newDefaultFromPath(nil).From("a").
		Join(I("beer-sample")).As("b").
		OnKeys(P("a", "foreignKey"))

	assert.Equal(t, "FROM a JOIN `beer-sample` AS b ON KEYS a.foreignKey", statement.String())
}
