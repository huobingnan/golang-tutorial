package nowcoder

// link: https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	var next, head, ptr *ListNode
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val >= pHead2.Val {
			next = pHead2
			pHead2 = pHead2.Next
		} else {
			next = pHead1
			pHead1 = pHead1.Next
		}
		if head == nil {
			head = next
			ptr = head
		} else {
			ptr.Next = next
			ptr = next
		}
	}
	for pHead1 != nil {
		if head == nil {
			head = pHead1
			ptr = head
		} else {
			ptr.Next = pHead1
			ptr = ptr.Next
		}
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		if head == nil {
			head = pHead2
			ptr = head
		} else {
			ptr.Next = pHead2
			ptr = ptr.Next
		}
		pHead2 = pHead2.Next
	}
	return head
}
