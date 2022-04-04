package nowcoder

import (
	"fmt"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		fmt.Println("\tInput : [1 2 3 4 5] 2")
		fmt.Println("\tExpect: [4 5]")
		l1 := CreateList([]int{1, 2, 3, 4, 5})
		l2 := FindKthToTail(l1, 2)
		fmt.Println("\tOutput:", ToSlice(l2))
	})

	t.Run("Case2", func(t *testing.T) {
		fmt.Println("\tInput : [] 100")
		fmt.Println("\tExpect: []")
		l2 := FindKthToTail(nil, 100)
		fmt.Println("\tOutput:", ToSlice(l2))
	})

	t.Run("Case3", func(t *testing.T) {
		fmt.Println("\tInput : [1 2 3 4 5 6 7] 100")
		fmt.Println("\tExpect: []")
		l1 := CreateList([]int{1, 2, 3, 4, 5, 6, 7})
		l2 := FindKthToTail(l1, 100)
		fmt.Println("\tOutput:", ToSlice(l2))
	})

	t.Run("Case4", func(t *testing.T) {
		fmt.Println("\tInput : [1 2 3 4] 4")
		fmt.Println("\tExpect: [1 2 3 4]")
		l1 := CreateList([]int{1, 2, 3, 4})
		l2 := FindKthToTail(l1, 4)
		fmt.Println("\tOutput:", ToSlice(l2))
	})
}
