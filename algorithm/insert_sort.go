package algorithm

// DirectInsertSort 2022-03-20 尝试更改为泛型算法
// DirectInsertSort 直接插入排序
func DirectInsertSort(source []int) {
	var elem, j int
	for i := 1; i < len(source); i++ {
		for j = i; j > 0; j-- {
			if source[j] < source[j-1] {
				elem = source[j]
				source[j] = source[j-1]
				source[j-1] = elem
			} else {
				break
			}
		}
	}
}

func BinarySearch[T int](source []T, target T) int {
	h, l := len(source), 0
	for l < h {
		mid := (h + l) / 2
		if source[mid] == target {
			return mid
		} else if source[mid] > target {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return h
}

func BinaryInsertSort(source []int) {
	var elem, j, k int
	for i := 1; i < len(source); i++ {
		elem = source[i]
		j = BinarySearch[int](source[:i], elem)
		// j是elem要插入的位置
		for k = i; k > j; k-- {
			source[k] = source[k-1]
		}
		source[j] = elem
	}
}
