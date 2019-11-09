package heap

import (
	"strconv"
	"strings"
)

// Heap is a representation of binary heap
type Heap struct {
	heap []int
	size int
}

// NewHeap creates new heap with all elements in array.
// If you need an empty heap, send arr as nil
func NewHeap(arr []int) *Heap {
	h := Heap{
		heap: make([]int, len(arr)),
		size: len(arr),
	}

	copy(h.heap, arr)
	for i := h.size / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return &h
}

// Insert adds new element in the heap
func (h *Heap) Insert(v int) {
	h.heap = append(h.heap, v)
	h.size++

	for i := h.size - 1; i != 0; {
		parent := (i - 1) / 2
		if h.heap[i] > h.heap[parent] {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
			i = parent
			continue
		}
		break
	}
}

// Max returns value of the max element in the heap
// Returns false if the heap is empty
func (h *Heap) Max() (int, bool) {
	if h.size == 0 {
		return 0, false
	}
	return h.heap[0], true
}

// ExtractMax returns value of the max element and removes it
// Returns false if the heap is empty
func (h *Heap) ExtractMax() (int, bool) {
	if h.size == 0 {
		return 0, false
	}

	max := h.heap[0]
	h.heap[0] = h.heap[h.size-1]
	h.heap = h.heap[:h.size-1]
	h.size--

	h.heapify(0)
	return max, true
}

// ConvertToSlice returns sorted slice with all elements of the heap.
// The heap will be empty in the end
func (h *Heap) ConvertToSlice() []int {
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

	var heapLayer, newLineInd int
	for i, el := range h.heap {
		sb.WriteString(strconv.Itoa(el))
		if i == newLineInd && i != h.size-1 {
			sb.WriteRune('\n')
			heapLayer++
			newLineInd += 2 * heapLayer
			continue
		}
		if i != h.size-1 {
			sb.WriteRune(' ')
		}
	}

	sb.WriteString("]")
	return sb.String()
}

// normalize vertex i
func (h *Heap) heapify(i int) {
	for {
		// find biggest among element i and his children
		biggest := i
		left := 2*i + 1
		right := 2 * (i + 1)
		if left < h.size && h.heap[i] < h.heap[left] {
			biggest = left
		}

		if right < h.size && h.heap[biggest] < h.heap[right] {
			biggest = right
		}

		if biggest == i {
			break // there is nothing more to normalize
		}

		// change biggest element with element i
		// and proceed the same procedure for new index
		h.heap[biggest], h.heap[i] = h.heap[i], h.heap[biggest]
		i = biggest
	}
}
