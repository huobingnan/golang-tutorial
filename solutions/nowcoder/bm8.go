package nowcoder

// link: https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9
// ACCEPT

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	var fast, slow *ListNode
	fast, slow = pHead, pHead
	for fast != nil {
		fast = fast.Next
		if k > 0 {
			k--
			continue
		}
		slow = slow.Next
	}
	if k == 0 {
		return slow
	} else {
		return nil
	}
}
