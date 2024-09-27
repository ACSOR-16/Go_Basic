package myPackage

import "fmt"

// CarPublic Car with access public
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// methods
// methods public
func PrintMessage(text string) {
	fmt.Println(text)
}

// method private
func printMessageOne(text string) {
	fmt.Println(text)
}
