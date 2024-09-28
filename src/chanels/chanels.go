package main

import "fmt"

func say(text string, character chan<- string) {
	character <- text
}
func main() {
	character := make(chan string, 1)
	fmt.Println("hello")

	go say("bye", character)
	fmt.Println(<-character)
}
