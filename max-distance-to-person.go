func maxDistToClosest(seats []int) int {
	exPers := make([]int, 0)
	for ind, val := range seats {
		if val == 1 {
			exPers = append(exPers, ind)
		}
	}
	fmt.Println(exPers)
	ans := 0
	if exPers[0] > 0 {
		ans = exPers[0]
	}
	if exPers[len(exPers)-1] < len(seats)-1 {
		ans = max(ans, len(seats)-exPers[len(exPers)-1]-1)
		fmt.Println(exPers[len(exPers)-1])
	}
	for i := 0; i < len(exPers)-1; i++ {
		ans = max(ans, (exPers[i+1]-exPers[i])/2)
	}

	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}