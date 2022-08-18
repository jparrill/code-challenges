/*
👋 Welcome Gophers! In this challenge, you will be tasked with finding out how many rotations an ordered int slice has undergone and shifted by.

Examples:
arr := []int{1, 2, 3, 4, 5}
MinRotations(arr) // returns 0

arr := []int{3, 4, 5, 1, 2}
MinRotations(arr) // returns 3
*/

package main

import (
	"fmt"
	"sort"
)

func MinRotations(array []int) int {
	var rotations int
	var tmpArr = make([]int, len(array))
	copy(tmpArr, array)
	sort.Ints(tmpArr)
	min := tmpArr[0]

	for i := range array {
		if min == array[i] {
			break
		}
		rotations++
	}
	return rotations
}

func main() {
	fmt.Println("Min Rotation Challenge")

	arr := []int{15, 18, 2, 3, 6, 12}
	fmt.Println(MinRotations(arr)) // returns 2

	arr = []int{7, 9, 11, 12, 5}
	fmt.Println(MinRotations(arr)) // returns 2

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(MinRotations(arr)) // returns 2

	arr = []int{3, 4, 5, 1, 2}
	fmt.Println(MinRotations(arr)) // returns 2

}
