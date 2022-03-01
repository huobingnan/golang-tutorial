package main

import "fmt"

// channel的底层是一种指针类型

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occur ", r)
		}
	}()
	// 这样的阻塞操作会在运行时抛出异常，但是编译期并不会检查
	// 不带缓冲的channel不能在同一个goroutine中使用，因为这样会造成死锁
	c := make(chan int) // 不带缓冲的channel
	c <- 1              // 由于channel不带缓冲，在向channel中放入一个元素后，goroutine便会阻塞，直到channel中的元素被取走
	_ = <-c             // 这段代码是不可达的，因为上一行代码会一直阻塞住直到channel中的元素被消费
}
