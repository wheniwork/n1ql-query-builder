package nqb

/**
 * Returned expression results in True if the string expression contains the substring.
 */
func Contains(expression interface{}, substring string) *Expression {
	return X("CONTAINS(" + toString(expression) + ", \"" + substring + "\")")
}

/**
 * Returned expression results in the conversion of the string so that the first letter
 * of each word is uppercase and every other letter is lowercase.
 */
func InitCap(expression interface{}) *Expression {
	return X("INITCAP(" + toString(expression) + ")")
}

/**
 * Returned expression results in the conversion of the string so that the first letter
 * of each word is uppercase and every other letter is lowercase.
 */
func Title(expression interface{}) *Expression {
	return X("TITLE(" + toString(expression) + ")")
}

/**
 * Returned expression results in the length of the string expression.
 */
func Length(expression interface{}) *Expression {
	return X("LENGTH(" + toString(expression) + ")")
}

/**
 * Returned expression results in the given string expression in lowercase
 */
func Lower(expression interface{}) *Expression {
	return X("LOWER(" + toString(expression) + ")")
}

/**
 * Returned expression results in the string with all leading white spaces removed.
 */
func Ltrim(expression interface{}) *Expression {
	return X("LTRIM(" + toString(expression) + ")")
}

/**
 * Returned expression results in the string with all leading chars removed (any char in the characters string).
 */
func LTRIM(expression interface{}, characters string) *Expression {
	return X("LTRIM(" + toString(expression) + ", \"" + characters + "\")")
}

/**
 * Returned expression results in the first position of the substring within the string, or -1.
 * The position is zero-based, i.e., the first position is 0.
 */
func Position(expression interface{}, substring string) *Expression {
	return X("POSITION(" + toString(expression) + ", \"" + substring + "\")")
}

/**
 * Returned expression results in the string formed by repeating expression n times.
 */
func Repeat(expression interface{}, n int) *Expression {
	return X("REPEAT(" + toString(expression) + ", " + toString(n) + ")")
}

/**
 * Returned expression results in a string with all occurrences of substr replaced with repl.
 */
func Replace(expression interface{}, substring, repl string) *Expression {
	return X("REPLACE(" + toString(expression) + ", \"" + substring + "\", \"" + repl + "\")")
}

/**
 * Returned expression results in a string with at most n occurrences of substr replaced with repl.
 */
func REPLACE(expression interface{}, substring, repl string, n int) *Expression {
	return X("REPLACE(" + toString(expression) + ", \"" + substring + "\", \"" + repl + "\", " + toString(n) + ")")
}

/**
 * Returned expression results in the string with all trailing white spaces removed.
 */
func Rtrim(expression interface{}) *Expression {
	return X("RTRIM(" + toString(expression) + ")")
}

/**
 * Returned expression results in the string with all trailing chars removed (any char in the characters string).
 */
func RTRIM(expression interface{}, characters string) *Expression {
	return X("RTRIM(" + toString(expression) + ", \"" + characters + "\")")
}

/**
 * Returned expression results in a split of the string into an array of substrings
 * separated by any combination of white space characters.
 */
func Split(expression interface{}) *Expression {
	return X("SPLIT(" + toString(expression) + ")")
}

/**
 * Returned expression results in a split of the string into an array of substrings separated by sep.
 */
func SPLIT(expression interface{}, sep string) *Expression {
	return X("SPLIT(" + toString(expression) + ", \"" + sep + "\")")
}

/**
 * Returned expression results in a substring from the integer position of the given length.
 *
 * The position is zero-based, i.e. the first position is 0.
 * If position is negative, it is counted from the end of the string -1 is the last position in the string.
 */
func Substr(expression interface{}, position, length int) *Expression {
	return X("SUBSTR(" + toString(expression) + ", " + toString(position) + ", " + toString(length) + ")")
}

/**
 * Returned expression results in a substring from the integer position to the end of the string.
 *
 * The position is zero-based, i.e. the first position is 0.
 * If position is negative, it is counted from the end of the string -1 is the last position in the string.
 */
func SUBSTR(expression interface{}, position int) *Expression {
	return X("SUBSTR(" + toString(expression) + ", " + toString(position) + ")")
}

/**
 * Returned expression results in the string with all leading and trailing white spaces removed.
 */
func Trim(expression interface{}) *Expression {
	return X("TRIM(" + toString(expression) + ")")
}

/**
 * Returned expression results in the string with all leading and trailing chars removed
 * (any char in the characters string).
 */
func TRIM(expression interface{}, characters string) *Expression {
	return X("TRIM(" + toString(expression) + ", \"" + characters + "\")")
}

/**
 * Returned expression results in uppercase of the string expression.
 */
func Upper(expression interface{}) *Expression {
	return X("UPPER(" + toString(expression) + ")")
}
