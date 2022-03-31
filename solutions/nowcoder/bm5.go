package nowcoder

import "container/heap"

// link: https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6
// ACCEPT

func MergeKLists(lists []*ListNode) *ListNode {
	// write code here
	if lists == nil {
		return nil
	}
	var head, ptr *ListNode
	var listNodeHeap ListNodeHeap
	heap.Init(&listNodeHeap)
	for _, ptr := range lists {
		heap.Push(&listNodeHeap, ptr)
	}
	for {
		min := heap.Pop(&listNodeHeap).(*ListNode)
		if min == nil {
			break
		}
		if head == nil {
			head = min
			ptr = head
		} else {
			ptr.Next = min
			ptr = ptr.Next
		}
		heap.Push(&listNodeHeap, ptr.Next)
	}
	return head
}
