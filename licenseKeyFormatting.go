package solution1

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	i := len(s)
	output := ""

	for i > k {
		output = "-" + s[i-k:i] + output

		i = i - k
	}
	output = s[0:i] + output
	return strings.ToUpper(output)

}
