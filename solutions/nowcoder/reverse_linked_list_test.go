package nowcoder

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	l := CreateList([]int{1, 2, 3, 4, 5})
	t.Log(ToSlice(l))
	l = ReverseList(l)
	t.Log(ToSlice(l))
}
