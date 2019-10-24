package list

import (
	"strconv"
	"strings"
)

// Node of list
type Node struct {
	value int
	next  *Node
}

// Merge merges two separate lists in a new one
// [3, 4] + [1, 2] -> [3, 4, 1, 2]
func Merge(a, b List) (newList List) {
	if a.root != nil {
		curNode := a.root
		for curNode != nil {
			newList.Add(curNode.value)
			curNode = curNode.next
		}
	}
	if b.root != nil {
		curNode := b.root
		for curNode != nil {
			newList.Add(curNode.value)
			curNode = curNode.next
		}
	}
	return newList
}

// List struct
type List struct {
	root *Node
	end  *Node
	size int
}

// Size returns the size of the list.
func (l *List) Size() int {
	return l.size
}

// Insert inserts the given element into a list at the given position.
// Returns true if the insertion succeded.
func (l *List) Insert(i, value int) (ok bool) {
	if i > l.Size() || i < 0 {
		return false
	}

	if i == l.Size() {
		l.Add(value)
		return true
	}

	node := &Node{value, nil}
	l.size++
	if i == 0 {
		node.next = l.root
		l.root = node
	} else {
		prevNode := l.root
		for j := 1; j < i; j++ {
			prevNode = prevNode.next
		}
		node.next = prevNode.next
		prevNode.next = node
	}
	return true
}

// Extract removes the node with index i from the list, if there is one.
// Returns the removed node or nil
func (l *List) Extract(i int) *Node {
	if i >= l.Size() || i < 0 {
		return nil
	}
	l.size--
	node := l.root
	if i == 0 {
		l.root = l.root.next
		if l.Size() == 0 {
			l.end = nil
		}
		return node
	}

	prevNode := l.root
	node = l.root.next
	for j := 1; j < i; j++ {
		prevNode = node
		node = node.next
	}

	prevNode.next = node.next
	if l.end == node {
		l.end = prevNode
	}

	return node
}

// Add to new item to the list
func (l *List) Add(newVal int) {
	l.size++
	if l.root == nil {
		l.root = &Node{newVal, nil}
		l.end = l.root
		return
	}

	l.end.next = &Node{newVal, nil}
	l.end = l.end.next
}

// String returns all values of the list in one string
func (l *List) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	if l.root != nil {
		curNode := l.root
		for curNode != nil {
			sb.WriteString(strconv.Itoa(curNode.value))
			if curNode != l.end {
				sb.WriteString(" ")
			}
			curNode = curNode.next
		}
	}

	sb.WriteString("]")
	return sb.String()
}

// Get returns the value of the i-th element
// Return false if there is no i-th element
func (l *List) Get(i int) (value int, ok bool) {
	if i >= l.Size() || i < 0 {
		return 0, false
	}

	node := l.root
	for j := 0; j < i; j++ {
		node = node.next
	}
	return node.value, true
}

// Begin returns the pointer to the root node of the list
func (l *List) Begin() *Node {
	return l.root
}

// End returns the pointer to the end node of the list
func (l *List) End() *Node {
	return l.end
}

// Value returns the value of the given node
func (n *Node) Value() int {
	return n.value
}

// Next returns the pointer to the next node
func (n *Node) Next() *Node {
	return n.next
}
