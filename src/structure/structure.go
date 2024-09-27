package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Honda", year: 2022}
	fmt.Println(myCar)

	// another way
	var otherCar car
	otherCar.brand = "Dodge"
	// otherCar.year = 2020
	fmt.Println(otherCar)
}
