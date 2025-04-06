package main

import "fmt"

func main(){

	slice := []string {"hola", "que", "hace"}

	for i, _ := range slice {
		fmt.Println(i)
	}

}