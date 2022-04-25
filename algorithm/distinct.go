package algorithm

import "sort"

func Distinct(src []int) []int {
	sort.Ints(src)
	k := 0
	for idx, e := 1, len(src); idx < e; idx++ {
		if src[idx] == src[idx-1] {
			k++
		} else {
			src[idx-k] = src[idx]
		}
	}
	return src[:len(src)-k]
}
