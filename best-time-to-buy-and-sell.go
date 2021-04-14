func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	lastLowest := prices[0]
	lastDiff := 0
	for i := 1; i <= len(prices)-1; i++ {
		if prices[i] <= lastLowest {
			lastLowest = prices[i]
		}
		if prices[i]-lastLowest >= lastDiff {
			lastDiff = prices[i] - lastLowest
		}
	}
	return lastDiff
}