package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("I have %d GB RAM, %d GB Disk and it's a %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "lenovo", disk: 500}
	// myPc.String()
	fmt.Println(myPc)
}
