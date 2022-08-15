package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("The Stack is empty")
	}

	Flight := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]

	return Flight, nil
}

func (s *Stack) Push(f Flight) {
	s.Items = append(s.Items, f)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("The Stack is empty")
	}

	return s.Items[len(s.Items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return false
}

func main() {
	fmt.Println("Go Stack Implementation")
}
