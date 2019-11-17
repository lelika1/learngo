package segtree_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/segtree"
)

func TestNewZeroCountTree(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			[]int{0, 0, -2, 0, 5, 6, 8, 0},
			tree(
				"4",
				"3 1",
				"2 1 0 1",
				"1 1 0 1 0 0 0 1",
			),
		},
		{
			[]int{0, 0, 0, 4, 0, 0},
			tree(
				"5",
				"3 2",
				"2 1 2 -",
				"1 1 1 0 1 1",
			),
		},
		{
			[]int{0, 0, 1, 4, 0, 6},
			tree(
				"3",
				"2 1",
				"2 0 1 -",
				"1 1 0 0 1 0",
			),
		},
		{
			[]int{0, 0, 3, 10},
			tree(
				"2",
				"2 0",
				"1 1 0 0",
			),
		},
		{
			[]int{3, 0, 0},
			tree(
				"2",
				"1 1",
				"0 1 1",
			),
		},
		{
			[]int{1, 0},
			tree(
				"1",
				"0 1",
			),
		},
		{
			[]int{1},
			tree(
				"0",
			),
		},
		{
			[]int{0},
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
		tree := segtree.NewZeroCountTree(ts.input)
		got := tree.String()

		if got != ts.want {
			t.Errorf("NewZeroCountTree(%v) = %q, want %q", ts.input, got, ts.want)
		}
	}
}

func TestZeroCount(t *testing.T) {
	tests := []struct {
		input []int
		i, j  int
		want  int
		fail  bool
	}{
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     0, j: 7,
			want: 4,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     0, j: 5,
			want: 3,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     0, j: 4,
			want: 2,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     0, j: 1,
			want: 1,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     0, j: 3,
			want: 2,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     4, j: 5,
			want: 1,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     2, j: 6,
			want: 2,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0, 8, 0},
			i:     3, j: 4,
			want: 0,
		},
		{
			input: []int{2, 0, 0, 4, 5, 10, 8},
			i:     3, j: 6,
			want: 0,
		},
		{
			input: []int{2, 0, 0, 4, 5, 0},
			i:     5, j: 5,
			want: 1,
		},
		{
			input: []int{0, 0, 0, 0, 0, 0},
			i:     0, j: 5,
			want: 6,
		},
		{
			input: []int{0, 0},
			i:     0, j: 1,
			want: 2,
		},
		{
			input: []int{0},
			i:     0, j: 0,
			want: 1,
		},
		{
			input: []int{1},
			i:     0, j: 0,
			want: 0,
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
		tree := segtree.NewZeroCountTree(ts.input)
		got, ok := tree.ZeroCount(ts.i, ts.j)

		if ok == ts.fail {
			// ZeroCount failed, but test doesn't expect this (or otherwise)
			t.Errorf("ZeroCount(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}

		if got != ts.want {
			t.Errorf("ZeroCount(%v; [%v, %v]) = %v, want %v", ts.input, ts.i, ts.j, got, ts.want)
		}
	}
}

func TestZeroCountTreeSet(t *testing.T) {
	tests := []struct {
		input    []int
		idx, val int
		want     string // ZeroCountTree.String() or '?' if idx is wrong
	}{
		{
			input: []int{2, 0, -2, 0, -5, 0},
			idx:   4, val: 5,
			want: tree(
				"3",
				"2 1",
				"1 1 1 -",
				"0 1 0 1 0 1",
			),
		},
		{
			input: []int{2, 0, -2, 0, 5},
			idx:   0, val: 0,
			want: tree(
				"3",
				"3 0",
				"2 1 0 -",
				"1 1 0 1 0",
			),
		},
		{
			input: []int{2, 3, -2, 4, 5},
			idx:   2, val: 0,
			want: tree(
				"1",
				"1 0",
				"0 1 0 -",
				"0 0 1 0 0",
			),
		},
		{
			input: []int{0, 0, -2, 4, 0},
			idx:   4, val: 1,
			want: tree(
				"2",
				"2 0",
				"2 0 0 -",
				"1 1 0 0 0",
			),
		},
		{
			input: []int{0, 0, 0},
			idx:   1, val: 4,
			want: tree(
				"2",
				"1 1",
				"1 0 1",
			),
		},
		{
			input: []int{3, 4, 2},
			idx:   2, val: 0,
			want: tree(
				"1",
				"0 1",
				"0 0 1",
			),
		},
		{
			input: []int{0},
			idx:   0, val: 3,
			want: tree(
				"0",
			),
		},
		{
			input: []int{1},
			idx:   0, val: 0,
			want: tree(
				"1",
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
		tree := segtree.NewZeroCountTree(ts.input)
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
