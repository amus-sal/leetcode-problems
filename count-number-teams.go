func numTeams(rating []int) int {
	ans, n := 0, len(rating)
	leftSmall := make([]int, n)
	leftLarge := make([]int, n)
	rightSmall := make([]int, n)
	rightLarge := make([]int, n)
	for i := 1; i < n-1; i++ {
		for j := i - 1; j >= 0; j-- {
			if rating[j] < rating[i] {
				leftSmall[i]++
			}
			if rating[j] > rating[i] {
				leftLarge[i]++
			}
		}
		for j := i + 1; j < n; j++ {
			if rating[j] < rating[i] {
				rightSmall[i]++
			}
			if rating[j] > rating[i] {
				rightLarge[i]++
			}
		}
	}

	for i := 1; i < n-1; i++ {
		ans += leftSmall[i]*rightLarge[i] + leftLarge[i]*rightSmall[i]
	}
	return ans
}