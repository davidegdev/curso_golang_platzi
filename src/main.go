package main

import "fmt"

func main() {

	//---------------------------------------------------------------------------------------------- FIRSTS CLASSES

	//-- Declaring a constant
	const pi1 float64 = 3.14
	const pi2 = 3.1415

	//-- Declaring variables
	base := 12 				//-- := is used to declare and assign a value to a variable
	var altura int = 14		//-- declare variable with type and instance
	var area int			//-- declare variable with type with zero value

	fmt.Println("Hello world")
	fmt.Println("Pi1: ", pi1)
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

	//---------------------------------------------------------------- CLASS #8 - Operadores aritméticos
	//-- Aritmetic Operations
	//-- Add
	x:=10
	y:=40

	result:= x + y
	fmt.Println("Add Result:", result)

	//-- Subtract
	result = y - x
	fmt.Println("Subtract Result:", result)

	//-- Multiply
	result = x * y
	fmt.Println("Multiply Result:", result)

	//-- Divide
	result = y / x
	fmt.Println("Divide Result:", result)

	//-- Modulus
	result = y % x
	fmt.Println("Modulus Result:", result)

	//-- Calculate Area of a rectangle
	baseRect := 20
	alturaRect := 10
	areaRect := baseRect * alturaRect
	fmt.Println("Area of rectangle: ", areaRect)

	//-- Calculate Area of a circle
	radio := 10
	pi := 3.1416
	areaCircle := pi * float64(radio * radio)
	fmt.Println("Area of circle: ", areaCircle)

	//-- Calculate Area of a Trapeze
	baseBig := 10
	baseSmall := 6
	heightTrapeze := 4
	areaTrapeze := ((baseBig + baseSmall) * heightTrapeze) / 2
	fmt.Println("Area of trapeze: ", areaTrapeze)

	//---------------------------------------------------------------- CLASS #10 - Paquete fmt: algo más que imprimir en consola
	helloMessage := "Hello"
	worldMessage := "world"
	fmt.Println(helloMessage, worldMessage)
	
	//-- 														Printf
	//-- %s string
	//-- %d int
	//-- %f float
	//-- %v any type
	//-- %T type of variable
	name := "Platzi"
	courses := 500
	fmt.Printf("%s has more than %d courses\n", name, courses)
	fmt.Printf("%v has more than %v courses\n", name, courses)

	//--														Sprintf

	message := fmt.Sprintf("%s has more than %d courses", name, courses)
	fmt.Println(message)		

	//--														Type of variables
	fmt.Printf("%T\n", helloMessage)		
	fmt.Printf("%T\n", courses)	

}
