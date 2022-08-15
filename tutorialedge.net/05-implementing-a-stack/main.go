/*
👋 Welcome Gophers! In this challenge, we are going to be implementing some of the basic functionality of the Stack data structure in Go.

This is going to be the first of a number of data-structure questions which may come in handy if you are about to go in for an interview!

We’ll be carrying on the theme of flying from the previous challenge here and implementing 3 crucial methods needed to support a basic implementation of a stack.

Push
The first challenge will be to implement the Push function of our Stack interface.

This method will take in a Flight and push the flight onto the top of our Items stack.

Peek
The second part of this challenge will be implementing the Peek function.

This method will allow us to view what item is at the top of our stack but not modify the underlying stack values.

Pop
The third and final part of this challenge will be implementing the Pop function.

This method will allow us to pop an element off the top of our Items stack and return to us the top flight.
*/

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
