package main

import "fmt"

func DP(t [][]int) int {
	ans := 0
	// process from the second row from the last row and keep getting up
	for row := len(t) - 2; row > -1; row-- {
		for col := 0; col < row+1; col++ {
			// update each [row][col] element to store the maximum sum
			// from the below row (row + 1) by look at both left and right
			t[row][col] += findMax(t[row+1][col], t[row+1][col+1])
		}
	}

	// after the loop if its correct we assume
	// we've grabed maximum from each row and store at the [0][0]
	ans = t[0][0]
	return ans
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	// if a == b we assume that we can return either a or b
	return b
}

func main() {

	result := DP(getTriangle())

	fmt.Println("Maximum path sum:", result)
}
