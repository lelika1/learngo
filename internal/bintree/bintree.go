package bintree

import (
	"strconv"
	"strings"
)

// Node of tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// Order of tree traversal.
type Order int

// Different options to traverse a tree.
const (
	InOrder Order = iota
	PreOrder
	PostOrder
)

// Tree structure
type Tree struct {
	root *Node
}

// Add to tree
func (t *Tree) Add(val int) {
	newNode := &Node{value: val}
	if t.root == nil {
		t.root = newNode
		return
	}

	for curNode := t.root; ; {
		if curNode.value > val {
			if curNode.left == nil {
				curNode.left = newNode
				break
			}
			curNode = curNode.left
			continue
		}
		if curNode.value < val {
			if curNode.right == nil {
				curNode.right = newNode
				break
			}
			curNode = curNode.right
			continue
		}
		return
	}
}

// Consumer ...
type Consumer func(v int)

// InOrder traverses the tree in order and calls fn for every entry in the tree.
func (t Tree) InOrder(fn Consumer) {
	if t.root == nil {
		return
	}

	curNode := t.root
	nodes := []*Node{curNode}
	for len(nodes) != 0 {
		for curNode.left != nil {
			curNode = curNode.left
			nodes = append(nodes, curNode)
		}

		curNode = nodes[len(nodes)-1]
		fn(curNode.value)
		nodes = nodes[:len(nodes)-1]

		for len(nodes) != 0 || curNode.right != nil {
			if curNode.right != nil {
				curNode = curNode.right
				nodes = append(nodes, curNode)
				break
			}
			curNode = nodes[len(nodes)-1]
			fn(curNode.value)
			nodes = nodes[:len(nodes)-1]
		}
	}
}

// PreOrder ...
func (t Tree) PreOrder(fn Consumer) {
	if t.root == nil {
		return
	}

	fn(t.root.value)
	Tree{t.root.left}.PreOrder(fn)
	Tree{t.root.right}.PreOrder(fn)
}

// PostOrder ...
func (t Tree) PostOrder(fn Consumer) {
	if t.root == nil {
		return
	}
	Tree{t.root.left}.PostOrder(fn)
	Tree{t.root.right}.PostOrder(fn)
	fn(t.root.value)
}

func (t Tree) String() string {
	return PrintContainer(t.InOrder)
}

// TraverseFn ...
type TraverseFn func(consumer Consumer)

// PrintContainer formats the contents of the tree in the given order.
func PrintContainer(traverse TraverseFn) string {
	var sb strings.Builder
	sb.WriteRune('[')
	traverse(func(v int) {
		if sb.Len() != 1 {
			sb.WriteRune(' ')
		}
		sb.WriteString(strconv.Itoa(v))
	})
	sb.WriteRune(']')
	return sb.String()
}
