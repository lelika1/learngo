package main

import (
	"fmt"
)

// Key ...
type Key int

// Value ...
type Value interface{}

// Node ...
type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
}

// Map implements a map from Key to Value.
// The underlying datastructure is BST (binary search tree).
// It is not guaranteed to be auto-balanced.
type Map struct {
	root *Node
}

// Insert returns an error if the key already exists.
// Problems with v *Value ??
func (m *Map) Insert(k Key, v Value) {
	newNode := Node{
		key:   k,
		value: v,
	}
	if m.root == nil {
		m.root = &newNode
		return
	}

	curNode := m.root
	for {
		if curNode.key == k {
			curNode.value = v
			break
		}

		if curNode.key < k {
			if curNode.right == nil {
				curNode.right = &newNode
				break
			}

			curNode = curNode.right
			continue
		}

		if curNode.left == nil {
			curNode.left = &newNode
			break
		}

		curNode = curNode.left
	}
}

// Find returns a pointer on value associated with the key.
// Returns an false if the key is not found.
func (m *Map) Find(k Key) (*Value, bool) {
	curNode := m.root
	for curNode != nil {
		if curNode.key == k {
			return &curNode.value, true
		}
		if curNode.key > k {
			curNode = curNode.left
			continue
		}
		curNode = curNode.right
	}
	return nil, false
}

func (m *Map) printPartTree(r *Node) {
	fmt.Printf("(k: %v, v: %v)\n", r.key, r.value)
	if r.left != nil {
		m.printPartTree(r.left)
	}
	if r.right != nil {
		m.printPartTree(r.right)
	}
}

// Print map in prefix traverse
func (m *Map) Print() {
	if m.root == nil {
		fmt.Println("Map is empty")
		return
	}

	m.printPartTree(m.root)
}

// Rm removes a given key if it is present in the map.
func (m *Map) Rm(k Key) {
	if m.root == nil {
		return
	}

	parent := m.root
	curNode := m.root
	for curNode != nil {
		if curNode.key == k {
			break
		}

		parent = curNode
		if curNode.key > k {
			curNode = curNode.left
			continue
		}

		curNode = curNode.right
	}

	if curNode == nil {
		return // there is no key in map
	}

	if curNode.left == nil && curNode.right == nil {
		if curNode == m.root {
			m.root = nil
			return
		}

		if parent.left == curNode {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}

	if curNode.left == nil {
		if curNode == m.root {
			m.root = curNode.right
			return
		}

		if parent.left == curNode {
			parent.left = curNode.right
			return
		}

		parent.right = curNode.right
		return
	}

	if curNode.right == nil {
		if curNode == m.root {
			m.root = curNode.left
			return
		}

		if parent.left == curNode {
			parent.left = curNode.left
			return
		}

		parent.right = curNode.left
		return
	}

	if curNode.right.left == nil {
		curNode.key = curNode.right.key
		curNode.value = curNode.right.value
		curNode.right = curNode.right.right
		return
	}

	leastNode := curNode.right
	for leastNode.left.left != nil {
		leastNode = leastNode.left
	}
	curNode.key = leastNode.left.key
	curNode.value = leastNode.left.value
	leastNode.left = nil
}

func main() {
	var m Map
	m.Insert(100, 100)
	m.Insert(150, 150)
	m.Insert(130, 130)
	m.Insert(200, 200)
	m.Insert(160, 160)
	m.Insert(230, 230)
	m.Insert(155, 155)
	m.Insert(170, 170)
	m.Insert(60, 60)
	m.Insert(70, 70)
	m.Insert(40, 40)
	m.Insert(80, 80)
	m.Insert(50, 50)
	m.Insert(55, 55)
	m.Insert(30, 30)
	m.Insert(20, 20)

	fmt.Println("Initial map")
	m.Print()

	var key Key = 60
	if v, ok := m.Find(key); ok {
		fmt.Printf("\nValue for key %v is %v.\n", key, *v)
		*v = 666
	}

	fmt.Printf("\nMap after changing value for key %v\n", key)
	m.Print()

	key = 10
	if _, ok := m.Find(key); !ok {
		fmt.Printf("\nValue for key %v is not found!\n", key)
	}

	key = 150
	fmt.Printf("\nMap after removing key %v\n", key)
	m.Rm(key)
	m.Print()
}
