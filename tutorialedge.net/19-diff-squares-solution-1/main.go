/*
In this challenge, you are going to implement the DiffSquares function so that it returns the difference between the first number squared minus the second number squared.
*/
package main

import "fmt"

func DiffSquares(n, m int) int {
	sq1 := n * n
	sq2 := m * m
	if m > n {
		return sq2 - sq1
	} else if n > m {
		return sq1 - sq2
	}
	return 0
}

func main() {
	fmt.Println("Calculate The Difference of Squares")
}
