package nowcoder

import "math"

// link: https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e

func GetDis(A []int, n int) int {
	max := 0
	min := A[0]
	for i := 1; i < n; i++ {
		min = int(math.Min(float64(min), float64(A[i])))
		max = int(math.Max(float64(A[i]-min), float64(max)))
	}
	return max
}
