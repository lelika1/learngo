package reverse

// Digit reverses integer number.
// Here are some examples: 453 -> 354, -120 -> -21.
func Digit(number int) (res int) {
	isNeg := false
	if number < 0 {
		isNeg = true
		number *= -1
	}

	for ; number > 0; number /= 10 {
		res = 10*res + number%10
	}

	if isNeg {
		res *= -1
	}
	return res
}
