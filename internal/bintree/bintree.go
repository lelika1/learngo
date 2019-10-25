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

func (t Tree) inOrder() string {
	if t.root == nil {
		return ""
	}

	var result string
	result += Tree{t.root.left}.inOrder()
	result += " "
	result += strconv.Itoa(t.root.value)
	result += Tree{t.root.right}.inOrder()
	return result
}

func (t Tree) preOrder() string {
	if t.root == nil {
		return ""
	}

	var result string
	result += " "
	result += strconv.Itoa(t.root.value)
	result += Tree{t.root.left}.preOrder()
	result += Tree{t.root.right}.preOrder()
	return result
}

func (t Tree) postOrder() string {
	if t.root == nil {
		return ""
	}

	var result string
	result += Tree{t.root.left}.postOrder()
	result += Tree{t.root.right}.postOrder()
	result += " "
	result += strconv.Itoa(t.root.value)
	return result
}

func (t Tree) String() string {
	return t.ToString(InOrder)
}

// ToString formats the contents of the tree in the given order.
func (t *Tree) ToString(order Order) string {
	switch order {
	case InOrder:
		return "[" + strings.TrimSpace(t.inOrder()) + "]"
	case PreOrder:
		return "[" + strings.TrimSpace(t.preOrder()) + "]"
	case PostOrder:
		return "[" + strings.TrimSpace(t.postOrder()) + "]"
	}
	return ""
}
