package algorithm

import (
	"testing"
)

func TestDistinct(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		src := []int{3, 2, 1, 1, 1, 2, 4, 5, 7, 0}
		t.Log(Distinct(src))
	})
}
