package nowcoder

import (
	"testing"
)

func TestAddInList(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("Input : [9 3 7] [6 3]")
		t.Log("Expect: [1 0 0 0]")
		head1 := CreateList([]int{9, 3, 7})
		head2 := CreateList([]int{6, 3})
		ret := AddInList(head1, head2)
		t.Log("Output:", ToSlice(ret))
	})

	t.Run("Case 2", func(t *testing.T) {
		t.Log("Input : [0] [6 3]")
		t.Log("Expect: [6 3]")
		head1 := CreateList([]int{0})
		head2 := CreateList([]int{6, 3})
		ret := AddInList(head1, head2)
		t.Log("Output:", ToSlice(ret))
	})
}
