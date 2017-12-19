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
	clause := newDefaultGroupByClause(nil).GroupBy(X("relation"))
	assert.Equal(t, "GROUP BY relation", clause.String())
}

func TestGroupByWithHaving(t *testing.T) {
	clause := newDefaultGroupByClause(nil).GroupBy(X("relation")).Having(X("count(*) > 1"))
	assert.Equal(t, "GROUP BY relation HAVING count(*) > 1", clause.String())
}

func TestGroupByWithLetting(t *testing.T) {
	clause := newDefaultGroupByClause(nil).GroupBy(X("relation")).Letting(NewAliasExpr("foo", X("bar")))
	assert.Equal(t, "GROUP BY relation LETTING foo = bar", clause.String())
}

func TestGroupByWithLettingAndHaving(t *testing.T) {
	clause := newDefaultGroupByClause(nil).
		GroupBy(X("relation")).
		Letting(NewAliasExpr("foo", X("bar")), NewAliasExpr("hello", S("world"))).
		Having(X("count(*) > 1"))
	assert.Equal(t, "GROUP BY relation LETTING foo = bar, hello = \"world\" HAVING count(*) > 1", clause.String())
}

func TestWhere(t *testing.T) {
	clause := newDefaultWhereClause(nil).Where(X("age").Gt(X("20")))
	assert.Equal(t, "WHERE age > 20", clause.String())

	clause = newDefaultWhereClause(nil).Where("age > 20")
	assert.Equal(t, "WHERE age > 20", clause.String())
}

func TestWhereWithGroupBy(t *testing.T) {
	clause := newDefaultWhereClause(nil).Where(X("age > 20")).GroupBy(X("age"))
	assert.Equal(t, "WHERE age > 20 GROUP BY age", clause.String())
}

func TestWhereWithGroupByAndHaving(t *testing.T) {
	clause := newDefaultWhereClause(nil).Where(X("age > 20")).GroupBy(X("age")).Having(X("count(*) > 10"))
	assert.Equal(t, "WHERE age > 20 GROUP BY age HAVING count(*) > 10", clause.String())
}

func TestLet(t *testing.T) {
	clause := newDefaultLetClause(nil).Let(NewAliasExpr("count", X("COUNT(*)")))
	assert.Equal(t, "LET count = COUNT(*)", clause.String())

	clause = newDefaultLetClause(nil).Let(NewAliasExpr("a", X("x > 5")), NewAliasExpr("b", S("foobar")))
	assert.Equal(t, "LET a = x > 5, b = \"foobar\"", clause.String())
}

func TestLetWithWhere(t *testing.T) {
	clause := newDefaultLetClause(nil).
		Let(NewAliasExpr("a", X("x > 5")), NewAliasExpr("b", S("foobar"))).
		Where(X("foo").Eq(S("bar")))
	assert.Equal(t, "LET a = x > 5, b = \"foobar\" WHERE foo = \"bar\"", clause.String())
}

func TestJoins(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	clauseString := newDefaultLetClause(nil).Join(sToken).String()
	clauseExpression := newDefaultLetClause(nil).Join(eToken).String()
	assert.Equal(t, "JOIN a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftJoin(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftJoin(eToken).String()
	assert.Equal(t, "LEFT JOIN a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).InnerJoin(sToken).String()
	clauseExpression = newDefaultLetClause(nil).InnerJoin(eToken).String()
	assert.Equal(t, "INNER JOIN a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftOuterJoin(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftOuterJoin(eToken).String()
	assert.Equal(t, "LEFT OUTER JOIN a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)
}

func TestNests(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	clauseString := newDefaultLetClause(nil).Nest(sToken).String()
	clauseExpression := newDefaultLetClause(nil).Nest(eToken).String()
	assert.Equal(t, "NEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftNest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftNest(eToken).String()
	assert.Equal(t, "LEFT NEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).InnerNest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).InnerNest(eToken).String()
	assert.Equal(t, "INNER NEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftOuterNest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftOuterNest(eToken).String()
	assert.Equal(t, "LEFT OUTER NEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)
}

func TestUnNests(t *testing.T) {
	eToken := X("a")
	sToken := "a"

	clauseString := newDefaultLetClause(nil).Unnest(sToken).String()
	clauseExpression := newDefaultLetClause(nil).Unnest(eToken).String()
	assert.Equal(t, "UNNEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftUnnest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftUnnest(eToken).String()
	assert.Equal(t, "LEFT UNNEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).InnerUnnest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).InnerUnnest(eToken).String()
	assert.Equal(t, "INNER UNNEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)

	clauseString = newDefaultLetClause(nil).LeftOuterUnnest(sToken).String()
	clauseExpression = newDefaultLetClause(nil).LeftOuterUnnest(eToken).String()
	assert.Equal(t, "LEFT OUTER UNNEST a", clauseString)
	assert.Equal(t, clauseString, clauseExpression)
}

//
// ====================================
// General Select Tests (select-clause)
// ====================================
//

func TestSelect(t *testing.T) {
	statement := newDefaultSelectClause(nil).Select(X("firstname"), X("lastname"))
	assert.Equal(t, "SELECT firstname, lastname", statement.String())

	statement = newDefaultSelectClause(nil).SelectAll(X("firstname"))
	assert.Equal(t, "SELECT ALL firstname", statement.String())
}

func TestSelectWithUnion(t *testing.T) {
	statement := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		Union().
		Select(X("a"), X("b"))
	assert.Equal(t, "SELECT firstname, lastname UNION SELECT a, b", statement.String())
}

func TestSelectWithUnionAll(t *testing.T) {
	statement := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		UnionAll().
		Select(X("a"), X("b"))

	assert.Equal(t, "SELECT firstname, lastname UNION ALL SELECT a, b", statement.String())
}

func TestSelectWithIntersect(t *testing.T) {
	statement := newDefaultSelectClause(nil).
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
	statement := newDefaultSelectClause(nil).
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
	statement := newDefaultSelectClause(nil).
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
	statement := newDefaultSelectClause(nil).
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
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.UnionClause(statement2).String())
}

func TestSelectChainedWithUnionAll(t *testing.T) {
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.UnionAllClause(statement2).String())
}

func TestSelectChainedWithIntersect(t *testing.T) {
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.IntersectClause(statement2).String())
}

func TestSelectChainedWithIntersectAll(t *testing.T) {
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.IntersectAllClause(statement2).String())
}

func TestSelectChainedWithExcept(t *testing.T) {
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.ExceptClause(statement2).String())
}

func TestSelectChainedWithExceptAll(t *testing.T) {
	statement1 := newDefaultSelectClause(nil).
		Select(X("firstname"), X("lastname")).
		From("foo").
		Where(X("lastname").Eq(S("foo")))

	statement2 := newDefaultSelectClause(nil).
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

	assert.Equal(t, expected, statement1.ExceptAllClause(statement2).String())
}

func TestOrderBy(t *testing.T) {
	statement := newDefaultOrderByClause(nil).OrderBy(Asc("firstname"))
	assert.Equal(t, "ORDER BY firstname ASC", statement.String())

	statement = newDefaultOrderByClause(nil).OrderBy(Asc("firstname"), Desc("lastname"))
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC", statement.String())

	statement = newDefaultOrderByClause(nil).OrderBy(DefaultSort("firstname"), DefaultSort(X("lastname")))
	assert.Equal(t, "ORDER BY firstname, lastname", statement.String())
}

func TestOrderByWithLimit(t *testing.T) {
	statement := newDefaultOrderByClause(nil).OrderBy(Asc("firstname")).Limit(5)
	assert.Equal(t, "ORDER BY firstname ASC LIMIT 5", statement.String())
}

func TestOrderByWithLimitAndOffset(t *testing.T) {
	statement := newDefaultOrderByClause(nil).
		OrderBy(Asc("firstname"), Desc("lastname")).
		Limit(5).
		Offset(10)
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC LIMIT 5 OFFSET 10", statement.String())
}

func TestOrderByWithOffset(t *testing.T) {
	statement := newDefaultOrderByClause(nil).
		OrderBy(Asc("firstname"), Desc("lastname")).
		Offset(3)
	assert.Equal(t, "ORDER BY firstname ASC, lastname DESC OFFSET 3", statement.String())
}

func TestOffset(t *testing.T) {
	statement := newDefaultOffsetClause(nil).Offset(3)
	assert.Equal(t, "OFFSET 3", statement.String())
}

func TestLimitWithOffset(t *testing.T) {
	statement := newDefaultLimitClause(nil).Limit(4).Offset(3)
	assert.Equal(t, "LIMIT 4 OFFSET 3", statement.String())
}

func TestUseIndexClauseSingle(t *testing.T) {
	hint1 := newDefaultHintClause(nil).UseIndexRef(IndexRef("test"))
	hint2 := newDefaultHintClause(nil).UseIndex("test")

	assert.Equal(t, "USE INDEX (`test`)", hint1.String())
	assert.Equal(t, hint1.String(), hint2.String())

	typedHint1 := newDefaultHintClause(nil).UseIndexRef(IndexRefType("test", GSI))
	typedHint2 := newDefaultHintClause(nil).UseIndexRef(IndexRefType("test", View))

	assert.Equal(t, "USE INDEX (`test` USING GSI)", typedHint1.String())
	assert.Equal(t, "USE INDEX (`test` USING VIEW)", typedHint2.String())
}

func TestUseIndexClauseMultiple(t *testing.T) {
	hint1 := newDefaultHintClause(nil).UseIndexRef(IndexRef("test"), IndexRef("test2"))
	hint2 := newDefaultHintClause(nil).UseIndex("test", "test2")

	assert.Equal(t, "USE INDEX (`test`,`test2`)", hint1.String())
	assert.Equal(t, hint1.String(), hint2.String())

	typedHint1 := newDefaultHintClause(nil).UseIndexRef(
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
	statement := newDefaultFromClause(nil).From("default")
	assert.Equal(t, "FROM default", statement.String())

	statement2 := newDefaultFromClause(nil).From("beer-sample").As("b")
	assert.Equal(t, "FROM beer-sample AS b", statement2.String())
}

func TestFromWithKeys(t *testing.T) {
	statement := newDefaultFromClause(nil).
		From("beer-sample").
		As("b").
		UseKeys("a.id")
	assert.Equal(t, "FROM beer-sample AS b USE KEYS a.id", statement.String())

	statement = newDefaultFromClause(nil).
		From("beer-sample").
		As("b").
		UseKeysValues("my-brewery")
	assert.Equal(t, "FROM beer-sample AS b USE KEYS \"my-brewery\"", statement.String())

	//statement = newDefaultFromClause(nil).
	//	From("beer-sample").
	//	UseKeys(JsonArray.From("key1", "key2")) //fixme
	//assert.Equal(t, "FROM beer-sample USE KEYS [\"key1\",\"key2\"]", statement.String())

	statement = newDefaultFromClause(nil).
		From("beer-sample").
		UseKeysValues("key1", "key2")
	assert.Equal(t, "FROM beer-sample USE KEYS [\"key1\",\"key2\"]", statement.String())
}

func TestUnNest(t *testing.T) {
	statement := newDefaultFromClause(nil).
		From("tutorial").As("contact").
		Unnest("contact.children").
		Where(X("contact.fname").Eq(S("Dave")))
	assert.Equal(t, "FROM tutorial AS contact UNNEST contact.children WHERE contact.fname = \"Dave\"",
		statement.String())

	statement = newDefaultFromClause(nil).
		From("default").
		LeftOuterUnnest("foo.bar").
		LeftUnnest("bar.baz").
		InnerUnnest("x.y")
	assert.Equal(t, "FROM default LEFT OUTER UNNEST foo.bar LEFT UNNEST bar.baz INNER UNNEST x.y", statement.String())
}

func TestNest(t *testing.T) {
	statement := newDefaultFromClause(nil).
		From("users_with_orders").As("user").
		Nest("orders_with_users").As("orders")
	assert.Equal(t, "FROM users_with_orders AS user NEST orders_with_users AS orders", statement.String())

	statement = newDefaultFromClause(nil).
		From("default").
		LeftOuterNest("foo.bar").
		LeftNest("bar.baz").
		InnerNest("x.y")
	assert.Equal(t, "FROM default LEFT OUTER NEST foo.bar LEFT NEST bar.baz INNER NEST x.y", statement.String())
}

//func TestNestWithKeys(t *testing.T) {
//	statement := newDefaultFromClause(nil).
//		From("users_with_orders").As("user").
//		Nest("orders_with_users").As("orders").
//		OnKeys(X(JsonArray.From("key1", "key2"))) //fixme
//	assert.Equal(t, "FROM users_with_orders AS user NEST orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
//		statement.String())
//}

func TestJoin(t *testing.T) {
	statement := newDefaultFromClause(nil).
		From("users_with_orders").As("user").
		Join("orders_with_users").As("orders")
	assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders", statement.String())

	statement = newDefaultFromClause(nil).
		From("default").
		LeftOuterJoin("foo.bar").
		LeftJoin("bar.baz").
		InnerJoin("x.y")
	assert.Equal(t, "FROM default LEFT OUTER JOIN foo.bar LEFT JOIN bar.baz INNER JOIN x.y", statement.String())
}

func TestJoinWithKeys(t *testing.T) {
	//statement := newDefaultFromClause(nil).
	//	From("users_with_orders").As("user").
	//	Join("orders_with_users").As("orders").
	//	OnKeys(X(JsonArray.From("key1", "key2"))) //fixme
	//assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
	//	statement.String())

	//statement = newDefaultFromClause(nil).
	//	From("users_with_orders").As("user").
	//	Join("orders_with_users").As("orders").
	//	OnKeys(JsonArray.From("key1", "key2")) //fixme
	//assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS [\"key1\",\"key2\"]",
	//	statement.String())

	statement := newDefaultFromClause(nil).
		From("users_with_orders").As("user").
		Join("orders_with_users").As("orders").
		OnKeys("orders.id")
	assert.Equal(t, "FROM users_with_orders AS user JOIN orders_with_users AS orders ON KEYS orders.id",
		statement.String())
}

func TestJoinWithEscapedNamespace(t *testing.T) {
	statement := newDefaultFromClause(nil).From("a").
		Join(I("beer-sample")).As("b").
		OnKeys(Path("a", "foreignKey"))

	assert.Equal(t, "FROM a JOIN `beer-sample` AS b ON KEYS a.foreignKey", statement.String())
}
