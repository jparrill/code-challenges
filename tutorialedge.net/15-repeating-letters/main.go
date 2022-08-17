/*
👋 Welcome Gophers! In this challenge, you are tasked with implementing a function DoubleChars which will take in a string and then return another string which has every letter in the word doubled.
*/
package main

import "fmt"

func DoubleChars(original string) string {
	var doubled string

	for _, ch := range original {
		doubled += fmt.Sprintf("%c%c", ch, ch)
	}
	return doubled
}

func main() {
	fmt.Println("Smallest Difference Challenge")

	original := "gophers"
	doubled := DoubleChars(original)
	fmt.Println(doubled) // ggoopphheerrss
}
