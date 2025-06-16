package main

import (
	"fmt"
	"sync"
	// "time"
)

func say(text string, wg *sync.WaitGroup) {

	// Cuando la ejecucion de la funcin "say" termine, libera esa goroutine del waitGroup
	defer wg.Done()

	fmt.Println(text)
}

func main(){
	// importamos un paquete sync que nos permite interactuar de forma primitiva con las go routines
	var wg sync.WaitGroup // <-- Esta variable "wg" nos sirve para agrupar goroutines.

	fmt.Println("Hello")

	// Agregamos una goroutine al waitGroup para que la goroutine base (main) la espere antes de que termine.
	wg.Add(1)
	go say("World", &wg)

	// Le decimos a la goroutine principal "main" que espere a la ejecucion de las gotoutines agregadas al waitGroup.
	wg.Wait()

	// time.Sleep(time.Second * 1)
}