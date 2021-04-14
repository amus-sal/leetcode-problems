package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	count := 0
	// count the freq of each number and store it in the map
	freqMap := make(map[int]int)
	for _, val := range A {
		if _, found := freqMap[val]; found {
			freqMap[val] += 1
		}
		freqMap[val] = 1
	}

	for _, val := range freqMap {
		// get value of number of valid pairs for each req number
		count += (val * (val - 1)) / 2
	}

	return count
}

// Time Complexity is O(N)

/*
 we can also solve it with the brute force sol by iterating over the (i : n)array and  iterating   (i+1: n) array and check if the number is repeated we will count ++

    Time Complexity os O(N^2)

*/
