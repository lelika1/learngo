package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/mylist"
)

func main() {
	var list mylist.List
	list.Add(5)
	list.Add(4)
	list.Add(3)
	list.Add(7)
	fmt.Println(list.String())
}
