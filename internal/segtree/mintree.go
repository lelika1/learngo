package segtree

// MinTree is a segment tree for finding min on intervals.
type MinTree struct {
	*SegmentTree
}

// NewMinTree creates MinTree for the given array.
func NewMinTree(arr []int) *MinTree {
	fn := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	return &MinTree{NewSegmentTree(arr, fn)}
}

// Min returns the min element in the given range of indices.
// Returns false if indices are incorrect.
func (t *MinTree) Min(i, j int) (int, bool) {
	return t.SegmentTree.Aggregate(i, j)
}
