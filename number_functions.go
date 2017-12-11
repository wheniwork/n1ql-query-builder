package nqb

// Abs returns expression results in the absolute value of the number.
func Abs(expression *Expression) *Expression {
	return X("ABS(" + toString(expression) + ")")
}

// Acos returns expression results in the arccosine in radians.
func Acos(expression *Expression) *Expression {
	return X("ACOS(" + toString(expression) + ")")
}

// Asin returns expression results in the arcsine in radians.
func Asin(expression *Expression) *Expression {
	return X("ASIN(" + toString(expression) + ")")
}

// Atan returns expression results in the arctangent in radians.
func Atan(expression *Expression) *Expression {
	return X("ATAN(" + toString(expression) + ")")
}

// ATAN returns expression results in the arctangent of expression2/expression1.
func ATAN(expression1, expression2 interface{}) *Expression {
	return X("ATAN(" + toString(expression1) + ", " + toString(expression2) + ")")
}

// Ceil returns expression results in the smallest integer not less than the number.
func Ceil(expression *Expression) *Expression {
	return X("CEIL(" + toString(expression) + ")")
}

// Cos returns expression results in the cosine.
func Cos(expression *Expression) *Expression {
	return X("COS(" + toString(expression) + ")")
}

// Degrees returns expression results in the conversion of radians to degrees.
func Degrees(expression *Expression) *Expression {
	return X("DEGREES(" + toString(expression) + ")")
}

// E returns expression results in the base of natural logarithms.
func E() *Expression {
	return X("E()")
}

// Exp returns expression results in the exponential of expression.
func Exp(expression *Expression) *Expression {
	return X("EXP(" + toString(expression) + ")")
}

// Ln returns expression results in the log base e.
func Ln(expression *Expression) *Expression {
	return X("LN(" + toString(expression) + ")")
}

// LOG returns expression results in the log base 10.
func LOG(expression *Expression) *Expression {
	return X("LOG(" + toString(expression) + ")")
}

// Floor returns expression results in the largest integer not greater than the number.
func Floor(expression *Expression) *Expression {
	return X("FLOOR(" + toString(expression) + ")")
}

// Pi returns expression results in Pi.
func Pi() *Expression {
	return X("PI()")
}

// Power returns expression results in expression1 to the power of expression2.
func Power(expression1, expression2 interface{}) *Expression {
	return X("POWER(" + toString(expression1) + ", " + toString(expression2) + ")")
}

// Radians returns expression results in the conversion of degrees to radians.
func Radians(expression *Expression) *Expression {
	return X("RADIANS(" + toString(expression) + ")")
}

// Random returns expression results in a pseudo-random number with optional seed.
func Random(seed interface{}) *Expression {
	if seed == nil {
		return X("RANDOM()")
	}

	return X("RANDOM(" + toString(seed) + ")")
}

// Round returns expression results in the value rounded to the given number of integer digits to the right
// of the decimal point (left if digits is negative).
func Round(expression interface{}, digits int) *Expression {
	if digits == 0 {
		return X("ROUND(" + toString(expression) + ")")
	}

	return X("ROUND(" + toString(expression) + ", " + toString(digits) + ")")
}

// Sign returns expression results in the sign of the numerical expression,
// represented as -1, 0, or 1 for negative, zero, or positive numbers respectively.
func Sign(expression *Expression) *Expression {
	return X("SIGN(" + toString(expression) + ")")
}

// Sin returns expression results in the sine.
func Sin(expression *Expression) *Expression {
	return X("SIN(" + toString(expression) + ")")
}

// SquareRoot returns expression results in the square root.
func SquareRoot(expression *Expression) *Expression {
	return X("SQRT(" + toString(expression) + ")")
}

// Tan returns expression results in the tangent.
func Tan(expression *Expression) *Expression {
	return X("TAN(" + toString(expression) + ")")
}

// Trunc returns expression results in a truncation of the number to the given number of integer digits
// to the right of the decimal point (left if digits is negative).
func Trunc(expression interface{}, digits int) *Expression {
	if digits == 0 {
		return X("TRUNC(" + toString(expression) + ")")
	}

	return X("TRUNC(" + toString(expression) + ", " + toString(digits) + ")")
}
