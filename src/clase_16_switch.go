package main

import "fmt"

func main() {

	//-- Con condicion
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//--Sin condicion
	result := 1
	switch {
	case result > 5:
		fmt.Println("Es mayor a 5")
	case result < 0:
		fmt.Println("Es menor a cero")
	default:
		fmt.Println("Fuera de condiciones")
	}
}