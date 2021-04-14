func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := 987898767
	for i := 0; i < len(nums)-2; i++ {
		fix := nums[i]
		ptr1 := i + 1
		ptr2 := len(nums) - 1
		for ptr1 < ptr2 {
			sum := fix + nums[ptr1] + nums[ptr2]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closest)) {
				closest = sum
			}

			if target > sum {
				ptr1 += 1
			} else {
				ptr2 -= 1
			}
		}

	}
	return closest

}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}