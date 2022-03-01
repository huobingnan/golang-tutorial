package test

import (
	"../algorithm"
	"testing"
)

func TestLinkedListNew(t *testing.T) {
	list := algorithm.NewList()
	t.Logf("list's size is %d\n", list.Size())
}

func TestLinkedListPush(t *testing.T) {
	list := algorithm.NewList()
	for i := 0; i < 15; i++ {
		list.Push(i * 2)
	}
}

func TestLinkedListPushAndPop(t *testing.T) {
	list := algorithm.NewList()
	for i := 0; i < 15; i++ {
		list.Push(2*i + 1)
	}
	for list.IsNotEmpty() {
		elem, _ := list.Pop()
		t.Logf("%d ", elem)
	}
}

func TestLinkedListPushAndShift(t *testing.T) {
	list := algorithm.NewList()
	for i := 0; i < 15; i++ {
		list.Push(i*2 + 1)
	}
	for list.IsNotEmpty() {
		elem, _ := list.Shift()
		t.Logf("%d ", elem)
	}
}

func TestLinkedListInsertAndRemove(t *testing.T) {
	list := algorithm.NewList()
	for i := 0; i < 15; i++ {
		list.Push(i*2 + 1)
	}
	list.Insert(100, 3)
	list.Insert(200, 5)
	var (
		elem interface{}
		err  error
	)
	elem, err = list.RemoveAt(5)
	if err != nil {
		t.Fatal("remove at 5")
	}
	t.Logf("index @ %d is %d", 5, elem)
	elem, err = list.RemoveAt(3)
	if err != nil {
		t.Fatal("remove at 3")
	}
	t.Logf("index @ %d is %d", 3, elem)
}
