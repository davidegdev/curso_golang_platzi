package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["David"] = 27
	m["Daniel"] = 24

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// encontrar valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)

}
