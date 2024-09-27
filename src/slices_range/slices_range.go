package main

import (
	"fmt"
	"strings"
)

func main() {
	var mySlice = []string{"cat", "dog", "bug", "pork"}
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	// ama = it's palindrome
	// amor to roma = it isn't palindrome
	isPalindrome("dabale arroz a la zorra el abad")
}

func isPalindrome(text string) {
	var textReverse string

	// convert text to lower
	text = strings.ToLower(text)
	var comparedTextOne = strings.Replace(text, " ", "", -1)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
		fmt.Println(textReverse)
	}

	var comparedText string = strings.Replace(textReverse, " ", "", -1)
	fmt.Println(comparedText)

	if comparedTextOne == comparedText {
		fmt.Println("it's palindrome")
	} else {
		fmt.Println("it isn't palindrome")
	}
}
