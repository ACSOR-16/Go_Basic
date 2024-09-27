package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("hello")
	fmt.Println("world")

	// continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// continue
		if i == 2 {
			fmt.Println("it's 2")
		}
		// break
		if i == 8 {
			break
		}
	}
}
