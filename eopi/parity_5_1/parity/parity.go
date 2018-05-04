package parity

// Simple ...
func Simple(n uint64) int8 {
	var parity int8
	for n > 0 {
		parity ^= 1
		n = n & (n - 1)
	}
	return parity
}

// Cached ...
func Cached(n uint64) int8 {
	var parity int8
	for n > 0 {
		parity ^= cache[uint8(n&0xff)]
		n >>= 8
	}
	return parity
}

var cache []int8

func init() {
	cache = make([]int8, 1<<8)
	for i := uint64(0); i < 1<<8; i++ {
		cache[i] = Simple(i)
	}
}
