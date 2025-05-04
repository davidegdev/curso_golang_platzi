package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("PC brand: %s, with RAM: %d, and Disk: %d", myPC.brand, myPC.ram, myPC.disk)
}

func main() {

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)

}
