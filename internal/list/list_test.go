package list_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/list"
)

func TestMergeList(t *testing.T) {
	var list1 list.List
	var list2 list.List
	mergedList := list.Merge(list1, list2)

	expected := "[]"
	got := mergedList.String()
	if got != expected {
		t.Errorf("Merge two empty list failed. Got:%v Expected: %v", got, expected)
	}
}

func TestListString(t *testing.T) {
	var list list.List
	expected := "[]"
	if list.String() != expected {
		t.Errorf("List.String for empty list failed. Got:%v Expected: %v", list.String(), expected)
	}

	list.Add(5)
	expected = "[5]"
	if list.String() != expected {
		t.Errorf("List.String failed. Got:%v Expected: %v", list.String(), expected)
	}

	list.Add(4)
	expected = "[5 4]"
	if list.String() != expected {
		t.Errorf("Add to list failed. Got:%v Expected: %v", list.String(), expected)
	}
}
