
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if Find(coins, i) {
					dp[i] = 1
				} else {
					dp[i] = min(dp[i], 1+dp[i-coins[j]])
				}
			}
		}
	}
	fmt.Println(dp)

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}