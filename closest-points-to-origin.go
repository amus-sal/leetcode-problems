type PointsDup struct {
	Number int
	Potint []int
}

func kClosest(points [][]int, K int) [][]int {
	// [(0,1), (2,3), (2,1), (1,1)]

	var output [][]int
	mapOut := make(map[int][][]int)
	var dummy []int
	for i := 0; i < len(points); i++ {
		dummy = append(dummy, (points[i][0]*points[i][0] + points[i][1]*points[i][1]))
		mapOut[(points[i][0]*points[i][0] + points[i][1]*points[i][1])] = append(mapOut[(points[i][0]*points[i][0]+points[i][1]*points[i][1])], points[i])

	}
	sort.Ints(dummy)
	prev := -1
	for i := 0; i < K; i++ {
		if dummy[i] != prev {
			if K-len(output) < len(mapOut[dummy[i]]) {
				output = append(output, mapOut[dummy[i]][0:K-len(output)]...)
				break
			} else {
				output = append(output, mapOut[dummy[i]]...)

			}
			prev = dummy[i]
		}
	}

	return output
}

