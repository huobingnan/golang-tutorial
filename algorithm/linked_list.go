package algorithm

import "errors"

type __Node struct {
	Data interface{}
	Next *__Node
}

type List struct {
	head *__Node
	tail *__Node
	size int
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list *List) Insert(element interface{}, index int) bool {
	if index > list.size || index < 0 {
		return false
	}
	node := new(__Node)
	node.Data = element
	if index == 0 {
		// head insert
		if list.head == nil {
			// initialize
			list.head = node
			list.tail = node
		} else {
			node.Next = list.head
			list.head = node
		}
	} else if index == list.size {
		// tail insert
		list.tail.Next = node
		list.tail = node
	} else {
		ptr := list.head
		for i := 1; i < index; i++ {
			ptr = ptr.Next
		}
		node.Next = ptr.Next
		ptr.Next = node
	}
	list.size++
	return true
}

func (list *List) Push(element interface{}) {
	list.Insert(element, list.size)
}

func (list *List) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of bound")
	}
	var ret interface{}
	if index == 0 {
		// head delete
		ret = list.head.Data
		list.head = list.head.Next
	} else {
		ptr := list.head
		for i := 1; i < index; i++ {
			ptr = ptr.Next
		}
		if index == list.size-1 {
			list.tail = ptr
		}
		ret = ptr.Next.Data
		ptr.Next = ptr.Next.Next

	}
	list.size--
	return ret, nil
}

func (list *List) Pop() (interface{}, error) {
	return list.RemoveAt(list.size - 1)
}

func (list *List) Shift() (interface{}, error) {
	return list.RemoveAt(0)
}

func (list *List) Size() int {
	return list.size
}

func (list *List) IsNotEmpty() bool {
	return list.size != 0
}

func (list *List) IsEmpty() bool {
	return list.size == 0
}

func BinarySearch(array []int, target int) int {
	// binary search
	return 0
}
