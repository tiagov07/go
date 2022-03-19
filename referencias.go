package main

import "fmt"

func equivalenciaEnPies(altura *float32) float32 { // un puntero es indicado con * esto para referenciar una variable
	// y optimizar espacio en memoria
	*altura = *altura * 3.28
	return *altura
}

func main() {
	var altura float32
	fmt.Println("Cual es tu altura en mts ? ")
	fmt.Scanf("%f", &altura)
	fmt.Println(" la altura es : ", altura, "mts")
	fmt.Println(" la altura es en pies : ", equivalenciaEnPies(&altura), "ft") //pero al momento de pasarle el parametro
	//a la funcion se le debe pasar es la referencia de la variable que se va a utilizar
}
