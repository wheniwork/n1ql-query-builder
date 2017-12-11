package nqb

// ArrayAgg returns expression results in a array of the non-MISSING values in the group, including NULLs.
func ArrayAgg(expression interface{}) *Expression {
	return X("ARRAY_AGG(" + toString(expression) + ")")
}

// Avg returns expression results in the arithmetic mean (average) of all the distinct number values in the group.
func Avg(expression interface{}) *Expression {
	return X("AVG(" + toString(expression) + ")")
}

// Count returns expression results in count of all the non-NULL and non-MISSING values in the group.
func Count(expression interface{}) *Expression {
	return X("COUNT(" + toString(expression) + ")")
}

// CountAll returns expression results in a count of all the input rows for the group, regardless of value (including NULL).
func CountAll() *Expression {
	return X("COUNT(*)")
}

// Max returns expression results in the maximum non-NULL, non-MISSING value in the group in N1QL collation order.
func Max(expression interface{}) *Expression {
	return X("MAX(" + toString(expression) + ")")
}

// Min returns expression results in the minimum non-NULL, non-MISSING value in the group in N1QL collation order.
func Min(expression interface{}) *Expression {
	return X("MIN(" + toString(expression) + ")")
}

// Sum returns expression results in the sum of all the number values in the group.
func Sum(expression interface{}) *Expression {
	return X("SUM(" + toString(expression) + ")")
}

// Distinct prefixes an expression with DISTINCT, useful for example for distinct count "COUNT(DISTINCT expression)".
func Distinct(expression interface{}) *Expression {
	return X("DISTINCT " + toString(expression))
}
