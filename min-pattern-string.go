// var res []string
package main

import (
	"fmt"
	"strings"
)

var loWestPatt string

func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	if len(t) > len(s) {
		return ""
	}
	//res = make([]string, 0)
	full, _ := contain(s, &t)
	if full == true {
		loWestPatt = s
	} else {
		return ""
	}
	getSt(s, "", &t)
	// fmt.Println(res)
	return loWestPatt
}

func getSt(s string, prefix string, min *string) {
	if s == "" {
		return
	}
	// res = append(res, prefix+string(s[0]))
	full, part := contain(prefix+string(s[0]), min)

	if full == true {
		// res = append(res, prefix+string(s[0]))
		if len(prefix+string(s[0])) < len(loWestPatt) {
			loWestPatt = prefix + string(s[0])
		}
		if string(s[0]) != *min {
			if fl, pt := contain(string(s[0]), min); (pt || fl) && string(s[0]) != *min {
				getSt(s, "", min)
			}
		}
		getSt(s[1:], "", min)

	} else if part == true {
		getSt(s[1:], prefix+string(s[0]), min)
	} else {
		getSt(s[1:], "", min)
	}
}

func contain(s string, pattern *string) (bool, bool) {
	total := 0
	if s == *pattern {
		return true, true
	}
	for _, pat := range *pattern {
		for ind, char := range s {
			if pat == char {
				total += 1
				s = strings.Replace(s, string(s[ind]), "", -1)
				break
			}
		}
	}

	if len(*pattern) == total {
		return true, true
	}
	if total > 0 {
		return false, true
	}

	return false, false

}

func main() {
	fmt.Println(minWindow("ab", "a"))
	fmt.Println(minWindow("bba", "ab"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))

}
