package nowcoder

import "math"

func _bm19GetHeight(k int, nums []int) int {
	if k < 0 || k >= len(nums) {
		return math.MinInt
	}
	return nums[k]
}

func _bm19Process1(nums []int) int {
	for idx, size := 0, len(nums); idx < size; idx++ {
		if nums[idx] > _bm19GetHeight(idx-1, nums) &&
			nums[idx] > _bm19GetHeight(idx+1, nums) {
			return idx
		}
	}
	return len(nums) - 1
}

func _bm19Process2(nums []int, start, end int) int {
	return 0
}

func FindPeakElement(nums []int) int {
	// write code here
	return _bm19Process1(nums)
}
