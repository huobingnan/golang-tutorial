package nowcoder

import "testing"

func TestBm13IsPail(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("\tInput : [1 2 2 1]")
		t.Log("\tExpect: true")
		l := CreateList([]int{1, 2, 2, 1})
		ret := Bm13IsPail(l)
		t.Log("\tOutput:", ret)
	})
}
