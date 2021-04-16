import (
	"fmt"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 2; i <= n; i++ {
		for st := 1; st <= n-i+1; st++ {

			end := st + i - 1
			curMin := 245343234
			fmt.Println("II: ", i)
			fmt.Println("ST: ", st)
			fmt.Println("EN: ", end)
			fmt.Println("--------------------")

			for pi := st; pi <= end; pi++ {
				curMin = min(curMin, pi+max(dp[st][pi-1], dp[pi+1][end]))
			}
			dp[st][end] = curMin
		}
	}
	return dp[1][n]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}