package nowcoder

import "testing"

func TestBm16DeleteDuplicates(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("\tInput : [1 2 2]")
		t.Log("\tExpect: [1]")
		l := Bm16DeleteDuplicates(CreateList([]int{1, 2, 2}))
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 2", func(t *testing.T) {
		t.Log("\tInput : [1 2 3 3 3 4 4]")
		t.Log("\tExpect: [1 2]")
		l := Bm16DeleteDuplicates(CreateList([]int{1, 2, 3, 3, 3, 4, 4}))
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 4", func(t *testing.T) {
		t.Log("\tInput : [1 1 1 1 2 3 3 4 5 5]")
		t.Log("\tExpect: [2 4]")
		l := Bm16DeleteDuplicates(CreateList([]int{1, 1, 1, 1, 2, 3, 3, 4, 5, 5}))
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 5", func(t *testing.T) {
		t.Log("\tInput : [1 2]")
		t.Log("\tExpect: [1 2]")
		l := Bm16DeleteDuplicates(CreateList([]int{1, 2}))
		t.Log("\tOutput:", ToSlice(l))
	})
}
