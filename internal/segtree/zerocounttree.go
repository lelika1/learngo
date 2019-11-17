package segtree

// ZeroCountTree is a segment tree for calculalating amount of 0-elements on intervals.
type ZeroCountTree struct {
	original []int
	st       SegmentTree
}

// NewZeroCountTree creates ZeroCountTree for the given array.
func NewZeroCountTree(arr []int) *ZeroCountTree {
	zeroArr := make([]int, len(arr))
	for i, el := range arr {
		if el == 0 {
			zeroArr[i] = 1
		} else {
			zeroArr[i] = 0
		}
	}

	st := NewSegmentTree(zeroArr, func(i, j int) int { return i + j })
	return &ZeroCountTree{
		original: arr,
		st:       *st,
	}
}

// ZeroCount returns the amount of 0-elements in the given range of indices.
// Returns false if indices are incorrect.
func (t *ZeroCountTree) ZeroCount(i, j int) (int, bool) {
	return t.st.F(i, j)
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *ZeroCountTree) Set(idx, val int) bool {
	if idx < 0 || idx >= len(t.original) || len(t.original) == 0 {
		return false // wrong index
	}

	t.original[idx] = val
	var isZero int
	if val == 0 {
		isZero = 1
	}

	return t.st.Set(idx, isZero)
}

// String prints the segment tree, layer by layer, one layer per line.
// First line - the sum of the whole array.
// Last line - elements of the original array.
func (t *ZeroCountTree) String() string {
	return t.st.String()
}
