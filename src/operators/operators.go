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

	//Integers
	//int = Depends del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimize memory when we know that the data will always be positive
	//uint = Depends del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//decimal
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//texts and booleans
	//string = ""
	//bool = true or false

	//complex
	//Complex64 = Real and Imaginary float32
	//Complex128 = Real and Imaginary float64
	//example : c:=10 + 8i
}
