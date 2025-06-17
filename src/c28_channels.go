package main

import "fmt"

// "c<-" ENTRADA
// "<-c" SALIDA
func say(text string, c chan <- string) {
	c <- text // Lo que llega en el parametro text, lo asignamos al channel c
}

func main() {
	c := make(chan string, 1) // Creacion de un channel (tipo de dato, cantidad de datos simultaneos a manejar)

	fmt.Println("Hello")

	go say("bye", c)

	fmt.Println(<-c) // Extraemos el contenido del channel
}