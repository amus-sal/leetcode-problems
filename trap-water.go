func trap(height []int) int {
	var ans int
	var leftMost, rightMost int
	for l, r := 0, len(height)-1; l < r; {
		if height[l] < height[r] {
			if height[l] > leftMost {
				leftMost = height[l]
			} else {
				ans += leftMost - height[l]
			}
			l++
		} else {
			if height[r] > rightMost {
				rightMost = height[r]
			} else {
				ans += rightMost - height[r]
			}
			r--
		}
	}
	return ans
}