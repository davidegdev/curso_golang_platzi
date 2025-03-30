package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValuePerTwo(a int) int{
	return a * 2
}

func doubleReturn(a int) (b, c int){
	return a, a*2
}


func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValuePerTwo(3)
	fmt.Println("Value: ", value)
	value1, value2 := doubleReturn(3)
	fmt.Println("Value1 & value2: ", value1, value2)

	// Usar solamente una de las variables del return de una funcion
	_, value3 := doubleReturn(4)
	fmt.Println("Value3:", value3)
}