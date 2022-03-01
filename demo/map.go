package main

import "fmt"

func main() {
	// 创建map的两种方式
	// 1. 通过字面量的形式初始化
	m1 := map[string]int{
		"Bob":   120,
		"Jenny": 100,
	}
	// 2. 通过make来创建
	_ = make(map[string]int)
	// map put element
	m1["Danny"] = 130
	// map delete element, no return
	delete(m1, "Danny")
	// is map contains key?
	if v, ok := m1["Bob"]; ok {
		fmt.Printf("map contains key 'Bob', value is %d\n", v)
	} else {
		fmt.Printf("map doesn't contains key 'Bob'\n")
	}
	func(m map[string]int) {
		// map作为函数参数传递时是引用传递
		// 所以在函数体内部的对于传入map的操作是会影响函数体外部的实际变量
		delete(m, "Jenny")
	}(m1)
	if _, ok := m1["Bob"]; ok {
		fmt.Println("map contains key 'Bob'")
	}
	for k, v := range m1 {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}

}
