package nowcoder

func _bm18BinarySearch(nums []int, target int) int {
	l, h := 0, len(nums)
	for l < h {
		mid := (h + l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return h
}

// T: O(M * logN)
func _bm18Process1(target int, array [][]int) bool {
	for rowNumber, eachRow := range array {
		c := _bm18BinarySearch(eachRow, target)
		if c < len(array[rowNumber]) && array[rowNumber][c] == target {
			return true
		}
	}
	return false
}

func Find(target int, array [][]int) bool {
	// write code here
	// row, column
	return _bm18Process1(target, array)
}
