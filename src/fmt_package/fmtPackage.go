package main

import "fmt"

func main() {
	// declaration of variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(worldMessage, helloMessage)

	// Printf
	name := "Plat-zi"
	Courses := 500
	fmt.Printf("%s have most of %d courses\n", name, Courses)
	fmt.Printf("%v have most of %v courses\n", Courses, name)

	// Sprintf
	message := fmt.Sprintf("%s have most of %d courses", name, Courses)
	fmt.Println(message)

	// date type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("courses: %T\n", Courses)

}
