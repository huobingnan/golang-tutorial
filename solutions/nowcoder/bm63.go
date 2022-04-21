package nowcoder

// link: https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4
// Accepted

func JumpFloor(number int) int {
	dp := make([]int, number+1)
	dp[0] = 1
	dp[1] = 1
	for floor := 2; floor <= number; floor++ {
		dp[floor] = dp[floor-1] + dp[floor-2]
	}
	return dp[number]
}
