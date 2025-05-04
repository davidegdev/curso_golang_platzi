package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2

}

func main() {

	a := 50
	b := &a // Save memory direction

	fmt.Println(a)
	fmt.Println(b)  // Print memory direction
	fmt.Println(*b) // get memory value

	myPC := pc {ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)
	myPC.ping()

	fmt.Println(myPC.ram)
	myPC.duplicateRAM()

	fmt.Println(myPC.ram)
	myPC.duplicateRAM()

	fmt.Println(myPC.ram)
}
