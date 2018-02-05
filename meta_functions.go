package nqb

// Metadata for the document expression
func Meta(expression interface{}) *Expression {
	if expression == nil {
		return X("META()")
	}

	return X("META(" + toString(expression) + ")")
}

// Base64 encoding of the expression, on the server side
func Base64(expression interface{}) *Expression {
	return X("BASE64(" + toString(expression) + ")")
}

// UUID Universally Unique Identifier version 4, generated on the server side
func UUID() *Expression {
	return X("UUID()")
}
