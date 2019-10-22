package main

import (
	"fmt"
	"math"
)

const Fault = 1e-10

func Sqrt(x float64) (sq float64, iter int) {
	sq_old := x
	for ; math.Abs(sq - sq_old) > Fault; iter++ {
		sq = sq_old - (sq_old * sq_old - x) / (2 * sq_old)
		sq, sq_old = sq_old, sq
	}
	return
}

func main() {
	sq, iter := Sqrt(2)
	fmt.Println(sq * sq, iter)
	
	fmt.Println(math.Abs(math.Sqrt(2) - sq))
}
