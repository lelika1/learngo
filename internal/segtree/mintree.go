package segtree

// MinTree is a segment tree for finding min on intervals.
type MinTree struct {
	st SegmentTree
}

// NewMinTree creates MinTree for the given array.
func NewMinTree(arr []int) *MinTree {
	f := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	st := NewSegmentTree(arr, f)
	return &MinTree{st: *st}
}

// Min returns the min element in the given range of indices.
// Returns false if indices are incorrect.
func (t *MinTree) Min(i, j int) (int, bool) {
	return t.st.F(i, j)
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *MinTree) Set(idx, val int) bool {
	return t.st.Set(idx, val)
}

// String prints the segment tree, layer by layer, one layer per line.
// First line - the min element of the whole array.
// Last line - elements of the original array.
func (t *MinTree) String() string {
	return t.st.String()
}
