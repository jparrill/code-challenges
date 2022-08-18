/*
In this challenge, you are going to implement the DiffSquares function so that it returns the difference between the first number squared minus the second number squared.

Less efficient than Solution 1:
- This methods hold floating point operations and casting
*/
package main

import (
	"fmt"
	"math"
)

func DiffSquares(n, m int) int {
	x := math.Pow(float64(n), 2)
	y := math.Pow(float64(m), 2)
	return int(x) - int(y)
}

func main() {
	fmt.Println("Calculate The Difference of Squares")
	result := DiffSquares(5, 4)
	fmt.Println(result)
}
