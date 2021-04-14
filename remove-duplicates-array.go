func removeDuplicates(nums []int) int {
	var writeIdx, prev, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == prev && count >= 2 {
			continue
		}
		if nums[i] != prev {
			count = 0
		}
		nums[writeIdx] = nums[i]
		writeIdx++
		count++
		prev = nums[i]
	}
	return writeIdx
}