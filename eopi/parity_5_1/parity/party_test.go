package parity_test

import (
	"eopi/parity_5_1/parity"
	"testing"
)

func TestParity(t *testing.T) {
	for ones := 0; ones < 64; ones++ {
		want := int8(ones & 1)
		var num uint64
		for i := 0; i < ones; i++ {
			num = (num << 1) ^ 1
		}
		if got := parity.Simple(num); got != want {
			t.Errorf("parity.Simple(%v) = %v, want %v", num, got, want)
		}
		if got := parity.Cached(num); got != want {
			t.Errorf("parity.Cached(%v) = %v, want %v", num, got, want)
		}
	}
}

func BenchmarkParitySimple(b *testing.B) {
	for i := uint64(0); i < 1<<24; i++ {
		parity.Simple(i)
	}
}

func BenchmarkParityCached(b *testing.B) {
	for i := uint64(0); i < 1<<24; i++ {
		parity.Cached(i)
	}
}
