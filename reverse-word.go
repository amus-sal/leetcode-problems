func reverseWords(s string) string {
	words := strings.Fields(s)
	i := 0
	j := len(words) - 1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}