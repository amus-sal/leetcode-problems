package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"strings"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(message string, K int) string {
	// write your code in Go 1.4
	words := strings.Fields(message)
	res := ""

	for i := 0; i < len(words); i++ {
		if len(res)+len(words[i]) <= K-1 {
			res += words[i] + " "
		} else {
			break
		}
	}
	// trim func
	res = strings.Trim(res, " ")
	return res

}
