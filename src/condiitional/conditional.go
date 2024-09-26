package main

import (
	"fmt"
	"log"
	"strconv"
)

func isPair(number int) bool {
	if module := number % 2; module == 0 {
		return true
	}
	return false
}

func getLogin(username, password string) {
	/* This function compares the username and password and checks if they match
	logging in this way (username and password in the code) is a very bad practice
	usually frameworks use libraries to do this check
	with passwords encrypted from the database.
	*/
	if username == "plat" && password == "1234" {
		fmt.Println("you're logged")
	} else {
		fmt.Println("username or password are incorrect")
	}
}
func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("is 1")
	} else {
		fmt.Println("Isn't 1")
	}

	// With and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("it's true")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("it's true, OR")
	}

	// convert text to number
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	//function is pair and getLogin
	fmt.Println(isPair(3))
	getLogin("plat", "1234")
}
