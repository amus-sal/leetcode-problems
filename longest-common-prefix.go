func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	out := strs[0]
	for i := 1; i < len(strs); i++ {
		out = getCommonPrefix(out, strs[i])
		if out == "" {
			return out
		}
	}
	return out
}

func getCommonPrefix(str string, str2 string) string {
	out := ""
	if len(str) > len(str2) {
		str, str2 = str2, str
	}
	for i := range str {
		if str[i] == str2[i] {
			out += string(str[i])
		} else {
			return out
		}
	}
	return out
}