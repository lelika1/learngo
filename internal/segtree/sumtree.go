package segtree

// SumTree is a segment tree for calculalating sum on intervals.
type SumTree struct {
	*SegmentTree
}

// NewSumTree creates SumTree for the given array.
func NewSumTree(arr []int) *SumTree {
	fn := func(i, j int) int { return i + j }
	return &SumTree{NewSegmentTree(arr, fn)}
}

// Sum returns the sum of elements in the given range of indices.
// Returns false if indices are incorrect.
func (t *SumTree) Sum(i, j int) (int, bool) {
	return t.Aggregate(i, j)
}
