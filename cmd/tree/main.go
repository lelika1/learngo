package main

import (
	"fmt"

	"github.com/lelika1/learngo/internal/bintree"
)

func main() {
	var tree bintree.Tree
	tree.Add(5)
	tree.Add(10)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(12)
	tree.Add(13)
	fmt.Printf("PreOrder: %v\n", tree.ToString(bintree.PreOrder))
	fmt.Printf("InOrder: %v\n", tree.ToString(bintree.InOrder))
	fmt.Printf("PostOrder: %v\n", tree.ToString(bintree.PostOrder))
	fmt.Printf("V: %v\n", tree)
}
