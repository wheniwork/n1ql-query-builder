package nqb

// Select starts a SELECT statement
func Select(expressions ...interface{}) FromPath {
	return newDefaultSelectPath(nil).Select(expressions...)
}

// SelectAll starts a SELECT ALL statement
func SelectAll(expressions ...interface{}) FromPath {
	return newDefaultSelectPath(nil).SelectAll(expressions...)
}

// SelectDistinct starts a SELECT DISTINCT statement
func SelectDistinct(expressions ...interface{}) FromPath {
	return newDefaultSelectPath(nil).SelectDistinct(expressions...)
}

// SelectRaw starts a SELECT RAW statement
func SelectRaw(expression interface{}) FromPath {
	return newDefaultSelectPath(nil).SelectRaw(expression)
}
