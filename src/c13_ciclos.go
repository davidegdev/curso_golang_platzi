package main

import "fmt"

func main(){

	// For condicional
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For while
	counter := 0
	for counter < 10{
		fmt.Println(counter)
		counter++
	}

	// For while
	counterMax := 10
	for counterMax > 0{
		fmt.Println("CounterMax:", counterMax)
		counterMax--
	}
}