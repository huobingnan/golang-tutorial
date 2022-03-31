package nowcoder

func doPrevorderTraversal(root *TreeNode, collector *[]int) {
	if root == nil {
		return
	}
	*collector = append(*collector, root.Val)
	doPrevorderTraversal(root.Left, collector)
	doPrevorderTraversal(root.Right, collector)
}

func PreorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0, 10)
	doPrevorderTraversal(root, &ret)
	return ret
}
