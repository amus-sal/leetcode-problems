func numIslands(grid [][]byte) int {
	// loop over ones
	count := 0
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return count
	}
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if string(grid[i][j]) == "1" {
				helper(grid, i, j)
				count++
			}
		}
	}
	return count
}

func helper(grid [][]byte, i, j int) {
	row, col := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= row || j >= col || string(grid[i][j]) != "1" {
		return
	}
	grid[i][j] = byte('#')
	helper(grid, i-1, j)
	helper(grid, i+1, j)
	helper(grid, i, j-1)
	helper(grid, i, j+1)
}