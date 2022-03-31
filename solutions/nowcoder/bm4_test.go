package nowcoder

import "testing"

func TestMerge(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		t.Log("\t\tInput : [1 3 5] [2 4 6]")
		t.Log("\t\tExpect: [1 2 3 4 5 6]")
		l1 := CreateList([]int{1, 3, 5})
		l2 := CreateList([]int{2, 4, 6})
		l3 := Merge(l1, l2)
		t.Log("\tOutput:", ToSlice(l3))
	})
	t.Run("Case2", func(t *testing.T) {
		t.Log("\t\tInput : [] []")
		t.Log("\t\tExpect: []")
		l3 := Merge(nil, nil)
		t.Log("\t\tOutput:", ToSlice(l3))
	})
}
