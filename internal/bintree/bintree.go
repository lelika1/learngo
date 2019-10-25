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
	newNode := &Node{val, nil, nil}
	if t.root == nil {
		t.root = newNode
		return
	}
	curNode := t.root
	var stop bool = false
	for !stop {
		if curNode.value > val {
			if curNode.left == nil {
				curNode.left = newNode
				stop = true
			}
			curNode = curNode.left
		} else if curNode.value < val {
			if curNode.right == nil {
				curNode.right = newNode
				stop = true
			}
			curNode = curNode.right
		}

	}
}

func (t *Tree) inString() string {
	curNode := t.root
	if curNode == nil {
		return ""
	}
	var result string
	if curNode.left != nil {
		leftTree := Tree{curNode.left}
		result += leftTree.inString()
	}
	result += " "
	result += strconv.Itoa(curNode.value)
	if curNode.right != nil {
		rightTree := Tree{curNode.right}
		result += rightTree.inString()
	}
	return result
}

func (t *Tree) preString() string {
	curNode := t.root
	if curNode == nil {
		return ""
	}
	var result string

	result += " "
	result += strconv.Itoa(curNode.value)

	if curNode.left != nil {
		leftTree := Tree{curNode.left}
		result += leftTree.preString()
	}

	if curNode.right != nil {
		rightTree := Tree{curNode.right}
		result += rightTree.preString()
	}
	return result
}

func (t *Tree) postString() string {
	curNode := t.root
	if curNode == nil {
		return ""
	}
	var result string

	if curNode.left != nil {
		leftTree := Tree{curNode.left}
		result += leftTree.postString()
	}

	if curNode.right != nil {
		rightTree := Tree{curNode.right}
		result += rightTree.postString()
	}

	result += " "
	result += strconv.Itoa(curNode.value)

	return result
}

func (t Tree) String() string {
	return t.ToString(InOrder)
}

// ToString formats the contents of the tree in the given order.
func (t *Tree) ToString(order Order) string {
	switch order {
	case InOrder:
		return "[" + strings.TrimSpace(t.inString()) + "]"
	case PreOrder:
		return "[" + strings.TrimSpace(t.preString()) + "]"
	case PostOrder:
		return "[" + strings.TrimSpace(t.postString()) + "]"
	}
	return ""
}
