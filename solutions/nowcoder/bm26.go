package nowcoder

import (
	"container/list"
	"math"
)

// LevelOrder 二叉树的层序遍历
func LevelOrder(root *TreeNode) [][]int {
	// write code here
	que := list.New()
	que.PushBack(root)
	ret := make([][]int, 0, 5)
	prevLevel, curLevel, curSize := -1, 0, 0
	for que.Len() != 0 {
		curSize += 1
		curLevel = int(math.Ceil(math.Log2(float64(curSize + 1))))
		if prevLevel != curLevel {
			cur := make([]int, 0, int(math.Pow(2, float64(curLevel-1))))
			ret = append(ret, cur)
			prevLevel = curLevel
		}
		r := que.Remove(que.Front()).(*TreeNode)
		if r != nil {
			ret[curLevel] = append(ret[curLevel], r.Val)
			que.PushBack(r.Left)
			que.PushBack(r.Right)
		}
	}
	return ret
}
