/*
In this challenge, you are going to implement the TriangularNumbers function which takes in n and returns the nth triangular number.

n = 1
result = 1

n = 3
result = 6

Clarification: https://en.wikipedia.org/wiki/Triangular_number
n ==  number of stars in the piramid base

if n = 5

*
**
***
****
*****

result = 15
*/

package main

import "fmt"

func TriangularNumber(n int) int {
	var result int = 0

	for i := n; i > 0; i-- {
		// first round
		if i == n {
			result = n
		} else {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println("Returning the 'nth' triangular number")

	fmt.Println(TriangularNumber(3))
	fmt.Println(TriangularNumber(5))
	fmt.Println(TriangularNumber(6))
	fmt.Println(TriangularNumber(9))
	fmt.Println(TriangularNumber(1))
}
