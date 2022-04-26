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

// 递归形式的归并排序的实现
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

// 对于递归形式的归并排序的一点改进
func _doMergeSortV2(src []int, left, right int, temp []int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	_doMergeSortV2(src, left, mid, temp)
	_doMergeSortV2(src, mid+1, right, temp)
	// copy source to temp
	copy(temp[left:right+1], src[left:right+1])
	// merge
	i, j, k := left, mid+1, 0
	for k = left; k <= right; k++ {
		if i == mid+1 {
			src[k], j = temp[j], j+1
		} else if j == right+1 || temp[i] <= temp[j] {
			src[k], i = temp[i], i+1
		} else {
			src[k], j = temp[j], j+1
		}
	}
}

func MergeSortV2(src []int) {
	_doMergeSortV2(src, 0, len(src)-1, make([]int, len(src)))
}
