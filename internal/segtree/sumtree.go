package segtree

// SumTree is a segment tree for calculalating sum on intervals.
type SumTree struct {
	st SegmentTree
}

// NewSumTree creates SumTree for the given array.
func NewSumTree(arr []int) *SumTree {
	st := NewSegmentTree(arr, func(i, j int) int { return i + j })
	return &SumTree{st: *st}
}

// Sum returns the sum of elements in the given range of indices.
// Returns false if indices are incorrect.
func (t *SumTree) Sum(i, j int) (int, bool) {
	return t.st.F(i, j)
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *SumTree) Set(idx, val int) bool {
	return t.st.Set(idx, val)
}

// String prints the segment tree, layer by layer, one layer per line.
// First line - the sum of the whole array.
// Last line - elements of the original array.
func (t *SumTree) String() string {
	return t.st.String()
}
