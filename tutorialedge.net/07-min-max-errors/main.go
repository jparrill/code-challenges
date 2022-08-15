/*
👋 Welcome Gophers! In this challenge, we’ll be testing your ability to handle edge cases and work with the errors package.

In this challenge, you have the task of implementing the GetMinMax function which will take in a slice of type Flight.

The function will return 3 values:

min - The cheapest price from the list passed in.
max - The most expensive price from the list passed in.
error - Any error values should there be issues with the list passed in.
*/
package main

import (
	"errors"
	"fmt"
	"sort"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Flights []Flight

func (fs Flights) Len() int           { return len(fs) }
func (fs Flights) Swap(x, y int)      { fs[x], fs[y] = fs[y], fs[x] }
func (fs Flights) Less(x, y int) bool { return fs[x].Price < fs[y].Price }

func GetMinMax(flights Flights) (int, int, error) {
	var min, max int
	if len(flights) < 3 {
		return min, max, errors.New("The slice does not have enough elements to return the desired values")
	}
	sort.Sort(flights)
	fmt.Println(flights)

	min = flights[0].Price
	max = flights[len(flights)-1].Price

	return min, max, nil
}

func main() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
}
