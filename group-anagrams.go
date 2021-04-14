func groupAnagrams(strs []string) [][]string {
	total := make(map[string][]string)
	for _, word := range strs {
		wordint := calcWord(word)
		total[wordint] = append(total[wordint], word)

	}
	var result [][]string
	for _, element := range total {
		result = append(result, element)
	}
	return result
}

func calcWord(str string) string {
	var ar []string
	for _, wordChar := range str {
		s := fmt.Sprintf("%c", wordChar)
		ar = append(ar, s)

	}
	sort.Strings(ar)
	return strings.Join(ar, "")
}
