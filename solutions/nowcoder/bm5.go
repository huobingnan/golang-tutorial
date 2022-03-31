package nowcoder

import "container/heap"

// link: https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6
// ACCEPT

type ListNodeHeap []*ListNode

func (self ListNodeHeap) Len() int { return len(self) }

func (self ListNodeHeap) Less(i, j int) bool {
	if self[i] != nil && self[j] != nil {
		return self[i].Val < self[j].Val
	} else if self[i] == nil {
		return false
	} else {
		return true
	}
}

func (self ListNodeHeap) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func (self *ListNodeHeap) Push(x interface{}) { *self = append(*self, x.(*ListNode)) }

func (self *ListNodeHeap) Pop() interface{} {
	old := *self
	n := len(old)
	x := old[n-1]
	*self = old[0 : n-1]
	return x
}

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
