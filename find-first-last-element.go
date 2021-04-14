func searchRange(nums []int, target int) []int {
	lo := 0
	hi := len(nums) - 1
	first := -1
	last := -1
	for lo <= hi {
		if nums[lo] != target && nums[hi] != target {
			lo++
			hi--
			continue
		}
		if nums[lo] == target && nums[hi] == target {
			return []int{lo, hi}
		}
		if nums[lo] == target {
			first = lo
			hi--
		} else {
			last = hi
			lo++
		}

	}
	return []int{first, last}
}
