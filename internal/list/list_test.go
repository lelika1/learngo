package list_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/list"
)

func TestMerge(t *testing.T) {
	var list1 list.List
	var list2 list.List
	mergedList := list.Merge(list1, list2)

	want := "[]"
	got := mergedList.String()
	if got != want {
		t.Fatalf("Merge two empty list failed. Got:%v Want: %v", got, want)
	}

	list1.Add(2)
	mergedList = list.Merge(list1, list2)
	want = "[2]"
	got = mergedList.String()
	if got != want {
		t.Fatalf("Merge list failed. Got:%v Want: %v", got, want)
	}

	list2.Add(3)
	mergedList = list.Merge(list1, list2)
	want = "[2 3]"
	got = mergedList.String()
	if got != want {
		t.Fatalf("Merge list failed. Got:%v Want: %v", got, want)
	}
}

func TestString(t *testing.T) {
	var list list.List
	want := "[]"
	if list.String() != want {
		t.Fatalf("List.String for empty list failed. Got:%v Want: %v", list.String(), want)
	}

	list.Add(5)
	want = "[5]"
	if list.String() != want {
		t.Fatalf("List.String failed. Got:%v Want: %v", list.String(), want)
	}

	list.Add(4)
	want = "[5 4]"
	if list.String() != want {
		t.Fatalf("Add to list failed. Got:%v Want: %v", list.String(), want)
	}
}

func TestExtract(t *testing.T) {
	var l list.List
	l.Add(0)

	got, ok := l.Extract(0)
	if !ok || got != 0 {
		t.Fatalf("List.Extract from [0]-list failed.")
	}

	l.Add(0)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	got, ok = l.Extract(5)
	if ok {
		t.Fatalf("List.Extract non-existent element failed.")
	}

	got, ok = l.Extract(4)
	if !ok || got != 4 {
		t.Fatalf("List.Extract last element failed.")
	}

	got, ok = l.Extract(0)
	if !ok || got != 0 {
		t.Fatalf("List.Extract first element failed.")
	}

	got, ok = l.Extract(1)
	if !ok || got != 2 {
		t.Fatalf("List.Extract middle element failed.")
	}
}

func TestInsert(t *testing.T) {
	var l list.List
	l.Insert(0, 0)
	l.Insert(1, 2)
	l.Insert(1, 1)
	l.Insert(0, -1)
	l.Insert(2, 100)
	want := "[-1 0 100 1 2]"
	got := l.String()
	if got != want {
		t.Fatalf("List.Insert failed. Got:%v Want: %v", got, want)
	}
	ok := l.Insert(100, 5)
	if ok {
		t.Fatalf("List.Insert out-of-bounds failed.")
	}
}

func TestGet(t *testing.T) {
	var l list.List
	l.Add(0)
	l.Add(1)
	l.Add(2)
	l.Add(3)

	_, ok := l.Get(5)
	if ok {
		t.Fatalf("List.Get out-of-bounds failed.")
	}

	got, ok := l.Get(0)
	if !ok || got != 0 {
		t.Fatalf("List.Get first element failed.")
	}

	got, ok = l.Get(2)
	if !ok || got != 2 {
		t.Fatalf("List.Get from middle failed.")
	}

	got, ok = l.Get(3)
	if !ok || got != 3 {
		t.Fatalf("List.Get last element failed.")
	}
}

func TestIteration(t *testing.T) {
	var l list.List
	l.Add(0)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	iter := l.Begin()
	for i := 0; i < l.Size(); i++ {
		if iter == nil {
			t.Fatalf("List iteration failed on %v.", i)
		}
		if got := iter.Value(); got != i {
			t.Errorf("List iteration failed on %v: wrong value %v", i, got)
		}
		iter = iter.Next()
	}
	if iter != nil {
		t.Errorf("List iteration failed. Should reach the end.")
	}
}
