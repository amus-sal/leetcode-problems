
var res int
var moved [][]bool

func findCircleNum(M [][]int) int {
	res = 0
	moved = make([][]bool, len(M))
	for i, _ := range M {
		moved[i] = make([]bool, len(M[i]))
		for c, _ := range moved[i] {
			moved[i][c] = false
		}
	}
	for r, _ := range M {
		for c, _ := range M {
			if M[r][c] == 1 && !moved[r][c] {
				res += 1
				move(M, r, c)

			}
		}
	}

	return res
}

func move(M [][]int, row int, col int) {

	for co, _ := range M[row] {
		if !moved[row][co] {
			if M[row][co] == 1 {
				moved[row][co] = true
				move(M, row, co)
			}
		}
	}
	for i := 0; i < len(M); i++ {
		if !moved[i][col] {
			if M[i][col] == 1 {
				moved[i][col] = true
				move(M, i, col)
			}
		}
	}

}