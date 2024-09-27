package main

import (
	pk "example/helloWorld/src/myPackage"
	"fmt"
)

func main() {
	var myOtherCar pk.CarPublic
	myOtherCar.Brand = "Chevrolet"
	fmt.Println(myOtherCar)

	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("124")
	// pk.printMessageOne("356")
}
