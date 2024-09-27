package main

import "fmt"

func main() {
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	// without condition
	value := 50
	switch {
	case value > 100:
		fmt.Println("It is greater than 100")
	case value < 0:
		fmt.Println("It is less than 0")
	default:
		fmt.Println("no condition")
	}
}
