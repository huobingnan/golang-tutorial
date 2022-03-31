package nowcoder

import (
	"testing"
)

func TestReverseLinkedListSpecialRange(t *testing.T) {
	t.Run("Case 1: {1, 2, 3, 4, 5}; 2, 4", func(t *testing.T) {
		l := CreateList([]int{1, 2, 3, 4, 5})
		ReverseBetween(l, 2, 4)
		t.Log(ToSlice(l))
	})
	t.Run("Case 2: {5}; 1, 1", func(t *testing.T) {
		l := CreateList([]int{5})
		ReverseBetween(l, 1, 1)
		t.Log(ToSlice(l))
	})
}
