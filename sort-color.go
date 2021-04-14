func sortColors(nums []int) {
	left, curr, right := 0, 0, len(nums)-1
	for curr <= right {
		if nums[curr] == 0 {
			nums[left], nums[curr] = nums[curr], nums[left]
			left++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[right] = nums[right], nums[curr]
			right--
		} else {
			curr++
		}
	}
}