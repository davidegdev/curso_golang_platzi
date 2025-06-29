package main

import (
	"github.com/labstack/echo"
)

func main(){
	
	// Instance echo
	e := echo.New()

	// Create a new route
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323")) // Start the server
}

// COMANDOS TERMINAL

// go get github.com/labstack/echo