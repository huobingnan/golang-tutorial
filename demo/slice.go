package main

import "fmt"

// golang slice demo
// 这里是中文的注释

func printSliceEachElementAddress(slice []int) {
	fmt.Printf("\t[ ")
	for idx := range slice {
		fmt.Printf("%p ", &slice[idx])
	}
	fmt.Printf("]\n")
}

func sliceExpand() {
	var slice []int
	slice = make([]int, 3, 5)
	for idx := range slice {
		slice[idx] = idx * idx
	}
	fmt.Printf("head address: %p, len: %d, capacity: %d\n", &slice[0], len(slice), cap(slice))
	printSliceEachElementAddress(slice)
	slice = append(slice, 10)
	slice = append(slice, 100)
	fmt.Printf("head address: %p, len: %d, capacity: %d\n", &slice[0], len(slice), cap(slice))
	printSliceEachElementAddress(slice)
	// expand occur
	// create a new memory space and copy original data to new memory space
	slice = append(slice, 1000)
	fmt.Printf("head address: %p, len: %d, capacity: %d\n", &slice[0], len(slice), cap(slice))
	printSliceEachElementAddress(slice)
}

// pass slice as a param for function. Actually this way is pass a reference
func passBySlice(slice []int) {
	if slice != nil && len(slice) > 0 {
		slice[0] = 100
	}
}

func main() {

	sliceExpand()
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice = %v\n", s)
	passBySlice(s)
	fmt.Printf("slice = %v\n", s)

	lowLevelArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array: %v, len: %d, capacity: %d\n", lowLevelArray, len(lowLevelArray), cap(lowLevelArray))
	s1 := lowLevelArray[:3]                                             // get a slice from index 0 until to 3(except index 3)
	fmt.Printf("s1: %v, len: %d, capacity: %d\n", s1, len(s1), cap(s1)) // len: 3 capacity: 5
	s1 = append(s1, 100)
	fmt.Printf("array: %v, len: %d, capacity: %d\n", lowLevelArray, len(lowLevelArray), cap(lowLevelArray))
	fmt.Printf("s1: %v, len: %d, capacity: %d\n", s1, len(s1), cap(s1)) // len: 4  capacity: 5
	s1 = append(s1, 1000)
	s1 = append(s1, 10000)
	s1 = append(s1, 1000000)
	fmt.Printf("array: %v, len: %d, capacity: %d\n", lowLevelArray, len(lowLevelArray), cap(lowLevelArray))
	fmt.Printf("s1: %v, len: %d, capacity: %d\n", s1, len(s1), cap(s1)) // len: 4  capacity: 5

	func() {
		var slice []int // slice 结构体
		// type slice {
		//	len int
		// 	cap int
		//  array [0]int
		//}
		fmt.Printf("slice addr : %p slice len : %d slice cap : %d\n", &slice, len(slice), cap(slice))
		slice = append(slice, 0)
		fmt.Printf("slice addr : %p slice len : %d slice cap : %d\n", &slice[0], len(slice), cap(slice))
		slice = append(slice, 1)
		fmt.Printf("slice addr : %p slice len : %d slice cap : %d\n", &slice[0], len(slice), cap(slice))
		slice = append(slice, 2)
		fmt.Printf("slice addr : %p slice len : %d slice cap : %d\n", &slice[0], len(slice), cap(slice))
	}()
}
