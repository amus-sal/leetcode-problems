import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if len(matrix) == 0 {
		return 0
	}
	cols := len(matrix[0])

	arr := make([]int, cols)
	max := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cur := int(matrix[i][j] - '0')
			if cur == 0 {
				arr[j] = 0
			} else {
				arr[j] = arr[j] + cur
			}
		}
		fmt.Println(arr)
		area := largestRectangleArea(arr)
		if area > max {
			max = area
		}
	}
	return max
}

func largestRectangleArea(heights []int) int {
	var stack []int
	stack = append(stack, -1)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) != 1 && heights[i] <= heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop

			maxArea = maxVal(maxArea, heights[top]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)

	}

	for len(stack) != 1 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxArea = maxVal(maxArea, heights[top]*(len(heights)-stack[len(stack)-1]-1))
	}

	return maxArea
}

func maxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}