package main

import "fmt"

func main() {
	//-- Declaring a constant
	const pi float64 = 3.14
	const pi2 = 3.1415

	//-- Declaring variables
	base := 12 				//-- := is used to declare and assign a value to a variable
	var altura int = 14		//-- declare variable with type and instance
	var area int			//-- declare variable with type with zero value

	fmt.Println("Hello world")
	fmt.Println("Pi: ", pi)
	fmt.Println("Pi2: ", pi2)
	fmt.Println("Base: ", base)
	fmt.Println("Altura: ", altura)
	fmt.Println("Area: ", area)

	//-- Zero value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	//-- Area of square
	const baseSquare int = 10
	areaSquare := baseSquare * baseSquare
	fmt.Println("Area of square: ", areaSquare)

}
