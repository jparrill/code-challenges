/*
In this challenge, you are going to implement the necessary methods needed to satisfy the provided Go interface.

On the left hand screen, you have a simple Go application that features an interface called Employee.

In order to complete this challenge, you will have to complete the code and satisfy this interface.
*/

package main

import "math/rand"

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func Random() int {
	return rand.Int()
}

func (e Engineer) Age() int {
	return Random()
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func main() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
}
