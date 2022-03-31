package nowcoder

import (
	"fmt"
	"testing"
)

func TestSortInList(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		fmt.Println("\tInput : [1 3 2 4 5]")
		fmt.Println("\tExpect: [1 2 3 4 5]")
		ret := SortInList(CreateList([]int{1, 3, 2, 4, 5}))
		fmt.Println("\tOutput:", ToSlice(ret))
	})

	t.Run("Case2", func(t *testing.T) {
		fmt.Println("\tInput : [-426572,-406609,724427,-157789,-132713,720732,-39078,-348926,-693458,559374,114739,-748249,428207,-767407,401563,646432,-682870,483610,-608888,94840,749542,359115,131739,935294,347325,80573,-869091,-757897,-860166,-227942,-484068,-170790,-362441,-860466,819197,817675,886101,463504,-100482,496406,-860791]")
		l := CreateList([]int{-426572, -406609, 724427, -157789, -132713, 720732, -39078, -348926, -693458, 559374, 114739, -748249, 428207, -767407, 401563, 646432, -682870, 483610, -608888, 94840, 749542, 359115, 131739, 935294, 347325, 80573, -869091, -757897, -860166, -227942, -484068, -170790, -362441, -860466, 819197, 817675, 886101, 463504, -100482, 496406, -860791})
		ret := SortInList(l)
		fmt.Println("\tOutput:", ToSlice(ret))
	})
}
