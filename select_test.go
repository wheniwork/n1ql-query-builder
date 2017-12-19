package nqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/wheniwork/n1ql-query-builder"
)

func ExampleSelect() {
	Select(S("Hello World").
		As("Greeting"))
}

func Test1(t *testing.T) {
	statement := Select(S("Hello World").
		As("Greeting"))
	assert.Equal(t, "SELECT \"Hello World\" AS Greeting", statement.String())
}

func Test2(t *testing.T) {
	statement := Select("*").
		From("tutorial").
		Where(X("fname").Eq(S("Ian")))
	assert.Equal(t, "SELECT * FROM tutorial WHERE fname = \"Ian\"", statement.String())
}

func Test3(t *testing.T) {
	statement := Select(X("children[0].fname").As("cname")).
		From("tutorial").
		Where(X("fname").Eq(S("Dave")))
	assert.Equal(t, "SELECT children[0].fname AS cname FROM tutorial WHERE fname = \"Dave\"", statement.String())
}

func Test4(t *testing.T) {
	statement := Select(Meta("tutorial").As("meta")).
		From("tutorial")

	assert.Equal(t, "SELECT META(tutorial) AS meta FROM tutorial", statement.String())
}

func Test5(t *testing.T) {
	statement := Select(X("fname"), X("age"), X("age/7").As("age_dog_years")).
		From("tutorial").
		Where(X("fname").Eq(S("Dave")))

	assert.Equal(t, "SELECT fname, age, age/7 AS age_dog_years FROM tutorial WHERE fname = \"Dave\"",
		statement.String())
}

func Test6(t *testing.T) {
	statement := Select(X("fname"), X("age"), Round(X("age/7"), 0).As("age_dog_years")).
		From("tutorial").
		Where(X("fname").Eq(S("Dave")))

	assert.Equal(t, "SELECT fname, age, ROUND(age/7) AS age_dog_years FROM tutorial WHERE fname = \"Dave\"",
		statement.String())
}

func Test7(t *testing.T) {
	statement := Select(X("fname").Concat(S(" ")).Concat(X("lname")).As("full_name")).
		From("tutorial")

	assert.Equal(t, "SELECT fname || \" \" || lname AS full_name FROM tutorial", statement.String())
}

func Test8(t *testing.T) {
	statement := Select("fname", "age").
		From("tutorial").
		Where(X("age").Gt(X("30")))

	assert.Equal(t, "SELECT fname, age FROM tutorial WHERE age > 30", statement.String())
}

func Test9(t *testing.T) {
	statement := Select("fname", "email").
		From("tutorial").
		Where(X("email").Like(S("%@yahoo.com")))

	assert.Equal(t, "SELECT fname, email FROM tutorial WHERE email LIKE \"%@yahoo.com\"", statement.String())
}

func ExampleSelectDistinct() {
	SelectDistinct("orderlines[0].productId").
		From("orders")
}

func Test10(t *testing.T) {
	statement := SelectDistinct("orderlines[0].productId").
		From("orders")

	assert.Equal(t, "SELECT DISTINCT orderlines[0].productId FROM orders", statement.String())
}

func Test11(t *testing.T) {
	statement := Select("fname", "children").
		From("tutorial").
		Where(X("children").IsNull())

	assert.Equal(t, "SELECT fname, children FROM tutorial WHERE children IS NULL", statement.String())
}

// todo
//func Test12(t *testing.T) {
//	statement := Select("fname", "children").
//		From("tutorial").
//		Where(AnyIn("child", X("tutorial.children")).Satisfies(X("child.age").Gt(10)))
//
//	assert.Equal(t, "SELECT fname, children FROM tutorial WHERE ANY child IN tutorial.children "+
//		"SATISFIES child.age > 10 END", statement.String())
//}

// todo
//func Test13(t *testing.T) {
//	statement := Select("fname", "email", "children").
//		From("tutorial").
//		Where(ArrayLength("children").Gt(X("0")).And(X("email")).Like(S("%@gmail.com")))
//
//	assert.Equal(t, "SELECT fname, email, children FROM tutorial WHERE ARRAY_LENGTH(children) > 0 AND email"+
//		" LIKE \"%@gmail.com\"", statement.String())
//}

func Test14(t *testing.T) {
	statement := Select("fname", "email").
		From("tutorial").
		UseKeysValues("dave", "ian")

	assert.Equal(t, "SELECT fname, email FROM tutorial USE KEYS [\"dave\",\"ian\"]", statement.String())
}

func Test15(t *testing.T) {
	statement := Select("children[0:2]").
		From("tutorial").
		Where(X("children[0:2]").IsNotMissing())

	assert.Equal(t, "SELECT children[0:2] FROM tutorial WHERE children[0:2] IS NOT MISSING", statement.String())
}

// todo
//func Test16(t *testing.T) {
//	statement := Select(X("fname").Concat("\" \"").Concat("lname").As("full_name"),
//		X("email"), X("children[0:2]").As("offsprings")).
//		From("tutorial").
//		Where(
//		X("email").Like(S("%@yahoo.com")).
//			Or(AnyIn("child", X("tutorial.children")).Satisfies(X("child.age").Gt(10))))
//
//	assert.Equal(t, "SELECT fname || \" \" || lname AS full_name, email, children[0:2] AS offsprings "+
//		"FROM tutorial WHERE email LIKE \"%@yahoo.com\" "+
//		"OR ANY child IN tutorial.children SATISFIES child.age > 10 END",
//		statement.String())
//}

func Test17(t *testing.T) {
	statement := Select("fname", "age").From("tutorial").OrderBy(DefaultSort("age"))

	assert.Equal(t, "SELECT fname, age FROM tutorial ORDER BY age", statement.String())
}

func Test18(t *testing.T) {
	statement := Select("fname", "age").
		From("tutorial").
		OrderBy(DefaultSort("age")).Limit(2)

	assert.Equal(t, "SELECT fname, age FROM tutorial ORDER BY age LIMIT 2", statement.String())
}

func Test19(t *testing.T) {
	statement := Select(Count("*").As("count")).
		From("tutorial")

	assert.Equal(t, "SELECT COUNT(*) AS count FROM tutorial", statement.String())
}

func Test20(t *testing.T) {
	statement := Select(X("relation"), Count("*").As("count")).
		From("tutorial").GroupBy(X("relation"))

	assert.Equal(t, "SELECT relation, COUNT(*) AS count FROM tutorial GROUP BY relation", statement.String())
}

func Test21(t *testing.T) {
	statement := Select(X("relation"), Count("*").As("count")).
		From("tutorial").GroupBy(X("relation")).
		Having(Count("*").Gt(1))

	assert.Equal(t, "SELECT relation, COUNT(*) AS count FROM tutorial GROUP BY relation "+
		"HAVING COUNT(*) > 1", statement.String())
}

// todo
//func Test22(t *testing.T) {
//	statement := Select(ArrayIn(X("child.fname"), "child", X("tutorial.children")).End().As("children_names")).
//		From("tutorial").Where(X("children").IsNotNull())
//
//	assert.Equal(t, "SELECT ARRAY child.fname FOR child IN tutorial.children END AS children_names "+
//		"FROM tutorial WHERE children IS NOT NULL", statement.String())
//}

func Test23(t *testing.T) {
	statement := Select(X("t.relation"), Count("*").As("count"), Avg("c.age").As("avg_age")).
		From("tutorial").As("t").
		Unnest("t.children").As("c").
		Where(X("c.age").Gt(10)).
		GroupBy(X("t.relation")).
		Having(Count("*").Gt(1)).
		OrderBy(Desc("avg_age")).
		Limit(1).Offset(1)

	// NOTE: the AS clause in the tutorial uses the shorthand "tutorial t"
	// we only support the extended syntax "tutorial AS t"
	// (the other one brings no real value in the context of the DSL)
	assert.Equal(t, "SELECT t.relation, COUNT(*) AS count, AVG(c.age) AS avg_age "+
		"FROM tutorial AS t "+
		"UNNEST t.children AS c "+
		"WHERE c.age > 10 "+
		"GROUP BY t.relation "+
		"HAVING COUNT(*) > 1 "+
		"ORDER BY avg_age DESC "+
		"LIMIT 1 OFFSET 1", statement.String())
}

// todo
//func Test24(t *testing.T) {
//	statement := Select("usr.personal_details", "orders").
//		From("users_with_orders").As("usr").
//		UseKeysValues("Elinor_33313792").
//		Join("orders_with_users").As("orders").
//		OnKeys(ArrayIn(X("s.order_id"), "s", X("usr.shipped_order_history")).End())
//
//	assert.Equal(t, "SELECT usr.personal_details, orders "+
//		"FROM users_with_orders AS usr "+
//		"USE KEYS \"Elinor_33313792\" "+
//		"JOIN orders_with_users AS orders "+
//		"ON KEYS ARRAY s.order_id FOR s IN usr.shipped_order_history END",
//		statement.String())
//}

// todo
//func Test25(t *testing.T) {
//	statement := Select("usr.personal_details", "orders").
//		From("users_with_orders").As("usr").
//		UseKeysValues("Tamekia_13483660").
//		LeftJoin("orders_with_users").As("orders").
//		OnKeys(ArrayIn(X("s.order_id"), "s", X("usr.shipped_order_history")).End())
//
//	assert.Equal(t, "SELECT usr.personal_details, orders "+
//		"FROM users_with_orders AS usr "+
//		"USE KEYS \"Tamekia_13483660\" "+
//		"LEFT JOIN orders_with_users AS orders "+
//		"ON KEYS ARRAY s.order_id FOR s IN usr.shipped_order_history END", statement.String())
//}

// todo
//func Test26(t *testing.T) {
//	statement := Select("usr.personal_details", "orders").
//		From("users_with_orders").As("usr").
//		UseKeysValues("Elinor_33313792").
//		Nest("orders_with_users").As("orders").
//		OnKeys(ArrayIn(X("s.order_id"), "s", X("usr.shipped_order_history")).End())
//
//	assert.Equal(t, "SELECT usr.personal_details, orders "+
//		"FROM users_with_orders AS usr "+
//		"USE KEYS \"Elinor_33313792\" "+
//		"NEST orders_with_users AS orders "+
//		"ON KEYS ARRAY s.order_id FOR s IN usr.shipped_order_history END", statement.String())
//}

func Test27(t *testing.T) {
	statement := Select("*").From("tutorial").As("contact").
		Unnest("contact.children").Where(X("contact.fname").Eq(S("Dave")))

	assert.Equal(t, "SELECT * FROM tutorial AS contact UNNEST contact.children "+
		"WHERE contact.fname = \"Dave\"", statement.String())
}

func Test28(t *testing.T) {
	statement := Select(X("u.personal_details.display_name").As("name"),
		X("s").As("order_no"), X("o.product_details")).
		From("users_with_orders").As("u").
		UseKeys(S("Aide_48687583")).
		Unnest("u.shipped_order_history").As("s").
		Join("users_with_orders").As("o").OnKeys("s.order_id")

	assert.Equal(t, "SELECT u.personal_details.display_name AS name, s AS order_no, o.product_details "+
		"FROM users_with_orders AS u USE KEYS \"Aide_48687583\" "+
		"UNNEST u.shipped_order_history AS s "+
		"JOIN users_with_orders AS o ON KEYS s.order_id", statement.String())
}

func Test29(t *testing.T) {
	t.SkipNow()
	//todo EXPLAIN clause + INSERT DML
	statement := Select("TODO")

	assert.Equal(t, "EXPLAIN INSERT INTO tutorial (KEY, VALUE) "+
		"VALUES (\"baldwin\", {\"name\":\"Alex Baldwin\", \"type\":\"contact\"})", statement.String())
}

func Test30(t *testing.T) {
	t.SkipNow()
	//todo EXPLAIN clause + DELETE DML
	statement := Select("*")

	assert.Equal(t, "EXPLAIN DELETE FROM tutorial t USE KEYS \"baldwin\" RETURNING t", statement.String())
}

func Test31(t *testing.T) {
	t.SkipNow()
	//todo EXPLAIN clause + UPDATE DML
	statement := Select("*")

	assert.Equal(t, "EXPLAIN UPDATE tutorial USE KEYS \"baldwin\" "+
		"SET type = \"actor\" RETURNING tutorial.type", statement.String())
}

func Test32(t *testing.T) {
	statement := Select(Count("*").As("product_count")).From("product")

	assert.Equal(t, "SELECT COUNT(*) AS product_count FROM product", statement.String())
}

//func Test33(t *testing.T) {
//	statement := Select("*").From("product").
//		Unnest("product.categories").As("cat").
//		Where(Lower("cat").In(JsonArray.From("golf"))). //fixme
//		Limit(10).Offset(10)
//	assert.Equal(t, "SELECT * FROM product "+
//		"UNNEST product.categories AS cat "+
//		"WHERE LOWER(cat) IN [\"golf\"] LIMIT 10 OFFSET 10", statement.String())
//}

func Test34(t *testing.T) {
	statement := SelectDistinct("categories").From("product").
		Unnest("product.categories").As("categories")

	assert.Equal(t, "SELECT DISTINCT categories FROM product "+
		"UNNEST product.categories AS categories", statement.String())
}

func Test35(t *testing.T) {
	statement := Select("productId", "name").From("product").
		Where(Lower("name").Like(S("%cup%")))

	assert.Equal(t, "SELECT productId, name FROM product WHERE LOWER(name) LIKE \"%cup%\"", statement.String())
}

func Test36(t *testing.T) {
	statement := Select("product").From("product").
		Unnest("product.categories").As("categories").
		Where(X("categories").Eq(S("Appliances")))

	assert.Equal(t, "SELECT product FROM product UNNEST product.categories AS categories "+
		"WHERE categories = \"Appliances\"", statement.String())
}

func Test37(t *testing.T) {
	statement := Select(
		X("product.name"),
		Count("reviews").As("reviewCount"),
		Round(Avg("reviews.rating"), 1).As("AvgRating"),
		X("category")).
		From("reviews").As("reviews").
		Join("product").As("product").OnKeys("reviews.productId")

	assert.Equal(t, "SELECT product.name, COUNT(reviews) AS reviewCount, "+
		"ROUND(AVG(reviews.rating), 1) AS AvgRating, category "+
		"FROM reviews AS reviews "+
		"JOIN product AS product ON KEYS reviews.productId", statement.String())
}

func Test38(t *testing.T) {
	statement := Select(X("product.name"), X("product.dateAdded"), Sum("items.count").As("unitsSold")).
		From("purchases").Unnest("purchases.lineItems").As("items").
		Join("product").OnKeys("items.product").GroupBy("product").
		OrderBy(DefaultSort("product.dateAdded"), Desc("unitsSold")).
		Limit(10)

	assert.Equal(t, "SELECT product.name, product.dateAdded, SUM(items.count) AS unitsSold "+
		"FROM purchases UNNEST purchases.lineItems AS items "+
		"JOIN product ON KEYS items.product GROUP BY product "+
		"ORDER BY product.dateAdded, unitsSold DESC LIMIT 10", statement.String())
}

func Test39(t *testing.T) {
	statement := Select("product.name", "product.unitPrice", "product.categories").From("product").
		Unnest("product.categories").As("categories").
		Where(X("categories").Eq(S("Appliances")).
			And(X("product.unitPrice").Lt(6.99)))

	assert.Equal(t, "SELECT product.name, product.unitPrice, product.categories FROM product "+
		"UNNEST product.categories AS categories WHERE categories = \"Appliances\" AND product.unitPrice < 6.99",
		statement.String())
}

func Test40(t *testing.T) {
	statement := Select(X("product.name"), Sum("items.count").As("unitsSold")).From("purchases").
		Unnest("purchases.lineItems").As("items").
		Join("product").OnKeys("items.product").
		GroupBy("product").
		OrderBy(Desc("unitsSold")).Limit(10)

	assert.Equal(t, "SELECT product.name, SUM(items.count) AS unitsSold FROM purchases "+
		"UNNEST purchases.lineItems AS items JOIN product ON KEYS items.product "+
		"GROUP BY product ORDER BY unitsSold DESC LIMIT 10", statement.String())
}

func Test41(t *testing.T) {
	statement := Select(X("product.name"), Round(Avg("reviews.rating"), 1).As("avg_rating")).
		From("reviews").
		Join("product").OnKeys("reviews.productId").
		GroupBy("product").
		OrderBy(Desc(Avg("reviews.rating"))).Limit(5)

	assert.Equal(t, "SELECT product.name, ROUND(AVG(reviews.rating), 1) AS avg_rating FROM reviews "+
		"JOIN product ON KEYS reviews.productId GROUP BY product "+
		"ORDER BY AVG(reviews.rating) DESC LIMIT 5", statement.String())
}

func Test42(t *testing.T) {
	statement := Select("purchases", "product", "customer").From("purchases").
		UseKeysValues("purchase0").
		Unnest("purchases.lineItems").As("items").
		Join("product").OnKeys("items.product").
		Join("customer").OnKeys("purchases.customerId")

	assert.Equal(t, "SELECT purchases, product, customer FROM purchases USE KEYS \"purchase0\" "+
		"UNNEST purchases.lineItems AS items JOIN product ON KEYS items.product "+
		"JOIN customer ON KEYS purchases.customerId", statement.String())
}

func Test43(t *testing.T) {
	statement := Select(X("customer.firstName"), X("customer.lastName"), X("customer.emailAddress"),
		Sum("items.count").As("purchaseCount"),
		Round(Sum(X("product.unitPrice").Multiply("items.count")), 0).As("totalSpent")).
		From("purchases").
		Unnest("purchases.lineItems").As("items").
		Join("product").OnKeys("items.product").
		Join("customer").OnKeys("purchases.customerId").
		GroupBy("customer")

	assert.Equal(t, "SELECT customer.firstName, customer.lastName, customer.emailAddress, "+
		"SUM(items.count) AS purchaseCount, ROUND(SUM(product.unitPrice * items.count)) AS totalSpent "+
		"FROM purchases UNNEST purchases.lineItems AS items "+
		"JOIN product ON KEYS items.product JOIN customer ON KEYS purchases.customerId GROUP BY customer",
		statement.String())
}

func Test44(t *testing.T) {
	statement := Select(Count("customer").As("customerCount"), X("state")).From("customer").
		GroupBy("state").OrderBy(Desc("customerCount"))

	assert.Equal(t, "SELECT COUNT(customer) AS customerCount, state FROM customer "+
		"GROUP BY state ORDER BY customerCount DESC", statement.String())
}

// todo
//func Test45(t *testing.T) {
//	statement := Select(Count(Distinct("purchases.customerId"))).From("purchases").
//		Where(strToMillis("purchases.purchasedAt").
//		Between(strToMillis(S("2014-02-01")).
//		And(strToMillis(S("2014-03-01")))))
//
//	assert.Equal(t, "SELECT COUNT(DISTINCT purchases.customerId) FROM purchases "+
//		"WHERE STR_TO_MILLIS(purchases.purchasedAt) BETWEEN STR_TO_MILLIS(\"2014-02-01\") "+
//		"AND STR_TO_MILLIS(\"2014-03-01\")", statement.String())
//}

func Test46(t *testing.T) {
	statement := Select(X("product"), Avg("reviews.rating").As("avgRating"),
		Count("reviews").As("numReviews")).
		From("product").
		Join("reviews").OnKeys("product.reviewList").
		GroupBy("product").Having(Avg("reviews.rating").Lt(1))

	assert.Equal(t, "SELECT product, AVG(reviews.rating) AS avgRating, COUNT(reviews) AS numReviews "+
		"FROM product JOIN reviews ON KEYS product.reviewList "+
		"GROUP BY product HAVING AVG(reviews.rating) < 1", statement.String())
}

func Test47(t *testing.T) {
	statement := Select(Substr("purchases.purchasedAt", 0, 7).As("month"),
		Round(Sum(X("product.unitPrice").Multiply("items.count")).Divide(1000000), 3).As("revenueMillion")).
		From("purchases").
		Unnest("purchases.lineItems").As("items").
		Join("product").OnKeys("items.product").
		GroupBy(Substr("purchases.purchasedAt", 0, 7)).
		OrderBy(DefaultSort("month"))

	assert.Equal(t, "SELECT SUBSTR(purchases.purchasedAt, 0, 7) AS month, "+
		"ROUND(SUM(product.unitPrice * items.count) / 1000000, 3) AS revenueMillion "+
		"FROM purchases UNNEST purchases.lineItems AS items JOIN product ON KEYS items.product "+
		"GROUP BY SUBSTR(purchases.purchasedAt, 0, 7) "+
		"ORDER BY month", statement.String())
}

// todo
//func Test48(t *testing.T) {
//	statement := Select("purchases.purchaseId", "l.product").From("purchases").
//		Unnest("purchases.lineItems").As("l").
//		Where(DatePartStr("purchases.purchasedAt", month).Eq(4).
//		And(DatePartStr("purchases.purchasedAt", year).Eq(2014)).
//		And(Sub(
//		Select("product.productId").
//			From("product").
//			UseKeys("l.product").
//			Where(X("product.unitPrice").Gt(500))).
//		Exists()))
//
//	assert.Equal(t, "SELECT purchases.purchaseId, l.product FROM purchases UNNEST purchases.lineItems AS l "+
//		"WHERE DATE_PART_STR(purchases.purchasedAt, \"month\") = 4 "+
//		"AND DATE_PART_STR(purchases.purchasedAt, \"year\") = 2014 "+
//		"AND EXISTS (SELECT product.productId "+
//		"FROM product USE KEYS l.product WHERE product.unitPrice > 500)", statement.String())
//}

func Test49(t *testing.T) {
	statement := Select("*").From("jungleville_inbox").Limit(1)

	assert.Equal(t, "SELECT * FROM jungleville_inbox LIMIT 1", statement.String())
}

func Test50(t *testing.T) {
	statement := Select("*").From("jungleville").As("`game-data`").
		Join("jungleville_stats").As("stats").OnKeysValues("zid-jungle-stats-0001").
		Nest("jungleville_inbox").As("inbox").OnKeysValues("zid-jungle-inbox-0001").
		Where(Path(I("game-data"), "uuid").Eq(S("zid-jungle-0001")))

	assert.Equal(t, "SELECT * FROM jungleville AS `game-data` JOIN jungleville_stats AS stats "+
		"ON KEYS \"zid-jungle-stats-0001\" NEST jungleville_inbox AS inbox ON KEYS \"zid-jungle-inbox-0001\" "+
		"WHERE `game-data`.uuid = \"zid-jungle-0001\"", statement.String())
}

func Test51(t *testing.T) {
	statement := Select("player.name", "inbox.messages").
		From("jungleville").As("player").
		UseKeysValues("zid-jungle-0001").
		LeftJoin("jungleville_inbox").As("inbox").
		OnKeys(S("zid-jungle-inbox-").Concat(SUBSTR("player.uuid", 11)))

	assert.Equal(t, "SELECT player.name, inbox.messages FROM jungleville AS player USE KEYS \"zid-jungle-0001\" "+
		"LEFT JOIN jungleville_inbox AS inbox ON KEYS \"zid-jungle-inbox-\" || SUBSTR(player.uuid, 11)", statement.String())
}

// todo
//func Test52(t *testing.T) {
//	statement := Select(X("stats.uuid").As("player"), X("hist.uuid").As("opponent"),
//		Sum(caseSearch().when(X("hist.result").Eq(S("won"))).Then(X(1)).ElseReturn(X(0))).As("wins"),
//		Sum(caseSearch().when(X("hist.result").Eq(S("lost"))).Then(X(1)).ElseReturn(X(0))).As("losses")).
//		From("jungleville_stats").As("stats").UseKeysValues("zid-jungle-stats-0004").
//		Unnest("stats.`pvp-hist`").As("hist").
//		GroupBy("stats.uuid", "hist.uuid")
//
//	assert.Equal(t, "SELECT stats.uuid AS player, hist.uuid AS opponent, "+
//		"SUM(CASE WHEN hist.result = \"won\" THEN 1 ELSE 0 END) AS wins, "+
//		"SUM(CASE WHEN hist.result = \"lost\" THEN 1 ELSE 0 END) AS losses "+
//		"FROM jungleville_stats AS stats USE KEYS \"zid-jungle-stats-0004\" "+
//		"UNNEST stats.`pvp-hist` AS hist GROUP BY stats.uuid, hist.uuid", statement.String())
//}

// todo
//func Test53(t *testing.T) {
//	statement := Select(X("player.name"), X("player.level"), X("stats.loadtime"),
//		Sum(caseSearch().when(X("hist.result").Eq(S("won"))).Then(X(1)).ElseReturn(X(0))).As("wins")).
//		From("jungleville_stats").As("stats").
//		Unnest("stats.`pvp-hist`").As("hist").
//		Join("jungleville").As("player").
//		OnKeys("stats.uuid").
//		GroupBy("player", "stats")
//
//	assert.Equal(t, "SELECT player.name, player.level, stats.loadtime, "+
//		"SUM(CASE WHEN hist.result = \"won\" THEN 1 ELSE 0 END) AS wins "+
//		"FROM jungleville_stats AS stats UNNEST stats.`pvp-hist` AS hist "+
//		"JOIN jungleville AS player ON KEYS stats.uuid GROUP BY player, stats", statement.String())
//}

func Test54(t *testing.T) {
	statement := Select("jungleville.level", "friends").
		From("jungleville").UseKeysValues("zid-jungle-0002").
		Join("jungleville.friends").OnKeys("jungleville.friends")

	assert.Equal(t, "SELECT jungleville.level, friends FROM jungleville USE KEYS \"zid-jungle-0002\" "+
		"JOIN jungleville.friends ON KEYS jungleville.friends", statement.String())
}
