func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	output := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(max(dp[i-1], nums[i-1])+nums[i], nums[i])
		output = max(dp[i], output)
	}
	return output
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}