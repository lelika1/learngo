package segtree

import (
	"math/bits"
	"strconv"
	"strings"
)

// TreeFunc should be an associative function.
// That means: f(f(x, y), z) = f(x, f(y, z)).
type TreeFunc func(int, int) int

// SegmentTree can be created and used for any associative function.
type SegmentTree struct {
	tree   []int
	size   int         // size of original array
	levels []levelDesc // description for every tree's layer (begin, end, etc)
	f      TreeFunc    // function that will be calculated for segments
}

// NewSegmentTree creates SegmentTree for the given array and the function.
func NewSegmentTree(arr []int, f TreeFunc) *SegmentTree {
	if len(arr) == 0 {
		return &SegmentTree{}
	}

	// find the nearest degree of 2 bigger than len(arr)
	size := 1 << (bits.Len(uint(len(arr)) - 1))
	nodes := make([]int, 2*size-1)

	levels := make([]levelDesc, bits.Len(uint(size)))
	levels[len(levels)-1] = levelDesc{
		begin: size - 1,
		end:   (size - 1) + len(arr) - 1,
		cap:   size,
		free:  size - len(arr),
	}
	for i := len(levels) - 2; i >= 0; i-- {
		levels[i].cap = levels[i+1].cap / 2
		levels[i].begin = levels[i+1].begin - levels[i].cap
		levels[i].free = levels[i+1].free / 2
		levels[i].end = levels[i+1].begin - levels[i].free - 1
	}

	copy(nodes[levels[len(levels)-1].begin:], arr) // put original array in the leafs of the SumTree
	for l := len(levels) - 2; l >= 0; l-- {
		for i := levels[l].begin; i <= levels[l].end; i++ {
			left := 2*i + 1
			right := 2*i + 2
			if right > levels[l+1].end {
				nodes[i] = nodes[left]
			} else {
				nodes[i] = f(nodes[left], nodes[right])
			}
		}
	}

	return &SegmentTree{
		tree:   nodes,
		size:   len(arr),
		levels: levels,
		f:      f,
	}
}

// String prints the segment tree, layer by layer, one layer per line.
// First line - is the result of functions 'f' on the whole array.
// Last line - elements of the original array.
// If the current node does not correspond to any element of the original array, then it will be '-'.
func (t *SegmentTree) String() string {
	var sb strings.Builder

	for l, level := range t.levels {
		for i := level.begin; i <= level.end; i++ {
			sb.WriteString(strconv.Itoa(t.tree[i]))
			if i != level.end {
				sb.WriteRune(' ')
			}
		}

		if l != len(t.levels)-1 {
			for i := 0; i < level.free; i++ {
				sb.WriteString(" -")
			}
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

// F returns the result of function 'f' on the given range of indices.
// Returns false if indices are incorrect.
func (t *SegmentTree) F(i, j int) (int, bool) {
	if i < 0 || i >= t.size || j < 0 || j >= t.size || i > j {
		return 0, false
	}

	lastIdx := t.levels[len(t.levels)-1].cap - 1
	return t.recursiveF(0, segmentRange{i, j}, segmentRange{0, lastIdx}), true
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *SegmentTree) Set(idx, val int) bool {
	if idx < 0 || idx >= t.size || len(t.tree) == 0 {
		return false // wrong index
	}

	if t.tree[t.levels[len(t.levels)-1].begin+idx] == val {
		return true // the same value, nothing to do
	}

	cur := t.levels[len(t.levels)-1].begin + idx
	t.tree[cur] = val

	for l := len(t.levels) - 2; l >= 0; l-- {
		cur = (cur - 1) / 2
		left := 2*cur + 1
		right := 2*cur + 2

		if right > t.levels[l+1].end {
			t.tree[cur] = t.tree[left]
		} else {
			t.tree[cur] = t.f(t.tree[left], t.tree[right])
		}
	}
	return true
}

type segmentRange struct {
	l int
	r int
}

// recursiveF checks whether the segment is in one of the children, or intersects with both.
// fSeg - segment where function 'f' should be calculated.
// treeSeg - ranges of the node.
func (t *SegmentTree) recursiveF(node int, fSeg, treeSeg segmentRange) int {
	if fSeg.l == treeSeg.l && fSeg.r == treeSeg.r {
		return t.tree[node]
	}

	left := 2*node + 1
	right := 2*node + 2
	mid := (treeSeg.l + treeSeg.r) / 2

	if fSeg.l <= mid && fSeg.r <= mid {
		return t.recursiveF(left, fSeg, segmentRange{treeSeg.l, mid})
	}

	if fSeg.l > mid && fSeg.r > mid {
		return t.recursiveF(right, fSeg, segmentRange{mid + 1, treeSeg.r})
	}

	return t.f(t.recursiveF(left, segmentRange{fSeg.l, mid}, segmentRange{treeSeg.l, mid}),
		t.recursiveF(right, segmentRange{mid + 1, fSeg.r}, segmentRange{mid + 1, treeSeg.r}))
}

// levelDesc is description for every segment tree's layer.
type levelDesc struct {
	begin int // index of the first element of this layer in the full tree
	end   int // index of the last actual element of this layer in the full tree
	free  int // unused element in this layer
	cap   int // amount of potential elements on the current layer in the full tree
}
