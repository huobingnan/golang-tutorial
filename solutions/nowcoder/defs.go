package nowcoder

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(src []int) *ListNode {
	head := new(ListNode)
	ptr := head
	for _, each := range src[:len(src)-1] {
		ptr.Val = each
		ptr.Next = new(ListNode)
		ptr = ptr.Next
	}
	ptr.Val = src[len(src)-1]
	return head
}

func ToSlice(head *ListNode) []int {
	ret := make([]int, 0, 10)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
