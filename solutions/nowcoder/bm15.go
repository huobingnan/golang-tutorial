package nowcoder

// link: https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79
// ACCEPT

func Bm15DeleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	ans := head
	cur := head.Next
	for {
		if cur == nil {
			break
		}
		if cur.Val != ans.Val {
			ans.Next = cur
			ans = cur
		}
		if cur.Next == nil {
			// tail
			if ans.Val == cur.Val {
				ans.Next = cur.Next
			}
		}
		cur = cur.Next
	}
	return head
}
