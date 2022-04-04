package nowcoder

// link: https://www.nowcoder.com/practice/1291064f4d5d4bdeaefbf0dd47d78541t
// ACCEPT

func doPostorderTraversal(root *TreeNode, collector *[]int) {
	if root == nil {
		return
	}
	doPostorderTraversal(root.Left, collector)
	doPostorderTraversal(root.Right, collector)
	*collector = append(*collector, root.Val)
}

func PostorderTraversal(root *TreeNode) []int {
	// write code here
	ret := make([]int, 0, 10)
	doPostorderTraversal(root, &ret)
	return ret
}
