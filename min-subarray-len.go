const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func minSubArrayLen(s int, nums []int) int {
	res := MaxInt
	lft := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= s {
			return 1
		}
		sum += nums[i]
		for sum >= s {
			res = min(res, i+1-lft)
			sum -= nums[lft]
			lft += 1
		}

	}

	if res == MaxInt {
		return 0
	}
	return res
}
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
