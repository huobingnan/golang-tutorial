package main

import "fmt"

// 传递一个容量为10的int数组
// 数组传递时会发生拷贝
func passParamWithArray(temp [10]int) {
	for each := range temp {
		fmt.Printf("%d ", each)
	}
	fmt.Println()
}

func main() {
	// 声明一个大小为10的数组, 数组的初始值会自动的赋值为0
	var array [10]int
	// range 返回一个数组或者切片的索引值和元素值
	for idx, elem := range array {
		fmt.Printf("index = %d, element = %d\n", idx, elem)
	}
	// initialize array
	for idx := range array {
		array[idx] = idx*2 + 1
	}
	// 可以使用len获取数组的长度
	fmt.Printf("array's length = %d\n", len(array))
	// 使用cap获取数组的容量大小
	fmt.Printf("array's capacity = %d\n", cap(array))
	arrayCopy := array // array之间的赋值是copy, 并不会发生引用的传递
	arrayCopy[0] = 100 // 并不会影响原array的值
	fmt.Printf("element at index 0 = %d\n", array[0])
	fmt.Printf("array address => %p\n", &array[0])
	fmt.Printf("array copy address => %p\n", &arrayCopy[0])
	arrayRef := array[:] //获取array的切片，这个引用的是原数组
	arrayRef[0] = 100
	fmt.Printf("original array is %v\n", array)
	fmt.Printf("array address => %p\n", &array[0])
	fmt.Printf("array reference address => %p\n", &arrayRef[0])

	// 获取array的切片
	slice := array[1:]
	fmt.Printf("slice from index 1 to tail => %v\n", slice)
	slice = array[2:3] // [2, 3) 左闭右开
	fmt.Printf("slice from index 2 to 3 => %v\n", slice)

	passParamWithArray(array)
}
