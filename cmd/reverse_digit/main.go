package main

import (
	"fmt"

	"github.com/lelika1/learngo/cmd/reverse_digit/reverse"
)

func main() {
	fmt.Println(152, reverse.Digit(152))
	fmt.Println(-52, reverse.Digit(-52))
	fmt.Println(-120, reverse.Digit(-120))
	fmt.Println(-2, reverse.Digit(-2))
}
