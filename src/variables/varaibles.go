package main

import "fmt"

func main() {
	// declaration of constants
	const pi float64 = 3.14
	const pi2 = 3.12

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	// declaration of integer variables
	// Go does not compile if variables are not used
	base := 12
	var height = 10
	var area int

	fmt.Println(base, height, area)

	// Zero values
	// GO assigns values ​​to empty variables
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// exercise
	const baseSquare = 10
	areaSquare := baseSquare * baseSquare

	fmt.Println("the area of square is:", areaSquare)
}
