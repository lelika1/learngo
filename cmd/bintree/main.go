package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/bintree"
)

func main() {
	var tree bintree.Tree
	tree.Add(2)
	tree.Add(1)
	tree.Add(3)
	tree.Add(0)

	var arr []int
	tree.InOrder(func(v int) {
		arr = append(arr, v)
	})

	fmt.Println(arr)
}
