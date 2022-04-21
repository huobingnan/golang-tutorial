package demo

import "testing"

func TestCoinsChange(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		res := CoinsChange([]int{20, 5, 7}, 27)
		t.Log("Result =", res)
	})
}
