package main

import "fmt"

func main() {
	// conditional for
	var count int = 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < count {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever < 5 {
			break
		}
	}

	// for range
	numberList := []int{1, 2, 2, 4}
	for index, value := range numberList {
		fmt.Println(index, value)
	}
}
