package nowcoder

// link: https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024
// ACCEPT

func Bm16DeleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = head
	var prev, ans, cur *ListNode
	prev = newHead
	ans = head
	cur = head.Next
	for cur != nil {
		// handle "tail node" boundary case
		if cur.Next == nil {
			if cur.Val == ans.Val {
				prev.Next = nil
			}
		}
		if ans.Val != cur.Val {
			if ans.Next != cur {
				prev.Next = cur
			} else {
				prev = ans
			}
			ans = cur
		}
		cur = cur.Next
	}
	return newHead.Next
}
