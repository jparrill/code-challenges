/*
In this challenge, we are going to be looking at how you can effectively filter out the duplicate entries from a slice in Go.

The task will be to implement the FilterDuplicates function so that it returns a slice of type string which contains only the unique names of developers retrieved from the inputted list.

Example:

// input
[]Developers{
  Developer{Name: "Elliot"},
  Developer{Name: "Alan"},
  Developer{Name: "Jennifer"},
  Developer{Name: "Graham"},
  Developer{Name: "Paul"},
  Developer{Name: "Alan"},
}

// output
[]string{
  "Elliot",
  "Alan",
  "Jennifer",
  "Graham",
  "Paul",
}
*/
package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	uniqFiltr := make(map[string]bool)
	var uniqList []string

	for _, dev := range developers {
		if _, exists := uniqFiltr[dev.Name]; !exists {
			uniqFiltr[dev.Name] = true
			uniqList = append(uniqList, dev.Name)
		}
	}
	return uniqList
}

func main() {
	fmt.Println("Filter Unique Challenge")
}
