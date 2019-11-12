package heap_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/lelika1/learngo/internal/heap"
)

func pictHeap(row string, rows ...string) string {
	var sb strings.Builder
	sb.WriteRune('[')
	sb.WriteString(row)
	for _, r := range rows {
		sb.WriteRune('\n')
		sb.WriteString(r)
	}
	sb.WriteRune(']')
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
	tests := [][]int{
		[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
		[]int{3, 3, 4, 5, 2, 8, 0, 8},
		[]int{1, 2, 3, 4, 5},
		[]int{8, 7, 5, 3},
		[]int{3, 5, 7, 8},
		[]int{8, 3},
		[]int{3, 8},
		[]int{8},
		nil,
	}

	for _, ts := range tests {
		h := heap.NewHeap(ts)
		got := heap.Drain(h)

		want := make([]int, len(ts))
		copy(want, ts)
		sort.Slice(want, func(i, j int) bool { return want[i] > want[j] })

		if !equal(got, want) {
			t.Errorf("Heap(%v) = %v, want %v", ts, got, want)
		}
	}
}

func TestInsertMax(t *testing.T) {
	tests := []struct {
		// We add elements from `add`, one by one, to an initially empty heap,
		// and check that heap.Max() has a corresponding value from wantMax.
		add     []int
		wantMax []int
	}{
		{
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
			[]int{4, 4, 5, 7, 7, 7, 10, 10, 10},
		},
		{
			[]int{3, 3, 4, 5, 2, 8, 0, 8},
			[]int{3, 3, 4, 5, 5, 8, 8, 8},
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
	}

	for _, ts := range tests {
		h := heap.NewHeap(nil)
		for i, v := range ts.add {
			if len(ts.add) != len(ts.wantMax) {
				t.Fatalf("Test [add=%v, wantMax=%v] is incorrect.", ts.add, ts.wantMax)
			}

			h.Insert(v)
			got, ok := h.Max()
			if !ok {
				t.Errorf("Max(%v) = (%v, %v), want %v", ts.add[:i+1], got, ok, ts.wantMax[i])
			}
			if got != ts.wantMax[i] {
				t.Errorf("Max(%v) = %v, want %v", ts.add[:i+1], got, ts.wantMax[i])
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
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0, 12, 9, 6, 100, -5, 11},
			pictHeap(
				"100",
				"10 12",
				"5 9 7 11",
				"3 0 1 8 2 6 -5 4",
			),
		},
		{
			[]int{4, 3, 5, 7, 1, 2, 10, 8, 0},
			pictHeap(
				"10",
				"8 7",
				"5 1 2 4",
				"3 0",
			),
		},
		{
			[]int{3, 3, 4, 5, 2, 8, 0, 8},
			pictHeap(
				"8",
				"8 5",
				"4 2 3 0",
				"3",
			),
		},
		{
			[]int{1, 2, 3, 4, 5},
			pictHeap(
				"5",
				"4 2",
				"1 3",
			),
		},
		{
			[]int{8, 3},
			pictHeap(
				"8",
				"3",
			),
		},
		{
			[]int{3, 8},
			pictHeap(
				"8",
				"3",
			),
		},
		{
			[]int{8},
			pictHeap("8"),
		},
		{
			nil,
			pictHeap(""),
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
