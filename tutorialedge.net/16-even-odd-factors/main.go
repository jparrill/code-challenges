/*
👋 Welcome Gophers! In this challenge, you will be tasked with implementing a function that will return either “odd” or “even” depending on whether or not a number has an odd or an even number of factors.

Aclaration:
- Find the factors of an int
- Check the lenght of the total Factors Found
- Check if that lenght is even or odd
*/

package main

import (
	"fmt"
	"math"
)

func OddEvenFactors(num int) string {
	factors := CheckFactors(num)
	fmt.Println("Factors:", factors)
	if len(factors)%2 == 0 {
		return "even"
	}
	return "odd"
}

func CheckFactors(num int) []int {
	factors := []int{1, num}
	limit := int(math.Ceil(float64(num) / 2))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			fmt.Printf("Found: %d\n", i)
			factors = append(factors, i)
		}
	}
	return factors
}

func main() {
	fmt.Println("Odd or Even Factors")
	numFactors := OddEvenFactors(23)
	fmt.Println("Result: ", numFactors)
	fmt.Println("------------")
	numFactors = OddEvenFactors(24)
	fmt.Println("Result: ", numFactors)
	fmt.Println("------------")
	numFactors = OddEvenFactors(36)
	fmt.Println("Result: ", numFactors)
	fmt.Println("------------")
}
