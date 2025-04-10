package main

import (
	"fmt"
)

func main(){

	//--- arrays
	var array [4]int
	array[0] = 0
	array[1] = 1
	fmt.Println(array, len(array), cap(array))

	//--- slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//--- metodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	fmt.Println(slice[:0])

	//--- append
	slice = append(slice, 7)
	fmt.Println(slice)

	//--- agregar nuevo slice al slice anterior
	nuevoSlice := []int{8,9,10}
	slice = append(slice, nuevoSlice...)
	fmt.Println(slice)

}