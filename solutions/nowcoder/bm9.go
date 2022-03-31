package nowcoder

// link: https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	size := 0
	ncpy := n
	// 使用快慢指针的思想
	fast, slow := head, head
	for fast != nil {
		size += 1
		fast = fast.Next
		if n >= 0 {
			n--
			continue
		}
		slow = slow.Next
	}
	// 判断删除的是否是头结点
	if ncpy == size {
		return head.Next
	} else {
		slow.Next = slow.Next.Next
		return head
	}
}
