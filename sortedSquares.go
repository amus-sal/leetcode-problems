import (
	"math"
	"fmt"
)

func sortedSquares(nums []int) []int {
	nwArr := make([]int, len(nums))
	lo, hi := 0, len(nums)-1
	cursor := hi
	for lo <= len(nums)-1 && (nums[lo] <= 0 || lo <= hi) {

		if int(math.Abs(float64(nums[lo]))) < nums[hi] {
			nwArr[cursor] = nums[hi] * nums[hi]
			cursor--
			hi--
		} else {
			nwArr[cursor] = nums[lo] * nums[lo]
			cursor--
			lo++
		}
	}
	return nwArr
}