package nowcoder

import (
	"unsafe"
)

// link: https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46
// ACCEPT

func FindFirstCommonNodeViaHash(pHead1, pHead2 *ListNode) *ListNode {
	// 哈希法，空间复杂度高
	// ACCEPT
	table := make(map[uintptr]bool)
	var ret *ListNode
	for pHead1 != nil {
		table[uintptr(unsafe.Pointer(pHead1))] = true
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		if _, ok := table[uintptr(unsafe.Pointer(pHead2))]; ok {
			ret = pHead2
			break
		}
		pHead2 = pHead2.Next
	}
	return ret
}

func FindFirstCommonNodeViaDoublePtr(pHead1, pHead2 *ListNode) *ListNode {
	var ptr1, ptr2 = pHead1, pHead2
	for ptr1 != ptr2 {
		if ptr1 == nil {
			ptr1 = pHead2
		} else {
			ptr1 = ptr1.Next
		}
		if ptr2 == nil {
			ptr2 = pHead1
		} else {
			ptr2 = ptr2.Next
		}
	}
	return ptr2
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	return FindFirstCommonNodeViaDoublePtr(pHead1, pHead2)
}
