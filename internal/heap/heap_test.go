package heap_test

import (
	"strings"
	"testing"

	"github.com/lelika1/learngo/internal/heap"
)

func pictHeap(row string, rows ...string) string {
	var sb strings.Builder
	sb.WriteString(row)
	for _, r := range rows {
		sb.WriteRune('\n')
		sb.WriteString(r)
	}
	return sb.String()
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNewHeap(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
			[]int{10, 8, 7, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{8},
			[]int{8},
		},
		{
			[]int{8, 3},
			[]int{8, 3},
		},
		{
			[]int{3, 8},
			[]int{8, 3},
		},
		// {
		// 	[]int{3, 3, 4, 5, 2, 8, 0, 8},
		// 	[]int{8, 8, 5, 4, 3, 3, 2, 0},
		// },
		{
			nil,
			nil,
		},
	}

	for _, ts := range tests {
		h := heap.NewHeap(ts.arr)
		got := h.ConvertToSlice()
		if !equal(got, ts.want) {
			t.Errorf("Heap(%v) = %v, want %v", ts.arr, got, ts.want)
		}
	}
}

func TestInsertMax(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int // max element at the moment
	}{
		{
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
			[]int{4, 4, 5, 7, 7, 7, 10, 10, 10},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{8},
			[]int{8},
		},
		{
			[]int{8, 3},
			[]int{8, 8},
		},
		{
			[]int{3, 8},
			[]int{3, 8},
		},
		// {
		// 	[]int{3, 3, 4, 5, 2, 8, 0, 8},
		// 	[]int{3, 3, 4, 5, 5, 8, 8, 8},
		// },
	}

	for _, ts := range tests {
		h := heap.NewHeap(nil)
		for i, v := range ts.arr {
			h.Insert(v)
			got, ok := h.Max()
			if !ok {
				t.Errorf("Max(%v) = (%v, %v), want %v", ts.arr[:i+1], got, ok, ts.want[i])
			}
			if got != ts.want[i] {
				t.Errorf("Max(%v) = %v, want %v", ts.arr[:i+1], got, ts.want[i])
			}
		}
	}
}

func TestEmptyHeap(t *testing.T) {
	h := heap.NewHeap(nil)
	max, ok := h.Max()
	if ok {
		t.Errorf("Max of empty heap should return false. Got=(%v; %v)", max, ok)
	}
	max, ok = h.ExtractMax()
	if ok {
		t.Errorf("ExtractMax in empty heap should return false. Got=(%v; %v)", max, ok)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		arr  []int
		want string
	}{
		{
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
			pictHeap(
				"[10",
				"8 7",
				"5 1 2 4",
				"3 0]",
			),
		},
		{
			[]int{1, 2, 3, 4, 5},
			pictHeap(
				"[5",
				"4 2",
				"1 3]",
			),
		},
		{
			[]int{8},
			"[8]",
		},
		{
			[]int{8, 3},
			pictHeap(
				"[8",
				"3]",
			),
		},
		{
			[]int{3, 8},
			pictHeap(
				"[8",
				"3]",
			),
		},
		{
			nil,
			"[]",
		},
	}

	for _, ts := range tests {
		h := heap.NewHeap(nil)
		for _, v := range ts.arr {
			h.Insert(v)
		}
		got := h.String()
		if got != ts.want {
			t.Errorf("String(%v) = %q, want %q", ts.arr, got, ts.want)
		}
	}
}
