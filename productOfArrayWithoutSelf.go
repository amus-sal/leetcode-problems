func productExceptSelf(nums []int) []int {
	total := 1
	zeroVal := -1
	for i, v := range nums {
		if v != 0 {
			total = total * v
		} else {
			if zeroVal != -1 {
				total = 0
				break
			}
			zeroVal = i
		}
	}
	out := make([]int, len(nums))
	for i, v := range nums {
		if zeroVal == -1 {
			out[i] = total / v
		} else {
			if i != zeroVal {
				out[i] = 0
			} else {
				out[i] = total
			}
		}
	}
	return out

}