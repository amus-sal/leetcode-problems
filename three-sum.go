var res [][]int
var m map[[3]int]bool
var a [3]int

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 0)
	m = make(map[[3]int]bool)

	for index, _ := range nums {
		get2sum(nums, 0, index)
	}

	return res
}

func get2sum(nums []int, target int, index int) {
	ptr1 := index + 1
	ptr2 := len(nums) - 1

	for ptr1 < ptr2 {
		if nums[ptr1]+nums[ptr2]+nums[index] == target {
			a = [3]int{nums[index], nums[ptr1], nums[ptr2]}
			if _, ok := m[a]; !ok {
				m[a] = true
				res = append(res, []int{nums[index], nums[ptr1], nums[ptr2]})
			}
			ptr1++
		} else if nums[ptr1]+nums[ptr2]+nums[index] > target {
			ptr2--
		} else {
			ptr1++
		}
	}
}
