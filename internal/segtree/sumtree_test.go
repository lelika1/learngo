package segtree_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/lelika1/learngo/internal/segtree"
)

func pictTree(row string, rows ...string) string {
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

func TestNewSumtree(t *testing.T) {
	tests := []struct {
		arr  []int
		want string
	}{
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			pictTree(
				"27",
				"7 20",
				"5 2 11 9",
				"2 3 -2 4 5 6 8 1",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8},
			pictTree(
				"26",
				"7 19",
				"5 2 11 8",
				"2 3 -2 4 5 6 8",
			),
		},
		{
			[]int{2, 3, -2, 4, 5, 6},
			pictTree(
				"18",
				"7 11",
				"5 2 11 0",
				"2 3 -2 4 5 6",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			pictTree(
				"12",
				"7 5",
				"5 2 5 0",
				"2 3 -2 4 5",
			),
		},
		{
			[]int{-2, 2, 3, 10},
			pictTree(
				"13",
				"0 13",
				"-2 2 3 10",
			),
		},
		{
			[]int{3, 4, 2},
			pictTree(
				"9",
				"7 2",
				"3 4 2",
			),
		},
		{
			[]int{1, 2},
			pictTree(
				"3",
				"1 2",
			),
		},
		{
			[]int{1},
			pictTree(
				"1",
			),
		},
		{
			nil,
			pictTree(""),
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.arr)
		got := tree.String()

		if got != ts.want {
			t.Errorf("NewSumtree(%v) = %q, want %q", ts.arr, got, ts.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		arr  []int
		i, j int
		want string // sum or '?' if indexes are wrong
	}{
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			0, 7, "27",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			0, 5, "18",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			0, 4, "12",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			0, 1, "5",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			0, 3, "7",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			4, 5, "11",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			2, 6, "21",
		},
		{
			[]int{2, 3, -2, 4, 5, 6, 8, 1},
			3, 4, "9",
		},
		{
			[]int{2, 3},
			0, 0, "2",
		},
		{
			[]int{2},
			0, 0, "2",
		},
		{
			[]int{2, 3},
			-1, 1, "?",
		},
		{
			[]int{2, 3},
			0, 2, "?",
		},
		{
			[]int{2, 3},
			1, 0, "?",
		},
		{
			nil, 0, 0, "?",
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.arr)
		sum, ok := tree.Sum(ts.i, ts.j)
		if !ok {
			if ts.want != "?" {
				t.Errorf("Sum(%v; [%v, %v]) = %v, want %v", ts.arr, ts.i, ts.j, sum, ts.want)
			}
			continue
		}

		if got := strconv.Itoa(sum); got != ts.want {
			t.Errorf("Sum(%v; [%v, %v]) = %v, want %v", ts.arr, ts.i, ts.j, sum, ts.want)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		arr      []int
		idx, val int
		want     string // SumTree.String() or '?' if idx is wrong
	}{
		{
			[]int{2, 3, -2, 4, 5},
			0, -2,
			pictTree(
				"8",
				"3 5",
				"1 2 5 0",
				"-2 3 -2 4 5",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			2, 0,
			pictTree(
				"14",
				"9 5",
				"5 4 5 0",
				"2 3 0 4 5",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			4, 1,
			pictTree(
				"8",
				"7 1",
				"5 2 1 0",
				"2 3 -2 4 1",
			),
		},
		{
			[]int{-2, 2, 3, 10},
			3, -3,
			pictTree(
				"0",
				"0 0",
				"-2 2 3 -3",
			),
		},
		{
			[]int{3, 4, 2},
			1, 4,
			pictTree(
				"9",
				"7 2",
				"3 4 2",
			),
		},
		{
			[]int{1},
			0, 3,
			pictTree(
				"3",
			),
		},
		{
			[]int{2, 3, -2, 4, 5},
			5, 5,
			"?",
		},
		{
			[]int{2, 3, -2, 4, 5},
			-1, 1,
			"?",
		},
		{
			nil,
			0, 2,
			"?",
		},
	}

	for _, ts := range tests {
		tree := segtree.NewSumTree(ts.arr)
		ok := tree.Update(ts.idx, ts.val)
		if !ok {
			if ts.want != "?" {
				t.Errorf("Update returns false. Want: want %q", ts.want)
			}
			continue
		}

		if got := tree.String(); got != ts.want {
			t.Errorf("Update(arr=%v, i=%v, v=%v) = %q, want %q",
				ts.arr, ts.idx, ts.val, got, ts.want)
		}
	}
}
