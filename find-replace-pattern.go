func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, wrd := range words {
		if validWord(wrd, pattern) {
			res = append(res, wrd)
		}
	}
	return res
}

func validWord(word string, pattern string) bool {
	chrMap := make(map[string]string)
	wrMap := make(map[string]string)

	if len(pattern) != len(word) {
		return false
	}
	for i, s := range pattern {
		if v, ok := chrMap[string(s)]; ok {
			if string(word[i]) != v {
				return false
			}
		} else {
			if _, ok := wrMap[string(word[i])]; !ok {
				chrMap[string(s)] = string(word[i])
				wrMap[string(word[i])] = string(s)
			} else {
				return false
			}
		}
	}
	return true
}

func nextGreaterElement(n int) int {
	if n < 10 {
		return -1
	}
	s := strconv.Itoa(n)
	if n > 2147483647 {
		return -1
	}
	r := []rune(s)
	done := false
	for i := len(s) - 1; i > 0; i-- {
		val1, _ := strconv.Atoi(string(s[i]))
		for j := i - 1; j > 0; j-- {
			val2, _ := strconv.Atoi(string(s[j]))
			if val1 > val2 {
				r[i], r[j] = r[j], r[i]
				// get the smallest opt
				r = []rune(string(r[:i]) + getSmallest(string(r[i:])))
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

	out, _ := strconv.Atoi(s)
	out = reverse_int(out)
	s = strconv.Itoa(out)
	return s

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