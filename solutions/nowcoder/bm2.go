package nowcoder

// link: https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c
// ACCEPT

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n {
		return head
	}

	var sp, ptr, prev, temp *ListNode
	ptr = head
	for i := 1; i < m; i++ {
		sp = ptr
		ptr = ptr.Next
	}
	if ptr == nil {
		return head
	} else {
		prev = ptr
		ptr = ptr.Next
	}
	// reverse
	for i := 0; i < n-m; i++ {
		temp = ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = temp
	}
	// boundary case handle
	if sp == nil {
		head.Next = ptr
		head = prev
	} else {
		sp.Next.Next = ptr
		sp.Next = prev
	}

	return head
}
