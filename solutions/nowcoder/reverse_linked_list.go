package nowcoder

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var cur, prev, temp *ListNode
	cur = pHead
	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}
