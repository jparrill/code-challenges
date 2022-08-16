/*
👋 Welcome Gophers! In this challenge, you are tasked with trying to implement a function that checks to see if a set is a sub-set of another set.

We’ll be carrying on the flying theme where the function takes in a `slice` of Flights and then checks to see if they exist within another `slice` of flights.

MORE EFFICIENT SOLUTION THAN SOLUTION 1
*/
package main

import (
	"fmt"
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
	// implement
	set := make(map[Flight]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
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
