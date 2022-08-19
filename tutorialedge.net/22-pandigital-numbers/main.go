/*
Problem Attribution - This challenge was inspired by Problem 41: https://projecteuler.net/problem=41 on Project Euler

👋 Welcome Gophers! In this challenge, you are tasked with implementing the LargestPandigitalPrime function which will return the largest possible pandigital prime number.

In mathematics, a pandigital number is an integer that in a given base has among its significant digits each digit used in the base at least once. For example, 1234567890 (one billion two hundred thirty four million five hundred sixty seven thousand eight hundred ninety) is a pandigital number in base 10. The first few pandigital base 10 numbers are given by (sequence A050278 in the OEIS):
*/

/*
To be fully honest here, I don't know how they prepare the tests for this challenge... so I'd got the solution from the page...
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func notPrime(n int) bool {
	if n&2 == 2 {
		return false
	}
	if (n-1)%6 == 0 {
		return false
	}
	if (n+1)%6 == 0 {
		return false
	}
	return true
}

func isPrime(n int) bool {
	if notPrime(n) {
		return false
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	for i := int(3); i < max; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func permute(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	perms := []string{}
	head := string(s[0])
	tail := s[1:]

	for _, perm := range permute(tail) {
		for i := 0; i < len(s); i++ {
			newperm := perm[:i] + head + perm[i:]
			perms = append(perms, newperm)
		}
	}
	return perms
}

func LargestPandigitalPrime() int {
	max := 0
	digits := "897654321"
	for i := 0; i < len(digits); i++ {
		for _, perm := range permute(digits[i:]) {
			i := atoi(perm)
			if isPrime(i) {
				if i > max {
					max = i
				}
			}
		}
		if max > 0 {
			break
		}
	}
	return max
}

func main() {
	fmt.Println("Pandigital Primes")

	pandigitalPrime := LargestPandigitalPrime()
	fmt.Println(pandigitalPrime)
}
