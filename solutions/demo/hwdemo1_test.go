package demo

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		ret := Solution([]int{2, 3, 1, 1, 4})
		t.Log(ret)
	})
	t.Run("Case2", func(t *testing.T) {
		ret := Solution([]int{3, 2, 1, 0, 4})
		t.Log(ret)
	})
	t.Run("Case3", func(t *testing.T) {
		ret := Solution([]int{2, 3, 2, 1, 4})
		t.Log(ret)
	})
	t.Run("Case4", func(t *testing.T) {
		t.Log(Solution([]int{0, 8, 2, 0, 9}))
	})
}
