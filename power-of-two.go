func isPowerOfTwo(n int) bool {
	i := 1
	for i < n {
		i *= 2
	}
	return i == n
}
