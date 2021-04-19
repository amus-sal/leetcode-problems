
func removeDuplicates(nums []int) int {
	a := len(nums)
	if a < 1 {
		return 0
	}
	i, j := 0, 1
	for j < a {
		if nums[i] == nums[j] {
			j++
			continue
		}
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}