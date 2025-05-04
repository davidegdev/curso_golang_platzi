package main

import "fmt"

func main(){

	//--defer
	defer fmt.Println("Esto se ejecuta al final de la ejecucion de main() ")
	fmt.Println("Hola Mundo")


	fmt.Println("Por acá, esto es lo ultimo del bloque de codigo")
}