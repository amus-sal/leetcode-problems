
var modBoard [][]int

func gameOfLife(board [][]int) {
	modBoard = make([][]int, len(board))
	for r, _ := range board {
		modBoard[r] = make([]int, len(board[r]))
		for c, val := range board[r] {
			modBoard[r][c] = val
		}
	}
	for indR, row := range board {
		for indC, _ := range row {
			modify(board, indR, indC)
		}
	}

	for r, _ := range board {
		for c, _ := range board[r] {
			board[r][c] = modBoard[r][c]
		}
	}
}

func modify(board [][]int, indR int, indC int) {
	lives := getNumbers(board, indR, indC)
	if board[indR][indC] == 1 && (lives > 3 || lives < 2) {
		modBoard[indR][indC] = 0
	}
	if board[indR][indC] == 0 && lives == 3 {
		modBoard[indR][indC] = 1
	}

}

func getNumbers(board [][]int, indR int, indC int) int {
	lives := 0
	if indC-1 < len(board[indR]) && indC-1 >= 0 && indR < len(board) && indR >= 0 {
		if board[indR][indC-1] == 1 {
			lives += 1
		}
	}

	if indC+1 < len(board[indR]) && indC+1 >= 0 && indR < len(board) && indR >= 0 {
		if board[indR][indC+1] == 1 {
			lives += 1
		}
	}
	if indC < len(board[indR]) && indC >= 0 && indR-1 < len(board) && indR-1 >= 0 {
		if board[indR-1][indC] == 1 {
			lives += 1
		}
	}
	if indC < len(board[indR]) && indC >= 0 && indR+1 < len(board) && indR+1 >= 0 {
		if board[indR+1][indC] == 1 {
			lives += 1
		}
	}
	if indC-1 < len(board[indR]) && indC-1 >= 0 && indR-1 < len(board) && indR-1 >= 0 {
		if board[indR-1][indC-1] == 1 {
			lives += 1
		}
	}
	if indC+1 < len(board[indR]) && indC+1 >= 0 && indR+1 < len(board) && indR+1 >= 0 {
		if board[indR+1][indC+1] == 1 {
			lives += 1
		}
	}
	if indC-1 < len(board[indR]) && indC-1 >= 0 && indR+1 < len(board) && indR+1 >= 0 {
		if board[indR+1][indC-1] == 1 {
			lives += 1
		}
	}
	if indC+1 < len(board[indR]) && indC+1 >= 0 && indR-1 < len(board) && indR-1 >= 0 {
		if board[indR-1][indC+1] == 1 {
			lives += 1
		}
	}

	return lives
}

