package main

import "fmt"

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	const generation int = 40
	number := Fib(generation)
	fmt.Println("This is Golang hello world demo")
	fmt.Printf("Fib(%d) = %d", generation, number)
}
