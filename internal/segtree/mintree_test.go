package segtree_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/segtree"
)

func TestNewMinTree(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			tree(
				"-2",
				"-2 1",
				"2 -2 5 1",
				"2 3 -2 4 5 6 8 1",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8},
			tree(
				"-2",
				"-2 5",
				"2 -2 5 8",
				"2 3 -2 4 5 6 8",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6},
			tree(
				"-2",
				"-2 5",
				"2 -2 5 -",
				"2 3 -2 4 5 6",
			),
		},
		{
			[]int{2, 3, -2, 4, -5, 6},
			tree(
				"-5",
				"-2 -5",
				"2 -2 -5 -",
				"2 3 -2 4 -5 6",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			tree(
				"-2",
				"-2 5",
				"2 -2 5 -",
				"2 3 -2 4 5",
			),
		},
		{
			[]int{-2, 2, 3, 10},
			tree(
				"-2",
				"-2 3",
				"-2 2 3 10",
			),
		},
		{
			[]int{3, 4, 2},
			tree(
				"2",
				"3 2",
				"3 4 2",
			),
		},
		{
			[]int{1, 2},
			tree(
				"1",
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
		tree := segtree.NewMinTree(ts.input)
		got := tree.String()

		if got != ts.want {
			t.Errorf("NewMinTree(%v) = %q, want %q", ts.input, got, ts.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		i, j  int
		want  int
		fail  bool
	}{
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 7,
			want: -2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 5,
			want: -2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 4,
			want: -2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 1,
			want: 2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     0, j: 3,
			want: -2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     4, j: 5,
			want: 5,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     2, j: 6,
			want: -2,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8, 1},
			i:     3, j: 4,
			want: 4,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6, 8},
			i:     6, j: 6,
			want: 8,
		},
		{
			input: []int{2, 3, -2, 4, 5, 6},
			i:     5, j: 5,
			want: 6,
		},
		{
			input: []int{2, 3, 2, 4, 5, 6},
			i:     0, j: 5,
			want: 2,
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
		tree := segtree.NewMinTree(ts.input)
		got, ok := tree.Min(ts.i, ts.j)

		if ok == ts.fail {
			// Min failed, but test doesn't expect this (or otherwise)
			t.Errorf("Min(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}

		if got != ts.want {
			t.Errorf("Min(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}
	}
}

func TestMinTreeSet(t *testing.T) {
	tests := []struct {
		input    []int
		idx, val int
		want     string // MinTree.String() or '?' if idx is wrong
	}{
		{
			input: []int{2, 3, -2, 4, -5, 6},
			idx:   4, val: 5,
			want: tree(
				"-2",
				"-2 5",
				"2 -2 5 -",
				"2 3 -2 4 5 6",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   0, val: -2,
			want: tree(
				"-2",
				"-2 5",
				"-2 -2 5 -",
				"-2 3 -2 4 5",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   2, val: 0,
			want: tree(
				"0",
				"0 5",
				"2 0 5 -",
				"2 3 0 4 5",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   4, val: 1,
			want: tree(
				"-2",
				"-2 1",
				"2 -2 1 -",
				"2 3 -2 4 1",
			),
		},
		{
			input: []int{-2, 2, 3, 10},
			idx:   3, val: -3,
			want: tree(
				"-3",
				"-2 -3",
				"-2 2 3 -3",
			),
		},
		{
			input: []int{3, 4, 2},
			idx:   1, val: 4,
			want: tree(
				"2",
				"3 2",
				"3 4 2",
			),
		},
		{
			input: []int{3, 4, 2},
			idx:   2, val: 4,
			want: tree(
				"3",
				"3 4",
				"3 4 4",
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
		tree := segtree.NewMinTree(ts.input)
		ok := tree.Set(ts.idx, ts.val)
		if !ok {
			if ts.want != "?" {
				t.Errorf("Set(arr=%v, i=%v, v=%v) = %v, want %q",
					ts.input, ts.idx, ts.val, ok, ts.want)
			}
			continue
		}

		if got := tree.String(); got != ts.want {
			t.Errorf("Set(arr=%v, i=%v, v=%v) = %q, want %q",
				ts.input, ts.idx, ts.val, got, ts.want)
		}
	}
}
