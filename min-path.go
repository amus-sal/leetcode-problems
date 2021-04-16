// func minPathSum(grid [][]int) int {
// 	row := len(grid)
// 	col := len(grid[0])
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			switch {
// 			case i == 0 && j == 0:
// 				continue
// 			case i == 0:
// 				grid[i][j] += grid[i][j-1]
// 			case j == 0:
// 				grid[i][j] += grid[i-1][j]
// 			default:
// 				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
// 			}
// 		}
// 	}
// 	return grid[row-1][col-1]
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

var store [][]int
var mainData [][]int
var cols, rows int

func minPathSum(grid [][]int) int {
	mainData = grid
	cols = len(grid[0])
	rows = len(grid)
	store = make([][]int, rows)
	for row := range store {
		store[row] = make([]int, cols)
	}
	cols--
	rows--
	return GetMinPath(0, 0)

}
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
func GetMinPath(row int, col int) int {
	if store[row][col] != 0 {
		return store[row][col]
	}
	if row == rows && col == cols {
		store[row][col] = mainData[row][col]
	} else if row < rows && col == cols {
		store[row][col] = GetMinPath(row+1, col) + mainData[row][col]
	} else if col < cols && row == rows {
		store[row][col] = GetMinPath(row, col+1) + mainData[row][col]
	} else {
		store[row][col] = min(GetMinPath(row, col+1), GetMinPath(row+1, col)) + mainData[row][col]
	}
	return store[row][col]
}

