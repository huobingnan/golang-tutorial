package main

import (
	"fmt"
	"math/rand"
)

func RandomInt() int {
	fmt.Println("Random int function call")
	return rand.Intn(10)
}

func AcceptAndEcho(number int) {
	fmt.Println("Accept: ", number)
}

func DeferAndReturn() int {
	defer fmt.Println("After main function 1")
	defer fmt.Println("After main function 2")
	// defer function is an FIFO call form
	fmt.Println("Main function")
	defer fmt.Println("Before return")
	return RandomInt()
}

func main() {

}
