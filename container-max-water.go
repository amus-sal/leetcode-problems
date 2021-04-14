func maxArea(height []int) int {

	maxArea := 0
	low := 0
	high := len(height) - 1

	for low < high {
		maxArea = max(min(height[low], height[high])*(high-low), maxArea)
		if height[low] > height[high] {
			high--
		} else {
			low++
		}
	}
	return maxArea
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}