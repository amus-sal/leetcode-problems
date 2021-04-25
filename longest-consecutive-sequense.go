func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	res := 1
	seq := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			seq++
			continue
		}
		if nums[i] == nums[i-1] {
			continue
		}
		res = max(res, seq)
		seq = 1
	}
	res = max(res, seq)
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}