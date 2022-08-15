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
