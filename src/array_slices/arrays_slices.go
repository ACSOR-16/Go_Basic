package main

import "fmt"

func main() {
	// array
	var myArray [4]int
	myArray[0] = 1
	myArray[3] = 3
	fmt.Println(myArray, len(myArray), cap(myArray))

	// slices
	var slices = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slices, len(slices), cap(slices))

	// slices methods
	fmt.Println(slices[0])
	fmt.Println(slices[:3])
	fmt.Println(slices[2:4])
	fmt.Println(slices[4:])

	// append
	slices = append(slices, 7)
	fmt.Println(slices)

	// append in new list
	var newList = []int{8, 9, 10}
	slices = append(slices, newList...)
	fmt.Println(slices)
}
