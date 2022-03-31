package nowcoder

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		fmt.Println("\tInput : [1 2] [1 4 5] [1]")
		fmt.Println("\tExpect: [1 1 1 2 4 5]")
		l1 := CreateList([]int{1, 2})
		l2 := CreateList([]int{1, 4, 5})
		l3 := CreateList([]int{1})
		l4 := MergeKLists([]*ListNode{l1, l2, l3})
		fmt.Println("\tOutput:", ToSlice(l4))
	})

	t.Run("Case2", func(t *testing.T) {
		fmt.Println("\tInput : [] [1 3 5] [2 4 6]")
		fmt.Println("\tExpect: [1 2 3 4 5 6]")
		l2 := CreateList([]int{1, 3, 5})
		l3 := CreateList([]int{2, 4, 6})
		l4 := MergeKLists([]*ListNode{nil, l2, l3})
		fmt.Println("\tOutput:", ToSlice(l4))
	})

	t.Run("Case3", func(t *testing.T) {
		fmt.Println("\tInput : [] [] []")
		fmt.Println("\tExpect: []")
		l4 := MergeKLists([]*ListNode{nil, nil, nil})
		fmt.Println("\tOutput:", ToSlice(l4))
	})
}
