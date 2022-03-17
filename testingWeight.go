package main

import "fmt"

func main() {
	var nombre string
	var weight float32
	fmt.Println("compartenos tu nombre")
	fmt.Scanf("%s", &nombre)
	fmt.Println("compartenos tu peso")
	fmt.Scanf("%f", &weight)
	if weight > 80 {
		fmt.Println("estas empezando a tener exceso de peso", nombre, "te recomendar bajar al menos a ", weight-5.0, "kg")
	}

}
