func mostCommonWord(paragraph string, banned []string) string {

	paragraph = strings.ToLower(paragraph)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	paragraph = reg.ReplaceAllString(paragraph, " ")
	wordsMap := make(map[string]int)
	for i := 0; i < len(banned); i++ {
		test := strings.ToLower(banned[i]) + " "
		paragraph = strings.ReplaceAll(paragraph, test, " ")
	}
	paragraph = reg.ReplaceAllString(paragraph, " ")
	paragraph = strings.Trim(paragraph, " ")
	words := strings.Split(paragraph, " ")
	mostCommon := 1
	var mosTcommonW string
	for _, word := range words {
		if _, found := wordsMap[word]; found {
			wordsMap[word]++
			if wordsMap[word] >= mostCommon {
				mosTcommonW = word
				mostCommon = wordsMap[word]
			}
		} else {
			wordsMap[word] = 1
			if mosTcommonW == "" {
				mosTcommonW = word
			}
		}
	}

	return mosTcommonW

}