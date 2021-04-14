func intersection(nums1 []int, nums2 []int) []int {
	out := make(map[int]bool)
	res := make([]int, 0)
	for _, val := range nums1 {
		out[val] = false
	}
	for _, val := range nums2 {
		if _, found := out[val]; found {
			if !out[val] {
				res = append(res, val)
			}
			out[val] = true
		}
	}

	return res

}