func findPeakElement(nums []int) int {
	return getPeek(nums, 0, len(nums)-1)

	// if len(nums) == 1 {
	//     return 0
	// }
	// nums = append(nums, -98767898767876)
	// for i:= 1; i <= len(nums)-2;i++{
	//     if nums[i-1]< nums[i] && nums[i]> nums[i +1]{
	//         return i
	//     }
	// }
	// return 0
}

func getPeek(nums []int, start int, end int) int {

	if start == end {
		return start
	}
	mid := (end + start) / 2
	if nums[mid] > nums[mid+1] {
		return getPeek(nums, start, mid)
	}
	return getPeek(nums, mid+1, end)

}