package nowcoder

// link: https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b
// ACCEPT

func bm11Process0(head1 *ListNode, head2 *ListNode) *ListNode {
	// 整型计算溢出, 输入的数据太大了
	// write code here
	var num1, num2, sum int
	var ret *ListNode
	for head1 != nil {
		num1 *= 10
		num1 += head1.Val
		head1 = head1.Next
	}
	for head2 != nil {
		num2 *= 10
		num2 += head2.Val
		head2 = head2.Next
	}
	sum = num1 + num2
	ret = new(ListNode)
	for {
		ret.Val = sum % 10
		sum /= 10
		if sum == 0 {
			break
		}
		next := new(ListNode)
		next.Next = ret
		ret = next
	}
	return ret
}

func AddInList(head1 *ListNode, head2 *ListNode) *ListNode {
	var prev, temp *ListNode
	var carry int
	// reverse list1
	for head1 != nil {
		temp = head1.Next
		head1.Next = prev
		prev = head1
		head1 = temp
	}
	head1 = prev
	// reverse list2
	prev = nil
	for head2 != nil {
		temp = head2.Next
		head2.Next = prev
		prev = head2
		head2 = temp
	}
	head2 = prev
	ret := new(ListNode)
	for {
		if head1 == nil && head2 == nil {
			break
		} else if head1 != nil && head2 != nil {
			ret.Val = (head1.Val + head2.Val + carry) % 10
			if head1.Val+head2.Val+carry >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			head1 = head1.Next
			head2 = head2.Next
		} else if head1 != nil {
			ret.Val = (head1.Val + carry) % 10
			if head1.Val+carry >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			head1 = head1.Next
		} else {
			ret.Val = (head2.Val + carry) % 10
			if head2.Val+carry >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			head2 = head2.Next
		}
		next := new(ListNode)
		next.Next = ret
		ret = next
	}
	if carry != 0 {
		// 最高位加完之后，依然有进位存在
		if ret.Val == 0 {
			ret.Val += carry
		} else {
			next := new(ListNode)
			next.Val += carry
			next.Next = ret
			ret = next
		}
	} else {
		if ret.Val == 0 {
			ret = ret.Next
		}
	}
	return ret
}
