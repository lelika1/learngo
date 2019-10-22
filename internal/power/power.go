package power

// Simple ...
func Simple(x float64, y int) float64 {
	if y < 0 {
		x = 1.0 / x
		y *= -1
	}

	ret := 1.0
	for ; y > 0; y-- {
		ret *= x
	}
	return ret
}

// Optimized ...
func Optimized(x float64, y int) float64 {
	if y < 0 {
		x = 1.0 / x
		y *= -1
	}

	ret := 1.0
	for y > 0 {
		if y&1 == 1 {
			ret *= x
		}
		x *= x
		y >>= 1
	}
	return ret
}
