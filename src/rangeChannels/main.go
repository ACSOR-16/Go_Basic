package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}
func main() {
	channel := make(chan string, 2)
	channel <- "message 1"
	channel <- "message 2"

	fmt.Println(len(channel), cap(channel))
	// range and close
	close(channel)

	// channel <- message 3

	for message := range channel {
		fmt.Println(message)
	}

	// select
	emailOne := make(chan string)
	emailTwo := make(chan string)

	go message("message One", emailOne)
	go message("message two", emailTwo)

	for index := 0; index < 2; index++ {
		select {
		case mes1 := <-emailOne:
			fmt.Println("Email received to emailOne", mes1)
		case mes2 := <-emailTwo:
			fmt.Println("Email received to emailTwo", mes2)
		}
	}
}
