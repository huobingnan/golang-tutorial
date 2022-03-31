package nowcoder

import "testing"

func TestEntryNodeOfLoop(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		l1 := CreateList([]int{1, 2})
		l2 := CreateList([]int{3, 4, 5})
		l1.Next = l2
		ptr := l2
		for ptr != nil && ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = l2
		ret := EntryNodeOfLoop(l1)
		t.Log("\tOutput:", ret.Val)
	})
}
