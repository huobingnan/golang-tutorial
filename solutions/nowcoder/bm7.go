package nowcoder

import (
	"unsafe"
)

// link: https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4
// ACCEPT

func Bm7HashSolution(pHead *ListNode) *ListNode {
	var ret *ListNode
	table := make(map[uintptr]bool)
	for pHead != nil {
		ptr := uintptr(unsafe.Pointer(pHead))
		if _, ok := table[ptr]; ok {
			ret = (*ListNode)(unsafe.Pointer(ptr))
			break
		} else {
			table[ptr] = true
		}
		pHead = pHead.Next
	}
	return ret
}

func Bm7FastAndSlowPtr(pHead *ListNode) *ListNode {
	var ret *ListNode
	var fast, slow = pHead, pHead
	for slow != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == slow {
			break
		}
	}
	fast = pHead
	for slow != nil {
		if fast == slow {
			ret = fast
			break
		}
		slow = slow.Next
		fast = fast.Next
	}
	return ret
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	return Bm7FastAndSlowPtr(pHead)
}
