package main

import "fmt"

func main() {
	var name string
	var weight float32
	fmt.Println("compartenos tu name")
	fmt.Scanf("%s", &name)
	fmt.Println("compartenos tu peso")
	fmt.Scanf("%f", &weight)
	if weight > 80 {
		weight = weight - 8.0
		fmt.Printf("estas empezando a tener exceso de peso %v \n te recomendamos bajar al menos a %v kg", name, weight)
	}

}
