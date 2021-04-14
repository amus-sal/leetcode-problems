func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch {
			case i == 0 && j == 0:
				continue
			case i == 0:
				grid[i][j] += grid[i][j-1]
			case j == 0:
				grid[i][j] += grid[i-1][j]
			default:
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][col-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}