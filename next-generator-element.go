
func nextGreaterElement(n int) int {
	if n < 10 {
		return -1
	}
	s := strconv.Itoa(n)
	if n > 2147483647 {
		return -1
	}
	done := false
	r := []rune(s)
	for i := len(s) - 2; i >= 0; i-- {
		val1, _ := strconv.Atoi(string(s[i]))
		for j := len(s) - 1; j > i; j-- {
			val2, _ := strconv.Atoi(string(s[j]))
			if val2 > val1 {

				r[i], r[j] = r[j], r[i]
				// get the smallest opt
				fmt.Println(string(r))
				r = []rune(string(r[:i+1]) + getSmallest(string(r[i+1:])))
				fmt.Println(string(r))
				done = true
				break
			}
		}
		if done {
			break
		}

	}

	out, _ := strconv.Atoi(string(r))
	if out == n {
		return -1
	}
	fmt.Println(out)
	if out > 2147483647 {
		return -1
	}
	return out
}

func getSmallest(s string) string {

	fmt.Println(s)
	fmt.Println("Before: ", s)
	s = reverse(s)
	fmt.Println("After: ", s)
	return s

}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}