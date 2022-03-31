package nowcoder

import (
	"testing"
)

func TestReverseLinkedListKGroup(t *testing.T) {
	l := CreateList([]int{1, 2, 3, 4, 5})
	t.Log(ToSlice(l))
	l = ReverseKGroup(l, 2)
	t.Log(ToSlice(l))
}
