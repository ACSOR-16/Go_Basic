package main

import (
	"fmt"
)

func functions(message string) {
	fmt.Println(message)
}

func tripleArgs(numberOne, numberTwo int, text string) {
	fmt.Println(numberOne, numberTwo, text)
}

func returnValue(number int) int {
	return number * 2
}

func doubleReturn(number int) (numberReturned, numberReturnedTwo int) {
	return number, number * 2
}

func main() {
	functions("Hello world")
	tripleArgs(2, 5, "sock")

	var secretValue = returnValue(4)
	fmt.Println("value: ", secretValue)

	var valueOne, _ = doubleReturn(2)
	fmt.Println("first value: ", valueOne)

	valueA, valueB := doubleReturn(3)
	fmt.Println(valueA, valueB)
}
