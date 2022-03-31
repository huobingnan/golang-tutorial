package nowcoder

import (
	"testing"
)

func TestRemoveNthFromEndLinkedList(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("Input: 1 -> 2 -> 3 -> 4 -> 5; n = 2")
		t.Log("Expect: [1, 2, 3, 5]")
		l := CreateList([]int{1, 2, 3, 4, 5})
		l = RemoveNthFromEnd(l, 2)
		t.Log(ToSlice(l))
	})

	t.Run("Case 2", func(t *testing.T) {
		t.Log("Input: 1 -> 2; n = 2")
		t.Log("Expect: [2]")
		l := CreateList([]int{1, 2})
		l = RemoveNthFromEnd(l, 2)
		t.Log(ToSlice(l))
	})
}
