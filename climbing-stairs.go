func climbStairs(n int) int {
	step := make([]int, n+1)
	step[0] = 1
	step[1] = 1
	for i := 2; i <= n; i++ {
		step[i] = step[i-1] + step[i-2]
	}
	return step[n]
}
