package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["Jose"] = 26
	myMap["Jefferson"] = 24

	fmt.Println(myMap)

	// loop
	for index, value := range myMap {
		fmt.Println(index, value)
	}

	// find value
	value, ok := myMap["ose"]
	fmt.Println(value, ok)
}
