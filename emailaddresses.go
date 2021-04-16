package solution

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]bool)

	for _, val := range emails {
		vals := strings.Split(val, "@")
		vals[0] = strings.ReplaceAll(vals[0], ".", "")

		valWithoutPlus := strings.Split(vals[0], "+")[0]

		uniqueEmails[valWithoutPlus+"@"+vals[1]] = true
	}

	return len(uniqueEmails)
}
