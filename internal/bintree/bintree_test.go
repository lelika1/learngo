package bintree_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/bintree"
)

func TestToStrint(t *testing.T) {
	var tree bintree.Tree
	tree.Add(5)
	tree.Add(10)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(12)
	tree.Add(13)

	want := "[1 2 4 5 10 12 13]"
	got := bintree.PrintContainer(tree.InOrder)
	if got != want {
		t.Fatalf("ToString(InOrder) failed. Got:%v Want: %v", got, want)
	}

	if str := tree.String(); str != want {
		t.Fatalf("ToString(InOrder) != String() (%v != %v)", want, str)
	}

	want = "[5 2 1 4 10 12 13]"
	got = bintree.PrintContainer(tree.PreOrder)
	if got != want {
		t.Fatalf("ToString(PreOrder) failed. Got:%v Want: %v", got, want)
	}

	want = "[1 4 2 13 12 10 5]"
	got = bintree.PrintContainer(tree.PostOrder)
	if got != want {
		t.Fatalf("ToString(PostOrder) failed. Got:%v Want: %v", got, want)
	}
}
