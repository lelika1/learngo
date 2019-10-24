package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/list"
)

func main() {
	var l list.List
	// l.Add(5)
	// l.Add(4)
	// l.Add(3)
	// l.Add(7)
	// fmt.Println(l.String())
	var list2 list.List
	newlist := list.Merge(l, list2)
	fmt.Println(newlist.String())
}
