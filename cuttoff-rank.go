func cutoffRank(ranks []int, cuttof int) int {
	lastRank := 1
	res := 0
	for i := 1; i < ranks; i++ {
		if ranks[i] == ranks[i-1] {
			res += 1
		} else {
			lastRank += 1
			if lastRank > cuttof {
				break
			}
		}
	}

	return out

}