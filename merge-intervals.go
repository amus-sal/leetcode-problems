func merge(intervalsCop [][]int) [][]int {
	out := make([][]int, 0)
	sort.Slice(intervalsCop, func(i, k int) bool {
		return intervalsCop[i][0] < intervalsCop[k][0]
	})
	start := false
	for i := 0; i <= len(intervalsCop)-1; i++ {
		if intervalsCop[i] == nil {
			continue
		}
		if !start {
			out = append(out, intervalsCop[i])
			start = true
			continue
		}

		if out[len(out)-1][1] >= intervalsCop[i][0] && out[len(out)-1][1] <= intervalsCop[i][1] {
			out[len(out)-1] = []int{
				out[len(out)-1][0],
				intervalsCop[i][1],
			}
		} else if out[len(out)-1][1] < intervalsCop[i][0] {
			out = append(out, intervalsCop[i])
		}

	}
	return out

}