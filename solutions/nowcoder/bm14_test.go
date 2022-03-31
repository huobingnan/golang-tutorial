package nowcoder

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		t.Log("Input : [1 2 3 4 5 6]")
		t.Log("Expect: [1 3 5 2 4 6]")
		l := CreateList([]int{1, 2, 3, 4, 5, 6})
		ret := OddEvenList(l)
		t.Log("Output:", ToSlice(ret))
	})

	t.Run("Case 2", func(t *testing.T) {
		t.Log("Input : [2 1 3 5 6 4 7]")
		t.Log("Expect: [2 3 6 7 1 5 4]")
		l := CreateList([]int{2, 1, 3, 5, 6, 4, 7})
		ret := OddEvenList(l)
		t.Log("Output:", ToSlice(ret))
	})

	t.Run("Case 3", func(t *testing.T) {
		t.Log("Input : [4 1]")
		t.Log("Expect: [4 1]")
		l := CreateList([]int{4, 1})
		ret := OddEvenList(l)
		t.Log("Output:", ToSlice(ret))
	})
}
