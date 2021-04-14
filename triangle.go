var res [][]int

func minimumTotal(triangle [][]int) int {
	res = make([][]int, len(triangle))
	for i := range res {
		res[i] = make([]int, i+1)
		for ind := range res[i] {
			res[i][ind] = 1234567
		}
	}
	out := get(triangle, 0, 0)
	fmt.Println(res)
	return out
}

func get(triangle [][]int, layer int, index int) int {
	if index > layer {
		return 1234567
	}
	if res[layer][index] != 1234567 {
		return res[layer][index]
	}
	if len(triangle)-1 == layer {
		return triangle[layer][index]
	}
	total := min(get(triangle, layer+1, index), get(triangle, layer+1, index+1)) + triangle[layer][index]
	res[layer][index] = total
	return total
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}