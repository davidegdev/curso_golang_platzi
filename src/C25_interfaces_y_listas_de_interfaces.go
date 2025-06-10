package main

import "fmt"


type figuras2D interface{
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base float64
	altura float64
}

func (c cuadrado) area() float64{
	return c.base * c.base
}

func (r rectangulo) area() float64{
	return r.altura * r.base
}

func calcularArea(f figuras2D){
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcularArea(myCuadrado)
	calcularArea(myRectangulo)

	// Lista interfaces. Listas multitipo o de multiples tipos.
	myInterface := []interface{}{"Hola", 12, 4.25}
	fmt.Println(myInterface...)
}