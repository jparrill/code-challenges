/*
👋 Welcome Gophers! In this challenge, we are going to be implementing some of the basic functionality of the Queue data structure in Go.

This is going to be the first of a number of data-structure questions which may come in handy if you are about to go in for an interview!

We’ll be carrying on the theme of flying from the previous challenge here and implementing 3 crucial methods needed to support a basic implementation of a Queue.

Push
The first challenge will be to implement the Push function of our Queue data structure.

This method will take in a Flight and push the flight onto the back of our Items queue.

Peek
The second part of this challenge will be implementing the Peek function.

This method will allow us to view what item is at the front of our queue but not modify the underlying stack values.

Pop
The third and final part of this challenge will be implementing the Pop function.

This method will allow us to pop an element off the front of our Items queue and return to us the first flight.
*/
package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("The Queue is empty")
	}
	flight := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]

	return flight, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("The Queue is empty")
	}

	return q.Items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return false
}

func main() {
	fmt.Println("Go Queue Implementation")
}
