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
