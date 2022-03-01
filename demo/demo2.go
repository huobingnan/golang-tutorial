package main

import (
	"fmt"
	"time"
)

var x [3]int

func a(xRef []int) {
	for {
		xRef[0]++
	}
}

func b() {
	fmt.Printf("%p\n", &x[0])
	for i := 0; i < 10; i++ {
		x[0]++
	}
	fmt.Printf("%p\n", &x[0])
}

func c() {
	for {
		x[0]++
	}
}

func main() {
	go a(x[:])
	go b()
	go c()
	fmt.Printf("%p\n", &x[0])
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(x)
}
