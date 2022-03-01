package main

import (
	"fmt"
	"time"
)

// 交替打印
// 准备连个goroutine一个goroutine打印数字，另一个goroutine打印字母
// 一个goroutine打印两个数字之后，另一个goroutine开始打印字母，循环交替进行
func main() {
	printNumber := make(chan int, 1)
	printLetter := make(chan int, 1)
	printNumber <- 1
	go func() {
		for {
			<-printNumber
			fmt.Printf("1")
			fmt.Printf("2")
			time.Sleep(200 * time.Millisecond)
			printLetter <- 1
		}
	}()
	go func() {
		for {
			<-printLetter
			fmt.Printf("A")
			fmt.Printf("B")
			time.Sleep(200 * time.Millisecond)
			printNumber <- 1
		}
	}()
	select {}
}
