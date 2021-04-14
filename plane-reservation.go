package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	out := Solution(2, "1A 2F 1C")
	fmt.Println(out)
}
func Solution(N int, S string) int {
	var SeatsNum = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"J": 8,
		"K": 9,
	}
	seats := strings.Split(S, " ")

	resSeats := make([][]int, N)
	for row := range resSeats {
		resSeats[row] = make([]int, 10)
		for col := range resSeats[row] {
			resSeats[row][col] = 0 // not reserved
		}
	}

	for i := 0; i < len(seats); i++ {
		rowNum, _ := strconv.ParseInt(string(seats[i][0]), 0, 32)
		colNum := SeatsNum[string(seats[i][1])]
		resSeats[rowNum-1][colNum] = 1
	}

	/*
	   four family  :
	   B to C   and F to J
	   D to G
	   F to J
	*/
	count := 0
	for row := range resSeats {
		prob1 := checkSeq(row, 1, 4, resSeats)
		prob2 := checkSeq(row, 3, 6, resSeats)
		prob3 := checkSeq(row, 5, 8, resSeats)

		if prob1 && prob3 {
			count += 2
		} else if prob2 {
			count += 1
		} else if prob1 {
			count += 1
		} else if prob3 {
			count += 1
		}
		fmt.Println(count)
	}
	return count

}

func checkSeq(row int, start int, end int, resSeats [][]int) bool {
	for i := start; i <= end; i++ {
		if resSeats[row][i] != 0 {
			return false
		}
	}
	return true
}
