package nqb

// Select starts a SELECT statement
func Select(expressions ...interface{}) FromClause {
	return newDefaultSelectClause(nil).Select(expressions...)
}

// SelectAll starts a SELECT ALL statement
func SelectAll(expressions ...interface{}) FromClause {
	return newDefaultSelectClause(nil).SelectAll(expressions...)
}

// SelectDistinct starts a SELECT DISTINCT statement
func SelectDistinct(expressions ...interface{}) FromClause {
	return newDefaultSelectClause(nil).SelectDistinct(expressions...)
}

// SelectRaw starts a SELECT RAW statement
func SelectRaw(expression interface{}) FromClause {
	return newDefaultSelectClause(nil).SelectRaw(expression)
}
