package algorithm

func _doMerge(src []int, left, right, mid int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	// merge
	for i <= mid && j <= right {
		if src[i] < src[j] {
			temp[k] = src[i]
			i++
		} else {
			temp[k] = src[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ {
		temp[k] = src[i]
		k++
	}
	for ; j <= right; j++ {
		temp[k] = src[j]
		k++
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
