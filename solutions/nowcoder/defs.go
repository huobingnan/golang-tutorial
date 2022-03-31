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

type ListNodeHeap []*ListNode

func (self ListNodeHeap) Len() int { return len(self) }

func (self ListNodeHeap) Less(i, j int) bool {
	if self[i] != nil && self[j] != nil {
		return self[i].Val < self[j].Val
	} else if self[i] == nil {
		return false
	} else {
		return true
	}
}

func (self ListNodeHeap) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func (self *ListNodeHeap) Push(x interface{}) { *self = append(*self, x.(*ListNode)) }

func (self *ListNodeHeap) Pop() interface{} {
	old := *self
	n := len(old)
	x := old[n-1]
	*self = old[0 : n-1]
	return x
}
