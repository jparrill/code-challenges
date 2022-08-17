/*
👋 Welcome Gophers, in this challenge, you are tasked with implementing a function that returns whether or not a year is in fact a leap year.
*/
package main

import (
	"fmt"
	"time"
)

func CheckLeapYear(year int) bool {
	leap := time.Date(year, 2, 29, 0, 0, 0, 0, time.UTC)
	if leap.Month() == time.February {
		return true
	}

	return false
}

func main() {
	fmt.Println("Check Leap Year")

	year := 2020
	fmt.Println(CheckLeapYear(year))
	year = 2021
	fmt.Println(CheckLeapYear(year))
}
