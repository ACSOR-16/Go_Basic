package main

import "fmt"

func main() {
	// Square
	baseRectangle := 20
	heightRectangle := 10

	areaRectangle := baseRectangle * heightRectangle
	fmt.Println("the area of the rectangle is:", areaRectangle)

	// circle
	const PI float64 = 3.14 // Constant
	var radioCircle float64 = 10

	areaCircle := PI * radioCircle * radioCircle

	fmt.Println("The aArea of the circle is: ", areaCircle)

	// trapezoid
	var baseOne float64 = 6
	var baseTwo float64 = 15
	var heightTrapezoid float64 = 25

	areaTrapezoid := ((baseOne + baseTwo) * heightTrapezoid) / 2

	fmt.Println("The area of the Trapezoid is:", areaTrapezoid)

}
