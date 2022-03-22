package main

import "fmt"

func mountainEco(message string, iterations uint) {
	if iterations > 1 {
		mountainEco(message, iterations-1)
	}
	fmt.Println(message, iterations)
}

func main() {
	mountainEco("Eco!!!", 7)
}
