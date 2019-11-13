package segtree_test

import (
	"strings"
	"testing"

	"github.com/lelika1/learngo/internal/segtree"
)

func tree(row string, rows ...string) string {
	var sb strings.Builder
	sb.WriteString(row)
	for _, r := range rows {
		sb.WriteRune('\n')
		sb.WriteString(r)
	}
	return sb.String()
}

func TestNewSumtree(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			tree(
				"27",
				"7 20",
				"5 2 11 9",
				"2 3 -2 4 5 6 8 1",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8},
			tree(
				"26",
				"7 19",
				"5 2 11 8",
				"2 3 -2 4 5 6 8",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6},
			tree(
				"18",
				"7 11",
				"5 2 11 0",
				"2 3 -2 4 5 6",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			tree(
				"12",
				"7 5",
				"5 2 5 0",
				"2 3 -2 4 5",
			),
		},
		{
			[]int{-2, 2, 3, 10},
			tree(
				"13",
				"0 13",
				"-2 2 3 10",
			),
		},
		{
			[]int{3, 4, 2},
			tree(
				"9",
				"7 2",
				"3 4 2",
			),
		},
		{
			[]int{1, 2},
			tree(
				"3",
				"1 2",
			),
		},
		{
			[]int{1},
			tree(
				"1",
			),
		},
		{
			nil,
			tree(""),
		},
		{
			[]int{},
			tree(""),
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.input)
		got := tree.String()

		if got != ts.want {
			t.Errorf("NewSumtree(%v) = %q, want %q", ts.input, got, ts.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input []int
		i, j  int
		want  int
		fail  bool
	}{
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 7,
			want: 27,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 5,
			want: 18,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 4,
			want: 12,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 1,
			want: 5,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 3,
			want: 7,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     4, j: 5,
			want: 11,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     2, j: 6,
			want: 21,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     3, j: 4,
			want: 9,
		},
		{
			input: []int{2, 3},
			i:     0, j: 0,
			want: 2,
		},
		{
			input: []int{2},
			i:     0, j: 0,
			want: 2,
		},
		{
			input: []int{2, 3},
			i:     -1, j: 1,
			fail: true,
		},
		{
			input: []int{2, 3},
			i:     0, j: 2,
			fail: true,
		},
		{
			input: []int{2, 3},
			i:     1, j: 0,
			fail: true,
		},
		{
			input: nil,
			i:     0, j: 0,
			fail: true,
		},
		{
			input: []int{},
			i:     0, j: 0,
			fail: true,
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.input)
		got, ok := tree.Sum(ts.i, ts.j)

		if ok == ts.fail {
			// Sum failed, but test doesn't expect this (or otherwise)
			t.Errorf("Sum(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}

		if got != ts.want {
			t.Errorf("Sum(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		input    []int
		idx, val int
		want     string // SumTree.String() or '?' if idx is wrong
	}{
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   0, val: -2,
			want: tree(
				"8",
				"3 5",
				"1 2 5 0",
				"-2 3 -2 4 5",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   2, val: 0,
			want: tree(
				"14",
				"9 5",
				"5 4 5 0",
				"2 3 0 4 5",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   4, val: 1,
			want: tree(
				"8",
				"7 1",
				"5 2 1 0",
				"2 3 -2 4 1",
			),
		},
		{
			input: []int{-2, 2, 3, 10},
			idx:   3, val: -3,
			want: tree(
				"0",
				"0 0",
				"-2 2 3 -3",
			),
		},
		{
			input: []int{3, 4, 2},
			idx:   1, val: 4,
			want: tree(
				"9",
				"7 2",
				"3 4 2",
			),
		},
		{
			input: []int{1},
			idx:   0, val: 3,
			want: tree(
				"3",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   5, val: 5,
			want: "?",
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   -1, val: 1,
			want: "?",
		},
		{
			input: nil,
			idx:   0, val: 0,
			want: "?",
		},
		{
			input: []int{},
			idx:   0, val: 0,
			want: "?",
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.input)
		ok := tree.Set(ts.idx, ts.val)
		if !ok {
			if ts.want != "?" {
				t.Errorf("Set returns false. Want: want %q", ts.want)
			}
			continue
		}

		if got := tree.String(); got != ts.want {
			t.Errorf("Set(arr=%v, i=%v, v=%v) = %q, want %q",
				ts.input, ts.idx, ts.val, got, ts.want)
		}
	}
}
