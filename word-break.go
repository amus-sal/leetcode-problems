func wordBreak(s string, wordDict []string) bool {

	return testBreak(s, wordDict)
}

func testBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {

			if testBreak(strings.Replace(s, word, "", 1), wordDict) {
				return true
			}
		}
	}

	return false
}

// Optimama sol
// func wordBreak(s string, wordDict []string) bool {

//     dict := make(map[string]bool)
//     for _, v := range wordDict {
//         dict[v] = true
//     }
//     mem := make([]bool, len(s)+1)
//     mem[0] = true
//     for i := 1; i < len(mem); i++ {
//         for j := i-1; j >= 0; j-- {
//             if mem[j] && dict[s[j:i]] {
//                 mem[i] = true
//                 break
//             }
//         }
//     }
//     return mem[len(mem)-1]
//     //return testBreak(s, wordDict)
// }