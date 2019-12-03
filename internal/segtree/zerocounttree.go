package segtree

// ZeroCountTree is a segment tree for calculalating amount of 0-elements on intervals.
type ZeroCountTree struct {
	*SegmentTree
}

// NewZeroCountTree creates ZeroCountTree for the given array.
func NewZeroCountTree(arr []int) *ZeroCountTree {
	zeroArr := make([]int, 0, len(arr))
	for _, el := range arr {
		if el == 0 {
			zeroArr = append(zeroArr, 1)
		} else {
			zeroArr = append(zeroArr, 0)
		}
	}

	fn := func(i, j int) int { return i + j }
	return &ZeroCountTree{NewSegmentTree(zeroArr, fn)}
}

// ZeroCount returns the amount of 0-elements in the given range of indices.
// Returns false if indices are incorrect.
func (t *ZeroCountTree) ZeroCount(i, j int) (int, bool) {
	return t.Aggregate(i, j)
}

// Set changes value of element idx and refreshes all tree.
// Returns false if idx is out of bounds.
func (t *ZeroCountTree) Set(idx, val int) bool {
	if val == 0 {
		return t.SegmentTree.Set(idx, 1)
	}

	return t.SegmentTree.Set(idx, 0)
}
