
var rotten [][]int

func orangesRotting(grid [][]int) int {
	res := 0
	total := 0
	rotten = make([][]int, 0)
	for indexR, c := range grid {
		for indexC, val := range c {
			if val != 0 {
				total += 1
			}
			if val == 2 {
				rotten = append(rotten, []int{indexR, indexC})
			}
		}
	}
	prev := len(rotten)

	if prev == total {
		return 0
	}
	for {

		if len(rotten) < total {
			for _, pos := range rotten {
				getRotten(grid, pos[0], pos[1])
			}
			if len(rotten) == prev {
				return -1
			} else {
				prev = len(rotten)
			}
			res += 1
		} else {
			break
		}

	}

	return res
}

func getRotten(grid [][]int, sr int, sc int) {
	if sr >= len(grid) || sc >= len(grid[sr]) {
		return
	}

	if sr+1 <= len(grid)-1 && grid[sr+1][sc] == 1 {
		grid[sr+1][sc] = 2
		rotten = append(rotten, []int{sr + 1, sc})
	}
	if sr-1 >= 0 && grid[sr-1][sc] == 1 {
		grid[sr-1][sc] = 2
		rotten = append(rotten, []int{sr - 1, sc})
	}
	if sc+1 <= len(grid[sr])-1 && grid[sr][sc+1] == 1 {
		grid[sr][sc+1] = 2
		rotten = append(rotten, []int{sr, sc + 1})
	}
	if sc-1 >= 0 && grid[sr][sc-1] == 1 {
		grid[sr][sc-1] = 2
		rotten = append(rotten, []int{sr, sc - 1})
	}

}