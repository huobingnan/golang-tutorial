package nowcoder

import "container/heap"

func SortInList(head *ListNode) *ListNode {
	// write code here
	// boundary case
	if head == nil {
		return nil
	}
	var h ListNodeHeap
	var rhead, ptr, min, next *ListNode
	heap.Init(&h)
	for head != nil {
		//NOTICE: If don't make `Next = nil`, it would  contribute to a ring
		next = head.Next
		head.Next = nil
		heap.Push(&h, head)
		head = next
	}
	size := h.Len()
	for i := 0; i < size; i++ {
		min = heap.Pop(&h).(*ListNode)
		if rhead == nil {
			rhead = min
			ptr = rhead
		} else {
			ptr.Next = min
			ptr = ptr.Next
		}
	}
	return rhead
}
