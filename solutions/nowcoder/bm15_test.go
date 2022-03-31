package nowcoder

import "testing"

func TestBm15DeleteDuplicates(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("\tInput : [1 1 2]")
		t.Log("\tExpect: [1 2]")
		l := CreateList([]int{1, 1, 2})
		l = Bm15DeleteDuplicates(l)
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 2", func(t *testing.T) {
		t.Log("\tInput : []")
		t.Log("\tExpect: []")
		l := Bm15DeleteDuplicates(nil)
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 3", func(t *testing.T) {
		t.Log("\tInput : [1 1 2 2 3 3 4 4 5 5 6 7 7 8]")
		t.Log("\tExpect: [1 2 3 4 5 6 7 8]")
		l := CreateList([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 7, 8})
		l = Bm15DeleteDuplicates(l)
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 4", func(t *testing.T) {
		t.Log("\tInput : [1 1]")
		t.Log("\tExpect: [1]")
		l := Bm15DeleteDuplicates(CreateList([]int{1, 1}))
		t.Log("\tOutput:", ToSlice(l))
	})
	t.Run("Case 5", func(t *testing.T) {
		t.Log("\tInput : [1]")
		t.Log("\tExpect: [1]")
		l := Bm15DeleteDuplicates(CreateList([]int{1}))
		t.Log("\tOutput:", ToSlice(l))
	})
}
