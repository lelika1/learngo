package heap

import (
	"strconv"
	"strings"
)

// Heap is a representation of binary heap
type Heap struct {
	heap []int
}

// NewHeap creates new heap with all elements in array.
// If you need an empty heap, send arr as nil
func NewHeap(arr []int) *Heap {
	h := Heap{
		heap: make([]int, len(arr)),
	}

	if len(arr) == 0 {
		return &h
	}

	copy(h.heap, arr)
	for i := len(h.heap) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return &h
}

// Insert adds new element in the heap
func (h *Heap) Insert(v int) {
	h.heap = append(h.heap, v)

	for i := len(h.heap) - 1; i != 0; {
		parent := (i - 1) / 2
		if h.heap[i] <= h.heap[parent] {
			return
		}

		// swap with parent
		h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		i = parent
	}
}

// Max returns value of the max element in the heap
// Returns false if the heap is empty
func (h *Heap) Max() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	return h.heap[0], true
}

// ExtractMax returns value of the max element and removes it
// Returns false if the heap is empty
func (h *Heap) ExtractMax() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}

	max := h.heap[0]
	last := len(h.heap) - 1
	h.heap[0] = h.heap[last]
	h.heap = h.heap[:last]

	h.heapify(0)
	return max, true
}

// Drain iteratively extracts the max element from the heap and returns the resulting sequence of elements.
// The given heap will be empty after this call.
func Drain(h *Heap) []int {
	var result []int
	for v, ok := h.ExtractMax(); ok; v, ok = h.ExtractMax() {
		result = append(result, v)
	}
	return result
}

// String returns a string representation of the heap
func (h *Heap) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	var nodesInLayer, newLineIdx int = 1, 0
	for i, el := range h.heap {
		sb.WriteString(strconv.Itoa(el))
		if i == len(h.heap)-1 {
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

// heapify reorders the heap to make the element with the given index satisfy heap invariant
func (h *Heap) heapify(idx int) {
	for {
		// find biggest among element with index idx and his children
		biggest := idx
		left := 2*idx + 1
		right := 2 * (idx + 1)
		if left < len(h.heap) && h.heap[idx] < h.heap[left] {
			biggest = left
		}

		if right < len(h.heap) && h.heap[biggest] < h.heap[right] {
			biggest = right
		}

		if biggest == idx {
			break // there is nothing more to normalize
		}

		// swap biggest element with element idx
		// and perfome the same procedure for new index
		h.heap[biggest], h.heap[idx] = h.heap[idx], h.heap[biggest]
		idx = biggest
	}
}
