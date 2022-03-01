package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					x := a*1000 + b*100 + c*10 + d*1
					y := 10000 + d*1000 + c*100 + b*10 + a*1
					if x*9 == y {
						fmt.Printf("A = %d B = %d C = %d D = %d\n", a, b, c, d)
					}
				}
			}
		}
	}
}
