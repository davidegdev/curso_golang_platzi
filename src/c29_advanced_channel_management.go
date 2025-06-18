package main

import "fmt"

func message(text string, c chan string){
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// Tamaño ocupado y capacidad del channel
	fmt.Println(len(c), cap(c))

	// close: lo usamos para cerrar el channel, de esta manera no podrá recibir mas datos
	close(c)

	// range: lo usamos para recorrer los datos del channel. IMPORTANTE: el cannel debe estar cerrado(close) para recorrerlo
	for dato := range c{
		fmt.Println(dato)
	}

	// select: lo usamos para seleccionar cannels
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i:=0;i<2;i++{
		select{
		case m1 := <- email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <- email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}