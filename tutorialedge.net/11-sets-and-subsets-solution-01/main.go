/*
👋 Welcome Gophers! In this challenge, you are tasked with trying to implement a function that checks to see if a set is a sub-set of another set.

We’ll be carrying on the flying theme where the function takes in a `slice` of Flights and then checks to see if they exist within another `slice` of flights.

UNNEFFICIENT SOLUTION
*/

package main

import (
	"fmt"
	"reflect"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	val := 0
	checks := len(first)
	for _, ff := range first {
		for _, sf := range second {
			if reflect.DeepEqual(ff, sf) {
				val++
				break
			}
		}
	}
	if val == checks {
		return true
	}

	return false
}

func main() {
	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
		Flight{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	fmt.Println(subset)
}
