func maxProduct(nums []int) int {
	curMax, curMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := curMax
		curMax = max(nums[i], max(tmp*nums[i], curMin*nums[i]))
		curMin = min(nums[i], min(tmp*nums[i], curMin*nums[i]))
		if curMax > ans {
			ans = curMax
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}