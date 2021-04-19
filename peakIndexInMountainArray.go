func peakIndexInMountainArray(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		m := (hi-lo)/2 + lo
		if arr[m] < arr[m+1] {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo
}