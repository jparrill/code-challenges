/*
👋 Welcome Gophers! In this challenge, you are tasked with implementing a function that checks to see whether a number is an armstrong number.


Armstrong Numbers - An Armstrong number is a 3-digit number such that each of the sum of the cubes of its digits equal the number itself:
371 = 3^3 + 7^3 + 1^3
*/
package main

import "fmt"

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	result := 0
	p_number := int(*n)
	digits := []int{}

	for p_number > 0 {
		digits = append(digits, p_number%10)
		p_number = p_number / 10
	}

	for _, digit := range digits {
		result += digit * digit * digit

	}

	if result == int(*n) {
		return true
	}
	return false
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 1
	fmt.Println(num1.IsArmstrong())
}
