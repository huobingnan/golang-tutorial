package demo

import "math"

func _minInt(nums []int) (int, int) {
	min := nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			index = i
		}
	}
	return min, index
}

func CoinsChange(candidate []int, target int) []int {
	dp := make([]int, target+1)
	state := make([]int, len(candidate))
	result := make([]int, 0, 5)
	dp[0] = 0
	for i := 1; i <= target; i++ {
		for idx, coin := range candidate {
			if i-coin < 0 {
				state[idx] = math.MaxInt
			} else {
				state[idx] = dp[i-coin]
			}
		}
		min, index := _minInt(state)
		if min == math.MaxInt {
			dp[i] = math.MaxInt
		} else {
			dp[i] = min + 1
			result = append(result, candidate[index])
		}
	}
	return result
}
