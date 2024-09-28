package main

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc Pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

func (myPc *Pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	// memory space
	a := 50
	// pointer
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	// structure Pc
	myPc := Pc{ram: 16, disk: 200, brand: "Rog"}
	fmt.Println(myPc)

	myPc.ping()
	fmt.Println(myPc)

	// duplicate Ram method
	myPc.duplicateRAM()
	fmt.Println(myPc)

	myPc.duplicateRAM()
	fmt.Println(myPc)
}
