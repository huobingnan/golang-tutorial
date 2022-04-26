package nowcoder

// link: https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5

func _searchInversePairs(data []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	res := _searchInversePairs(data, left, mid) + _searchInversePairs(data, mid+1, right)
	res = res % 1000000007
	for i, j := left, mid+1; i <= mid && j <= right; {
		if data[i] > data[j] {
			res, j = res+1, j+1
		} else {
			i++
		}
	}
	return res
}

func InversePairs(data []int) int {
	// write code here
	return _searchInversePairs(data, 0, len(data)-1)
}
