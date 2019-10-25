package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/parity"
)

func main() {
	fmt.Printf("%v\n", parity.Cached(0x1aabcdef1234adf1))
	fmt.Printf("%v\n", parity.Simple(0x1aabcdef1234adf1))
}
