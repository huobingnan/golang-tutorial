package nowcoder

// link: https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3
// ACCEPT

func OddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	var oddPtr, evenPtr, evenHead, oddHead *ListNode
	ptr := head
	index := 1
	for ptr != nil {
		if index%2 != 0 {
			// 奇数
			if oddHead == nil {
				oddHead = ptr
				oddPtr = ptr
			} else {
				oddPtr.Next = ptr
				oddPtr = ptr
			}
		} else {
			// 偶数
			if evenHead == nil {
				evenHead = ptr
				evenPtr = ptr
			} else {
				evenPtr.Next = ptr
				evenPtr = ptr
			}
		}
		ptr = ptr.Next
		index += 1
	}
	// 拼接起来
	evenPtr.Next = nil
	oddPtr.Next = evenHead
	return oddHead
}
