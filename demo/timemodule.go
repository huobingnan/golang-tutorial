package main

import (
	"fmt"
	"time"
)

func SetInterval(action func(), duration time.Duration) *time.Ticker {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				action()
			}
		}
	}()
	return ticker
}

func main() {
	ticker := SetInterval(func() {
		fmt.Println("Hello world")
	}, 500*time.Millisecond)
	time.Sleep(2 * time.Second)
	ticker.Stop()
}
