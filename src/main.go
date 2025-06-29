package main

import (
	"fmt"

	//"github.com/labstack/echo"
)

func doubleReturn(a int) (c, d int) { return a, a * 2 };

func main(){

	var num uint8 = 200
	
	// Instance echo
	/*e := echo.New()

	// Create a new route
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323")) // Start the server*/

	fmt.Printf("2, %T", 2)
	fmt.Println()
	fmt.Println("num: ", num)
	x,y := doubleReturn(2)

	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)
	
}

// COMANDOS TERMINAL

// go get github.com/labstack/echo