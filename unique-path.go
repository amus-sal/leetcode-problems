var pos [][]int

func uniquePaths(m int, n int) int {
	pos = make([][]int, m)
	for indR := range pos {
		pos[indR] = make([]int, n)
		for ind := range pos[indR] {
			pos[indR][ind] = -1
		}
	}
	return getPath(0, 0, m-1, n-1)
}

func getPath(stRo int, stCo int, m int, n int) int {
	if pos[stRo][stCo] != -1 {
		return pos[stRo][stCo]
	}
	out := 0
	if stRo == m && stCo == n {
		out = 1
	} else if stRo == m {
		out = getPath(stRo, stCo+1, m, n)
	} else if stCo == n {
		out = getPath(stRo+1, stCo, m, n)
	} else {
		out = getPath(stRo, stCo+1, m, n) + getPath(stRo+1, stCo, m, n)
	}
	pos[stRo][stCo] = out
	return out
}

