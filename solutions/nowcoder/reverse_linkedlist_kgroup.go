package nowcoder

import (
	"errors"
)

// link: https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e
// !AC

func findKthNode(start *ListNode, k int) (*ListNode, error) {
	if start == nil || k < 0 {
		return nil, errors.New("error")
	}
	for k != 0 && start != nil {
		start = start.Next
		k -= 1
	}
	if start == nil || k != 0 {
		return nil, errors.New("error")
	} else {
		return start, nil
	}
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	var prev, cur, temp, s, ret *ListNode
	cur, s, prev = head, head, head
	if cur != nil {
		cur = cur.Next
	} else {
		return cur
	}
	for {
		e, err := findKthNode(cur, k-1)
		if err != nil {
			s.Next = cur
			break
		}
		s.Next = e
		s = cur
		// reverse
		for prev != e {
			temp = cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}
		// head pointer
		if ret == nil {
			ret = e
		}
	}
	return ret
}
