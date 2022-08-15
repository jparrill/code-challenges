package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Flights []Flight

func (fs Flights) Len() int           { return len(fs) }
func (fs Flights) Swap(x, y int)      { fs[x], fs[y] = fs[y], fs[x] }
func (fs Flights) Less(x, y int) bool { return fs[x].Price > fs[y].Price }

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights Flights) Flights {
	sort.Sort(flights)
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	var flights []Flight

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
