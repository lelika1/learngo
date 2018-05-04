package main

import (
	"eopi/parity_5_1/parity"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", parity.Cached(0x1aabcdef1234adf1))
	fmt.Printf("%v\n", parity.Simple(0x1aabcdef1234adf1))
}
