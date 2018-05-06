package power_test

import (
	"math"
	"testing"

	"bitbucket.org/lelika/learngo/eopi/degree_5_7/power"
)

func TestPowerSimple(t *testing.T) {
	x := 1.37
	eps := math.Pow(1, -10)
	maxDegree := 100
	for y := -1 * maxDegree; y <= maxDegree; y++ {
		simpl := power.Simple(x, y)
		lib := math.Pow(x, float64(y))
		if math.Abs(simpl-lib) >= eps {
			t.Errorf("power.Simple(%v, %v) = %v, want %v", x, y, simpl, lib)
		}
	}
}

func TestPowerOptimized(t *testing.T) {
	x := 1.37
	eps := math.Pow(1, -10)
	maxDegree := 100
	for y := -1 * maxDegree; y <= maxDegree; y++ {
		opt := power.Optimized(x, y)
		lib := math.Pow(x, float64(y))
		if math.Abs(opt-lib) >= eps {
			t.Errorf("power.Optimized(%v, %v) = %v, want %v", x, y, opt, lib)
		}
	}
}

func BenchmarkPowerSimple(b *testing.B) {
	x := 1.27
	maxDegree := 32000
	for i := -1 * maxDegree; i <= maxDegree; i++ {
		power.Simple(x, i)
	}
}

func BenchmarkPowerOptimized(b *testing.B) {
	x := 1.27
	maxDegree := 32000
	for i := -1 * maxDegree; i <= maxDegree; i++ {
		power.Optimized(x, i)
	}
}
