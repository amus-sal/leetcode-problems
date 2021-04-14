var res []string

func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	if len(s) < 4 {
		return []string{}
	}
	res = make([]string, 0)
	getRem(s, 4, "")
	fmt.Println(res)
	return res

}

func getRem(word string, rem int, prev string) {
	if rem*3 < len(word) {
		return
	}

	if rem == 0 && word == "" {
		return
	}
	if rem == 1 && word == "" {
		return
	}
	if rem == 1 {
		if len(word) > 3 {
			return
		}
		if getNum(word) <= 255 {
			if len(word) >= 1 && getNum(string(word[0])) != 0 {
				res = append(res, prev+word)
			}
			if getNum(string(word[0])) == 0 && len(word) == 1 {
				res = append(res, prev+word)
			}
			return
		}
	}

	getRem(string(word[1:]), rem-1, prev+string(word[0])+".")
	if rem < len(word) {
		if getNum(string(word[0])) != 0 {
			if getNum(string(word[0])) == 2 && getNum(string(word[1])) <= 5 {
				getRem(word[2:], rem-1, prev+string(word[0])+string(word[1])+".")
				if getNum(string(word[0])+string(word[1])+string(word[2])) <= 255 {
					getRem(string(word[3:]), rem-1, prev+string(word[0])+string(word[1])+string(word[2])+".")
				}
			} else if getNum(string(word[0])) <= 2 {
				getRem(word[2:], rem-1, prev+string(word[0])+string(word[1])+".")
				if getNum(string(word[0])+string(word[1])+string(word[2])) <= 255 {
					getRem(word[3:], rem-1, prev+string(word[0])+string(word[1])+string(word[2])+".")
				}
			} else if getNum(string(word[0])) > 2 {
				getRem(word[2:], rem-1, prev+string(word[0])+string(word[1])+".")
			}
		}
	}

}

func getNum(word string) int {
	i, _ := strconv.ParseInt(word, 0, 64)
	return int(i)
}
