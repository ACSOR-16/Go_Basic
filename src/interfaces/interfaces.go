package main

import (
	"fmt"
)

// interface
type figures2D interface {
	area() float64
}

// structures
type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

// methods
func (square square) area() float64 {
	return square.base * square.base
}

func (rectangle rectangle) area() float64 {
	return rectangle.base * rectangle.height
}

func calculate(figure figures2D) {
	fmt.Println("Areas:", figure.area())
}

// FUNCTION MAIN
func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 4, height: 3}

	calculate(mySquare)
	calculate(myRectangle)

	// interface list
	myInterface := []interface{}{"sheet", 12, 3.14}
	fmt.Println(myInterface...)
}
