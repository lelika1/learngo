package segtree

import (
	"math/bits"
	"strconv"
	"strings"
)

// SumTree is a segment tree for calculalating sum on intervals.
type SumTree struct {
	tree     []int
	size     int // size of original array
	firstIdx int // index of the first element
}

// NewSumTree creates SumTree for the given array.
func NewSumTree(arr []int) *SumTree {
	if len(arr) == 0 {
		return &SumTree{}
	}

	// find the nearest degree of 2 bigger than len(arr)
	size := 1 << (bits.Len(uint(len(arr)) - 1))
	if size < len(arr) {
		size *= 2
	}

	nodes := make([]int, 2*size-1)
	firstIdx := size - 1

	copy(nodes[firstIdx:], arr) // put origingal array in the leafs of the SumTree

	for i := firstIdx - 1; i >= 0; i-- {
		nodes[i] = nodes[2*i+1] + nodes[2*i+2]
	}

	return &SumTree{
		tree:     nodes,
		size:     len(arr),
		firstIdx: firstIdx,
	}
}

// Sum returns the sum of elements in the given range of indices.
// Returns false if indices are incorrect.
func (t *SumTree) Sum(i, j int) (int, bool) {
	if i < 0 || i >= t.size || j < 0 || j >= t.size || i > j {
		return 0, false
	}

	return t.recursiveSum(0, segRange{i, j}, segRange{0, t.size - 1}), true
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *SumTree) Set(idx, val int) bool {
	if idx < 0 || idx >= t.size {
		return false // wrong index
	}

	if t.tree[t.firstIdx+idx] == val {
		return true // the same value, nothing to do
	}

	diff := val - t.tree[t.firstIdx+idx]
	for i := t.firstIdx + idx; ; i = (i - 1) / 2 {
		t.tree[i] += diff
		if i == 0 {
			// stop after updating the root
			break
		}
	}
	return true
}

// String prints the segment tree, layer by layer, one layer per line.
// First line - the sum of the whole array.
// Last line - elements of the original array.
func (t *SumTree) String() string {
	var sb strings.Builder

	layerSize := 1
	var lineEnd int
	for i, el := range t.tree {
		sb.WriteString(strconv.Itoa(el))
		if i == t.firstIdx+t.size-1 {
			break
		}

		if i == lineEnd {
			sb.WriteRune('\n')
			layerSize *= 2
			lineEnd += layerSize
			continue
		}
		sb.WriteRune(' ')
	}

	return sb.String()
}

type segRange struct {
	l int
	r int
}

// recursiveSum checks whether the segment is in one of the children, or intersects with both.
// sumSeg - segment where sum should be calculated.
// treeSeg - ranges of the node.
func (t *SumTree) recursiveSum(node int, sumSeg, treeSeg segRange) int {
	if sumSeg.l == treeSeg.l && sumSeg.r == treeSeg.r {
		return t.tree[node]
	}

	left := 2*node + 1
	right := 2*node + 2

	mid := (treeSeg.l + treeSeg.r) / 2
	if sumSeg.l <= mid && sumSeg.r <= mid {
		return t.recursiveSum(left, sumSeg, segRange{treeSeg.l, mid})
	}

	if sumSeg.l > mid && sumSeg.r > mid {
		return t.recursiveSum(right, sumSeg, segRange{mid + 1, treeSeg.r})
	}
	ret := t.recursiveSum(left, segRange{sumSeg.l, mid}, segRange{treeSeg.l, mid})
	ret += t.recursiveSum(right, segRange{mid + 1, sumSeg.r}, segRange{mid + 1, treeSeg.r})
	return ret
}
