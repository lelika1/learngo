package segtree

import (
	"math"
	"strconv"
	"strings"
)

// SumTree is a segment tree for calculalating sum on intervals
type SumTree struct {
	tree    []int
	size    int // size of original array
	leafIdx int // index of the first element
}

// NewSumTree creates SumTree for the given array
func NewSumTree(arr []int) *SumTree {
	var t SumTree
	if arr == nil {
		return &t
	}

	t.size = len(arr)

	// find nearest dergree of 2, that is bigger than len(arr)
	deg2 := math.Round(math.Log2(float64(len(arr))))
	size := int(math.Pow(2.0, deg2))
	if size < len(arr) {
		size *= 2
	}

	t.tree = make([]int, 2*size-1)
	t.leafIdx = size - 1

	// put origingal array in the leafs of the SumTree
	for i := 0; i < len(arr); i++ {
		t.tree[t.leafIdx+i] = arr[i]
	}

	for i := t.leafIdx - 1; i >= 0; i-- {
		t.tree[i] = t.tree[2*i+1] + t.tree[(i+1)*2]
	}
	return &t
}

func (t *SumTree) recursiveSum(node, i, j, left, right int) int {
	if i == left && j == right {
		return t.tree[node]
	}

	mid := (left + right) / 2
	if i <= mid && j <= mid {
		return t.recursiveSum(2*node+1, i, j, left, mid)
	}

	if i > mid && j > mid {
		return t.recursiveSum(2*(node+1), i, j, mid+1, right)
	}

	return t.recursiveSum(2*node+1, i, mid, left, mid) +
		+t.recursiveSum(2*(node+1), mid+1, j, mid+1, right)
}

// Sum returns the amount on the given segment
// Returns false if indexes are incorrect
func (t *SumTree) Sum(i, j int) (int, bool) {
	if i < 0 || i >= t.size || j < 0 || j >= t.size || i > j {
		return 0, false
	}

	return t.recursiveSum(0, i, j, 0, t.size-1), true
}

// Update changes value of element idx and refreshes all tree
// Returns false if idx is out of bounds
func (t *SumTree) Update(idx, val int) bool {
	if idx < 0 || idx >= t.size {
		return false // wrong index
	}

	if t.tree[t.leafIdx+idx] == val {
		return true // the same value, nothing to do
	}

	diff := val - t.tree[t.leafIdx+idx]
	for i := t.leafIdx + idx; ; i = (i - 1) / 2 {
		t.tree[i] += diff
		if i == 0 {
			// stop after refreshing root
			break
		}
	}
	return true
}

// String returns a string representation of the whole SumTree
// Original array is stored in leafs
func (t *SumTree) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	var nodesInLayer, newLineIdx int = 1, 0
	for i, el := range t.tree {
		sb.WriteString(strconv.Itoa(el))
		if i == t.leafIdx+t.size-1 {
			break
		}

		if i == newLineIdx {
			sb.WriteRune('\n')
			nodesInLayer *= 2
			newLineIdx += nodesInLayer
			continue
		}

		sb.WriteRune(' ')
	}

	sb.WriteString("]")
	return sb.String()
}
