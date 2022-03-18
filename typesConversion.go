package main

import (
	"fmt"
	"strconv"
)

func main() {
	mayorEdad := "true"
	mayorEdadBool, _ := strconv.ParseBool(mayorEdad) // el _ se usa para que en caso de tener un error se go lo omita
	// eso lo hacemos cuando no le queremos dar manejo al error
	fmt.Printf(" ahora la variable mayorEdadBool es : %v y de tipo %T  \n", mayorEdadBool, mayorEdadBool)

	mayorEdadstring := strconv.FormatBool(mayorEdadBool) // esto con el fin de parsear un booleando a un string
	fmt.Printf(" ahora la variable mayorEdadstring es : %v y de tipo %T \n", mayorEdadstring, mayorEdadstring)
}
