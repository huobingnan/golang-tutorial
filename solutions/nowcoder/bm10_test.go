package nowcoder

import (
	"fmt"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		fmt.Println("\tInput : [1 2 3] [4 5] [6 7]")
		fmt.Println("\tExpect: [6 7]")
		l1 := CreateList([]int{1, 2, 3})
		l2 := CreateList([]int{4, 5})
		l3 := CreateList([]int{6, 7})
		ptr := l1
		for ptr != nil && ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = l3
		ptr = l2
		for ptr != nil && ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = l3
		l4 := FindFirstCommonNode(l1, l2)
		fmt.Println("\tOutput:", ToSlice(l4))
	})

	t.Run("Case2", func(t *testing.T) {
		fmt.Println("\tInput : [1 2 3] [4 5 6] []")
		fmt.Println("\tExpect: []")
		l1 := CreateList([]int{1, 2, 3})
		l2 := CreateList([]int{4, 5, 6})
		l4 := FindFirstCommonNode(l1, l2)
		fmt.Println("\tOutput:", ToSlice(l4))
	})

	t.Run("Benchmark for LCM", func(t *testing.T) {
		l1 := CreateList([]int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21})
		l2 := CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		l3 := CreateList([]int{23, 24})
		fmt.Println("\tInput :", ToSlice(l1), ToSlice(l2), ToSlice(l3))
		fmt.Println("\tExpect: [23 24]")
		ptr := l1
		for ptr != nil && ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = l3
		ptr = l2
		for ptr != nil && ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = l3
		l4 := FindFirstCommonNode(l1, l2)
		fmt.Println("\tOutput:", ToSlice(l4))
	})
}
