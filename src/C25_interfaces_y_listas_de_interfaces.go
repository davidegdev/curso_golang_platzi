package main

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64{
	return c.base * c.base
}

func main() {
	myCuadrado := cuadrado{base: 2}
}