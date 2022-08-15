/*
👋 Welcome Gophers! In this Go challenge, you are going to be implementing a function that takes in two string values and checks to see if they are permutations of one another.

Example
If I have 2 strings, “abc” and “cba”, when I pass these strings into the function, it should return true as these two strings are permutations of each other.
*/
package main

import "fmt"

func CheckPermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	var result1, result2 int
	for _, v := range str1 {
		result1 += int(v)
	}

	for _, v := range str2 {
		result2 += int(v)
	}

	if result1 != result2 {
		return false
	}

	return true
}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcme"
	str2 := "medac"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)
}
