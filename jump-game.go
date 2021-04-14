func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	output := make([]bool, len(nums)-1)

	minJumb := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			minJumb += 1
			output[i] = false
			continue
		}
		if nums[i] >= minJumb {
			minJumb = 1
			output[i] = true
		} else {
			output[i] = false
			minJumb += 1
		}

	}

	return output[0]
}