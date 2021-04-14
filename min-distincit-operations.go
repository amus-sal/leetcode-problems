package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	freq_map := make(map[int]int)
	for i := 0; i < len(A); i++ {
		freq_map[A[i]] += 1
	}
	res := 0
	for _, val := range freq_map {
		if val > 1 {
			res += val - 1
		}
	}
	return res
}
