package main

import (
	"fmt"
	"sync"
)

func RunWithoutMutexProtect() {
	number := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			number += 1
		}
		wait.Done()
	}(&wg)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			number -= 1
		}
		wait.Done()
	}(&wg)
	wg.Wait()
	fmt.Println("number = ", number)
}

func RunWithChannel() {
	number := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		number <- 0
		fmt.Println("number was push into channel")
	}()
	go func(number chan int, wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			cpy := <-number
			cpy += 1
			number <- cpy
		}
		wait.Done()
	}(number, &wg)

	go func(number chan int, wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			cpy := <-number
			cpy -= 1
			number <- cpy
		}
		wait.Done()
	}(number, &wg)
	wg.Wait()
	fmt.Println("number = ", <-number)
	close(number)
}

func main() {
	// RunWithoutMutexProtect()
	RunWithChannel()
}
