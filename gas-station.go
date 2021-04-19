func canCompleteCircuit(gas []int, cost []int) int {
	totalRound := 0
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		totalRound += gas[i]
	}
	if totalRound < 0 {
		return -1
	}

	start := 0
	sum := gas[0]
	for i := 1; i < len(gas); i++ {
		if sum+gas[i] < 0 || gas[start] < 0 {
			start = i
			sum = gas[i]
			continue
		}
		sum += gas[i]
	}
	return start

}