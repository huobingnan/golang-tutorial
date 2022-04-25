package algorithm

func _doMerge(src []int, left, right, mid int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	// merge
	for ; i <= mid && j <= right; k++ {
		if src[i] < src[j] {
			temp[k] = src[i]
			i++
		} else {
			temp[k] = src[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		temp[k] = src[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		temp[k] = src[j]
	}
	// copy
	copy(src[left:right+1], temp)
}

func _doMergeSort(src []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	_doMergeSort(src, left, mid)
	_doMergeSort(src, mid+1, right)
	_doMerge(src, left, right, mid)
}

func MergeSort(src []int) {
	_doMergeSort(src, 0, len(src)-1)
}
