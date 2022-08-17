/*
👋 Welcome Gophers! In this challenge, you are tasked with finding the smallest difference between two slices of int values.

Example: Let’s say I have 2 int arrays; [1, 2] and [4, 5]. The function should return the smallest difference which would be the difference between 2 and 4.

WARNING: The Page solution is not functional at all so this is my solution, is much computational complex, but works in all cases
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func CalcSmallestDifference(arr1, arr2 []int) float64 {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var minDiff, tmpDiff, min, max float64

	minDiff = math.Max(float64(arr1[len(arr1)-1]), float64(arr2[len(arr2)-1]))

	for o := range arr1 {
		for i := range arr2 {
			if arr1[o] == arr2[i] {
				return 0
			}
			min = math.Min(float64(arr1[o]), float64(arr2[i]))
			max = math.Max(float64(arr1[o]), float64(arr2[i]))
			tmpDiff = max - min
			minDiff = math.Min(tmpDiff, minDiff)
		}
	}

	return minDiff
}

func main() {
	fmt.Println("Smallest Difference Challenge")

	arr1 := []int{9, 8, 7, 6}
	arr2 := []int{7, 3, 2, 5}

	smallestDiff := CalcSmallestDifference(arr1, arr2)

	fmt.Println(smallestDiff)

	arr3 := []int{1, 2}
	arr4 := []int{5, 67, 8, 5, 6, 8}

	smallestDiff = CalcSmallestDifference(arr3, arr4)

	fmt.Println(smallestDiff)
}
