func printVertically(s string) []string {
	words := strings.Fields(s)
	//  get the max length
	if len(words) == 0 {
		return nil
	}
	//  find the max length of the words to difine the size of the two-dimensional array
	maxL := len(words[0])
	for i := 1; i < len(words); i++ {
		if maxL < len(words[i]) {
			maxL = len(words[i])
		}
	}
	res := make([]string, 0)
	for i := 0; i < maxL; i++ {
		wd := ""
		for _, word := range words {
			if i < len(word) {
				wd += string(word[i])
			} else {
				wd += " "
			}
		}
		res = append(res, wd)
	}
	// trim " " from the right
	for i := 0; i < len(res); i++ {
		p := -1
		for j := len(words) - 1; j >= 0; j-- {
			if res[i][j] != ' ' {
				p = j
				break
			}
		}
		res[i] = res[i][0 : p+1]
	}

	return res
}