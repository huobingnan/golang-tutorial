package demo

// link: https://www.lintcode.com/problem/116/
// Accepted

func _findNext(slice []int, start, end int) int {
	if start >= len(slice) {
		start = len(slice)
	}
	if end > len(slice) {
		end = len(slice)
	}
	max := slice[start]
	index := start
	offset := 1
	for idx := start + 1; idx < end; idx++ {
		if slice[idx]+offset > max {
			index = idx
			max = slice[idx] + offset
		}
		offset += 1
	}
	return index
}

func Solution(A []int) bool {
	start := 0
	for {
		energy := A[start]
		if start+energy >= len(A)-1 {
			return true
		}
		if energy == 0 {
			return false
		}
		start = _findNext(A, start+1, start+energy+1)
	}
}
