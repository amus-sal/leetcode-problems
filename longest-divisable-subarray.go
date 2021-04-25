func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums) // nlogn
	max, l, dp := 1, len(nums), make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}
	res, prev := []int{}, -1
	for i := l - 1; i >= 0; i-- {
		if dp[i] == max && (prev == -1 || prev%nums[i] == 0) {
			res = append(res, nums[i])
			prev = nums[i]
			max--
		}
	}
	return res

}

