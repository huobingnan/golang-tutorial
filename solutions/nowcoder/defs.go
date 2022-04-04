package nowcoder

import (
	"container/list"
	"strconv"
	"strings"
)

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

// ParseTree 将输入的字符串解析为一个树
// 字符串以','作为分隔符，字符序列按照完全二叉树的顺序排列，null值用#表示
func ParseTree(input string) *TreeNode {
	split := strings.Split(input, ",")
	if split == nil || len(split) == 0 {
		panic("Parse error, split by ','")
	}
	parentQue := list.New()
	var parent, root, node *TreeNode
	var err error
	var tag = true // true: left false : right
	for _, val := range split {
		val = strings.TrimSpace(val)
		if val == "#" {
			node = nil
		} else {
			node = new(TreeNode)
			node.Val, err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
		// 检查parent
		if parentQue.Len() == 0 {
			if root == nil {
				root = node
			}
		} else {
			parent = parentQue.Front().Value.(*TreeNode)
			if tag {
				if parent != nil {
					parent.Left = node
				}
			} else {
				if parent != nil {
					parent.Right = node
				}
				parentQue.Remove(parentQue.Front())
			}
			tag = !tag
		}
		parentQue.PushBack(node)
	}
	return root
}

// TreeLevelOrderTraversal 二叉树的中序遍历序列
func TreeLevelOrderTraversal(root *TreeNode) []string {
	ret := make([]string, 0, 10)
	que := list.New()
	que.PushBack(root)
	for que.Len() != 0 {
		top := que.Remove(que.Front())
		if top == nil {
			ret = append(ret, "#")
		} else {
			node := top.(*TreeNode)
			if node != nil {
				que.PushBack(node.Left)
				que.PushBack(node.Right)
				ret = append(ret, strconv.Itoa(node.Val))
			} else {
				que.PushBack(nil)
				que.PushBack(nil)
				ret = append(ret, "#")
			}
		}

	}
	var end int
	for end = len(ret) - 1; end >= 0; end-- {
		if ret[end] != "#" {
			break
		}
	}
	return ret[:end+1]
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
