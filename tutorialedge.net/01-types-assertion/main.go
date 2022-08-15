/*
Challenge
In this challenge, we are going to define a function that is called GetDeveloper which will take in 2 interface{} arguments.

Within this function, you will have to declare a new Developer instance and use type assertion to populate the values correctly before then returning this new Developer instance.
*/

package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	dev := Developer{
		Name: name.(string),
		Age:  age.(int),
	}
	return dev
}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
