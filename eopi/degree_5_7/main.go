package main

import (
	"fmt"

	"bitbucket.org/lelika/learngo/eopi/degree_5_7/power"
)

func main() {
	x := 2.2
	y := 13
	fmt.Printf("%v ^ %v = %v\n", x, y, power.Simple(x, y))
	fmt.Printf("%v ^ %v = %v\n", x, y, power.Optimized(x, y))
}
