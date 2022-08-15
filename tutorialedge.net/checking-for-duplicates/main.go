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
